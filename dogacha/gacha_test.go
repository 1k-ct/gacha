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

func TestPercentages(t *testing.T) {
	type args struct {
		all int64
		n   int64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "1%",
			args: args{all: 100, n: 1},
			want: 0.01,
		},
		{
			name: "0.1%",
			args: args{all: 1000, n: 1},
			want: 0.001,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Percentages(tt.args.all, tt.args.n); got != tt.want {
				t.Errorf("Percentages() = %v, want %v", got, tt.want)
			}
		})
	}
}
