package pcolor

import "testing"

func TestPrintSucc(t *testing.T) {
	type args struct {
		str string
		a   []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "PrintSucc",
			args: args{
				str: "Succ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintSucc(tt.args.str, tt.args.a...)
		})
	}
}
