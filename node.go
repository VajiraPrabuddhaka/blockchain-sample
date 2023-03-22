package main

import (
	"encoding/json"
	"log"
	"net"
	"time"
)

type Node struct {
	Address string
	// Blockchain *Blockchain
	Peers []string
}

type Message struct {
	Description string
	Timestamp time.Time
}

func (n *Node) listen() {
	listener, err := net.Listen("tcp", n.Address)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go n.handleConnection(conn)
	}
}

func (n *Node) handleConnection(conn net.Conn) {
	defer conn.Close()

	// Read incoming message
	var msg Message
	err := json.NewDecoder(conn).Decode(&msg)
	if err != nil {
		log.Println(err)
		return
	}

	// Update blockchain if message is valid
	if msg.isValid() {
		n.Blockchain.addBlock(msg.Data)
	}

	// Forward message to other nodes in the network
	for _, peer := range n.Peers {
		if peer != n.Address {
			go n.forwardMessage(msg, peer)
		}
	}
}

func (n *Node) forwardMessage(msg Message, peer string) {

}
