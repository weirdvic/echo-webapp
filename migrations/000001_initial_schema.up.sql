CREATE TABLE IF NOT EXISTS snippets (
    id SERIAL,
    title VARCHAR(128),
    content TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP + '1 day'::interval
)
;
