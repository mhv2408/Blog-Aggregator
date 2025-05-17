-- +goose UP
CREATE TABLE posts(
    id UUID PRIMARY KEY,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    title TEXT NOT NULL,
    url TEXT NOT NULL UNIQUE,
    description TEXT,
    published_at DATE NOT NULL,
    feed_id UUID NOT NULL,
    FOREIGN KEY(feed_id) REFERENCES feed(id) ON DELETE CASCADE
);

-- +goose DOWN
DROP TABLE posts;