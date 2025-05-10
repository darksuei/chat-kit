CREATE TABLE files (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    identifier TEXT NOT NULL,
    name TEXT NOT NULL,

    channel_id INTEGER NOT NULL,
    participant_id INTEGER NOT NULL,
    message_id INTEGER,
    version_id INTEGER,
    permission_id INTEGER,
    scope TEXT,

    FOREIGN KEY (channel_id) REFERENCES channels(id) ON DELETE CASCADE,
    FOREIGN KEY (participant_id) REFERENCES channel_participants(id) ON DELETE CASCADE,
    FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE SET NULL
);

CREATE TABLE file_versions (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP,

    file_id INTEGER NOT NULL,
    revision TEXT NOT NULL DEFAULT 'LATEST',
    is_active BOOLEAN NOT NULL DEFAULT TRUE,

    FOREIGN KEY (file_id) REFERENCES files(id) ON DELETE CASCADE
);
