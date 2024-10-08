-- +goose Up

create table rounds (
  currentBet int not null,
  totalBet int not null,
  playerNum int
);

-- +goose Down

drop table rounds;