CREATE TABLE refresh_tokens (
    user_id    UUID PRIMARY KEY,
    token      TEXT NOT NULL UNIQUE,
    expire_at  TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,

    CONSTRAINT fk_refresh_tokens_user
        FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
);