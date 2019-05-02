CREATE TABLE Actor(
    actor_id SERIAL PRIMARY KEY,
    actor_name VARCHAR (32),
    display_name VARCHAR (60),
    description VARCHAR (255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);