CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    email_verification_key VARCHAR(255),
    email_verification_time TIMESTAMP,
    password VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    cell_phone VARCHAR(20) NOT NULL,
    city VARCHAR(255),
    state VARCHAR(255),
    zip_code VARCHAR(20),
    district VARCHAR(255),
    street VARCHAR(255),
    street_number VARCHAR(20),
    complement VARCHAR(255),
    birthdate DATE,
    gender varchar,
    cpf VARCHAR(14),
    source SMALLINT,
    record_date TIMESTAMP ,
    level INTEGER
);

-- colmeia
CREATE TABLE Hive (
    IdExterno VARCHAR(255) NOT NULL,
    Name VARCHAR(255) NOT NULL,
    Description VARCHAR(255) NOT NULL,
    Slug VARCHAR(255) NOT NULL,
    PRIMARY KEY (IdExterno)
);

CREATE TABLE Moisture (
    IdExterno VARCHAR(255) NOT NULL,
    Temp VARCHAR(255) NOT NULL,
    Time TIMESTAMP NOT NULL,
    FOREIGN KEY (IdExterno) REFERENCES Hive(IdExterno)
);

CREATE TABLE OutsideTemperature (
    IdExterno VARCHAR(255) NOT NULL,
    Temp VARCHAR(255) NOT NULL,
    Time TIMESTAMP NOT NULL,
    FOREIGN KEY (IdExterno) REFERENCES Hive(IdExterno)
);

CREATE TABLE Temperature (
    IdExterno VARCHAR(255) NOT NULL,
    Temp VARCHAR(255) NOT NULL,
    Time TIMESTAMP NOT NULL,
    FOREIGN KEY (IdExterno) REFERENCES Hive(IdExterno)
);

CREATE TABLE WeightHive (
    IdExterno VARCHAR(255) NOT NULL,
    Value VARCHAR(255) NOT NULL,
    Time TIMESTAMP NOT NULL,
    FOREIGN KEY (IdExterno) REFERENCES Hive(IdExterno)
);


CREATE TABLE apiary(
    Name VARCHAR NOT NULL,
    Slug VARCHAR(255) NOT NULL,
    Local VARCHAR NOT NULL

);