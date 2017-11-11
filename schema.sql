USE HEEV;

DROP TABLE SensorData;
DROP TABLE CarTablet;
DROP TABLE Cars;

CREATE TABLE Cars (
    Id INT AUTO_INCREMENT,
    Name VARCHAR(256),
    WheelRadius DOUBLE,
    PRIMARY KEY(Id)
);

CREATE TABLE CarTablet (
    AndroidId INT,
    CarId INT,
    FOREIGN KEY(CarId) REFERENCES Cars(Id),
    PRIMARY KEY(AndroidId)
);

CREATE TABLE SensorData (
    Id INT AUTO_INCREMENT,
    CarId INT,
    LogTime DATETIME(6),
    WheelRpm DOUBLE,
    GroundSpeed DOUBLE,
    WindSpeed DOUBLE,
    BatteryVoltage DOUBLE,
    LKillSwitch BOOLEAN,
    MKillSwitch BOOLEAN,
    RKillSwitch BOOLEAN,
    SecondaryBatteryVoltage DOUBLE,
    CoolantTemperature DOUBLE,
    IntakeTemperature DOUBLE,
    SystemCurrent DOUBLE,
    FOREIGN KEY(CarId) REFERENCES Cars(Id),
    PRIMARY KEY(Id)
);