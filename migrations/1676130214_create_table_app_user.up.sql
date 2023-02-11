CREATE TABLE app_user (
    username varchar NOT NULL,
    email    varchar NOT NULL,
    salt     varchar NOT NULL,
    password varchar NOT NULL,
    phone    varchar
);