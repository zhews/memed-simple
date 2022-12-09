CREATE DATABASE memed;
\c memed;
CREATE TABLE memed_user
(
    id            UUID PRIMARY KEY,
    username      TEXT UNIQUE,
    name          TEXT,
    admin         BOOLEAN,
    password_hash BYTEA,
    created_at    INT,
    updated_at    INT
);