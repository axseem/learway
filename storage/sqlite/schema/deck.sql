CREATE TABLE deck (
    id TEXT PRIMARY KEY,
    user_id text NOT NULL,
    title TEXT NOT NULL,
    cards BLOB NOT NULL,
    subject TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES user (id)
);