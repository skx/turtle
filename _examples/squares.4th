1 pen

: square
    dup dup dup
    4 0 do
        forward
        90 turn
    loop
;

30 0 do
    150 150 move
    360 m /  i *  direction
    80 square
loop

save
