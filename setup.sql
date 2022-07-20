create table classification (
	classification_id int not null,
	primary key(id),
	title varchar(15) not null
)

insert into classification (id, title) 
values 
	(0, 'Красный'),
	(1, 'Белый'), 
	(2, 'Зеленый'),
	(3, 'Черный'),
	(4, 'Пуэр')

create table tea (
	tea_id int GENERATED ALWAYS AS identity,
	primary key(tea_id),
	title varchar(255) not null,
	teaTypeID int not null,
	constraint fk_teaType
		foreign key(teaTypeID)
			references classification(classification_id)
)

create table post (
	post_id int GENERATED ALWAYS AS identity,
	primary key(post_id),
	tea_id int not null,
	post_text text,
	rating int, 
	constraint fk_tea
		foreign key(tea_id)
			references tea(tea_id)
)