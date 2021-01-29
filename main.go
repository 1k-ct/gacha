package main

import (
	"fmt"

	"github.com/1k-ct/gacha/dogacha"
)

func main() {
	cou := 0
	var hit bool
	for !hit {
		hit = dogacha.Gacha(100, 1) // 1%のガチャ
		cou++
	}
	fmt.Println(cou, "回目で当たりました。")
}
