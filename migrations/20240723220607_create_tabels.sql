-- +goose Up
-- +goose StatementBegin
CREATE TABLE message(
    id         SERIAL   NOT NULL,
    text TEXT NOT NULL,
    key UUID NOT NULL,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE message;
-- +goose StatementEnd
