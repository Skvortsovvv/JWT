-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
       id uuid primary key,
       login text not null,
       password text not null,
       permission int not null
);

CREATE UNIQUE INDEX users_login_unique_index ON users (login);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
