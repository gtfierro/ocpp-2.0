package main

import (
	"context"
	"encoding/json"
	"github.com/gtfierro/ocpp2.0/messages"
	"sync"
)

type ChargingStationConfig struct {
}

type ChargingStation struct {
	variables map[string]string
	varlock   sync.RWMutex
}

func NewChargingStation(cfg ChargingStationConfig) (*ChargingStation, error) {
	return &ChargingStation{}, nil
}

func (station *ChargingStation) CallFunction(ctx context.Context, frame OCPPCallFrame) (*OCPPCallResultFrame, *OCPPCallErrorFrame) {
	// frame tells us which function to call

	// TODO: validate

	// TODO: track message id (in context?)
	ctx = context.WithValue(ctx, "messageid", frame.MessageId)

	var response interface{}
	var err error

	// for each possible action, we unmarshal the payload of the call frame into a struct
	// and then call the relevant function
	switch frame.Action {
	case "SetVariables":
		var request messages.SetVariablesRequest
		err = json.Unmarshal(frame.Payload, &request)
		if err != nil {
			break
		}
		response, err = station.SetVariables(ctx, request)
		if err != nil {
			break
		}
	case "GetVariables":
		var request messages.GetVariablesRequest
		err = json.Unmarshal(frame.Payload, &request)
		if err != nil {
			break
		}
		response, err = station.GetVariables(ctx, request)
		if err != nil {
			break
		}
	}

	if err != nil {
		errorframe := ErrorFrameWithError(err)
		return nil, &errorframe
	}

	var resultframe OCPPCallResultFrame

	resultframe.Payload, err = json.Marshal(response)

	return &resultframe, nil

}

func (station *ChargingStation) SetVariables(ctx context.Context, request messages.SetVariablesRequest) (messages.SetVariablesResponse, error) {

	return messages.SetVariablesResponse{}, nil
}

func (station *ChargingStation) GetVariables(ctx context.Context, request messages.GetVariablesRequest) (messages.GetVariablesResponse, error) {
	return messages.GetVariablesResponse{}, nil
}
