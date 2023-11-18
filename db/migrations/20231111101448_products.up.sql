CREATE TABLE products(
    id serial primary key,
    name varchar(50) not null unique,
    description text default 'not set',
    price decimal(10, 3) not null,
    quantity int not null,
    category int not null,
    created_at timestamp default current_timestmap
)