-- stolen from https://docs.docker.com/guides/pre-seeding/
CREATE DATABASE testdb;

-- connect to testdb
\c testdb

CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  name VARCHAR(50),
  email VARCHAR(100) UNIQUE
);

INSERT INTO users (name, email) VALUES
  ('Chika', 'chika@example.com'),
  ('Patrick', 'patrick@example.com'),
  ('Spongebob', 'spongebob@example.com')
ON CONFLICT (email) DO NOTHING;
