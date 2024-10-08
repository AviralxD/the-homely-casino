-- +goose Up

create table players (
  name text not null unique,
  totalBet int,
  pack bool default(false),
  show bool default(false)
);

-- +goose Down

drop table players;