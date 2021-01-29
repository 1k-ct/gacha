package dogacha

import (
	"fmt"
	"testing"
)

func TestGacha(t *testing.T) {
	type args struct {
		all               int64
		expectProbability int64
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "1/100_1%",
			args: args{all: 100, expectProbability: 1},
		},
		{
			name: "1/1000_0.1%",
			args: args{all: 1000, expectProbability: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cou := 0
			var hit bool
			for !hit {
				hit = Gacha(tt.args.all, tt.args.expectProbability)
				cou++
				if cou > 10000 {
					t.Errorf("Gacha()  %v", cou)
				}
			}
			fmt.Println(cou, "回目で当たりました。")
		})
	}
}
