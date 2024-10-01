package main

import (
	"bufio"
	"crypto/sha256"
	"fmt"
	"os"
)

func main() {
	// コンソールから入力を取得
	fmt.Println("Enter your text")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// ハッシュ値を計算
	hash := sha256.Sum256([]byte(input))

	// ハッシュ値を表示
	fmt.Printf("Hash: %x\n", hash)
}
