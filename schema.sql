create database lootbox;

create table users
(id serial,
username varchar(100)
);

create table press_history
(
    id serial,
    user_id bigint references users(id) ,
    unix_timestamp bigint

);

create table rewards
(
    id serial,
    user_id bigint references users(id),
    reward bigint

);
