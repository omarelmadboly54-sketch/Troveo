CREATE TABLE IF NOT EXISTS products(
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(300) NOT NULL,
    price NUMERIC(10,2) NOT NULL,
    photo VARCHAR(255) NOT NULL,
    quantity INT NOT NULL DEFAULT 0,
    category_id BIGINT NOT NULL,
    seller_id BIGINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Foreign Key (category_id) REFERENCES categories(id) ON DELETE RESTRICT,
    Foreign Key (seller_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT product_price_check
    CHECK (price>=0),
    CONSTRAINT product_quantity_check
    CHECK (quantity>=0)
);
