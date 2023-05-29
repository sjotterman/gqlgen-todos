CREATE TABLE restaurant (
	id serial4 NOT NULL,
	"name" varchar(140) NOT NULL,
	description varchar(1024) NOT NULL,
	phone_number varchar(40) NOT NULL,
	CONSTRAINT restaurant_pkey PRIMARY KEY (id)
);
