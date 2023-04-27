package chap3

import "testing"

func Test_commaIterative(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "5digits",
			args: args{
				"12345",
			},
			want: "12,345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := commaIterative(tt.args.s); got != tt.want {
				t.Errorf("commaIterative() = %v, want %v", got, tt.want)
			}
		})
	}
}
