package percentage

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/1k-ct/gacha/dogacha"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func TestArbitraryProbability(t *testing.T) {
	n := 10000
	cou := 0
	for i := 0; i < n; i++ {
		b, err := arbitraryProbability(0.2)
		if err != nil {
			return
		}
		if b {
			cou++
		}
	}
	log.Println((float64(cou) / float64(n)) * 100)
}
func TestPlace(t *testing.T) {
	var n float64 = 0.03
	// fmt.Println(float64(n) / 10)
	cou := 0
	for n != 0 {
		n /= 10
		cou++
	}
	fmt.Println(cou)
}
func TestLog(t *testing.T) {
	var n float64 = 0.03
	p := math.Floor(math.Log10(n) + 1)
	fmt.Println(p)
	// fmt.Println(int64(math.Pow10(int(p + 1))))
	// fmt.Println(rand.Int63n(2))
}

// func TestArbitraryProbability2(t *testing.T) {
// 	n := 100000
// 	cou := 0
// 	for i := 0; i < n; i++ {
// 		b, err := arbitraryProbabilityFromString(0.03)
// 		if err != nil {
// 			return
// 		}
// 		if b {
// 			cou++
// 		}
// 	}
// 	log.Println((float64(cou) / float64(n) * 100))
// }
func abilPro(n float64, par float64) bool {
	rand.Seed(time.Now().UnixNano())
	p := rand.Int63n(int64(n))
	if 0 <= p && float64(p) <= par-1 {
		return true
	}
	return false
}
func TestSome(t *testing.T) {
	// rand.Seed(time.Now().UnixNano())
	// p := rand.Int63n(int64(100))
	// fmt.Println(p)
	n := 1000000
	cou := 0
	for i := 0; i < n; i++ {
		b := abilPro(100, 1)
		if b {
			cou++
		}
	}
	fmt.Println(cou)
	fmt.Println(float64(cou) / float64(n))
}
func TestArbitraryProbabilityTrim(t *testing.T) {
	count := 1000000
	point := 0
	var all int64 = 100
	var n int64 = 1
	for i := 0; i < count; i++ {
		b, _ := dogacha.Gacha(all, n)
		if b {
			point++
		}
	}
	probability := float64(point) / float64(count)   // 計算した確率
	expectProbability := dogacha.Percentages(all, n) //理想確率
	//  errRate計算した確率誤差
	var errRate float64 = math.Abs(expectProbability - probability)
	fmt.Println(expectProbability, probability, errRate)
	// fmt.Println()
}
func percentageCheck(count int, all, n int64) (float64, float64, float64) {
	point := 0
	for i := 0; i < count; i++ {
		b, _ := dogacha.Gacha(all, n)
		if b {
			point++
		}
	}
	expectProbability := dogacha.Percentages(all, n) //理想確率
	probability := float64(point) / float64(count)   // 計算した確率
	//  errRate計算した確率誤差
	var errRate float64 = math.Abs(expectProbability - probability)
	return expectProbability, probability, errRate
}
func TestAllIn(t *testing.T) {
	var cou float64 = 100000000
	var allErrRate float64
	// for i := 0; i < int(cou); i++ {
	// 	expectProbability, probability, errRate := percentageCheck(10000, 100, 2)
	// 	fmt.Println(expectProbability, probability, errRate)
	// 	allErrRate += errRate
	// }
	// fmt.Println(allErrRate / cou)

	finished := make(chan bool)

	funcs := []func(){
		func() {
			for i := 0; i < int(cou/4); i++ {
				_, _, errRate := percentageCheck(100000, 100, 2)
				allErrRate += errRate
				finished <- true
			}
		},
		func() {
			for i := 0; i < int(cou/4); i++ {
				_, _, errRate := percentageCheck(100000, 100, 2)
				allErrRate += errRate
				finished <- true
			}
		},
		func() {
			for i := 0; i < int(cou/4); i++ {
				_, _, errRate := percentageCheck(100000, 100, 2)
				allErrRate += errRate
				finished <- true
			}
		},
		func() {
			for i := 0; i < int(cou/4); i++ {
				_, _, errRate := percentageCheck(100000, 100, 2)
				allErrRate += errRate
				finished <- true
			}
		},
	}

	for _, sleep := range funcs {
		go sleep()
	}

	for i := 0; i < len(funcs); i++ {
		<-finished
	}
	fmt.Println(allErrRate / cou)
}
func TestGacha(t *testing.T) {
	var cou float64
	var n float64 = 10000

	for i := 0; i < int(n); i++ {

		var p bool
		for !p {
			p, _ = dogacha.Gacha(100, 99)
			cou++
		}
	}
	fmt.Println(cou / n)

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "日本語"
	p.X.Label.Text = "japan"
	p.Y.Label.Text = "Y"

	p.Add(plotter.NewGrid())

	pts := make(plotter.XYs, 3)
	for i := range pts {
		pts[i].X = 99
		pts[i].Y = cou / n
	}
	// pts := make(plotter.XYs, 3)
	pts[2].X = 98
	pts[2].Y = cou / n

	// pts := plotter.XY{X: 99, Y: cou / n}
	plot, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	// plot.GlyphStyle.Color = color.RGBA{R: 119, G: 200, B: 39}
	p.Add(plot)
	// p.Legend.Add("p", plot)
	// p.X.Min = 0
	// p.X.Max = 100
	// p.Y.Min = 0
	// p.Y.Max = 100
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot2.png"); err != nil {
		log.Fatal(err)
	}
}
func TestCheckPro(t *testing.T) {
	var cou float64
	var n float64 = 10000

	for i := 0; i < int(n); i++ {
		var p bool
		for !p {
			p, _ = dogacha.Gacha(100, 1)
			cou++
		}
	}
	fmt.Println(cou / n)
}
