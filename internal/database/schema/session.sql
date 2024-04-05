CREATE TABLE session (
    id TEXT PRIMARY KEY,
    user_id TEXT NOT NULL,
    fingerprint BLOB NOT NULL,
    ip INTEGER NOT NULL,
    expires_at DATETIME NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id)
);