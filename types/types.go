package types

type SetVariableDataType struct {
	AttributeType  AttributeEnumType `json:"attributeType"`
	AttributeValue string            `json:"attributeValue"`
	Component      ComponentType     `json:"component"`
	Variable       VariableType      `json:"variable"`
}

type SetVariableResultType struct {
	AttributeType  AttributeEnumType         `json:"attributeType"`
	AttributeValue SetVariableStatusEnumType `json:"attributeStatus"`
	Component      ComponentType             `json:"component"`
	Variable       VariableType              `json:"variable"`
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
	AttributeType AttributeEnumType `json:"attributeType"`
	Component     ComponentType     `json:"component"`
	Variable      VariableType      `json:"variable"`
}

type GetVariableResultType struct {
	AttributeStatus GetVariableStatusEnumType `json:"attributeStatus"`
	AttributeType   AttributeEnumType         `json:"attributeType"`
	AttributeValue  string                    `json:"attributeValue"`
	Component       ComponentType             `json:"component"`
	Variable        VariableType              `json:"variable"`
}

type TransactionType struct {
	Id                string
	ChargingState     ChargingStateEnumType
	TimeSpentCharging int
	StoppedReason     ReasonEnumType
	RemoteStartId     int
}

type IdTokenType struct {
	IdToken string          `json:"idToken"`
	Type    IdTokenEnumType `json:"type"`
	//AdditionalInfo AdditionalInfoType
}

//type MeterValueType struct {
//	// rfc3339
//	Timestamp    string             `json:"timestamp"`
//	SampledValue []SampledValueType `json:"sampledValue"`
//}
