package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	var num uint32 = 0b11111111111111111111111111111101
	for idx := 0; idx < b.N; idx++ {
		hammingWeight(num)
	}
}
func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n = 00000000000000000000000000001011",
			args: args{num: 0b00000000000000000000000000001011},
			want: 3,
		},
		{
			name: "n = 00000000000000000000000010000000",
			args: args{num: 0b00000000000000000000000010000000},
			want: 1,
		},
		{
			name: "n = 11111111111111111111111111111101",
			args: args{num: 0b11111111111111111111111111111101},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
