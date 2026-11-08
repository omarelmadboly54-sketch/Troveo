CREATE TABLE IF NOT EXISTS users(
    id BIGSERIAL PRIMARY KEY,
    first_name VARCHAR(30) NOT NULL,
    last_name  VARCHAR(30) NOT NULL,
    email VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(255) NOT NULL,
    role VARCHAR(30) NOT NULL DEFAULT('user'),
    account_type VARCHAR(30) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    CONSTRAINT account_type_check
    CHECK (account_type IN ('admin','seller','customer')),
    CONSTRAINT role_check
    CHECK (role IN ('admin','user')),
    CONSTRAINT user_role_account_type_check
    CHECK (
        (role ='admin' AND account_type='admin')
        OR
        (role='user' AND account_type IN ('seller','customer'))
    )
);
