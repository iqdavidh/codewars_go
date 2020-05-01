package kata

import "testing"

//Expect(IsTriangle(5, 1, 2)).To(Equal(false))
//Expect(IsTriangle(1, 2, 5)).To(Equal(false))
//Expect(IsTriangle(2, 5, 1)).To(Equal(false))
//Expect(IsTriangle(4, 2, 3)).To(Equal(true))
//Expect(IsTriangle(5, 1, 5)).To(Equal(true))
//Expect(IsTriangle(2, 2, 2)).To(Equal(true))
//Expect(IsTriangle(-1, 2, 3)).To(Equal(false))

func TestIsTriangle(t *testing.T) {
	type args struct {
		a int
		b int
		c int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"t1", args{2, 5, 1}, false},
		{"t2", args{4, 2, 3}, true},
		{"t3", args{5, 1, 5}, true},
		{"t4", args{2, 2, 2}, true},
		{"t5", args{-1, 2, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsTriangle(tt.args.a, tt.args.b, tt.args.c); got != tt.want {
				t.Errorf("IsTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
