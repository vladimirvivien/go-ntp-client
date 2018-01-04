# A trivial NTP Client
This repository is an implementation of a (really) trivial NTP client in Go. It uses the `encoding/binary` package to encode and decode NTP packets sent to and received from a remote NTP server over UDP. You can learn more about NTP here, read the specs RFC5905, and find a (seemingly) way better Go NTP client, with many features implemented, [here](https://github.com/beevik/ntp).

The code is explained in this writeup titled [Let's make an NTP client in Go](https://medium.com/learning-the-go-programming-language/lets-make-an-ntp-client-in-go-287c4b9a969f). 
