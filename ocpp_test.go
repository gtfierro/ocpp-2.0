package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestFrames(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/csms/ocpp/test1", HandleOCPP)
	s := httptest.NewServer(mux)
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.1
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	url := u + "/csms/ocpp/test1"
	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

}

func TestMarshalUnmarshal(t *testing.T) {
	require := require.New(t)

	// create OCPPCallFrame
	frame := OCPPCallFrame{
		MessageId: "123",
		Action:    "TEST",
		Payload:   nil,
	}

	b, err := frame.MarshalJSON()
	require.NoError(err)

	fmt.Println(string(b))
	frame2 := new(OCPPCallFrame)
	err = frame2.UnmarshalJSON(b)
	require.NoError(err)

	require.Equal(frame.MessageId, frame2.MessageId)
	require.Equal(frame.Action, frame2.Action)
	require.Equal(frame.Payload, frame2.Payload)
}
