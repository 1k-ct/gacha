package percentage

import (
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/1k-ct/gacha/dogacha"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func checkPro(n float64) {
	var cou float64
	// var n float64 = 10000

	for i := 0; i < int(n); i++ {
		var p bool
		for !p {
			p = dogacha.Gacha(100, 1)
			cou++
		}
	}
	fmt.Println(cou / n)
}
func floatToNaturalNum(sf float64) (float64, int, error) {
	// 0.001% = 1万分の1 (年末ジャンボ 宝くじ 4等)
	n := strconv.FormatFloat(sf, 'f', -1, 64)
	if len(n) >= 7 {
		return -1, -1, errors.New("n >= 6 number over")
	}
	lIndex := strings.LastIndex(n, ".")
	placeNum := len(n) - lIndex
	if lIndex == -1 {
		return sf, placeNum, nil
	}
	s := sf * math.Pow10(placeNum-1)
	return s, placeNum - 1, nil
}

// trimFloatingNumber converts a float64 to an int(float64).
// If the return value is -1, the argument n is len(n) >= 7 or strings.parseFloat is err.
func trimFloatingNumber(n float64) (float64, error) {
	s := strconv.FormatFloat(n, 'f', -1, 64)
	if len(s) >= 7 {
		return -1, errors.New("n >= 7 number over")
	}
	s2 := strings.Replace(s, ".", "", 1)
	f, err := strconv.ParseFloat(s2, 64)
	if err != nil {
		return -1, err
	}
	return f, nil
}

// 0.02% パーセント表示
// expectProbability := 0.02 理想確率
func arbitraryProbability(expectProbability float64) (bool, error) {
	rand.Seed(time.Now().UnixNano())
	f, n, err := floatToNaturalNum(expectProbability)
	if err != nil {
		return false, err
	}
	// log.Println(f, n)
	p := rand.Int63n(int64(math.Pow10(n - 1)))
	if 0 <= p && p <= int64(f)-1 {
		return true, nil
	}
	return false, nil
}
func arbitraryProbabilityFromString(expectProbability float64) (bool, error) {
	rand.Seed(time.Now().UnixNano())
	f, err := trimFloatingNumber(expectProbability)
	if err != nil {
		return false, err
	}
	// log.Println(f)
	placeNum := math.Floor(math.Log10(expectProbability) + 1) //桁数
	// log.Println(placeNum)
	n := rand.Int63n(int64(math.Pow10(int(placeNum + 1))))
	if 0 <= n && float64(n) <= f-1 {
		return true, nil
	}
	return false, nil

}

// func arbitraryProbabilityTrim(all, expectProbability int64) bool {
// 	rand.Seed(time.Now().UnixNano())
// 	p := rand.Int63n(all)
// 	if 0 <= p && p <= expectProbability-1 {
// 		return true
// 	}
// 	return false
// }
func showGlyph() {
	var n float64 = 1000

	p, err := plot.New()
	if err != nil {
		log.Fatal(err)
	}
	// p.Title.Text = ""
	// p.X.Label.Text = "X"
	// p.Y.Label.Text = "Y"

	// pts := make(plotter.XYs, 1000)
	p.Add(plotter.NewGrid())

	var i int64
	pts := make(plotter.XYs, 1000)
	for i = 1; i < 1000; i++ {
		var cou float64 = 0
		for v := 0; v < int(n); v++ {
			var b bool
			for !b {
				b = dogacha.Gacha(1000, i)
				cou++
			}
		}
		pts[i].X = dogacha.Percentages(1000, i)
		pts[i].Y = cou / n
	}
	plot, err := plotter.NewScatter(pts)
	if err != nil {
		log.Fatal(err)
	}
	// plot.GlyphStyle.Color = color.RGBA{R: 119, G: 200, B: 39}
	p.Add(plot)
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "plot3.png"); err != nil {
		log.Fatal(err)
	}
}
