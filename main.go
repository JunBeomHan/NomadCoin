package main

import (
	"NomadCoin/blockchain"
	
)


func main() {
	chain := blockchain.GetBlockchain()
}


/*
#4.4
sync 패키지에 대해서 알아보자.
- sync 패키지는 우리가 동기적(어떤 일이나 행동을 일으키게 하는 계기가 되는. 또는 그런 것.)으로 처리해야 하는 부분을 제대로 처리하게 도와준다.
- 예를 들어 여러개의 쓰레드와 고루틴이 돌아가고 있는데, 이 코드는 단 한번만 실행시키고 싶다면 sync 패키지를 사용한다.
 
그리고 Sync 패키지의 Once 녀석을 사용할 것이다.


*/