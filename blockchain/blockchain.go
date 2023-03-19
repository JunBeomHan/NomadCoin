package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

type block struct {
	Data     string
	Hash     string
	PrevHash string
}

type blockchain struct {
	blocks []*block
}

var b *blockchain
var once sync.Once

func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.Data + b.PrevHash))
	b.Hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totlaBlocks := len(GetBlockchain().blocks)
	if totlaBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totlaBlocks-1].Hash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

// Export
func (b *blockchain) AddBlock(data string) {
	b.blocks = append(b.blocks, createBlock(data))
}

func GetBlockchain() *blockchain {
	if b == nil {
		once.Do(func() {
			b = &blockchain{}
			b.AddBlock("Genesis")
		})
	}
	return b

}

// Main에서 부르는 형식
//
//	func AllBlocks() []*block {
//		return GetBlockchain().blocks
//	}
//
// Main에서 GetBlockchain을 실행 후 사용할 함수
func (b *blockchain) AllBlocks() []*block {
	return b.blocks
}
