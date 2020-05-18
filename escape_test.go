package jerr

import "testing"

func Test_escapeJSON(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"backslash", "\"", "\\\""},
		{"new line in Windows", "\r\n", "\\n"},
		{"new line in Mac OS before X", "\r", "\\n"},
		{"new line in Unix/macOS", "\n", "\\n"},
		{"tab", "\t", "\\t"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := escapeJSON(tt.input)
			want := tt.output

			if got != want {
				t.Error("Got:", got, ",", "Want:", want)
			}
		})
	}
}
