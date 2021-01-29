# gacha
ガチャガチャ  

```
$ go get github.com/1k-ct/gacha/dogacha
```

```go
func main() {
    cou := 0
	var hit bool
	for !hit {
		hit = dogacha.Gacha(100, 1) // 1%のガチャ
		cou++
	}
	fmt.Println(cou, "回目で当たりました。")
}
```

ソシャゲと同じようなガチャが出来ます。  
Gacha(100, 1)は、100分の1のように指定します。(1%)  
0.1%のときは、Gacha(1000, 1)とします。  
2%は、Gacha(100, 2)とします。  
