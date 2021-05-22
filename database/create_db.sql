CREATE TABLE url_mapping (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    url_id VARCHAR(10) NOT NULL,
    origin_url TEXT NOT NULL,
    expire_at TIMESTAMP NOT NULL
);

CREATE INDEX idx_url_id ON url_mapping(url_id);
CREATE INDEX idx_origin_url ON url_mapping(origin_url);