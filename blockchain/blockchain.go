package blockchain

import (
	"crypto/sha256"
	"fmt"
	"sync"
)

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
	// blocks는 pointer의 slice가 된다.
	blocks []*block

	// +-----------------------+
	// |   포인터(Pointer)       |  ->  [0x100] (첫 번째 요소의 주소) (이 배열은 *block이다.)
	// +-----------------------+
	// |         길이(Length)   |  ->  3 (저장된 요소의 개수)
	// +-----------------------+
	// |         용량(Capacity) |  ->  4 (참조하는 배열의 크기)
	// +-----------------------+

}

/*
	여기서 blocks []block을 blocks []*block으로 변경한 이유는
	blocks는 block의 슬라이스 이기 때문에 슬라이스에 append할 때
*/

var b *blockchain // blockchain 패키지 내부에서만 이 변수에 접근가능하다.
/*
Singleton의 의미는, 이 변수의 instance를 직접 공유하지 않고, 그 대신, 이 변수의 instance를 우릴 대신해서 드러내주는 Function을 생성하는 것이다.
function을 생성하는 만큼, 다른 패키지에서 우리의 blockchain이 어떻게 드러날 지를 제어할 수 있다는 의미이기도 하다.
*/
var once sync.Once

// Once는  Do라는 function을 가지고 있다.
// Do 함수는 확실히 단 한 번만 호출되도록 해주는 함수이다.=>  단 한 번만 호출되도록 해주는 함수이다. (고루틴을 이용한 동시 자원 참조 방지)
// once는 Do라는 function을 가지고 잇다.
/*
type Once struct {
	// done indicates whether the action has been performed.
	// It is first in the struct because it is used in the hot path.
	// The hot path is inlined at every call site.
	// Placing done first allows more compact instructions on some architectures (amd64/386),
	// and fewer instructions (to calculate offset) on other architectures.
	done uint32
	m    Mutex
}
*/

// 좀더 모듈화된 방법으로 구현한다.
func (b *block) calculateHash() {
	hash := sha256.Sum256([]byte(b.data + b.prevHash))
	b.hash = fmt.Sprintf("%x", hash)
}

func getLastHash() string {
	totlaBlocks := len(GetBlockchain().blocks) // blockchain의 전체 길이
	if totlaBlocks == 0 {
		return ""
	}
	return GetBlockchain().blocks[totlaBlocks-1].prevHash
}

func createBlock(data string) *block {
	newBlock := block{data, "", getLastHash()}
	newBlock.calculateHash()
	return &newBlock
}

func GetBlockchain() *blockchain {
	if b == nil { // nil인지 아닌지 판단하여, blockchain인의 초기화를 한번만 하게 할 수 있다.(왜냐하면 닐일 경우 , 대입을 해주기때문에 더 이상 닐이 아니기 때문이다.)
		once.Do(func() {
			b = &blockchain{}
			b.blocks = append(b.blocks, createBlock("Genesis Block"))
		})
	}
	// 여기서 질문 하나, 그러면 once로 처리하면 무슨 의미가 있나?
	// 여기서 중요한 점은 지금은 고 루틴을 사용하지 않지만 향후 고 루틴을 이용해 코딩을 하게 된다.
	// 여기서 중요한 점은 고루틴을 사용하게 되면 한 자원에 동시에 접근할 수 있게 한다.
	// 근데 do 함수가 없으면 함수가 두 번 실행되면 싱글톤 패턴의 정의에 벗어나니깐, do 함수를 사용해서 미연에 방지한다.
	return b // 이 시점에세 blockchain이 이미 초기화 된 걸 알고 있음
	/*
		그리고 이것이 싱글톤 패턴이다. Single Singly Singleton 패턴이다.
	*/
}

/*

싱글톤 패턴 좀더 찾아본 내용
정의 : 객체의 인스턴스가 오직 1개만 생성되는 팬턴을 의미한다.
싱글톤 패턴을 구현하는 방법은 여러가지 있지만, 그 중 객체를 미리 생성하고 가져오는 가장 단순한 방법으로 설명하겠다고 함.

그러면 인스턴스를 오직 한 개로만 가져오면 어떤 이점이 있을까?
- 메모리 측면에서 너무 좋다!
	최초 한번의 인스턴스가 생성되므로, 인스턴스에 접근할 때 메모리 낭비를 방지할 수 있다.
	뿐만 아니라 이미 생성된 인스턴스를 활용하니 속도 측면에서 엄청난 이점이 있다.

- 그리고 blockchain을 다른 패키지에서 사용하기 용의하다.

- 초기화 작업을 언제든지 변경할 수 있다.

- 그리고 전역 변수로 생성했기 때문에, 한개임을 보증할 수 있다.


*/

/*

동기 비동기

동기 : 함수를 호출하면 함수가 다 실행될때까지 기다리는 것을 동기라고 한다.
동기 같은 경우는, 함수의 반환 값이 어떤 값으로 사용될때 필요한다.

비동기 : 동기와 반대로, 함수를 호출하면 메인은 기다리지 않고 코드가 실행되는 것을 뜻한다.
비동기 같은 경우는, 함수를 호출하면 기다리지 않고 가게 된다고 했는데, 그럼 만약 호출된 함수가 리턴 한다면 어떻게 처리할까?
=> 그래서 호출한 함수의 값을 반환해주는 콜백 함수가 존재한다. -> 그래서 콜백 함수로 들어오게 해준다.
그리고 메인에서는 이 콜백함수로 호출을 하여 마무리 한다.

Nodejs의 통신은 비동기 식이고, 안드로이드 통신도 비동기 이다.
*/
