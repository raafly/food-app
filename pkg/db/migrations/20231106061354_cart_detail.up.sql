CREATE TABLE carts_detail(
    id serial primary key,
    cart_id int not null,
    product_id int not null,
    quantity int not null,
    created_at timestamp default current_timestamp 
)