CREATE DATABASE IF NOT EXISTS products;

CREATE TABLE products (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255),
    qty int,
    PRIMARY KEY (id)
);