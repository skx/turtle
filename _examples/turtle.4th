\
\ This is an example of drawing some simple graphics with FORTH
\
\



\
\ Draw an eqilateral triangle with the sides having the length of the
\ value on the stack
\
: triangle
  dup dup
   3 0 do
     forward
     120 turn
   loop
;

: square
  dup dup dup dup
  4 0 do
    forward
    90 turn
  loop
;

: circle
360 0 do
 1 forward
 1 turn
loop
;



\
\ Pen is down, touching the paper, ensuring that movements
\ result in pixels being set.
\
1 pen


\
\ Move to 50,50 and draw a small triangle
\
50 50 move
30 triangle

\
\ Move to 100,100 and draw a larger triangle
\
100 100 move
90 triangle


\
\ Square time now!
\
250 150 move
10 square

\
\ And again
\
150 150 move
75 square


\
\ Circle time is upon us. Beware the Elves
\
20 200 move
circle


\
\ Save the result into `turtle.png`
\
save

\
\ Final newline is important..
\
