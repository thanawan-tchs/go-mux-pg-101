CREATE TABLE products (
	product_id   TEXT,
	name         TEXT,
	price        DECIMAL,
	amount       INT,
	date_created TIMESTAMP,
	date_updated TIMESTAMP,

	PRIMARY KEY (product_id)
);