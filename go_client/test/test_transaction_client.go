package main

import (
  "log"

  "go_client/client"
)

func main() {
  log.Printf("???")

  var client *resdb_client.TransactionClient
  var uid uint64

  client = resdb_client.MakeTransactionClient("127.0.0.1",30005)

  uid, _ = client.SendRawTransaction(1, "I", "you", 100)

  log.Printf("get uid %s\n",uid)
}

