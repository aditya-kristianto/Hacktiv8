CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name varchar(100) NOT NULL,
    contact varchar(20) NOT NULL,
    email varchar(100) NOT NULL,
    created_at timestamptz DEFAULT now()
);