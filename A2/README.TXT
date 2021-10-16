The directory contents are as follows:

A2/reservation/types.go

Contains type declarations for Reservation and Flight types,
including their member functions, of which the RPC implementations are among.

A2/rsvserver/rsvserver.go

Server that listens and handles incoming RPC requests. Run with "go run rsvserver.go"


A2/rsvclient/rsvclient.go

Client to make RPC requests (list, passengerlist, and reserve).

Run with "go run rsvclient.go list localhost 5999"
Run with "go run rsvclient.go passengerlist localhost 5999"
Run with "go run rsvclient.go reserve localhost 5999 economy michael 23"

