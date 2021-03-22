package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestOptimize(t *testing.T) {
	tests := []struct {
		name     string
		solution []string
		want     []string
	}{
		{name: "no puz", want: []string{}},
		{
			name: "normal",
			solution: []string{
				"movep1y(-1)	 # Move #1: {p1y:-1,p2y:0,s1x:3,s2x:3}",
				"movep2y(1)	 # Move #2: {p1y:-1,p2y:1,s1x:3,s2x:3}	",
				"moves1x(2)	 # Move #3: {p1y:-1,p2y:1,s1x:2,s2x:3}	",
				"movep2y(2)	 # Move #4: {p1y:-1,p2y:2,s1x:2,s2x:3}	",
				"moves2x(4)	 # Move #5: {p1y:-1,p2y:2,s1x:2,s2x:4}	",
				"movep1y(-2)	 # Move #6: {p1y:-2,p2y:2,s1x:2,s2x:4}",
				"moves1x(1)	 # Move #7: {p1y:-2,p2y:2,s1x:1,s2x:4}  ",
			},
			want: []string{
				"movep1y(-1)	 # Move #1: {p1y:-1,p2y:0,s1x:3,s2x:3}",
				"movep2y(1)	 # Move #2: {p1y:-1,p2y:1,s1x:3,s2x:3}	",
				"moves1x(2)	 # Move #3: {p1y:-1,p2y:1,s1x:2,s2x:3}	",
				"movep2y(2)	 # Move #4: {p1y:-1,p2y:2,s1x:2,s2x:3}	",
				"moves2x(4)	 # Move #5: {p1y:-1,p2y:2,s1x:2,s2x:4}	",
				"movep1y(-2)	 # Move #6: {p1y:-2,p2y:2,s1x:2,s2x:4}",
				"moves1x(1)	 # Move #7: {p1y:-2,p2y:2,s1x:1,s2x:4}  ",
			},
		},
		{
			name: "two s1x in a row",
			solution: []string{
				"movep1y(-2)	 # Move #6: {p1y:-2,p2y:2,s1x:2,s2x:4}",
				"moves1x(1)	 # Move #7: {p1y:-2,p2y:2,s1x:1,s2x:4}",
				"moves1x(0)	 # Move #8: {p1y:-2,p2y:2,s1x:0,s2x:4}",
			},
			want: []string{
				"movep1y(-2)	 # Move #6: {p1y:-2,p2y:2,s1x:2,s2x:4}",
				"moves1x(0)	 # Move #8: {p1y:-2,p2y:2,s1x:0,s2x:4}",
			},
		},
		{
			name: "string of p1 and p2",
			solution: []string{
				"moves2x(3)	 # Move #10: {p1y:-1,p2y:2,s1x:0,s2x:3}",
				"movep1y(0)	 # Move #11: {p1y:0,p2y:2,s1x:0,s2x:3} ",
				"movep1y(1)	 # Move #12: {p1y:1,p2y:2,s1x:0,s2x:3} ",
				"movep1y(2)	 # Move #13: {p1y:2,p2y:2,s1x:0,s2x:3} ",
				"movep2y(1)	 # Move #14: {p1y:2,p2y:1,s1x:0,s2x:3} ",
				"movep1y(1)	 # Move #15: {p1y:1,p2y:1,s1x:0,s2x:3} ",
				"movep1y(0)	 # Move #16: {p1y:0,p2y:1,s1x:0,s2x:3} ",
				"movep1y(-1)	 # Move #17: {p1y:-1,p2y:1,s1x:0,s2x:3}",
				"movep2y(0)	 # Move #18: {p1y:-1,p2y:0,s1x:0,s2x:3}",
				"movep1y(0)	 # Move #19: {p1y:0,p2y:0,s1x:0,s2x:3} ",
				"movep1y(1)	 # Move #20: {p1y:1,p2y:0,s1x:0,s2x:3} ",
				"movep1y(2)	 # Move #21: {p1y:2,p2y:0,s1x:0,s2x:3} ",
				"movep2y(-1)	 # Move #22: {p1y:2,p2y:-1,s1x:0,s2x:3}",
				"movep1y(1)	 # Move #23: {p1y:1,p2y:-1,s1x:0,s2x:3}",
				"movep1y(0)	 # Move #24: {p1y:0,p2y:-1,s1x:0,s2x:3}",
				"moves2x(2)	 # Move #25: {p1y:0,p2y:-1,s1x:0,s2x:2}",
			},
			want: []string{
				"moves2x(3)	 # Move #10: {p1y:-1,p2y:2,s1x:0,s2x:3}",
				"movep2y(-1)	 # Move #22: {p1y:2,p2y:-1,s1x:0,s2x:3}",
				"movep1y(0)	 # Move #24: {p1y:0,p2y:-1,s1x:0,s2x:3}",
				"moves2x(2)	 # Move #25: {p1y:0,p2y:-1,s1x:0,s2x:2}",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := optimize(tt.solution)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got =\n%v\n\nwant =\n%v", strings.Join(got, "\n"), strings.Join(tt.want, "\n"))
			}
		})
	}
}
