CREATE TABLE customers(
    id varchar(100) not null primary key,
    username varchar(50) not null,
    email varchar(50) not null,
    password varchar(50) not null,
    phone varchar(13),
    address text,
    created_at timestamp default current_timestamp
)
