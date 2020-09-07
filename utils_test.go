package utils

import "testing"

func TestContains(t *testing.T) {
	type args struct {
		a []string
		x string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "contains", args: args{a: []string{"a", "b", "c"}, x: "a"}, want: true},
		{name: "not contains", args: args{a: []string{"a", "b", "c"}, x: "a"}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}
