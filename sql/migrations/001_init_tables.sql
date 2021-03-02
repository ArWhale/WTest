-- +goose Up
-- +goose NO TRANSACTION
-- DROP SCHEMA IF EXISTS public CASCADE;
-- CREATE SCHEMA public;
--
-- GRANT ALL ON SCHEMA public TO postgres;
-- GRANT ALL ON SCHEMA public TO public;

CREATE TABLE IF NOT EXISTS customers(
    id SERIAL PRIMARY KEY,
    first_name text NOT NULL,
    last_name text NOT NULL,
    birthdate date NOT NULL,
    gender text NOT NULL,
    e_mail text NOT NULL,
    address text NOT NULL,
    UNIQUE(e_mail)
);

-- +goose Down
