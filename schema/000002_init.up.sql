CREATE TABLE users (
  id serial NOT NULL unique,
  name varchar(255) NOT NULL,
  username varchar(255) NOT NULL unique,
  passwordHash varchar(255) NOT NULL
); 

CREATE TABLE list (
  id serial NOT NULL unique,
  title varchar(255) NOT NULL,
  description varchar(255) NULL
);

CREATE TABLE item (
  id serial NOT NULL unique,
  title varchar(255) NOT NULL,
  description varchar(255) NULL,
  done boolean NOT NULL DEFAULT false
);

CREATE TABLE user_list (
  id serial NOT NULL unique,
  id_user int references users(id) on delete cascade not null,
  id_list int references list(id) on delete cascade not null  
); 

CREATE TABLE list_item (
  id serial NOT NULL unique,
  id_list int references list(id) on delete cascade not null,
  id_item int references item(id) on delete cascade not null
); 