package blockchain_sample

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
}

func calculateHash(block Block) string {
	record := string(block.Index) + block.Timestamp + block.Data + block.PrevHash + string(block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func generateBlock(oldBlock Block, data string) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Nonce = 0
	newBlock.Hash = calculateHash(newBlock)
	for !isHashValid(newBlock.Hash) {
		newBlock.Nonce++
		newBlock.Hash = calculateHash(newBlock)
	}
	return newBlock
}

func isHashValid(hash string) bool {
	return hash[:4] == "0000"
}

func replaceChain(newBlocks []Block) {
	if len(newBlocks) > len(Blockchain) {
		Blockchain = newBlocks
	}
}

var Blockchain []Block

func main() {
	t := time.Now()
	genesisBlock := Block{0, t.String(), "Genesis Block", "", "", 0}
	genesisBlock.Hash = calculateHash(genesisBlock)
	Blockchain = append(Blockchain, genesisBlock)

	fmt.Println("Blockchain:", Blockchain)

	block1 := generateBlock(Blockchain[len(Blockchain)-1], "Transaction 1")
	Blockchain = append(Blockchain, block1)
	fmt.Println("Blockchain:", Blockchain)

	block2 := generateBlock(Blockchain[len(Blockchain)-1], "Transaction 2")
	Blockchain = append(Blockchain, block2)
	fmt.Println("Blockchain:", Blockchain)

	replaceChain(Blockchain)
}
