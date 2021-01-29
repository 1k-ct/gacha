package main

import (
	"fmt"
	"log"

	"github.com/1k-ct/gacha/dogacha"
)

func main() {
	cou := 0
	var err error
	var hit bool
	for !hit {
		hit, err = dogacha.Gacha(100, 1) // 1%のガチャ
		if err != nil {
			log.Fatal(err)
		}
		cou++
	}
	fmt.Println(cou, "回目で当たりました。")
}
