package eth_client


import (
	//"bytes"
	"context"
  "time"
	"math/big"
	"sync"
  "log"
  "strconv"
  //  "encoding/hex"


  "github.com/xdb/go-xdb-sdk/proto"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PollblkTransactionConfirmer struct {
	client     *ethclient.Client
	ctx       context.Context
  data map[uint64]*resdb.Transaction
	lock     sync.Mutex
  min_v uint64
}

const chainId = 4

func newTransaction(from string, to string, seq uint64, amount uint64)(*resdb.Transaction){
  return &resdb.Transaction{
    From: from,
    To:to,
    Uid:seq,
    Amount:amount,
  }
}

func NewPollblkTransactionConfirmer(endpoint string) *PollblkTransactionConfirmer {
	var this PollblkTransactionConfirmer

	this.client, _ = ethclient.Dial(endpoint)
  log.Print("endpoint:",endpoint)
  this.data = make(map[uint64]*resdb.Transaction)
  this.min_v = 0
	this.ctx = context.Background()

	go this.run()

	return &this
}

func (this *PollblkTransactionConfirmer) GetData(seq uint64)(tx *resdb.Transaction) {
	this.lock.Lock()
  _,ok := this.data[seq]
  if (ok) {
    tx = this.data[seq]
  }
	this.lock.Unlock()
  return
}

func GetTransactionMessage(tx *types.Transaction) types.Message {
   msg, err := tx.AsMessage(types.LatestSignerForChainID(tx.ChainId()), nil)
   if err != nil {
    log.Fatal(err)
   }
   return msg
}

func (this *PollblkTransactionConfirmer) processBlock(number *big.Int) error {
	var stxs []*types.Transaction
	var stx *types.Transaction
	var block *types.Block
	//var hashes []string
	var err error
	//var i int
  var num_int int
  var seq uint64
  var start uint64

  log.Print("number:",number)
  num_int, _ = strconv.Atoi(number.String())
  if(this.min_v==0){
    start = uint64(1)
  } else {
    start = uint64(num_int)
  }

  for i:=start; i <= uint64(num_int); i++{
    a := big.NewInt(int64(i))
    block, err = this.client.BlockByNumber(this.ctx, a)
    if err != nil {
      return err
    }

    if (this.min_v == 0) {
       this.min_v = i
    }

    stxs = block.Transactions()
    if(len(stxs) == 0){
      seq = i - this.min_v+1
      log.Print("txn:",len(stxs), a, seq)
      this.data[seq] = newTransaction("", "", i, 1)
      continue
    }

    this.lock.Lock()
    for _, stx = range stxs {
      seq = i - this.min_v+1
      log.Print("txn:",len(stxs), a, seq)
      this.data[seq] = newTransaction(GetTransactionMessage(stx).From().Hex(), stx.To().Hex(), i, 1)
      break
    }
    this.lock.Unlock()
  }

	return nil
}

func (this *PollblkTransactionConfirmer) run() {
	var subcription ethereum.Subscription
	var events chan *types.Header
	var event *types.Header
	var err error

  eloop: for {
    this.ctx = context.Background()
	  events = make(chan *types.Header)
    subcription, err = this.client.SubscribeNewHead(this.ctx, events)
    if err != nil {
      log.Print("wait")
      time.Sleep(time.Second)
      continue
    }
    break eloop
  }

	loop: for {
		select {
		case event = <- events:
			err = this.processBlock(event.Number)
			if err != nil {
				break loop
			}
		case err = <- subcription.Err():
			break loop
		case <- this.ctx.Done():
			err = this.ctx.Err()
			break loop
		}
	}

	subcription.Unsubscribe()

	close(events)
}

