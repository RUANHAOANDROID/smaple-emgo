create table if not exists demo (
    id       BIGINT       PRIMARY KEY NOT NULL,
    name     VARCHAR (20),
    crt_date VARCHAR (20)
);
create table if not exists device_config (
id       VARCHAR (50)       PRIMARY KEY NOT NULL,
manufacturer_id1 VARCHAR (50) ,
manufacturer_id2 VARCHAR (50) ,
buffer VARCHAR (10) ,
delayed_time VARCHAR (10) ,
invalid_time VARCHAR (10) ,
check_url VARCHAR (200) ,
write_off_url VARCHAR (200) ,
num_up_url VARCHAR (200) ,
num_up_time VARCHAR (10) ,
heartbeat_url VARCHAR (200) ,
heartbeat_time VARCHAR (10) ,
true_voice1 VARCHAR (200) ,
false_voice1 VARCHAR (200) ,
true_voice2 VARCHAR (200) ,
false_voice2 VARCHAR (200) ,
true_voice3 VARCHAR (200) ,
false_voice3 VARCHAR (200) ,
true_voice4 VARCHAR (200) ,
false_voice4 VARCHAR (200) ,
config_url VARCHAR (200)
);
create table if not exists device (
id       VARCHAR (50)       PRIMARY KEY NOT NULL,
device_no VARCHAR (50) ,
device_ip VARCHAR (50) ,
serial_number VARCHAR (50) ,
device_version VARCHAR (10) ,
device_status VARCHAR (10) ,
status VARCHAR (10) ,
add_time VARCHAR (200) ,
update_time VARCHAR (200)
);
create table if not exists device_num_upload (
id       Integer      PRIMARY KEY AUTOINCREMENT,
device_no VARCHAR (50) ,
biz_date VARCHAR (50) ,
people_num VARCHAR (10) ,
is_upload VARCHAR (10)
);
create table if not exists sys_log (
id       Integer      PRIMARY KEY AUTOINCREMENT,
content TEXT ,
add_time TIMESTAMP default (datetime('now', 'localtime'))
);
create table if not exists sys_log_time (
id       Integer      PRIMARY KEY AUTOINCREMENT,
last_time TIMESTAMP default (datetime('now', 'localtime'))
);
create table if not exists sys_user (
id       Integer      PRIMARY KEY AUTOINCREMENT,
user_name VARCHAR (50) ,
user_pwd VARCHAR (50)
);