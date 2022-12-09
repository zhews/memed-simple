CREATE DATABASE memed;
\c memed;
CREATE TABLE memed_user
(
    id            UUID PRIMARY KEY,
    username      TEXT UNIQUE,
    name          TEXT,
    password_hash BYTEA,
    created_at    INT,
    updated_at    INT
);