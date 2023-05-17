CREATE DATABASE db_account_service;
USE db_account_service;

CREATE TABLE users(
	user_id varchar(50) primary key,
    name varchar(100),
	phone varchar(50) not null unique,
    password varchar(50) not null,
    balance int,
    last_login datetime,
    created_at timestamp default current_timestamp
);

CREATE TABLE balance(
	user_id varchar(50) primary key,
    total int,
    CONSTRAINT fk_balance_users FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE top_up(
	top_up_id varchar(50) primary key,
    total int,
    payment_method varchar(50),
    user_id varchar(50),
    created_at timestamp default current_timestamp,
    CONSTRAINT fk_topup_users FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE transfer(
	transfer_id varchar(50) primary key,
    receiver_id varchar(50),
    user_id varchar(50),
    total int, 
    method_transfer varchar(50),
    created_at timestamp default current_timestamp
);

CREATE TABLE login_activity(
	user_id varchar(50) primary key,
    login_at datetime default current_timestamp,
    constraint fk_login_activity_user foreign key (user_id) REFERENCES users(user_id) on delete cascade
);