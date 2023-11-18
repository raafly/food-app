CREATE TABLE products (
    id serial primary key,
    name varchar(50) not null,
    description text,
    img varchar(50),
    price int not null,
    quantity int not null,
    category int not null,
    constraint products_categories Foreign Key (category) REFERENCES categories(id),
    created_at timestamp default current_timestamp
)