DROP TABLE IF EXISTS user;
create table user
(
  id int auto_increment primary key,
  username varchar(255) not null,
  password varchar(255) not null,
  email    varchar(255) not null,
  phone char(11) not null,
  Identity char(18),
  client_id char(32),
  client_port char(6),
  login_time datetime,
  logout_time datetime,
  is_logout TINYINT(1),
  HeartbeatTime datetime not null,
  device_info varchar(255) not null
)
