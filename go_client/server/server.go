package main

import (
    "fmt"
    "net"
    "log"
    "encoding/binary"
    "sync"

    "github.com/resilientdb/go-resilientdb-sdk/proto"
    "github.com/resilientdb/go-resilientdb-sdk/diem_client"
    //"github.com/resilientdb/go-resilientdb-sdk/algorand_client"
    //"github.com/resilientdb/go-resilientdb-sdk/eth_client"
    "github.com/golang/protobuf/proto"
)

type Service struct {
  result_list map[uint64]resdb.BlockMiningInfo
  //confirmer *algorand_client.PollblkTransactionConfirmer
  //confirmer1 *algorand_client.PollblkTransactionConfirmer
  //confirmer2 *algorand_client.PollblkTransactionConfirmer
  confirmer *ndiem.PollblkTransactionConfirmer
  confirmer1 *ndiem.PollblkTransactionConfirmer
  confirmer2 *ndiem.PollblkTransactionConfirmer
  confirmer3 *ndiem.PollblkTransactionConfirmer
  //confirmer *eth_client.PollblkTransactionConfirmer
	lock     sync.Mutex
}

func MakeService() *Service{
  return &Service{
    result_list: make(map[uint64]resdb.BlockMiningInfo),
    confirmer: ndiem.NewPollblkTransactionConfirmer("http://127.0.0.1:9000"),
    confirmer1: ndiem.NewPollblkTransactionConfirmer("http://127.0.0.1:9001"),
    confirmer2: ndiem.NewPollblkTransactionConfirmer("http://127.0.0.1:9002"),
    confirmer3: ndiem.NewPollblkTransactionConfirmer("http://127.0.0.1:9003"),
    //confirmer: algorand_client.NewPollblkTransactionConfirmer("http://127.0.0.1:9001"),
    //confirmer1: algorand_client.NewPollblkTransactionConfirmer("http://127.0.0.1:9002"),
    //confirmer2: algorand_client.NewPollblkTransactionConfirmer("http://127.0.0.1:9003"),
    //confirmer: eth_client.NewPollblkTransactionConfirmer("ws://127.0.0.1:9001"),
  }
}

func (s* Service) Process(buf []byte) ([]byte, error) {
  var resdb_message resdb.ResDBMessage
  var request resdb.Request
  var resp resdb.ResDBMessage
  var data []byte
  var err error

  err = proto.Unmarshal(buf, &resdb_message)
  if err != nil{
    log.Fatalln("UnMashal data error:", err)
    return nil, err
  }

  err = proto.Unmarshal(resdb_message.Data, &request)
  if err != nil{
    log.Fatalln("UnMashal data error:", err)
    return nil, err
  }

  data, err = s.Dispatch(request.Data, request.Type)
  if (err !=nil) {
    return nil, err
  }
  if(data == nil){
    return nil,nil
  }
  resp.Data = data
  return proto.Marshal(&resp)
}

func (s* Service) Dispatch(buf []byte, request_type int32)([]byte, error){
  if(request_type == 18) {
    var query resdb.TxnQueryRequest
    var response resdb.CustomQueryResponse
    var resp []byte
    var err error

    err = proto.Unmarshal(buf, &query)
    if err != nil{
      log.Fatalln("UnMashal request error:", err)
        return nil, err
    }

    log.Print("get min: max result:",query)
    if(query.IsQueryResults == 1){
      resp = s.GetResult(query.MinSeq)
      if (resp != nil ){
        response.RespStr = resp
      }
      resp, _= proto.Marshal(&response)
    } else {
      resp = s.GetTransaction(query.MinSeq, query.MaxSeq)
      response.RespStr = resp
      resp, _= proto.Marshal(&response)
    }
    return resp,nil
  } else if(request_type == 19){
    s.SaveResult(buf)
    return nil,nil
  }
  return nil, nil
}

func MakeRequest(buf []byte) (*resdb.Request){
  return &resdb.Request {
    Data:buf,
  }
}

func MakeClientRequest(buf []byte)(*resdb.BatchClientRequest_ClientRequest){
  return &resdb.BatchClientRequest_ClientRequest{
    Request: MakeRequest(buf),
  }
}

func (s*Service) SaveResult(buf []byte){
  var result resdb.BlockMiningInfo
  var err error

  err = proto.Unmarshal(buf, &result)
  if err != nil{
    log.Fatalln("UnMashal request error:", err)
    return
  }
	s.lock.Lock()
  s.result_list[result.Header.MinSeq]=result
	s.lock.Unlock()
}

