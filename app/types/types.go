package server

import (
	"net"
	"sync"
	"time"
)

type ServerConfig struct {
	Port       string
	Role       string
	Replid     string
	ReplOffset int
	MasterHost string
	MasterPort string
}

type KVV struct {
	Value    string
	ExpireAt time.Time
}

type ServerState struct {
	Config  *ServerConfig
	Store   map[string]KVV
	StoreMu sync.RWMutex

	ReplicationOffset int
	LastWaitOffset    int

	Replicas        []*ClientState
	ReplicaMu       sync.RWMutex
	ReplicaAckBytes map[int]int
	PropagationChan chan []string
}

type ClientState struct {
	Server           *ServerState
	ConnectionId     net.Conn
	Id               int
	InTransaction    bool
	TransactionQueue [][]string
}

func (s *ServerState) AddOffsetToServer(offset int) {
	s.Config.ReplOffset += offset
}
