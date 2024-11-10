CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name VARCHAR(25) UNIQUE NOT NULL
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE url_tags (
    url_id INTEGER REFERENCES urls(id) ON DELETE CASCADE,
    tag_id INTEGER REFERENCES tags(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (url_id, tag_id)
);

CREATE INDEX idx_url_tags_url_id ON url_tags (url_id);
CREATE INDEX idx_url_tags_tag_id ON url_tags (tag_id);
CREATE INDEX idx_tags_name ON tags(name);