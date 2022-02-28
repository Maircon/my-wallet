module my-wallet

go 1.17

require (
	mywallet.com/db v0.0.0-00010101000000-000000000000
	mywallet.com/routers v0.0.0-00010101000000-000000000000
)

require (
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/lib/pq v1.10.4 // indirect
)

replace mywallet.com/db => ./database

replace mywallet.com/routers => ./routers
