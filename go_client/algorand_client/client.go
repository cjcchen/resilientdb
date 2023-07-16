package algorand_client


import (
  "log"
  "sync"
	"encoding/binary"
	"context"

  "github.com/xdb/go-xdb-sdk/proto"
	"github.com/algorand/go-algorand-sdk/client/v2/algod"
	"github.com/algorand/go-algorand-sdk/client/v2/common/models"
)

func noteToUid(note []byte) (uint64, bool) {
	if len(note) != 8 {
		return 0, false
	}

	return binary.LittleEndian.Uint64(note), true
}

type PollblkTransactionConfirmer struct {
	client    *algod.Client
  data map[uint64]*XXXX.Transaction
  min_v uint64
	lock     sync.Mutex
}

const benchmarkToken =
	"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"

func newTransaction(seq uint64)(*XXXX.Transaction){
  return &XXXX.Transaction{
    Uid:seq,
  }
}

func NewPollblkTransactionConfirmer(endpoint string) *PollblkTransactionConfirmer {
	var this PollblkTransactionConfirmer

	this.client, _ = algod.MakeClient(endpoint, benchmarkToken)

  //log.Print("chainid:",chainId)
  log.Print("endpoint:",endpoint)
  this.data = make(map[uint64]*XXXX.Transaction)
  this.min_v = 0

	go this.run()

	return &this
}

func (this *PollblkTransactionConfirmer) GetData(seq uint64)(tx *XXXX.Transaction) {
	this.lock.Lock()
  _,ok := this.data[seq]
  if (ok) {
    tx = this.data[seq]
  }
	this.lock.Unlock()
  return
}

func (this *PollblkTransactionConfirmer) parseBlock(round uint64) {
  
	  this.lock.Lock()
    this.data[uint64(round)] = newTransaction(uint64(round))
	  this.lock.Unlock()

	return
}

func (this *PollblkTransactionConfirmer) run() {
	var client *algod.Client = this.client
	var status models.NodeStatus
	var round uint64
	var err error

  round = 0
	loop: for {
		status, err = client.StatusAfterBlock(round).Do(context.Background())
		if err != nil {
			break loop
		}

		for round < status.LastRound {
			this.parseBlock(round)
			round += 1
		}
	}
}
