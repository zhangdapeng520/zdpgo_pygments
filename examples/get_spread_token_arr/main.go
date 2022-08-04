package main

import (
	"fmt"

	"github.com/zhangdapeng520/zdpgo_pygments"
)

func main() {
	tokens := []string{"defF(N,N):", "returnN-N", "defF(N,N):", "returnN+N"}
	newTokens := zdpgo_pygments.GetSpreadTokenArr(tokens, 2)
	fmt.Println(newTokens)
}
