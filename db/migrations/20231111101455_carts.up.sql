CREATE TABLE carts (
    id SERIAL PRIMARY KEY,
    user_id varchar(50) unique,
    Foreign Key (user_id) REFERENCES customers(username),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE carts_detail (
    id SERIAL PRIMARY KEY,
    cart_id int not null,
    product_id INT not null unique,
    quantity INT not null,
    price int default 0,
    FOREIGN KEY (cart_id) REFERENCES carts(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);