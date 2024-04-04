-- Create schema
CREATE SCHEMA vending_machine_service;

SET search_path TO vending_machine_service;

-- MachineType Table
CREATE TABLE MachineType (
                                             id SERIAL PRIMARY KEY,
                                             description TEXT NOT NULL
);

-- Product Table
CREATE TABLE Product (
                                         id SERIAL PRIMARY KEY,
                                         name VARCHAR(255) NOT NULL,
                                         price DECIMAL(10, 2) NOT NULL,
                                         details JSON NOT NULL
);

-- Machine Table
CREATE TABLE Machine (
                                         uuid UUID PRIMARY KEY,
                                         type_id INTEGER NOT NULL,
                                         location JSON NOT NULL,
                                         status VARCHAR(50) NOT NULL,
                                         storage_details JSON NOT NULL,
                                         FOREIGN KEY (type_id) REFERENCES MachineType(id)
);

-- User Table
CREATE TABLE USER_LOGIN (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(255) NOT NULL,
                                     password VARCHAR(255) NOT NULL
);