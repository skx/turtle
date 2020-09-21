\
\ Draw a grid
\

1 pen

\
\ Draw a series of vertical lines
\
: vert
  180 direction
  31 0 do
    10 i * 0 move
    300 forward
  loop
;

\
\ Horizontal lines
\
: horiz
  90 direction
  31 0 do
    0 10 i * move
    300 forward
  loop
;

vert 10
horiz 10
save
