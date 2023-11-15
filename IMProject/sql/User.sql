DROP TABLE IF EXISTS user;
create table user
(
  id       int auto_increment
    primary key,
  username varchar(255) not null,
  password varchar(255) not null,
  email    varchar(255) not null,
  phone char(11) not null,
  Identity char(18) not null,
  client_id char(32) not null

)
