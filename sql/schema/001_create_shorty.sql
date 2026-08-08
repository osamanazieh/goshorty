-- +goose Up 
CREATE TABLE shorty(
    id UUID PRIMARY KEY, 
    url TEXT UNIQUE NOT NULL,
    short_code TEXT UNIQUE NOT NULL, 
    created_at TIMESTAMP NOT NULL, 
    updated_at TIMESTAMP NOT NULL, 
    hits INT NOT NULL DEFAULT 0
);

-- +goose Down
DROP TABLE shorty;