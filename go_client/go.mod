module go_client

go 1.17

require (
	github.com/golang/protobuf v1.5.3
	github.com/resilientdb/go-resilientdb-sdk v0.0.0
	google.golang.org/protobuf v1.30.0
)

replace github.com/resilientdb/go-resilientdb-sdk => ./
