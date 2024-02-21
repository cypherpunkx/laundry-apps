CREATE DATABASE enigma_laundry

CREATE TABLE employee (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100),
    phone_number VARCHAR(20) UNIQUE,
    address TEXT
)

CREATE TABLE product (
    id VARCHAR(100) PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price BIGINT,
    uom VARCHAR(50)
)

CREATE TABLE customer (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(20) UNIQUE,
    address TEXT
)

CREATE TABLE bill (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    bill_date date, 
    entry_date date,   
    employee_id VARCHAR(100),
    customer_id VARCHAR(100),
    FOREIGN KEY(employee_id) REFERENCES employee(id),
    FOREIGN KEY(customer_id) REFERENCES customer(id)
)

CREATE TABLE bill_detail (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    bill_id VARCHAR(100),
    product_id VARCHAR(100),
    product_price BIGINT,
    qty int,
    finish_date date,
    FOREIGN KEY(bill_id) REFERENCES bill(id),
    FOREIGN KEY(product_id) REFERENCES product(id)
)

CREATE TABLE user_credential (
    id VARCHAR(100) PRIMARY KEY NOT NULL,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(200) NOT NULL,
    isActive BOOLEAN DEFAULT TRUE
)

CREATE TABLE user_picture (
    id VARCHAR(100) PRIMARY KEY,
    user_id VARCHAR(100) NOT NULL UNIQUE,
    file_location VARCHAR(250) NOT NULL,
    FOREIGN KEY(user_id) REFERENCES user_credential(id)
);