CREATE TABLE Users(
    actor   VARCHAR (35),
    pass    BYTEA,
    CONSTRAINT fc_user_actor
        FOREIGN KEY (actor) REFERENCES Actors(actor_name)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    CONSTRAINT pk_user PRIMARY KEY (actor)
);