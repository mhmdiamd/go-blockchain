package main

import (
	"time"

	"github.com/mhmdiamd/go-blockchain/network"
)

func main() {
  trLocal := network.NewLocalTransportImpl("LOCAL")
  trRemote := network.NewLocalTransportImpl("REMOTE")

  trLocal.Connect(trRemote)
  trRemote.Connect(trLocal)

  go func() {
    for {
      trRemote.SendMessage(trLocal.Addr(), []byte("Hello World"))
      time.Sleep(1 * time.Second)
    }
  }()

  opts := network.ServerOpts{
    Transports : []network.Transport{trLocal},
  }

  s := network.NewServer(opts)
  s.Start()
}
