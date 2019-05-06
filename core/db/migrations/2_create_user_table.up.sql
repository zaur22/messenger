CREATE TABLE Users(
    actor   integer,
    pass    BYTEA NOT NULL,
    CONSTRAINT fc_user_actor
        FOREIGN KEY (actor) REFERENCES Actors
        ON DELETE CASCADE,
    CONSTRAINT pk_user PRIMARY KEY (actor)
);