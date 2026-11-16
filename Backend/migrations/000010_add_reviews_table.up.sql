CREATE TABLE IF NOT EXISTS reviews(
    id BIGSERIAL PRIMARY KEY,
    product_id BIGINT NOT NULL,
    comment VARCHAR(500) NOT NULL,
    rating INT NOT NULL,
    user_id BIGINT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    Foreign Key (product_id) REFERENCES products(id) ON DELETE CASCADE,
    Foreign Key (user_id) REFERENCES users(id) ON DELETE CASCADE,
    CONSTRAINT rating_check
    CHECK(
        rating BETWEEN 1 AND 5
    ),
    CONSTRAINT unique_user_product_review UNIQUE(user_id,product_id)
);

