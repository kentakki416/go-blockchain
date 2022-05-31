// ../utils/neighbor.goが正しく動作するかを確認する
package main

import (
	"fmt"
	"go-blockchain/utils"
)

func main() {
	fmt.Println(utils.FindNeighbors("127.0.0.1", 5001, 0, 3, 5001, 5004))
}
