-- stolen from https://docs.docker.com/guides/pre-seeding/
CREATE DATABASE testdb;

-- connect to testdb
\c testdb

-- enable hashing for passwords
CREATE extension IF NOT EXISTS pgcrypto;

/*
* -- Style --
* Table name: PascalCase, singular form (except Users, not to clash with user keyword)
* Column name: snake_case
*/

/*
* GENERATED ALWAYS AS IDENTITY == то же самое что SERIAL, но строже (нельзя вставить account_id, будет ошибка)
* https://www.postgresql.org/docs/current/ddl-identity-columns.html

* CONSTRAINT positive_balance - дает название ограничению (balance >= 0), название будет видно название в логах
* https://www.postgresql.org/docs/current/ddl-constraints.html#DDL-CONSTRAINTS-CHECK-CONSTRAINTS

* CREATE INDEX - способ ускорить поиск подходящих строчек (rows) по запросам (query), нам пока не нужно
* https://www.postgresql.org/docs/current/indexes-intro.html
*/
CREATE TABLE IF NOT EXISTS Account (
    account_id   INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    account_type VARCHAR(10) NOT NULL CHECK (account_type IN ('regular', 'admin')),
    username     VARCHAR(50) UNIQUE NOT NULL,
    password     VARCHAR(100) NOT NULL,
    balance      FLOAT CONSTRAINT positive_balance CHECK (balance >= 0) DEFAULT 500.0
    --CREATE INDEX idx_username ON Users (username)
);

CREATE TABLE IF NOT EXISTS Item (
    item_id     INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    name        VARCHAR(50) NOT NULL,
    description TEXT
    --CREATE INDEX idx_name ON Item (name)
);

/*
* REFERENCES Account(account_id) == account_id в Inventory должен так же существовать в Account
* Например, если в Inventory есть account_id=528, а в Account нет, то будет ошибка
* https://www.postgresql.org/docs/current/tutorial-fk.html
*/
CREATE TABLE IF NOT EXISTS Inventory (
    account_id    INT REFERENCES Account(account_id),
    item_id       INT REFERENCES Item(item_id),
    item_quantity INT CONSTRAINT positive_quantity CHECK (item_quantity >= 0)
);

CREATE TABLE IF NOT EXISTS Transaction (
    transaction_id     INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    transaction_type   VARCHAR(10) NOT NULL CHECK (transaction_type IN ('purchase', 'sell', 'transfer')),
    transaction_date   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    transaction_amount FLOAT CONSTRAINT positive_transaction CHECK (transaction_amount >= 0),
    item_id            INT REFERENCES Item(item_id),
    item_quantity      INT CONSTRAINT positive_quantity CHECK (item_quantity >= 0),
    seller_id          INT REFERENCES Account(account_id),
    buyer_id           INT REFERENCES Account(account_id)
    --CREATE INDEX idx_transaction_type ON Transaction (transaction_type)
);


INSERT INTO Account (account_type, username, password, balance) VALUES
    ('regular', 'Chika', 'secret', 0),
    ('admin', 'Patrick', 'super_secret', 100),
    ('admin', 'Spongebob', 'ultra_secret', 6969.69)
ON CONFLICT (username) DO NOTHING;

INSERT INTO Item (name, description) VALUES
    ('Rock', 'Ooga booga'),
    ('Stick', 'OOGA BOOGA');

INSERT INTO Inventory (account_id, item_id, item_quantity) VALUES
    (2, 1, 10),
    (3, 2, 5);
