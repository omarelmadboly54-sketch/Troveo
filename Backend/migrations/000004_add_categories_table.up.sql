CREATE TABLE IF NOT EXISTS categories(
    id BIGSERIAL PRIMARY KEY,
    category VARCHAR(30) NOT NULL,
    parent_id BIGINT,
    Foreign Key (parent_id) REFERENCES categories(id) ON DELETE CASCADE
);



