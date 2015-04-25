
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE ping (
  id bigserial primary key not null,
  message text
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE ping;
