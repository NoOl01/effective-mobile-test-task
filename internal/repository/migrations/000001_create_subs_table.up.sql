CREATE TABLE IF NOT EXISTS subs (
    id SERIAL PRIMARY KEY,
    user_id VARCHAR(50) NOT NULL,
    service_name VARCHAR(255) NOT NULL,
    price INTEGER NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE NOT NULL
);