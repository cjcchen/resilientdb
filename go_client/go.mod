module go_client

go 1.17

require (
	github.com/golang/protobuf v1.5.3
	github.com/resilientdb/go-resilientdb-sdk v0.0.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/avast/retry-go v3.0.0+incompatible // indirect
	github.com/diem/client-sdk-go v1.2.1 // indirect
	github.com/novifinancial/serde-reflection/serde-generate/runtime/golang v0.0.0-20201214184956-1fd02a932898 // indirect
	golang.org/x/crypto v0.0.0-20200728195943-123391ffb6de // indirect
	golang.org/x/sys v0.0.0-20200812155832-6a926be9bd1d // indirect
)

replace github.com/resilientdb/go-resilientdb-sdk => ./
