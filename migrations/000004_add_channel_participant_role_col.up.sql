DO $$ BEGIN
    CREATE TYPE channel_participant_role AS ENUM ('creator', 'admin', 'participant');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

ALTER TABLE channel_participants
ADD COLUMN role channel_participant_role NOT NULL DEFAULT 'participant';
