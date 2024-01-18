package network

import (
	"fmt"
	"sync"
)

type LocalTransport struct {
  addr NetAddr
  consumeCh chan RPC
  lock sync.RWMutex
  peers map[NetAddr]*LocalTransport
}

func NewLocalTransportImpl(addr NetAddr) *LocalTransport {
  return &LocalTransport{
    addr : addr,
    consumeCh: make(chan RPC, 1024),
    peers: make(map[NetAddr]*LocalTransport),
  }
}

func (t *LocalTransport) Consume() <-chan RPC {
  return t.consumeCh
}

func (t *LocalTransport) Addr() NetAddr {
  return t.addr
}

func(t *LocalTransport) Connect(tr *LocalTransport) error {
  t.lock.Lock()
  defer t.lock.Unlock()

  t.peers[tr.Addr()] = tr
  
  return nil
}

func(t *LocalTransport) SendMessage(to NetAddr, payload []byte) error {
  t.lock.Lock()
  defer t.lock.Unlock()

  peer, ok := t.peers[to]

  if !ok {
    return fmt.Errorf("%s: could not send message to %s", t.addr, to)
  }

  peer.consumeCh <- RPC{
    From : t.addr,
    Payload : payload,
  }

  return nil
}


