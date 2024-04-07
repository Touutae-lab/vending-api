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
                                         img_url TEXT NOT NULL,
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

CREATE TABLE "order" (
                                     id SERIAL PRIMARY KEY,
                                     machine_id UUID NOT NULL,
                                     product_id INT NOT NULL,
                                     quantity INT NOT NULL,
                                     total_price DECIMAL(10, 2) NOT NULL,
                                     order_date TIMESTAMP NOT NULL
);

CREATE TABLE "order_status" (
                                     id SERIAL PRIMARY KEY,
                                     order_id INT NOT NULL,
                                     status VARCHAR(50) NOT NULL
);


ALTER TABLE "order" ADD FOREIGN KEY (machine_id) REFERENCES Machine(uuid);
ALTER TABLE "order" ADD FOREIGN KEY (product_id) REFERENCES Product(id);
ALTER TABLE "order_status" ADD FOREIGN KEY (order_id) REFERENCES "order"(id);

-- Insert data
WITH inserted_products AS (
    -- Insert products and return their IDs
    INSERT INTO product (name, price, img_url, details)
        VALUES
            ('ITEM 1', 10.00,  '/cass-fresh-lager-cans-korea-355.jpg','Details of ITEM 1'),
            ('ITEM 2', 15.00, '/Coca-Cola-Original-Soft-Drink.webp','Details of ITEM 2'),
            ('ITEM 3', 15.00, '/DRWK_75002459-1__01061.jpg','Details of ITEM 3'),
            ('ITEM 4', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4'),
            ('ITEM 6', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4'),
            ('ITEM 7', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4'),
            ('ITEM 8', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4'),
            ('ITEM 9', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4'),
            ('ITEM 10', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4'),
            ('ITEM 11', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4'),
            ('ITEM 12', 15.00, '/luxus-belgian-style-lager-cans-355ml__72410.jpg','Details of ITEM 4')

        RETURNING id
),
     products_array AS (
         -- Build a JSON array of product IDs
         SELECT json_agg(json_build_object('product_id', id, 'quantity', 10)) AS product_details
         FROM inserted_products
     )

-- Insert into Machine with the JSON object
INSERT INTO machine (uuid, name, location, status, storage_details)
SELECT
    'f47ac10b-58cc-4372-a567-0e02b2c3d479',
    'Machine A',
    'Location A',
    'ACTIVE',
    product_details
FROM products_array;
-- End of script