module server

go 1.22.2

require (
	cipher v0.0.0-00010101000000-000000000000
	google.golang.org/grpc v1.63.2
)

require (
	github.com/zenazn/pkcs7pad v0.0.0-20170308005700-253a5b1f0e03 // indirect
	golang.org/x/net v0.21.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240227224415-6ceb2ff114de // indirect
	google.golang.org/protobuf v1.33.0 // indirect
)

replace cipher => ./cipher
