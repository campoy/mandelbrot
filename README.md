mandelbrot
==========

An example on how to use different concurrency patterns with Go to compute Mandelbrot sets.

It turns out the simplest way is also the fastest!
Just run one goroutine per pixel and you get more than 50% speed-up

![A cool mandelbrot](/mandelbrot.png "Mandelbrot set")

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=seq
real	0m25.858s
user	0m25.260s
sys	0m0.219s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=px
real	0m11.575s
user	0m27.576s
sys	0m0.186s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=row
real	0m11.854s
user	0m27.890s
sys	0m0.188s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=1
real	0m36.673s
user	0m42.374s
sys	0m21.116s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=2
real	0m23.045s
user	0m37.755s
sys	0m9.540s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=4
real	0m16.897s
user	0m39.619s
sys	0m4.055s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=8
real	0m17.390s
user	0m41.210s
sys	0m3.086s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=16
real	0m16.923s
user	0m40.004s
sys	0m2.418s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=32
real	0m16.603s
user	0m40.724s
sys	0m2.727s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=64
real	0m16.379s
user	0m40.183s
sys	0m2.468s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=128
real	0m16.006s
user	0m39.889s
sys	0m2.148s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=256
real	0m16.055s
user	0m39.902s
sys	0m2.155s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=512
real	0m16.028s
user	0m39.938s
sys	0m2.354s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=1024
real	0m16.246s
user	0m39.893s
sys	0m3.084s

time mandelbrot -out=out.png -h=4096 -w=4096 -mode=workers -workers=2048
real	0m16.263s
user	0m40.196s
sys	0m3.253s
