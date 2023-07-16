package main

import (
  "log"

  "github.com/xdb/go-xdb-sdk/client"
  "github.com/xdb/go-xdb-sdk/proto"
)

func main() {
  log.Printf("???")

  var client *xxxx_client.TransactionClient
  var tx0 XXXX.Transaction
  var tx []*XXXX.Transaction
  var uid uint64
  var resp map[uint64]int32

  client = xxxx_client.MakeTransactionClient("127.0.0.1",10005)

  uid, _ = client.SendRawTransaction(1, "I", "you", 100)

  log.Printf("get uid %s\n",uid)

  tx = make([]*XXXX.Transaction, 1)

  log.Printf("????", len(tx))

  tx0.From = "I"
  tx0.To = "you"
  tx0.Amount = 100
  tx[0] = &tx0
  log.Printf("????")

  resp, _ = client.SendBatchTransaction(tx)
  log.Printf("get ret:",resp)
}

