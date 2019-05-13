CREATE TABLE Actors(
    actor_id SERIAL PRIMARY KEY,
    actor_name VARCHAR (35) NOT NULL,
    display_name VARCHAR (60) NOT NULL DEFAULT '',
    description VARCHAR (255) NOT NULL DEFAULT '',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT UC_Actor_Name UNIQUE(actor_name)
);