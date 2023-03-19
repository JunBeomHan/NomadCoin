package blockchain


// blockchain을 공유하고 초기화하는 부분을 구현할 것이다.
// 우리가 사용할 것은 singleton 패턴이다.
// 싱글톤패턴이란, 우리의 application 내에서 언제든지 blockchain의 단 하나의 instance만을 공유하는 방법을 말한다.

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

