CREATE TABLE "user"(
    actor   INTEGER,
    pass    BYTEA,
    CONSTRAINT fc_user_actor
        FOREIGN KEY (actor) REFERENCES Actor
        ON DELETE CASCADE,
    CONSTRAINT pk_user PRIMARY KEY (actor)
);