package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type MessageType uint

const (
	UNKNOWN = iota
	_
	CALL
	CALLRESULT
	CALLERROR
)

type OCPPCallFrame struct {
	MessageType MessageType
	MessageId   string
	Action      string
	Payload     json.RawMessage
}

func (frame *OCPPCallFrame) MarshalJSON() ([]byte, error) {
	var v = []interface{}{
		CALL,
		frame.MessageId,
		frame.Action,
		frame.Payload,
	}
	return json.Marshal(v)
}

func (frame *OCPPCallFrame) UnmarshalJSON(data []byte) error {
	tmp := []interface{}{&frame.MessageType, &frame.MessageId, &frame.Action, &frame.Payload}
	wantLen := len(tmp)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CALL frame: %d != %d", g, e)
	}

	return nil
}

type OCPPCallResultFrame struct {
	MessageType MessageType
	MessageId   string
	Payload     json.RawMessage
}

func (frame *OCPPCallResultFrame) MarshalJSON() ([]byte, error) {
	var v = []interface{}{
		CALLRESULT,
		frame.MessageId,
		frame.Payload,
	}
	return json.Marshal(v)
}

func (frame *OCPPCallResultFrame) UnmarshalJSON(data []byte) error {
	tmp := []interface{}{&frame.MessageType, &frame.MessageId, &frame.Payload}
	wantLen := len(tmp)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CALLRESULT frame: %d != %d", g, e)
	}

	return nil
}

type OCPPCallErrorFrame struct {
	MessageType      MessageType
	MessageId        string
	ErrorCode        string
	ErrorDescription string
	ErrorDetails     json.RawMessage
}

func ErrorFrameWithError(err error) OCPPCallErrorFrame {
	return OCPPCallErrorFrame{
		MessageType:      CALLERROR,
		ErrorDescription: err.Error(),
	}
}

func (frame *OCPPCallErrorFrame) MarshalJSON() ([]byte, error) {
	var v = []interface{}{
		CALLERROR,
		frame.MessageId,
		frame.ErrorCode,
		frame.ErrorDescription,
		frame.ErrorDetails,
	}
	return json.Marshal(v)
}

func (frame *OCPPCallErrorFrame) UnmarshalJSON(data []byte) error {
	tmp := []interface{}{&frame.MessageType, &frame.MessageId, &frame.ErrorCode, &frame.ErrorDescription, &frame.ErrorDetails}
	wantLen := len(tmp)
	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}
	if g, e := len(tmp), wantLen; g != e {
		return fmt.Errorf("wrong number of fields in CALLERROR frame: %d != %d", g, e)
	}

	return nil
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func HandleOCPP(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	for {
		// Read message from browser
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		// Print the message to the console
		fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

		// Write message back to browser
		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	}
}

func main() {
	http.HandleFunc("/csms/ocpp/test1", HandleOCPP)
	http.ListenAndServe(":8080", nil)
}
