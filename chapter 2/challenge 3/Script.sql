CREATE DATABASE chap2-challenge3;

create table book(
	id serial primary key,
	title varchar(255) not null,
	author varchar(255) not null,
	description varchar(255) not null
);