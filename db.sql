CREATE TABLE IF NOT EXISTS users(
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	email VARCHAR(255) NOT NULL UNIQUE,
	password VARCHAR(100) NOT NULL,
	token VARCHAR(255) NULL UNIQUE,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL
);

CREATE TABLE IF NOT EXISTS categories(
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	title VARCHAR(255) NOT NULL,
	icon VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT fk_user_category FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS companies(
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL,
	cnpj VARCHAR(255) NOT NULL,
	name VARCHAR(255) NOT NULL,
	title VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT fk_user_company FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS receipts(
	id SERIAL PRIMARY KEY,
	category_id INTEGER NOT NULL, 
	company_id INTEGER NOT NULL, 
	user_id INTEGER NOT NULL, 
	title VARCHAR(255) NOT NULL,
	tax REAL NULL,
	extra REAL NOT NULL,
	discount REAL NOT NULL,
	total REAL NOT NULL,
	url VARCHAR(255) NULL,
	access_key VARCHAR(255) NULL,
	issued_at TIMESTAMP NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT fk_category_receipt FOREIGN KEY (category_id) REFERENCES categories(id),
	CONSTRAINT fk_company_receipt FOREIGN KEY (company_id) REFERENCES companies(id),
	CONSTRAINT fk_user_receipt FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS items(
	id SERIAL PRIMARY KEY,
	receipt_id INTEGER NOT NULL, 
	title VARCHAR(255) NOT NULL,
	price REAL NULL,
	quantity REAL NULL,
	total REAL NULL,
	tax REAL NULL,
	measure VARCHAR(255) NOT NULL,
	created_at TIMESTAMP NULL,
	updated_at TIMESTAMP NULL,
	CONSTRAINT fk_receipt_item FOREIGN KEY (receipt_id) REFERENCES receipts(id)
);