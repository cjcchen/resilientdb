package xxxx_client

import (
	"fmt"
	"net"
  "log"
    "encoding/binary"

  "github.com/xdb/go-xdb-sdk/proto"
  "github.com/golang/protobuf/proto"
)

type TransactionClient struct {
  ip string
  port int
  conn net.Conn
}

func MakeTransactionClient(ip string, port int) *TransactionClient {
  return &TransactionClient{
    ip:ip,
    port:port+20000,
    conn:nil,
  }
}

func (c*TransactionClient)GetIp() string {
  return c.ip
}

func (c * TransactionClient)SendRawTransaction(uid uint64, from string, to string, amount uint64) (uint64, error){

  var tx resdb.Transaction
  var req resdb.TransactionsRequest
  var resp *resdb.TransactionsResponse
  var err error

  tx.Uid = uid
  tx.From = from
  tx.To =  to
  tx.Amount = amount

  req.Transactions = make([]*resdb.Transaction, 1)
  req.Transactions[0] = &tx

  for i:=0; i < 3; i++ {
    resp, err = c.SendTransaction(&req)
      if(err != nil){
        continue
      }
    break;
  }

  if (resp.Result[0].Ret <0) {
    return 0, err
  }
  return uid, nil
}

func (c * TransactionClient)SendBatchTransaction(txns []* resdb.Transaction) (map[uint64]int32, error){
  var req resdb.TransactionsRequest
  var resp *resdb.TransactionsResponse
  var resp_list map[uint64]int32
  var err error

  req.Transactions = txns

  resp, err = c.SendTransaction(&req)
  if(err != nil){
    return nil, err
  }

  resp_list = make(map[uint64]int32)

  for _, v := range resp.Result {
    if v.Ret >=0 {
      resp_list[v.Uid] = 1
    } else {
      resp_list[v.Uid] = -1
    }
  }

  return resp_list, nil
}


func (c * TransactionClient)SendTransaction(req *resdb.TransactionsRequest) (*resdb.TransactionsResponse, error){
  var data_len uint32
  var read_len uint32
  var bs []byte
  var err error
  var data []byte
  var response resdb.TransactionsResponse

  if c.conn == nil {
    c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.ip, c.port))
    if err != nil {
        log.Println("connect fail",err)
        return nil, err
    }
    log.Printf("connect to %s:%d\n",c.ip, c.port)
  }

  //defer conn.Close() // 关闭TCP连接

  data, err = proto.Marshal(req)
  if err != nil {
    log.Fatalln("Mashal data error:", err)
    return nil, err
  }

  data_len = uint32(len(data))
  bs = make([]byte, 8)
  binary.LittleEndian.PutUint32(bs, data_len)

  _, err = c.conn.Write(bs)
  if err != nil {
    c.conn.Close() // 关闭TCP连接
    c.conn=nil
    return nil, err
  }

  _, err = c.conn.Write([]byte(data))
  if err != nil {
    c.conn.Close() // 关闭TCP连接
    c.conn=nil
    return nil, err
  }

  _, err = c.conn.Read(bs)
  if err != nil {
    log.Println("recv failed, err:", err)
    c.conn.Close() // 关闭TCP连接
    c.conn=nil
    return nil, err
  }

  read_len = binary.LittleEndian.Uint32(bs)

  bs = make([]byte, read_len)
  _, err = c.conn.Read(bs)
  if err != nil {
    log.Println("recv failed, err:", err)
    c.conn.Close() // 关闭TCP连接
    c.conn=nil
      return nil, err
  }

  err = proto.Unmarshal(bs, &response)
  if err != nil{
    log.Fatalln("UnMashal data error:", err)
  }

  return &response, nil
}


