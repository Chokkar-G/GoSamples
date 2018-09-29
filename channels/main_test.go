// _Channels_ are the pipes that connect concurrent

// goroutines. You can send values into channels from one

// goroutine and receive those values into another

// goroutine.

package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
