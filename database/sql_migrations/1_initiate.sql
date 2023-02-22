-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE customer (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(256) NOT NULL,
    birth_date TIMESTAMP NOT NULL,
    address VARCHAR(256) NOT NULL,
    email VARCHAR(256) NOT NULL,
    phone_number VARCHAR(256) NOT NULL,
    password VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    isAdmin bool NOT NULL
);

CREATE TABLE wallet (
    id SERIAL PRIMARY KEY,
    balance FLOAT NOT NULL,
    account_name VARCHAR(256) NOT NULL,
    account_number VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    customer_id INT,

    CONSTRAINT FK_wallet_customer FOREIGN KEY(customer_id)
        REFERENCES customer(id)
);

CREATE TABLE event (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    description VARCHAR(256) NOT NULL,
    start_date TIMESTAMP NOT NULL,
    end_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    category_id SERIAL NOT NULL,

    CONSTRAINT FK_event_category FOREIGN KEY(category_id)
        REFERENCES category(id)
);

CREATE TABLE ticket (
    id SERIAL PRIMARY KEY,
    name VARCHAR(256) NOT NULL,
    date TIMESTAMP NOT NULL,
    quota INT NOT NULL,
    price VARCHAR(256) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    event_id SERIAL NOT NULL,

    CONSTRAINT FK_ticket_event FOREIGN KEY(event_id)
        REFERENCES event(id)
);

CREATE TABLE transaction (
    id SERIAL PRIMARY KEY,
    date TIMESTAMP,
    qr_code VARCHAR(256),
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    ticket_id SERIAL NOT NULL,
    customer_id SERIAL NOT NULL,

    CONSTRAINT FK_transaction_ticket FOREIGN KEY(ticket_id)
        REFERENCES ticket(id),
    CONSTRAINT FK_transaction_customer FOREIGN KEY(customer_id)
        REFERENCES customer(id)
);

-- +migrate StatementEnd
