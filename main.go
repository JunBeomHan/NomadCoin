package main

import (
	"crypto/sha256"
	"fmt"
)

// 우리가 지금 구현할 블럭체인은 오직 block 안에다가만 저장할 것임.
type block struct {
	data     string
	hash     string
	prevHash string // 왼래 블럭 간의 연결은 hash로 되어 있지만 지금은 전 블럭의 해쉬만 저장할 거임.
}

// 결정론적(같은 입력값은 항상 같은 출력값을 얻게 된다.)
// 단반향(해쉬값을 입력하면 본래의 값을 되찾을 수 없음) -> 이것을 one-way-function 이라고 한다.

/*
Bloack 1
	Block 1 Hash = (data + previous Hash value)
	Hash = (data + "")

Blaock 2
	Hash = data+Block1 Hash

*/

type blockchain struct {
	blocks []block
}

func (b *blockchain) getLastHash() string {
	if len(b.blocks) > 0 {
		return b.blocks[len(b.blocks)-1].hash
	}
	return ""
}

func (b *blockchain /*복사값이 아닌, 원본 값을 원함*/) addBlock(data string) {

	/*
		첫 번째 데이터와 두 번째 데이터 부터는 prevHash값이 존재한다는 차이 점이 존재한다.
	*/
	newblock := block{data, "", b.getLastHash()}
	hash := sha256.Sum256([]byte(newblock.data + newblock.prevHash))
	newblock.hash = fmt.Sprintf("%x", hash)
	b.blocks = append(b.blocks, newblock)
}

func (b *blockchain) listBlocks() {
	for _, block := range b.blocks {
		fmt.Printf("Data: %s\n", block.data)
		fmt.Printf("Hash: %s\n", block.hash)
		fmt.Printf("Prev Hash: %s\n", block.prevHash)

	}
}

func main() {
	chain := blockchain{}
	chain.addBlock("one block")
	chain.addBlock("tow block")
	chain.listBlocks()
}
