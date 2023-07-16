package resdb_client

import (
	"fmt"
	"net"
  "log"
    "encoding/binary"

  "github.com/xdb/go-xdb-sdk/proto"
  "github.com/golang/protobuf/proto"
)

type PoCTransactionClient struct {
  ip string
  port int
  conn net.Conn
}

func MakePoCTransactionClient(ip string, port int) *PoCTransactionClient {
  return &PoCTransactionClient{
    ip:ip,
    port:port+20000,
    conn:nil,
  }
}

func (c*PoCTransactionClient)GetIp() string {
  return c.ip
}

func (c * PoCTransactionClient)Query(uids []uint64) (map[uint64]int32, error){
  var req resdb.TransactionQuery
  var resp *resdb.TransactionQuery
  var resp_list map[uint64]int32
  var err error

  req.Uids = uids

  resp, err = c.QueryTransaction(&req)
  if(err != nil){
    return nil, err
  }


  resp_list = make(map[uint64]int32)

  for _, v := range resp.Uids {
      resp_list[v] = 1
  }

  return resp_list, nil
}

func (c * PoCTransactionClient)SendBatchTransaction(txns []* resdb.Transaction) (error){
  var req resdb.TransactionsRequest
  var err error

  req.Transactions = txns

  err = c.SendTransaction(&req)
  if(err != nil){
    return err
  }

  return nil
}

func (c * PoCTransactionClient)SendTransaction(req *resdb.TransactionsRequest) (error){
  var data_len uint32
  var bs []byte
  var err error
  var data []byte

  if c.conn == nil {
    c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.ip, c.port))
    if err != nil {
        log.Println("connect fail",err)
        return err
    }
    log.Printf("connect to %s:%d\n",c.ip, c.port)
  }

  data, err = proto.Marshal(req)
  if err != nil {
    log.Fatalln("Mashal data error:", err)
    return  err
  }

  data_len = uint32(len(data))
  bs = make([]byte, 8)
  binary.LittleEndian.PutUint32(bs, data_len+1)

  _, err = c.conn.Write(bs)
  if err != nil {
    log.Printf("write fail")
    c.conn.Close()
    c.conn=nil
    return err
  }

  bs = make([]byte, 1)
  bs[0] = 1

  _, err = c.conn.Write(bs)
  if err != nil {
    c.conn.Close()
    c.conn=nil
    return err
  }

  _, err = c.conn.Write([]byte(data))
  if err != nil {
    c.conn.Close()
    c.conn=nil
    return err
  }

  return nil
}


func (c * PoCTransactionClient)QueryTransaction(req *resdb.TransactionQuery) (*resdb.TransactionQuery, error){
  var data_len uint32
  var read_len uint32
  var bs []byte
  var err error
  var data []byte
  var response resdb.TransactionQuery
  //log.Printf("query txn ip %s port %d",c.ip, c.port)
  if c.conn == nil {
    c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", c.ip, c.port))
    if err != nil {
        log.Println("connect fail",err)
        return nil, err
    }
    log.Printf("connect to %s:%d\n",c.ip, c.port)
  }

  data, err = proto.Marshal(req)
  if err != nil {
    log.Fatalln("Mashal data error:", err)
    return nil, err
  }

  data_len = uint32(len(data))
  bs = make([]byte, 8)
  binary.LittleEndian.PutUint32(bs, data_len+1)

  _, err = c.conn.Write(bs)
  if err != nil {
    c.conn.Close()
    c.conn=nil
    log.Printf("query write fail")
    return nil, err
  }

  bs = make([]byte, 1)
  bs[0] = 0
  _, err = c.conn.Write(bs)
  if err != nil {
    c.conn.Close()
    c.conn=nil
    log.Printf("query write fail")
    return nil, err
  }

  _, err = c.conn.Write([]byte(data))
  if err != nil {
    c.conn.Close()
    c.conn=nil
    log.Printf("query write fail")
    return nil, err
  }

  bs = make([]byte, 8)
  _, err = c.conn.Read(bs)
  if err != nil {
    log.Println("query recv len failed, err:", err)
    c.conn.Close()
    c.conn=nil
    return nil, err
  }

  read_len = binary.LittleEndian.Uint32(bs)

  bs = make([]byte, read_len)
  _, err = c.conn.Read(bs)
  if err != nil {
    log.Println("query recv failed, err:", err)
    c.conn.Close()
    c.conn=nil
      return nil, err
  }

  err = proto.Unmarshal(bs, &response)
  if err != nil{
    log.Fatalln("UnMashal data error:", err)
  }
  return &response, nil
}


