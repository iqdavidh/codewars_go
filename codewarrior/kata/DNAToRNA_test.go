package kata

import (
	"strings"
	"testing"
)

type MyString string

func (s MyString) IsUpperCase() bool {
	// Your code here!

	sUpper := strings.ToUpper(string(s))
	return s == MyString(sUpper)
}

func MakeNegative(x int) int {
	if x <= 0 {
		return x
	}

	return -x
}

func TestDNAtoRNA(t *testing.T) {
	type args struct {
		dna string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"c1", args{"GCAT"}, "GCAU"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DNAtoRNA(tt.args.dna); got != tt.want {
				t.Errorf("DNAtoRNA() = %v, want %v", got, tt.want)
			}
		})
	}
}
