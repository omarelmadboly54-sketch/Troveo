CREATE TABLE IF NOT EXISTS user_coupons(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    coupon_id BIGINT NOT NULL,
    is_used BOOLEAN NOT NULL DEFAULT false,
    assigned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Foreign Key (user_id) REFERENCES users(id) ON DELETE CASCADE,
    Foreign Key (coupon_id) REFERENCES coupons(id) ON DELETE CASCADE,
    CONSTRAINT unique_user_coupon UNIQUE(user_id,coupon_id)
);
