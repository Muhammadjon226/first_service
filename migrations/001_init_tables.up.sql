CREATE TABLE IF NOT EXISTS posts (
    id INTEGER NOT NULL PRIMARY KEY,
    user_id INTEGER,
    title TEXT,
    body TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);