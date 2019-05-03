CREATE TABLE Actors(
    actor_id SERIAL PRIMARY KEY,
    actor_name VARCHAR (35),
    display_name VARCHAR (60),
    description VARCHAR (255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT UC_Actor_Name UNIQUE (actor_name)
);