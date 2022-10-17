CREATE TABLE users (
    id SERIAL PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    email varchar(100) NOT NULL,
    created_at timestamptz DEFAULT now()
);