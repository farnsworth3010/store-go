CREATE TABLE user
(
    id            serial primary key,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE admin
(
    user_id int references user(id) not null
);

CREATE TABLE blog
(
    id            serial primary key,
    title         varchar(255),
    text          text,
    creation_date timestamp not null default current_timestamp
);