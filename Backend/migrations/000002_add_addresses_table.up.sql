CREATE TABLE IF NOT EXISTS addresses(
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    city VARCHAR(30) NOT NULL,
    area VARCHAR(30) NOT NULL,
    street VARCHAR(40) NOT NULL,
    building_number INT NOT NULL,
    apartment_number INT,
    phone_number VARCHAR(15) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, 
    Foreign Key (user_id) REFERENCES users(id) ON DELETE CASCADE
);
