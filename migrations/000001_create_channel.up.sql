CREATE TABLE channels (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    "name" TEXT NOT NULL,
    is_direct BOOLEAN NOT NULL,
    "description" TEXT,
    image_id INTEGER
);

CREATE TABLE channel_participants (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    user_id INTEGER NOT NULL
);

CREATE TABLE channel_participants_channels (
    channel_id INTEGER NOT NULL,
    channel_participant_id INTEGER NOT NULL,
    PRIMARY KEY (channel_id, channel_participant_id),
    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (channel_participant_id) REFERENCES channel_participants(id) ON DELETE CASCADE
);