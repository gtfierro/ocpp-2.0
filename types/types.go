package types

import (
	"github.com/gtfierro/ocpp2.0/enums"
)

type SetVariableDataType struct {
	AttributeType  enums.AttributeEnumType `json:"attributeType"`
	AttributeValue string                  `json:"attributeValue"`
	Component      ComponentType           `json:"component"`
	Variable       VariableType            `json:"variable"`
}

type SetVariableResultType struct {
	AttributeType  enums.AttributeEnumType         `json:"attributeType"`
	AttributeValue enums.SetVariableStatusEnumType `json:"attributeStatus"`
	Component      ComponentType                   `json:"component"`
	Variable       VariableType                    `json:"variable"`
}

type ComponentType struct {
	Name     string   `json:"name"`
	Instance string   `json:"instance"`
	Evse     EVSEType `json:"evse"`
}

type EVSEType struct {
	Id          int `json:"id"`
	ConnectorId int `json:"connectorid"`
}

type VariableType struct {
	Name     string `json:"name"`
	Instance string `json:"instance"`
}

type GetVariableDataType struct {
	AttributeType enums.AttributeEnumType `json:"attributeType"`
	Component     ComponentType           `json:"component"`
	Variable      VariableType            `json:"variable"`
}

type GetVariableResultType struct {
	AttributeStatus enums.GetVariableStatusEnumType `json:"attributeStatus"`
	AttributeType   enums.AttributeEnumType         `json:"attributeType"`
	AttributeValue  string                          `json:"attributeValue"`
	Component       ComponentType                   `json:"component"`
	Variable        VariableType                    `json:"variable"`
}

type TransactionType struct {
	Id                string
	ChargingState     enums.ChargingStateEnumType
	TimeSpentCharging int
	StoppedReason     enums.ReasonEnumType
	RemoteStartId     int
}

type IdTokenType struct {
	IdToken string                `json:"idToken"`
	Type    enums.IdTokenEnumType `json:"type"`
	//AdditionalInfo AdditionalInfoType
}

//type MeterValueType struct {
//	// rfc3339
//	Timestamp    string             `json:"timestamp"`
//	SampledValue []SampledValueType `json:"sampledValue"`
//}
