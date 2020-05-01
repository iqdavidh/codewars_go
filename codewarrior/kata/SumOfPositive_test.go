package kata

import "testing"

func TestSumOfPositive(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"c1", args{[]int{1, 2, 3, 4, 5}}, 15},
		{"c2", args{[]int{1, -2, 3, 4, 5}}, 13},
		{"c3", args{[]int{}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumOfPositive(tt.args.numbers); got != tt.want {
				t.Errorf("SumOfPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
