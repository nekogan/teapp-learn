create table users (
	user_id int not null GENERATED ALWAYS AS identity,
	primary key(user_id),
	user_name varchar(30) not null,
	user_pass varchar(255) not null,
	user_avatar varchar,
	user_firstname varchar(100),
	user_secondname varchar(100)
);

create table post (
	post_id int not null GENERATED ALWAYS AS identity,
	user_id int,
	foreign key(user_id) REFERENCES users(user_id) on delete cascade on update cascade,
	post_title varchar(255),
	post_category varchar(60),
	post_tags varchar(255),
	post_text text
);