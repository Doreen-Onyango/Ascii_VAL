package ascii

import "testing"

func TestPrintChar(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintChar(tt.args.word)
		})
	}
}
