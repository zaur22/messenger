CREATE TABLE Message(
    message_id BIGSERIAL,
    value Text,
    sender_actor INTEGER,
    to_actor INTEGER,
    CONSTRAINT pk_message PRIMARY KEY (message_id),
    CONSTRAINT fc_message_sender
        FOREIGN KEY (sender_actor) REFERENCES actor
        ON DELETE CASCADE,
    CONSTRAINT fc_message_to
        FOREIGN KEY (to_actor) REFERENCES actor
        ON DELETE CASCADE
);  