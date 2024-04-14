CREATE TABLE user (
    id TEXT PRIMARY KEY,
    username TEXT UNIQUE NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password BLOB NOT NULL,
    name TEXT NOT NULL,
    description TEXT,
    picture TEXT,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);