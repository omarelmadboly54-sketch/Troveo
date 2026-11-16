ALTER TABLE coupons
ALTER COLUMN expired_at TYPE TIMESTAMP
USING expired_at AT TIME ZONE 'Africa/Cairo';
