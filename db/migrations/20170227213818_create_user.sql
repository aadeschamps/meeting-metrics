
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    display_name text,
    email text UNIQUE,
    password_digest text
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