func (s*Service) GetTransaction(min_seq uint64, max_seq uint64) (buf []byte){
  var data []byte
  var num int
  var resp resdb.TxnQueryResponse


  num = int(max_seq - min_seq)
  resp.Data = make([][]byte, num)
  resp.Seq = make([]uint64, num)

  var idx int
  idx = 0

  for i := 0; i < num; i++ {
    data = s.GetClientRequest(uint64(i) + min_seq)
    if( data == nil ){
      t :=0 
      for j:=1; j<=10;j++ {
        data = s.GetClientRequest(uint64(i) + min_seq+uint64(j))
        if ( data != nil ) {
          t = t+1
        }
      }
      //log.Print("check data",uint64(i) + min_seq," num:",t)
      if t >= 8 {
        log.Print("no data:",uint64(i) + min_seq)
        continue
      }
      break;
      continue;
    }
    resp.Data[idx] = data
    resp.Seq[idx] = uint64(i)+min_seq
    idx+=1
  }
  if(idx > 0){
    resp.Data = resp.Data[0:idx]
    resp.Seq = resp.Seq[0:idx]
  } else {
    resp.Data = nil
    resp.Seq = nil
  }

  log.Print("============== get data done:", len(resp.Data))

  data, _= proto.Marshal(&resp)
  return data;
}

func (s*Service) GetResult(seq uint64) ([]byte){
  var data []byte
  var resp resdb.TxnQueryResponse
  var result resdb.BlockMiningInfo

	s.lock.Lock()
  _,ok := s.result_list[seq];
  if(ok) {
    result = s.result_list[seq]
  } else {
    s.lock.Unlock()
    return nil
  }
	s.lock.Unlock()

  resp.Data = make([][]byte, 1)
  resp.Seq = make([]uint64, 1)

  data, _ = proto.Marshal(&result)
  resp.Data[0] = data
  resp.Seq[0] = seq

  data, _= proto.Marshal(&resp)
  return data;
}

func (s*Service) GetData(seq uint64) (tx *resdb.Transaction) {
  tx = s.confirmer.GetData(seq)
  if(tx !=nil) {
    return
  }
  tx = s.confirmer1.GetData(seq)
  if(tx !=nil) {
    return
  }
  tx = s.confirmer2.GetData(seq)
  if(tx !=nil) {
    return
  }
  tx = s.confirmer3.GetData(seq)
  if(tx !=nil) {
    return
  }
  return
}

func (s*Service) GetClientRequest(seq uint64) (buf []byte){
  var data []byte
  var err error
  var client_request resdb.BatchClientRequest
  var txns resdb.TransactionsRequest
  var txn *resdb.Transaction

  txn = s.GetData(seq)
  if (txn == nil){
    return nil
  }

  client_request.ClientRequests = make([]*resdb.BatchClientRequest_ClientRequest,1)
  txns.Transactions = make([]*resdb.Transaction,1)
  txns.Transactions[0] = txn

  data, err = proto.Marshal(&txns)
  if err != nil {
    log.Fatalln("Mashal data error:", err)
    return nil
  }
  client_request.ClientRequests[0] = MakeClientRequest(data)

  data, _= proto.Marshal(&client_request)
  return data;
}


func main() {
    fmt.Println("Starting the server ...")
    listener, err := net.Listen("tcp", "0.0.0.0:50000")
    if err != nil {
        fmt.Println("Error listening", err.Error())
        return
    }
    service := MakeService()
    for {
        conn, err := listener.Accept()
        if err != nil {
            fmt.Println("Error accepting", err.Error())
            return
        }
        go Do(conn, service)
    }
}

func Do(conn net.Conn, service *Service) {
  var buf []byte
  var err error

  for {
   buf, err = ReadData(conn)
   if (err != nil) {
     //fmt.Println("Error reading", err.Error())
       return
   }

   buf, err = service.Process(buf)
   if (err != nil) {
     //fmt.Println("Error reading", err.Error())
       return
   }
   if(buf == nil) {
    return;
   }

   err = WriteData(conn, buf)
   if (err != nil) {
     //fmt.Println("Error reading", err.Error())
       return
   }
  }
}

func ReadData(conn net.Conn) ([]byte, error){
  var length int
  var err error

  length, err = ReadLen(conn)
  if err != nil {
    //fmt.Println("Error reading", err.Error())
      return nil, err
  }

  buf := make([]byte, length)
  length, err = conn.Read(buf)
  if err != nil {
    //fmt.Println("Error reading", err.Error())
      return nil, err
  }
  return buf, nil
}

func ReadLen(conn net.Conn) (int, error){
   buf := make([]byte, 8)
   _, err := conn.Read(buf)
   if err != nil {
     //fmt.Println("Error reading", err.Error())
     return 0, err
   }
   return int(binary.LittleEndian.Uint32(buf)), nil
}

func WriteLen(conn net.Conn, length int) error {
  var bs []byte
  var err error

  bs = make([]byte, 8)
  binary.LittleEndian.PutUint32(bs, uint32(length))

  _, err = conn.Write(bs)
  return err
}

func WriteData(conn net.Conn, buf []byte)  error{
  var err error

  err = WriteLen(conn, len(buf))
  if err != nil {
    return err
  }

  _, err = conn.Write(buf)
  return err
}

