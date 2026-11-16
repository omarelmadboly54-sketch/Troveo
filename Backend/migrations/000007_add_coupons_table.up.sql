CREATE TABLE IF NOT EXISTS coupons(
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(15) UNIQUE NOT NULL,
    discount_type VARCHAR(30) NOT NULL,
    value NUMERIC(10,2) NOT NULL,
    minimum_order_amount NUMERIC(10,2) NOT NULL,
    expired_at TIMESTAMP NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT coupon_discount_type
    CHECK(
        discount_type IN ('percentage','fixed')
    ),
    CONSTRAINT coupon_value_check
    CHECK(
        value>0 AND (discount_type!='percentage' OR value<=100)
    ),
    CONSTRAINT minimum_order_amount_check
    CHECK(
        minimum_order_amount>=0
    )
);
