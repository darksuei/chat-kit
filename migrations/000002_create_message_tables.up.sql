-- Messages table
CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    channel_id INTEGER NOT NULL,
    participant_id INTEGER,
    content TEXT NOT NULL,
    is_child BOOLEAN DEFAULT FALSE,
    parent_id INTEGER,

    CONSTRAINT fk_messages_channel_id
        FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,

    CONSTRAINT fk_messages_participant_id
        FOREIGN KEY (participant_id) REFERENCES channel_participants(id) ON DELETE SET NULL,

    CONSTRAINT fk_messages_parent_id
        FOREIGN KEY (parent_id) REFERENCES messages(id) ON DELETE CASCADE
);

-- Message ReadBy join table
CREATE TABLE message_read_by (
    message_id INTEGER NOT NULL,
    channel_participant_id INTEGER NOT NULL,
    PRIMARY KEY (message_id, channel_participant_id),

    CONSTRAINT fk_readby_message
        FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,

    CONSTRAINT fk_readby_participant
        FOREIGN KEY (channel_participant_id) REFERENCES channel_participants(id) ON DELETE CASCADE
);

-- Message Mentions join table
CREATE TABLE message_mentions (
    message_id INTEGER NOT NULL,
    channel_participant_id INTEGER NOT NULL,
    PRIMARY KEY (message_id, channel_participant_id),

    CONSTRAINT fk_mentions_message
        FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,

    CONSTRAINT fk_mentions_participant
        FOREIGN KEY (channel_participant_id) REFERENCES channel_participants(id) ON DELETE CASCADE
);

-- Message Reactions table
CREATE TABLE message_reactions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    message_id INTEGER NOT NULL,
    emoji TEXT NOT NULL,

    CONSTRAINT fk_reactions_message
        FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);
