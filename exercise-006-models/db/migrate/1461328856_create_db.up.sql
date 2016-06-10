CREATE TABLE customers(
customer_id SERIAL PRIMARY KEY
, email text UNIQUE
, first_name varchar(70)
, last_name varchar(70)
, birth_date date
, created_at TIMESTAMPTZ default now()
, updated_at TIMESTAMPTZ default now()
);

CREATE INDEX customers__created_at ON customers(created_at);
CREATE INDEX customers__updated_at ON customers(updated_at);

CREATE TABLE products(
product_id SMALLINT PRIMARY KEY
, product_name varchar(70)
);

INSERT INTO products(product_id, product_name)
VALUES (1, 'kayak'), (2, 'canoe'), (3, 'paddle'), (4, 'vest');

CREATE TABLE orders(
order_id SERIAL PRIMARY KEY
, product_id INTEGER NOT NULL REFERENCES products
, quantity INTEGER
, customer_id INTEGER NOT NULL REFERENCES customers
, created_at TIMESTAMPTZ default now()
, updated_at TIMESTAMPTZ default now()
);

CREATE INDEX orders__product_id ON orders(product_id);
CREATE INDEX orders__customer_id ON orders(customer_id);
CREATE INDEX orders__created_at ON orders(created_at);
CREATE INDEX orders__updated_at ON orders(updated_at);
