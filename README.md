##  Mini-Soc-Clie (Minimalist Socket Client)

name inspired from 'miniscule'

Focus is to minimize a Golang binary for a client application

[Inspiration](https://hacktofinale.dyte.io/challenges/golf)

### Problem Statement

Develop a minimalistic client (in golang) using websocket to send a "hello" message to a server and print the server's response. The twist? Your goal is to make the executable binary as tiny as possible!

### Requirements Gathering

[gfg article on socket programming in C++](https://www.geeksforgeeks.org/socket-programming-cc/)

[PPT on Socket & Network in C](https://www.csd.uoc.gr/~hy556/material/tutorials/cs556-3rd-tutorial.pdf)

[Tutorial on creating websocket in Go](https://yalantis.com/blog/how-to-build-websockets-in-go/)

#### Some resources for Golang

[Go by Examples](https://gobyexample.com/)

From the above articles it seems that sockets are like files, there's a protocol which they have to
follow to talk with each other, C has sockets library, there are some
standards which we have to adhere to, to be able to talk with server from client

[Profiling and benchmarking golang code blog](https://go.dev/blog/pprof)

This is a very good blog, details on how we can analyze where our memory is being consumed

[Paper on Benchmark of C++/Java/Go/Scala](https://research.google/pubs/pub37122/)

One more way to reduce Golang Binary size:
[Blog](https://gophercoding.com/reduce-go-binary-size/#:~:text=In%20short%2C%20adding%20the%20ldflags,commands%20like%20go%20tool%20nm%20.)








