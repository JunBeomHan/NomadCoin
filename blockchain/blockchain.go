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

var b *blockchain 	// blockchain 패키지 내부에서만 이 변수에 접근가능하다.
/*
	Singleton의 의미는, 이 변수의 instance를 직접 공유하지 않고, 그 대신, 이 변수의 instance를 우릴 대신해서 드러내주는 Function을 생성하는 것이다.
	function을 생성하는 만큼, 다른 패키지에서 우리의 blockchain이 어떻게 드러날 지를 제어할 수 있다는 의미이기도 하다.
*/

func GetBlockchain()*blockchain {
	if b == nil {	// nil인지 아닌지 판단하여, blockchain인의 초기화를 한번만 하게 할 수 있다.(왜냐하면 닐일 경우 , 대입을 해주기때문에 더 이상 닐이 아니기 때문이다.)
		b = &blockchain{}
		// blockchain이 어떻게 생성될 지 제어할 수 있다. 
		// 만약 초기화 단계에서 blockchain을 database에서 가져올 수 있다는 의미이다!(와 그래서 사용하는 거구나)
		 
	}
	return b	// 이 시점에세 blockchain이 이미 초기화 된 걸 알고 있음
}

