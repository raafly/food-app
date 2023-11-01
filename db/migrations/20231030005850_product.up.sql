CREATE TABLE products (
    id varchar(50) not null primary key,
    name varchar(50) not null,
    description text,
    price int not null,
    quantity int not null,
    category int not null,
    constraint products_categories Foreign Key (category) REFERENCES categories(id)
)