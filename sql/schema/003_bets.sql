-- +goose Up

create table bets (
  name text not null unique references players(name) on delete cascade,
  amount int not null
);

-- +goose Down

drop table bets;