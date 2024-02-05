CREATE TABLE users
(
    id            serial primary key,
    name          varchar(255) not null,
    username      varchar(255) not null,
    password_hash varchar(255) not null
);