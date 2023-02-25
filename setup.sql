create table users (
	user_id int not null GENERATED ALWAYS AS identity,
	primary key(user_id),
	user_name varchar(255) not null,
	user_pass varchar(255) not null,
	user_avatar varchar,
	user_firstname varchar(100),
	user_secondname varchar(100)
);

create table post (
	post_id int not null GENERATED ALWAYS AS identity,
	user_id int,
	foreign key(user_id) REFERENCES users(user_id) on delete cascade on update cascade,
	post_title varchar(40),
	post_category varchar(1000),
	post_text text,
	post_rating int
);