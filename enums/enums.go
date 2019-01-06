package enums

type AttributeEnumType string

var (
	ACTUAL_AttributeEnumType AttributeEnumType = "Actual"
	TARGET_AttributeEnumType AttributeEnumType = "Target"
	MINSET_AttributeEnumType AttributeEnumType = "MinSet"
	MAXSET_AttributeEnumType AttributeEnumType = "MaxSet"
)

type SetVariableStatusEnumType string

var (
	ACCEPTED_SetVariableStatusEnumType                     SetVariableStatusEnumType = "Accepted"
	REJECTED_SetVariableStatusEnumType                                               = "Rejected"
	INVALID_VALUE_SetVariableStatusEnumType                                          = "InvalidValue"
	UNKNOWN_COMPONENT_SetVariableStatusEnumType                                      = "UnknownComponent"
	UNKNOWN_VARIABLE_SetVariableStatusEnumType                                       = "UnknownVariable"
	NOT_SUPPORTED_ATTRIBUTE_TYPE_SetVariableStatusEnumType                           = "NotSupportedAttributeType"
	OUT_OF_RANGE_SetVariableStatusEnumType                                           = "OutOfRange"
	REBOOT_REQUIRED_SetVariableStatusEnumType                                        = "RebootRequired"
)

type GetVariableStatusEnumType string

var (
	ACCEPTED_GetVariableStatusEnumType                     GetVariableStatusEnumType = "Accepted"
	REJECTED_GetVariableStatusEnumType                     GetVariableStatusEnumType = "Rejected"
	UNKNOWN_COMPONENT_GetVariableStatusEnumType            GetVariableStatusEnumType = "UnknownComponent"
	UNKNOWN_VARIABLE_GetVariableStatusEnumType             GetVariableStatusEnumType = "UnknownVariable"
	NOT_SUPPORTED_ATTRIBUTE_TYPE_GetVariableStatusEnumType GetVariableStatusEnumType = "NotSupportedAttributeType"
)

type ConnectorStatusEnumType string

var (
	AVAILABLE_ConnectorStatusEnumType   ConnectorStatusEnumType = "Available"
	OCCUPIED_ConnectorStatusEnumType    ConnectorStatusEnumType = "Occupied"
	RESERVED_ConnectorStatusEnumType    ConnectorStatusEnumType = "Reserved"
	UNAVAILABLE_ConnectorStatusEnumType ConnectorStatusEnumType = "Unavailable"
	FAULTED_ConnectorStatusEnumType     ConnectorStatusEnumType = "Faulted"
)

type TransactionEventEnumType string

var (
	ENDED_TransactionEventEnumType   TransactionEventEnumType = "Ended"
	STARTED_TransactionEventEnumType TransactionEventEnumType = "Started"
	UPDATED_TransactionEventEnumType TransactionEventEnumType = "Updated"
)

type TriggerReasonEnumType string

var (
	AUTHORIZED_TriggerReasonEnumType           TriggerReasonEnumType = "Authorized"
	CABLEPLUGGEDIN_TriggerReasonEnumType       TriggerReasonEnumType = "CablePluggedIn"
	CHARGINGRATECHANGED_TriggerReasonEnumType  TriggerReasonEnumType = "ChargingRateChanged"
	CHARGINGSTATECHANGED_TriggerReasonEnumType TriggerReasonEnumType = "ChargingStateChanged"
	DEAUTHORIZED_TriggerReasonEnumType         TriggerReasonEnumType = "Deauthorized"
	ENERGYLIMITREACHED_TriggerReasonEnumType   TriggerReasonEnumType = "EnergyLimitReached"
	EVCOMMUNICATIONLOST_TriggerReasonEnumType  TriggerReasonEnumType = "EVCommunicationLost"
	EVCONNECTTIMEOUT_TriggerReasonEnumType     TriggerReasonEnumType = "EVConnectTimeout"
	METERVALUECLOCK_TriggerReasonEnumType      TriggerReasonEnumType = "MeterValueClock"
	METERVALUEPERIODIC_TriggerReasonEnumType   TriggerReasonEnumType = "MeterValuePeriodic"
	TIMELIMITREACHED_TriggerReasonEnumType     TriggerReasonEnumType = "TimeLimitReached"
	TRIGGER_TriggerReasonEnumType              TriggerReasonEnumType = "Trigger"
	UNLOCKCOMMAND_TriggerReasonEnumType        TriggerReasonEnumType = "UnlockCommand"
	STOPAUTHORIZED_TriggerReasonEnumType       TriggerReasonEnumType = "StopAuthorized"
	EVDEPARTED_TriggerReasonEnumType           TriggerReasonEnumType = "EVDeparted"
	EVDETECTED_TriggerReasonEnumType           TriggerReasonEnumType = "EVDetected"
	REMOTESTOP_TriggerReasonEnumType           TriggerReasonEnumType = "RemoteStop"
	REMOTESTART_TriggerReasonEnumType          TriggerReasonEnumType = "RemoteStart"
)

type ChargingStateEnumType string

var (
	CHARGING_ChargingStateEnumType      ChargingStateEnumType = "Charging"
	EVDETECTED_ChargingStateEnumType    ChargingStateEnumType = "EVDetected"
	SUSPENDEDEV_ChargingStateEnumType   ChargingStateEnumType = "SuspendedEV"
	SUSPENDEDEVSE_ChargingStateEnumType ChargingStateEnumType = "SuspendedEVSE"
)

type ReasonEnumType string

var (
	DEAUTHORIZED_ReasonEnumType       ReasonEnumType = "DeAuthorized"
	EMERGENCYSTOP_ReasonEnumType      ReasonEnumType = "EmergencyStop"
	ENERGYLIMITREACHED_ReasonEnumType ReasonEnumType = "EnergyLimitReached"
	EVDISCONNECTED_ReasonEnumType     ReasonEnumType = "EVDisconnected"
	GROUNDFAULT_ReasonEnumType        ReasonEnumType = "GroundFault"
	IMMEDIATERESET_ReasonEnumType     ReasonEnumType = "ImmediateReset"
	LOCAL_ReasonEnumType              ReasonEnumType = "Local"
	LOCALOUTOFCREDIT_ReasonEnumType   ReasonEnumType = "LocalOutOfCredit"
	MASTERPASS_ReasonEnumType         ReasonEnumType = "MasterPass"
	OTHER_ReasonEnumType              ReasonEnumType = "Other"
	OVERCURRENTFAULT_ReasonEnumType   ReasonEnumType = "OvercurrentFault"
	POWERLOSS_ReasonEnumType          ReasonEnumType = "PowerLoss"
	POWERQUALITY_ReasonEnumType       ReasonEnumType = "PowerQuality"
	REBOOT_ReasonEnumType             ReasonEnumType = "Reboot"
	REMOTE_ReasonEnumType             ReasonEnumType = "Remote"
	SOCLIMITREACHED_ReasonEnumType    ReasonEnumType = "SOCLimitReached"
)

type IdTokenEnumType string

var (
	Central_IdTokenEnumType         IdTokenEnumType = "Central"
	eMAID_IdTokenEnumType           IdTokenEnumType = "eMAID"
	ISO14443_IdTokenEnumType        IdTokenEnumType = "ISO14443"
	KeyCode_IdTokenEnumType         IdTokenEnumType = "KeyCode"
	Local_IdTokenEnumType           IdTokenEnumType = "Local"
	NoAuthorization_IdTokenEnumType IdTokenEnumType = "NoAuthorization"
	ISO15693_IdTokenEnumType        IdTokenEnumType = "ISO15693"
)
