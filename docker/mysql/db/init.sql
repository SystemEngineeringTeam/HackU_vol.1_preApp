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
