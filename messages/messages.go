package messages

import (
	"github.com/gtfierro/ocpp2.0/types"
	"time"
)

type SetVariablesRequest struct {
	SetVariableData []types.SetVariableDataType `json:"setVariableData"`
}

type SetVariablesResponse struct {
	SetVariableResult []types.SetVariableResultType `json:"setVariableResult"`
	//Csr               string                  `json:"csr"`
	// TODO: TypeOfCertificate
}

type GetVariablesRequest struct {
	GetVariableData []types.GetVariableDataType `json:"setVariableData"`
}

type GetVariablesResponse struct {
	GetVariableResult []types.GetVariableResultType `json:"setVariableResult"`
	Csr               string                        `json:"csr"`
	// TODO: TypeOfCertificate
}

type StatusNotificationRequest struct {
	// rfc3339
	Timestamp       time.Time                     `json:"timestamp"`
	ConnectorStatus types.ConnectorStatusEnumType `json:"connectorStatus"`
	EvseId          int                           `json:"evseId"`
	ConnectorId     int                           `json:"connectorId"`
}

type StatusNotificationResponse struct {
	// no fields
}

type TransactionEventRequest struct {
	EventType types.TransactionEventEnumType `json:"eventType"`
	// rfc3339
	Timestamp          time.Time                   `json:"timestamp"`
	TriggerReason      types.TriggerReasonEnumType `json:"triggerReason"`
	SeqNo              int                         `json:"seqno"`
	Offline            bool                        `json:"offline"`
	NumberOfPhasesUsed int                         `json:"numberOfPhasesUsed"`
	CableMaxCurrent    float64                     `json:"cableMaxCurrent"`
	ReservationId      int                         `json:"reservationId"`
	TransactionData    types.TransactionType       `json:"transactionData"`
	IdToken            types.IdTokenType           `json:"idToken"`
	Evse               types.EVSEType              `json:"evse"`
	//MeterValue         MeterValueType        `json:"meterValue"`
}

type TransactionEventResponse struct {
	// SHALL only be sent when charging has ended. Final total cost of this transaction, including taxes. In the currency configured with the Configuration Variable: Currency. When omitted, the transaction was NOT free. To indicate a free transaction, the CSMS SHALL send 0.00.
	TotalCost float64 `json:"totalCost"`
	// Priority from a business point of view. Default priority is 0, The range is from -9 to 9. Higher values indicate a higher priority. The chargingPriority in TransactionEventResponse is temporarily, so it may not be set in the IdTokenInfoType afterwards. Also the chargingPriority in TransactionEventResponse has a higher priority than the one in IdTokenInfoType.
	ChargingPriority int `json:"chargingPriority"`
	// This contains information about authorization status, expiry and group id. Is required when the transactionEventRequest contained an idToken.
	// TODO: add
	//IdTokenInfo IdTokenInfoType `json:"idTokenInfo"`
	// TODO: add
	// This can contain updated personal message that can be shown to the EV Driver. This can be used to provide updated tariff information .
	//UpdatedPersonalMessage MessageContentType `json:"updatedPersonalMessage"`
}

type RequestStartTransactionRequest struct {
}

type RequestStartTransactionResponse struct {
}
