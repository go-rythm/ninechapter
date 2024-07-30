package ds

import "testing"

func TestNthUglyNumber(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10",
			args: args{n: 10},
			want: 12,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NthUglyNumber2(tt.args.n); got != tt.want {
				t.Errorf("NthUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
