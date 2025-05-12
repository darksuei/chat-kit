-- Remove the column
ALTER TABLE channel_participants
DROP COLUMN role;

-- Drop the enum type (PostgreSQL only; safe to ignore if not used)
DROP TYPE IF EXISTS channel_participant_role;
