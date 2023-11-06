CREATE TABLE customers(
    id varchar(50) not null primary key,
    username varchar(50) not null,
    email varchar(50) not null,
    password varchar(50) not null,
    phone varchar(20),
    address varchar(100) not null,
    constraint username_unix UNIQUE(username),
    update_at timestamp, 
    created_at timestamp default current_timestamp,
)