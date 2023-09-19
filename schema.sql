create database lootbox;

create table users (
    id serial primary key not null,
    username varchar(100)
);

create table press_history (
    id serial primary key,
    user_id bigint references users(id),
    unix_timestamp bigint
);

create table rewards (
    id serial primary key,
    user_id bigint references users(id),
    reward bigint
);