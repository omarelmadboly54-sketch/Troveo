ALTER TABLE coupons
ALTER COLUMN expired_at TYPE TIMESTAMPTZ
USING expired_at AT TIME ZONE 'Africa/Cairo';