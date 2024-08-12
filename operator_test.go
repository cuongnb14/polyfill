package polyfill

import "testing"

func TestTernary(t *testing.T) {
	type args struct {
		condition   bool
		trueResult  interface{}
		falseResult interface{}
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		{
			name: "True Condition",
			args: args{condition: true, trueResult: "TrueCase", falseResult: "FalseCase"},
			want: "TrueCase",
		},
		{
			name: "False Condition",
			args: args{condition: false, trueResult: "TrueCase", falseResult: "FalseCase"},
			want: "FalseCase",
		},
		{
			name: "Zero True Case",
			args: args{condition: true, trueResult: 0, falseResult: 1},
			want: 0,
		},
		{
			name: "Zero False Case",
			args: args{condition: false, trueResult: 0, falseResult: 1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ternary(tt.args.condition, tt.args.trueResult, tt.args.falseResult)
			if got != tt.want {
				t.Errorf("Ternary() = %v, want %v", got, tt.want)
			}
		})
	}
}
