-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    email VARCHAR (128) NOT NULL DEFAULT '',
    password VARCHAR (128) NOT NULL DEFAULT '',
    name VARCHAR (32) NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (email)
);

INSERT INTO users ("email", "password", "name") VALUES ('admin@test.com', '$2a$10$z0aFtC6R0SU/AVaBLyW4cezo/r5wBdOH.7dsrgweAndBkijcwHE1m', 'Admin');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
