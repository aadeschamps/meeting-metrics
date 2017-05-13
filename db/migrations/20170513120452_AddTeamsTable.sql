
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    display_name text
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE teams;
