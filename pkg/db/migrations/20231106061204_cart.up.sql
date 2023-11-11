CREATE TABLE carts(
    id serial primary key,
    user_id int not null,
    total int,
    created_at timestamp default current_timestamp
)
