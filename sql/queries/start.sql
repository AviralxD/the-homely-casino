-- name: startGame :one
insert into rounds (currentBet, totalBet, playerNum)
values (1, 1, 0)
returning * ;

-- name clearGame :exec
delete * from rounds; 

-- name addPlayer :one
update rounds 
set playerNum = playerNum + 1;
insert into players (name, bet, totalBet, pack, show)
values ($1, [0], 0, false, false);


