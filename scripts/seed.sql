-- stolen from https://docs.docker.com/guides/pre-seeding/
CREATE DATABASE testdb;

-- connect to testdb
\c testdb

-- enable hashing for passwords
CREATE extension IF NOT EXISTS pgcrypto;

/*
 * -- Style --
 * Table name: PascalCase, singular form (except Users, and Cases not to clash
 * with reserved keywords)
 * Column name: snake_case
*/

/*
 * GENERATED ALWAYS AS IDENTITY - то же самое что SERIAL, но строже (нельзя
 * вручную вставить account_id, выдаст ошибку)
 * https://www.postgresql.org/docs/current/ddl-identity-columns.html

 * CONSTRAINT positive_balance - дает название ограничению (balance >= 0);
 * название ограничения будет видно в логах. Не гарантирует, что значение не NULL,
 * поэтому тоже проверяем.
 * https://www.postgresql.org/docs/current/ddl-constraints.html#DDL-CONSTRAINTS-CHECK-CONSTRAINTS

 * account_type IN ('regular', 'admin') - так можно проверить, чтобы
 * account_type был одним из разрешенных значений.

 * CREATE INDEX - способ ускорить поиск подходящих строчек (rows) по запросам (query), нам пока не нужно
 * https://www.postgresql.org/docs/current/indexes-intro.html
*/
CREATE TABLE Account (
    account_id   INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    account_type VARCHAR(10) NOT NULL CONSTRAINT valid_type CHECK (account_type IN ('regular', 'admin')),
    username     VARCHAR(50) UNIQUE NOT NULL,
    password     VARCHAR(100) NOT NULL,
    balance      FLOAT NOT NULL CONSTRAINT positive_balance CHECK (balance >= 0) DEFAULT 500.0
    --CREATE INDEX idx_username ON Users (username)
);

/*
 * Rarity. Используется для установки редкости предмета/кейса.
 * TODO: продумать политику ON DELETE (пока что ставит NULL).
*/
CREATE TABLE Rarity (
    rarity_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    rarity    VARCHAR(50) UNIQUE NOT NULL DEFAULT 'N/A'
);

/*
 * Item. Используется для хранения всех существующих предметов и их данных.
*/
CREATE TABLE Item (
    item_id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    item_name        VARCHAR(50) CONSTRAINT unique_item_name UNIQUE NOT NULL,
    item_rarity      VARCHAR(50) NOT NULL REFERENCES Rarity(rarity) ON DELETE SET NULL,
    item_description TEXT
    --CREATE INDEX idx_name ON Item (name)
);

/*
 * Cases. Используется для хранения данных о кейсах и их данных.
*/
CREATE TABLE Cases (
    case_id          INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    case_name        VARCHAR(50) CONSTRAINT unique_case_name UNIQUE NOT NULL,
    case_rarity      VARCHAR(50) NOT NULL REFERENCES Rarity(rarity) ON DELETE SET NULL,
    case_description TEXT
);

/*
 * CaseItem. Используется для хранения id предметов которые могут упасть с
 * кейса с конкретным названием. UNIQUE (case_name, item_id) для того, чтобы не
 * было дубликатов пар указанных колонн.
*/
CREATE TABLE CaseItem (
    item_case_id INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    case_name    VARCHAR(50) NOT NULL REFERENCES Cases(case_name) ON DELETE CASCADE,
    item_id      INT NOT NULL REFERENCES Item(item_id) ON DELETE CASCADE,
    CONSTRAINT unique_name_and_id UNIQUE (case_name, item_id)
);

/*
 * Inventory - таблица, где двойной PK: (id аккаунта, id предмета).
 * Используется для хранения количества конкретного предмета у конкретного
 * аккаунта.

 * REFERENCES Account(account_id) - account_id в таблице Inventory должен так же
 * существовать в таблице Account. Например: если в Inventory есть account_id=528, а в
 * Account нет, то выдаст ошибку, в то же время NULL принимается и не считается
 * ошибкой. https://www.postgresql.org/docs/current/tutorial-fk.html

 * ON DELETE CASCADE - если в ориг таблице удалим ряд с account_id/item_id, то
 * в этой таблице тоже (ибо нет смысла хранить инвентарь несуществующего
 * аккаунта или количество несуществующего предмета).
*/
CREATE TABLE Inventory (
    account_id    INT REFERENCES Account(account_id) ON DELETE CASCADE,
    item_id       INT REFERENCES Item(item_id) ON DELETE CASCADE,
    item_quantity INT NOT NULL CONSTRAINT positive_quantity CHECK (item_quantity >= 0),
    PRIMARY KEY (account_id, item_id)
);

CREATE TABLE Transaction (
    transaction_id     INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    transaction_type   VARCHAR(10) NOT NULL CHECK (transaction_type IN ('purchase', 'sell', 'transfer')),
    transaction_date   TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    transaction_amount FLOAT NOT NULL CONSTRAINT positive_transaction CHECK (transaction_amount >= 0),
    item_id            INT NOT NULL REFERENCES Item(item_id),
    item_quantity      INT NOT NULL CONSTRAINT positive_quantity CHECK (item_quantity >= 0),
    seller_id          INT NOT NULL REFERENCES Account(account_id),
    buyer_id           INT NOT NULL REFERENCES Account(account_id)
    --CREATE INDEX idx_transaction_type ON Transaction (transaction_type)
);


CREATE TABLE Listing (
    listing_id    INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
    seller_id     INT NOT NULL REFERENCES Account(account_id),
    item_id       INT NOT NULL REFERENCES Item(item_id),
    item_quantity INT NOT NULL CONSTRAINT positive_quantity CHECK (item_quantity > 0),
    lot_price     FLOAT NOT NULL CONSTRAINT positive_price CHECK (lot_price > 0)
);

INSERT INTO Account (account_type, username, password, balance) VALUES
    ('regular', 'Chika', 'secret', 0),
    ('admin', 'Patrick', 'super_secret', 100),
    ('admin', 'Spongebob', 'ultra_secret', 6969.69)
ON CONFLICT (username) DO NOTHING;

INSERT INTO Rarity (rarity) VALUES
    ('Common'),
    ('Rare'),
    ('Epic'),
    ('Mythic'),
    ('Legendary');

INSERT INTO Item (item_name, item_rarity, item_description) VALUES
    ('Rock', 'Common', 'Ooga booga'),
    ('Stick', 'Common', 'bad booga'),
    ('Sharp rock', 'Rare', 'Ow'),
    ('Long stick', 'Rare', 'OOO BOOGA BOOGA');

INSERT INTO Cases (case_name, case_rarity) VALUES
    ('Common Case', 'Common'),
    ('Rare Case', 'Rare');

INSERT INTO CaseItem (case_name, item_id) VALUES
    ('Common Case', 1),
    ('Common Case', 2),
    ('Rare Case', 3),
    ('Rare Case', 4);

INSERT INTO Inventory (account_id, item_id, item_quantity) VALUES
    (2, 1, 10),
    (3, 2, 5);
