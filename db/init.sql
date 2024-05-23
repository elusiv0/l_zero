CREATE TABLE IF NOT EXISTS delivery (
    delivery_id SERIAL PRIMARY KEY,
    delivery_name VARCHAR,
    phone VARCHAR,
    zip INT,
    city VARCHAR,
    delivery_address VARCHAR,
    region VARCHAR,
    email VARCHAR
);
CREATE TABLE IF NOT EXISTS payment (
    payment_id SERIAL PRIMARY KEY,
    payment_transaction VARCHAR, 
    request_id VARCHAR,
    currency VARCHAR,
    payment_provider VARCHAR,
    amount INT,
    payment_dt INT,
    bank VARCHAR,
    delivery_cost INT,
    goods_total INT,
    custom_fee INT
);
CREATE TABLE IF NOT EXISTS item (
    item_id SERIAL PRIMARY KEY,
    chrt_id int,
    track_number VARCHAR,
    price int,
    rid VARCHAR,
    item_name VARCHAR,
    sale int,
    size int,
    total_price int,
    nm_id int,
    brand VARCHAR,
    item_status int
);
CREATE TABLE IF NOT EXISTS order_ (
    order_uid VARCHAR PRIMARY KEY,
    track_number VARCHAR,
    order_entry VARCHAR,
    delivery_id int REFERENCES delivery (delivery_id), 
    payment_id int REFERENCES payment (payment_id),
    locale VARCHAR,
    internal_signature VARCHAR,
    customer_id VARCHAR,
    delivery_service VARCHAR,
    shard_key int,
    sm_id int,
    date_created timestamp,
    oof_shard int
);