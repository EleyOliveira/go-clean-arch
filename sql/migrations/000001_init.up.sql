CREATE TABLE IF NOT EXISTS orders (
    id varchar(36) NOT NULL PRIMARY KEY,
    price decimal(10,2) NOT NULL,
    tax decimal(10,2) NOT NULL,
    finalprice decimal(10,2) NOT NULL  
);