
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE teams_users (
    id SERIAL PRIMARY KEY,
    team_id integer REFERENCES teams,
    user_id integer REFERENCES users,
    admin_user boolean
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE teams_users;
