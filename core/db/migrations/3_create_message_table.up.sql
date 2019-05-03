CREATE TABLE Messages(
    message_id BIGSERIAL,
    value Text,
    sender_actor INTEGER,
    to_actor INTEGER,
    CONSTRAINT pk_message PRIMARY KEY (message_id),
    CONSTRAINT fc_message_sender
        FOREIGN KEY (sender_actor) REFERENCES Actors
        ON DELETE CASCADE,
    CONSTRAINT fc_message_to
        FOREIGN KEY (to_actor) REFERENCES Actors
        ON DELETE CASCADE
);  