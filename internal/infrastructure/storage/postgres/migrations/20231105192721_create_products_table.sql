-- +goose Up
-- +goose StatementBegin
CREATE TABLE products (
    id bigserial NOT NULL PRIMARY KEY,
    sku int NOT NULL,
    name varchar(80),
    price int,
    created_at timestamp NOT NULL DEFAULT now(),
    updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TRIGGER set_timestamp_trigger_products
BEFORE UPDATE ON products
FOR EACH ROW
EXECUTE FUNCTION trigger_set_timestamp();
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE products;
-- +goose StatementEnd
