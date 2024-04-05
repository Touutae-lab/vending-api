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
                                         details TEXT NOT NULL
);

-- Machine Table
CREATE TABLE Machine (
                                         uuid UUID PRIMARY KEY,
                                         name VARCHAR(255) NOT NULL,
                                         location TEXT NOT NULL,
                                         status VARCHAR(50) NOT NULL,
                                         storage_details JSON NOT NULL
);

-- User Table
CREATE TABLE USER_LOGIN (
                                     id SERIAL PRIMARY KEY,
                                     username VARCHAR(255) NOT NULL,
                                     password VARCHAR(255) NOT NULL
);

CREATE TABLE Transaction (
                                     id SERIAL PRIMARY KEY,
                                     machine_id UUID NOT NULL,
                                     product_id INT NOT NULL,
                                     quantity INT NOT NULL,
                                     total_price DECIMAL(10, 2) NOT NULL,
                                     order_date TIMESTAMP NOT NULL
);