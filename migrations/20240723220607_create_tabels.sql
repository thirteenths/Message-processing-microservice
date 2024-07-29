-- +goose Up
-- +goose StatementBegin
CREATE TABLE message(
    id         SERIAL   NOT NULL,
    text TEXT NOT NULL,
    key UUID NOT NULL,
    status BOOLEAN DEFAULT FALSE,
    PRIMARY KEY (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE message;
-- +goose StatementEnd
