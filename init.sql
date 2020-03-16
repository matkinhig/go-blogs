-- create database or table if not exist

create table if not exist categories(
    id bigint primary key auto_increment,
    description varchar(500) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)

engine = InnoDB
default charset = utf8;

-- create another table if not using auto migrate
