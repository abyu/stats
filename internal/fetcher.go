package internal

import (
	log "github.com/sirupsen/logrus"
	"time"
)

//FetchTask ...
type FetchTask struct {
	Store MetricsStore
	Frequency time.Duration
}

//NewFetchTask ...
func NewFetchTask(store MetricsStore, frequency time.Duration) FetchTask {
	return FetchTask{
		Store:     store,
		Frequency: frequency,
	}
}

func (f *FetchTask) fetchBalance() {
	balance := GetTokenBalance("0x922769620ecbf5805733abeb44825523609646fb", "0x8b3192f5eebd8579568a2ed41e6feb402f93f73f")
	f.Store.New(TokenBalance{Balance: balance, TokenName: "SaitamaInu", TimeStamp: time.Now()})
}

//RunBlocking ...
func (f *FetchTask) RunBlocking() {
	log.Info("Fetching data...")
	for {
		f.fetchBalance()
		time.Sleep(f.Frequency)
	}
}
