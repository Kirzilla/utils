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
		{name: "not contains", args: args{a: []string{"a", "b", "c"}, x: "x"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainsInt(t *testing.T) {
	type args struct {
		a []int
		x int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "contains", args: args{a: []int{1, 2, 3}, x: 1}, want: true},
		{name: "not contains", args: args{a: []int{1, 2, 3}, x: 10}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsInt(tt.args.a, tt.args.x); got != tt.want {
				t.Errorf("ContainsInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
