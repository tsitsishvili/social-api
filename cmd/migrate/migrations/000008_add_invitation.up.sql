CREATE TABLE IF NOT EXISTS user_invitations (
    token bytea PRIMARY KEY,
    user_id BIGINT NOT NULL,
    expires TIMESTAMP(0) WITH TIME ZONE NOT NULL
);