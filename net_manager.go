package gnet

import (
	"net"
	"sync"
)

var manager_mtx sync.Mutex

var server net.PacketConn
var peers map[string]*net.Addr = make(map[string]*net.Addr) // name, address

func GetPeers() *map[string]*net.Addr {
	return &peers
}
func GetPeer(name string) *net.Addr {
	manager_mtx.Lock()
	p, e := peers[name]
	if e == false {
		manager_mtx.Unlock()
		return nil
	}
	manager_mtx.Unlock()
	return p
}
func EnterPeer(name string, addr *net.Addr) {
	manager_mtx.Lock()
	peers[name] = addr
	manager_mtx.Unlock()
}
func LeavePeer(name string) {
	manager_mtx.Lock()
	delete(peers, name)
	manager_mtx.Unlock()
}
