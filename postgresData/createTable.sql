CREATE TABLE IF NOT EXISTS client(
    clientID int, password varchar(255), login varchar(255), email varchar(255)
);

CREATE TABLE IF NOT EXISTS access (
    accessID int, accessStatus varchar(15), accessLevel int
);

CREATE TABLE IF NOT EXISTS accessClient (clientID int, accessID int);

CREATE TABLE IF NOT EXISTS accessHome (accessID int, homeID int);

CREATE TABLE IF NOT EXISTS home (homeID int, name varchar(20));

CREATE TABLE IF NOT EXISTS device (
    deviceID int, name varchar(20), typeDevice varchar(20), status varchar(10), brand varchar(15), maxParametr int, minParametr int, powerConsumption int
);

CREATE TABLE IF NOT EXISTS deviceHome (homeID int, deviceID int);

CREATE TABLE IF NOT EXISTS historyDev (
    historyDevID int, timeWork time, AverageIndicator decimal, EnergyConsumed int
);

CREATE TABLE IF NOT EXISTS historyDevice (
    historyDevID int, deviceID int
);