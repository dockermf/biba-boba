-- stolen from https://docs.docker.com/guides/pre-seeding/
CREATE DATABASE testdb;

-- connect to testdb
\c testdb

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    login VARCHAR(50) UNIQUE,
    password VARCHAR(100),
    balance FLOAT
);

INSERT INTO users (login, password, balance) VALUES
  ('Chika', 'secret', 0),
  ('Patrick', 'super_secret', 100),
  ('Spongebob', 'ultra_secret', 6969.69)
ON CONFLICT (login) DO NOTHING;
