CREATE TABLE Messages(
    message_id BIGSERIAL,
    value Text NOT NULL,
    is_read Boolean NOT NULL DEFAULT FALSE,
    sender_actor INTEGER NOT NULL,
    to_actor INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT pk_message PRIMARY KEY (message_id),
    CONSTRAINT fc_message_sender
        FOREIGN KEY (sender_actor) REFERENCES Actors
        ON DELETE CASCADE,
    CONSTRAINT fc_message_to
        FOREIGN KEY (to_actor) REFERENCES Actors
        ON DELETE CASCADE
);