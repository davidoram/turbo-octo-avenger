
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table users (
  id bigserial primary key not null,
  email text not null,
  password text not null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
drop table users;
