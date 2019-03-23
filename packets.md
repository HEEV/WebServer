# Packets

All communication will be done via JSON, with the packet structure as described
below.

## Tablet's initial communication with server

```json
{
    "MessageType": "GetNextRunNumber",
    "AndroidId": network->macAddress
}
```

## Disconnect from server

Simple `quit` string message, non-JSON, over websocket

## Update from tablet

```json
{
    "AndroidId": network->macAddress,
    "MessageType": "Log",
    "RunNumber": runNum,
    "BatteryVoltage": uiInterface->getBatteryVoltage,
    "GroundSpeed": uiInterface->getGroundSpeed,
    "IntakeTemperature": uiInterface->getManifoldAirTemp,
    "LKillSwitch": 0,
    "Latitude": QJsonValue(currentCoordinate.latitude,
    "LogTime": dateStr,
    "Longitude": QJsonValue(currentCoordinate.longitude,
    "LapNumber": uiInterface->getCurrentLapNumber,
    "MKillSwitch": 0,
    "RKillSwitch": 0,
    "SecondaryBatteryVoltage": 0.0,
    "WheelRpm": uiInterface->getEngineRPM,
    "WindSpeed": uiInterface->getWindSpeed,
    "SystemCurrent": 1.02f,
    "CoolantTemperature": 42.42
}
```

## Update run number on tablet

```json
```