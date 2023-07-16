package main

import (
  "log"

  "github.com/xdb/go-xdb-sdk/client"
)

func main() {
  var client *resdb_client.PoCTransactionClient
  var resp map[uint64]int32
  var req []uint64

  req = append(req,1)

  client = resdb_client.MakePoCTransactionClient("127.0.0.1",10006)

  resp, _ = client.Query(req)

  log.Printf("get uid %s\n",resp)
}

