# OCPP Implementation

Charging management station ("back office") communicates with a Charging Station.
Charging Station has one or more EVSEs, which each have on or more connectors.

In our simple case, we have 1 EVSE with 1 connector.

## WebSockets

We are implementing "OCPP-J", which is the JSON-based section. This is the only encoding supported by OCPP 2.0.

Connection URL: `ws://<charging management station>/ocpp/CS001` for charging station 1.
The last component of the URL can be arbitrary string that matches the "identifierString" schema (48 chars max).

In `Sec-Websocket-Protocol` field, make sure we have the protocol name `ocpp2.0` there.

### RPC Framework

Essentially implementing an RPC framework here on top of WebSockets.

Messages are either a Request (CALL) or a Reply (CALLRESULT) which can also be a CALLERROR. Anyone interacting with the Charging Station should have at most one outstanding call to the charging station. Timeouts are encouraged.

Messages have a message type defined by a number:

- `CALL`: 2
- `CALLRESULT`: 3
- `CALLERROR`: 4


Messages also have mesage IDs which are unique for the same sender. The corresponding CALLRESULT or CALLERROR for a CALL must have the same messageid. This is max length of 36 characters.

Messages also have a payload field, which has a JSON value. Can be empty.

Messages also have an "Action" field, which is the OCPP message name without the "Request" suffix.
CALLRESULT does not contain the action field.o

---

Syntax:


**`CALL`**: 4 elements: `[<MessageTypeId>, "<MessageId>", "<Action>", {<Payload>}]`

example:

```json
[2,
    "19223201",
    "BootNotification",
    {
        "reason": "PowerUp",
        "chargingStation": {
            "model": "SingleSocketCharger",
            "vendorName": "VendorX"
        }
    }
]
```

---

**`CALLRESULT`**: 4 elements: `[<MessageTypeId>, "<MessageId>", {<Payload>}]`

example:

```json
[3,
 "19223201",
 {
    "currentTime": "2013-02-01T20:...",
    "interval": 300,
    "status": "Accepted"
 }
```

---

**`CALLERROR`**: 4 elements: `[<MessageTypeId>, "<MessageId>", "<errorCode>", "<errorDescription",{<errorDetails}]`

- errorCode should contain a string from the RPC Framework Error codes
- errorDetails is user-defined


## Minimum Device Model

Applies to any model of charging station.

Need to implement:

- B05 Set Variables
- B06 Get Variables
- B07 Get Base Report (all 3 report bases)

## JSON Schema

- `OCPP-2.0_part3_schemas` can validate with https://github.com/xeipuuv/gojsonschema
- validate the messages this way

## Transactions


### Charging

Probably needs to use "Remote Start Transaction"

What are the scenarios?

- charge induced by car when user plugs in cable
    1. plugging i ncable triggers this
    2. charging stations sends `StatusNotificationRequest` to CSMS to notify of "Occupied" connector
    3. CSMS responds with `StatusNotificationResponse`, confirming the request
    4. charging station sends `TransactionEventRequest` (eventType = Started) notifying the CSMS about started transaction
    5. CSMS responds with `TransactionEventResponse`
    6. Remote Start triggered (externally? via EV driver? app? user notification? charging station operator)
    7. CSMS sends `RequestStartTransactionRequest` to Charging Station
    8. Charging Station responds with `RequestStartTransactionResponse` with the transactionId of the already started transaction
        - contains status if charging station has accepted the request and will start a transaction
    9. Charging Station sends `TransactionEventRequest` (eventType = Updated, chargingState = Charging) to CSMS
    10. CSMS sends `TransactionEventResponse`
-  need to check `AuthorizeRemoteStart`

- charge induced by remote start
    1. CSMS sends `RequestStartTransactionRequest` to Charging Station
    2. Charging Station responds with `RequestStartTransactionResponse`
    3. EV driver authorized by CSMS dependent on config vars
    4. Charging Station sends `StatusNotificationRequest` to CSMS for occupied connector
    5. CSMS sends `StatusNotificationResponse`
    6. Charging Station sends `TransactionEventRequest` (eventType = Started) to notify CSMS
    7. Plug in Cable
    8. Energy offer started
    9. Charging Station sends `TransactionEventRequest` (eventType = Updated, chargingState = Charging) to inform CSMS
    10. CSMS sends `TransactionEventResponse`

For now, we are going to implement notification when it starts charging. This requires the following messages:

- `StatusNotificationRequest`
    - `timestamp (dateTime) x1`
    - `connectorStatus (ConnectorStatusEnumType) x1`
        - `Available`
        - `Occupied`
        - `Reserved`
        - `Unavailable`
        - `Faulted`
    - `evseId (integer) x1`
    - `connectorId (integer) x1`
- `StatusNotificationResponse` 
    - no fields; deprecated?
- `TransactionEventRequest`(page 334 of spec)
    - `eventType (TransactionEventEnumType x1)`
        - `Ended`
        - `Started`
        - `Updated`
    - `timestamp (dateTime) x1`
    - `triggerReason (TriggerReasonEnumType)`
        - Authorized
        - CablePluggedIn
        - ChargingRateChanged
        - ChargingStateChanged
        - Deauthorized
        - EnergyLimitReached
        - EVCommunicationLost
        - EVConnectTimeout
        - MeterValueClock
        - MeterValuePeriodic
        - TimeLimitReached
        - Trigger
        - UnlockCommand
        - StopAuthorized
        - EVDeparted
        - EVDetected
        - RemoteStop
        - RemoteStart
    - `seqNo (integer)` 
    - `offline (boolean)` (opt)
    - `numberOfPhasesUsed (integer)` (opt)
    - `cableMaxCurrent (decimal)` (opt)
    - `reservationId (integer)` (opt)
    - `transactionData (TransactionType)` (required)
    - `idToken (IdTokenType)` (opt)
    - `evse (EVSEType)` (opt)
    - `meterValue (MeterValueType)` (opt)
- `TransactionEventResponse`
- `RequestStartTransactionRequest`
- `RequestStartTransactionResponse`
- `TransactionType`:
    - `id (identifier string)` (required)
    - `chargingState (ChargingStateEnumType)` (opt)
        - `Charging`
        - `EVDetected`
        - `SuspendedEV`
        - `SuspendedEVSE`
    - `timeSpentCharging (int)` (opt)
    - `stoppedReason (ReasonEnumType)` (opt)
    - `remoteStartId (int)` (opt)
