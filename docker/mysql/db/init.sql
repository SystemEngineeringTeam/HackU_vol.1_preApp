SET CHARSET UTF8;
CREATE DATABASE IF NOT EXISTS pre_app_db DEFAULT CHARACTER SET utf8;

use pre_app_db;

create table users(
    id int auto_increment not null primary key,
    name varchar(128) not null
);

create table tasks(
    id int auto_increment not null primary key,
    title varchar(128) not null,
    deadline datetime
);

create table links_table(
    task_id int not null,
    user_id int not null,
    foreign key (task_id) references tasks (id) on delete cascade,
    foreign key (user_id) references users (id) on delete cascade
);



insert into 
    users(name)
values
    ("秦");

insert into 
    users(name)
values
    ("外山");
insert into 
    users(name)
values
    ("福田");
insert into 
    users(name)
values
    ("相畑");

insert into
    tasks(title,deadline)
values
    ("朝9時に起きる","2020-08-12 09:00:00");

insert into
    links_table(task_id,user_id)
values
    (1,3);
