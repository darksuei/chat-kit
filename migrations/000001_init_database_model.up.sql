CREATE TYPE channel_participant_role AS ENUM ('creator', 'admin', 'participant');

CREATE TABLE channels (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    is_direct BOOLEAN NOT NULL,
    description TEXT,
    image_id BIGINT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

CREATE TABLE channel_participants (
    id BIGSERIAL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    role channel_participant_role NOT NULL,
    channel_id BIGINT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP,
    CONSTRAINT fk_channel FOREIGN KEY (channel_id) REFERENCES channels(id)
);
