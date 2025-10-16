package pcolor

import "testing"

func TestPrintSucc(t *testing.T) {
	type args struct {
		prefix string
		str    string
		a      []any
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "PrintSucc",
			args: args{
				prefix: "color",
				str:    "Succ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintSucc(tt.args.prefix, tt.args.str, tt.args.a...)
		})
	}
}
