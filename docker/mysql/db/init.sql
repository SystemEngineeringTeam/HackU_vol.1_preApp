use pre_app_db;

create table users(
    user_id int auto_increment not null primary key,
    user_name varchar(128) not null
);

create table tasks(
    task_id int auto_increment not null primary key,
    title varchar(128) not null,
    deadline datetime not null
);

create table links_table(
    task_id int not null,
    user_id int not null,
    foreign key (task_id) references tasks (task_id),
    foreign key (user_id) references users (user_id)
);

insert into 
    users(user_name)
values
    ("秦");

insert into 
    users(user_name)
values
    ("外山");
insert into 
    users(user_name)
values
    ("福田");
insert into 
    users(user_name)
values
    ("相畑");
