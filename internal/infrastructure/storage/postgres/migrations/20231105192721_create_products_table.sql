-- +goose Up
-- +goose StatementBegin
CREATE TABLE production (
    id bigserial NOT NULL PRIMARY KEY,
    sku int NOT NULL,
    name varchar(80),
    price int,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TRIGGER set_timestamp_trigger_productions
BEFORE UPDATE ON productions
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE productions;
-- +goose StatementEnd
