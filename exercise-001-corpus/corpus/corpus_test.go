package corpus

import (
	"reflect"
	"testing"
)

func TestAnalyze(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want []WordCount
	}{
		{"Short string", args{data: "Hey hey hey what what!"}, []WordCount{{"hey", 3}, {"what", 2}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Analyze(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Analyze() = %v, want %v", got, tt.want)
			}
		})
	}
}
