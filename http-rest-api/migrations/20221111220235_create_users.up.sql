CREATE TABLE users(
    id generated always as identity primary key,
    email varchar not null unique,
    enrypted_password varchar not null

);