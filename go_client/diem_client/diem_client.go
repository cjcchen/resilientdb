package ndiem

import (
  "log"
  "sync"

  "github.com/resilientdb/go-resilientdb-sdk/proto"
	"github.com/diem/client-sdk-go/diemclient"
	"github.com/diem/client-sdk-go/diemjsonrpctypes"
)

type PollblkTransactionConfirmer struct {
	client   diemclient.Client
  data map[uint64]*resdb.Transaction
	lock     sync.Mutex
  min_v uint64
}

const chainId = 4

func newTransaction(from string, to string, seq uint64, version uint64, amount uint64)(*resdb.Transaction){
  return &resdb.Transaction{
    From: from,
    To:to,
    Uid:seq,
    Amount:amount,
    Version: version,
  }
}

func NewPollblkTransactionConfirmer(endpoint string) *PollblkTransactionConfirmer {
	var this PollblkTransactionConfirmer

	this.client = diemclient.New(chainId, endpoint)
  log.Print("chainid:",chainId)
  log.Print("endpoint:",endpoint)
  this.data = make(map[uint64]*resdb.Transaction)
  this.min_v = 0

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

func (this *PollblkTransactionConfirmer) parseTransaction(tx *diemjsonrpctypes.Transaction) bool {
	var sender string
  var receiver string
  var amount uint64
  var seq uint64
  var version uint64

  sender = tx.Transaction.Sender
  receiver = tx.Transaction.Script.Receiver
  amount = tx.Transaction.Script.Amount
  seq = tx.Transaction.SequenceNumber
  version = tx.Version
  if (len(receiver) == 0) {
    if(this.min_v>0){
      log.Print("get skip version:",version)
    }
    return false
  }
  if(sender == "000000000000000000000000000000dd"){
    if(this.min_v>0){
      log.Print("get skip version:",version)
    }
    return false
  }
  //log.Print("get txn:",sender)
  //log.Print("get receiver:",receiver)
  //log.Print("get amount:", amount)
  //log.Print("get seq:", seq)
  //log.Print("push version:",version)
  //log.Print("min v:",this.min_v)
  //log.Print("get txn:",tx)
  //log.Print("event:",tx.Events)
  //log.Print("metadata:",tx.Transaction.Script.Type)
	this.lock.Lock()
  this.data[version-this.min_v] = newTransaction(sender, receiver, version, seq,amount)
	this.lock.Unlock()
  return true
}

func (this *PollblkTransactionConfirmer) run() {
	var txs []*diemjsonrpctypes.Transaction
	var tx *diemjsonrpctypes.Transaction
	var meta *diemjsonrpctypes.Metadata
	var v, version uint64
	var err error
  var ok bool

  log.Print("run diem")

	meta, err = this.client.GetMetadata()
	if err != nil {
		log.Print("get meta: %s", err.Error())
		return
	}

	v = meta.Version

	for {
		meta, err = this.client.GetMetadata()
		if err != nil {
			log.Print("get meta: %s", err.Error())
			return
		}

		version = meta.Version

		for v < version {
			v += 1

			txs, err = this.client.GetTransactions(v, 100, true)
			if err != nil {
				continue
			}

			for _, tx = range txs {
				if tx.Version > v {
					v = tx.Version
				}

				if tx.Transaction.Type != "user" {
					continue
				}

        //log.Print("get version:",tx.Version)
				ok = this.parseTransaction(tx)
        if (ok){
          if ( this.min_v == 0){
            this.min_v = tx.Version
          }
        }
			}
		}
	}
}
