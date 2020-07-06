package jerr

import "testing"

func Test_escapeJSON(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"backslash", args{input: "\""}, "\\\""},
		{"new line in Windows", args{input: "\r\n"}, "\\n"},
		{"new line in Mac OS before X", args{input: "\r"}, "\\n"},
		{"new line in Unix/macOS", args{input: "\n"}, "\\n"},
		{"tab", args{input: "\t"}, "\\t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := escapeJSON(tt.args.input)

			if got != tt.want {
				t.Error("Got:", got, "Want:", tt.want)
			}
		})
	}
}
