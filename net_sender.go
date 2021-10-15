package gnet

import (
	"fmt"
	"net"
)

const MAX_BUFFSIZE = 1500

func send(to *net.Addr, buf []byte, buf_len int) (int, error) {
	total_send := 0
	leftsize := buf_len
	for leftsize >= MAX_BUFFSIZE {
		n, err := server.WriteTo(buf[total_send:total_send+MAX_BUFFSIZE], (*to))
		if err != nil {
			fmt.Println("Send error " + (*to).String() + " : " + err.Error())
			return n, err
		}
		total_send += MAX_BUFFSIZE
		leftsize -= MAX_BUFFSIZE
	}
	if leftsize > 0 {
		n, err := server.WriteTo(buf[total_send:total_send+leftsize], (*to))
		if err != nil {
			fmt.Println("Send error " + (*to).String() + " : " + err.Error())
			return n, err
		}
		total_send += n
		leftsize -= n
	}
	return total_send, nil
}

func Broadcast(buf []byte, buf_len int) {
	peers := *GetPeers()
	for name, peer := range peers {
		// Test
		_, err := server.WriteTo(buf[:buf_len], *peer) //
		// _, err := Send(player.Addr, buf[:buf_len], buf_len)
		if err != nil {
			LeavePeer(name)
		}
	}
}

func Unicast(name string, buf []byte, buf_len int) (int, error) {
	peer := GetPeer(name)
	s, err := send(peer, buf[:buf_len], buf_len)
	if err != nil {
		LeavePeer(name)
		return s, err
	}
	return s, nil
}

// func RunUDPServer(server_ip string, server_port int) net.Addr {
// 	serv_addr := server_ip + ":" + strconv.Itoa(server_port)
// 	var err error
// 	server, err = net.ListenPacket("udp", serv_addr)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("server address: ", server.LocalAddr().String())

// 	go ExecLockstep()

// 	// 200ms 마다 오브젝트 동기화
// 	//go SyncAllObjects(200)
// 	go SyncAllObjects(200)

// 	go func() {
// 		for {
// 			buf := make([]byte, MAX_BUFFSIZE)
// 			n, clientAddress, err := server.ReadFrom(buf)
// 			if (n == 0) || (err != nil) {
// 				fmt.Println("buffer size is 0...")
// 			}
// 			handleCommand(buf, n, clientAddress)
// 		}
// 	}()
// 	return server.LocalAddr()
// }
