CREATE TABLE users (
    user_id SERIAL PRIMARY KEY,
    _name VARCHAR(100) NOT NULL,
    last_name VARCHAR(100) NOT NULL,
    country VARCHAR(50),
    date_of_birth DATE
);

CREATE TABLE wallet (
    wallet_id SERIAL PRIMARY KEY,
    wallet_name VARCHAR(1000) NOT NUL,
);