// Go supports _methods_ defined on struct types.

package main

import "testing"

func Test_rect_area(t *testing.T) {
	tests := []struct {
		name string
		r    *rect
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.area(); got != tt.want {
				t.Errorf("rect.area() = %v, want %v", got, tt.want)
			}
		})
	}
}
