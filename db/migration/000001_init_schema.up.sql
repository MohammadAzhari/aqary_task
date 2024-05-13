CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255),
    password VARCHAR(255),
    phone_number VARCHAR(20) UNIQUE,
    otp VARCHAR(10),
    otp_expiration_time TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profile (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    first_name VARCHAR(100),
    last_name VARCHAR(100),
    address TEXT
);

CREATE INDEX idx_phone_number ON users (phone_number);
