CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    channel_id INTEGER NOT NULL,
    participant_id INTEGER NOT NULL,

    content TEXT NOT NULL,

    is_child BOOLEAN NOT NULL DEFAULT FALSE,
    parent_id INTEGER,

    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (participant_id) REFERENCES channel_participants(id) ON DELETE CASCADE,
    FOREIGN KEY (parent_id) REFERENCES messages(id) ON DELETE SET NULL
);

CREATE TABLE message_reactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    message_id INTEGER NOT NULL,
    emoji TEXT NOT NULL,

    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);

CREATE TABLE message_read_by (
    message_id INTEGER NOT NULL,
    channel_participant_id INTEGER NOT NULL,
    PRIMARY KEY (message_id, channel_participant_id),
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    FOREIGN KEY (channel_participant_id) REFERENCES channel_participants(id) ON DELETE CASCADE
);

CREATE TABLE message_mentions (
    message_id INTEGER NOT NULL,
    channel_participant_id INTEGER NOT NULL,
    PRIMARY KEY (message_id, channel_participant_id),
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
    FOREIGN KEY (channel_participant_id) REFERENCES channel_participants(id) ON DELETE CASCADE
);
