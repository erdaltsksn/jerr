package jerr_test

import (
	"errors"
	"testing"

	"github.com/erdaltsksn/jerr"
)

func TestNew(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{`empty`,
			args{message: ""},
			"{}",
		},
		{`simple`,
			args{message: "Hello World"},
			`{"message":"Hello World"}`,
		},
		{`i18n`,
			args{message: "ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә"},
			`{"message":"ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә"}`,
		},
		{`html`,
			args{message: `<p class='title'>Paragraph<hr /></p>`},
			`{"message":"<p class='title'>Paragraph<hr /></p>"}`,
		},
		{`html with double quote`,
			args{message: `<div class="title"></div>`},
			`{"message":"<div class=\"title\"></div>"}`,
		},
		{`newline`,
			// DO NOT remove the new line in this string literal
			args{message: `New
Line`},
			`{"message":"New\nLine"}`,
		},
		{`newline with \n`,
			args{message: `New\nLine`},
			`{"message":"New\nLine"}`,
		},
		{`tab`,
			args{message: `json	error`},
			`{"message":"json\terror"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jerr.New(tt.args.message).Error()

			if got != tt.want {
				t.Error("Got:", got, "Want:", tt.want)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	type args struct {
		message string
		wrapped error
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{`empty`,
			args{message: "", wrapped: errors.New("Hello from wrapped")},
			`{"message":"Hello from wrapped"}`,
		},
		{`simple`,
			args{message: "Failed", wrapped: errors.New("You have entered a wrong number")},
			`{"message":"Failed","details":"You have entered a wrong number"}`,
		},
		{`i18n`,
			args{message: "ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә", wrapped: errors.New("abc")},
			`{"message":"ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә","details":"abc"}`,
		},
		{`html`,
			args{message: `<p class='title'>Paragraph<hr /></p>`, wrapped: errors.New("abc")},
			`{"message":"<p class='title'>Paragraph<hr /></p>","details":"abc"}`,
		},
		{`html with double quote`,
			args{message: `<div class="title"></div>`, wrapped: errors.New("abc")},
			`{"message":"<div class=\"title\"></div>","details":"abc"}`,
		},
		{`newline`,
			// DO NOT remove the new line in this string literal
			args{message: `New
Line`, wrapped: errors.New("abc")},
			`{"message":"New\nLine","details":"abc"}`,
		},
		{`newline with \n`,
			args{message: `New\nLine`, wrapped: errors.New("abc")},
			`{"message":"New\nLine","details":"abc"}`,
		},
		{`tab`,
			args{message: `json	error`, wrapped: errors.New("abc")},
			`{"message":"json\terror","details":"abc"}`,
		},
		{`empty with a jerr.New() error`,
			args{message: "", wrapped: jerr.New("Hello from jerr.New()")},
			`{"message":"Hello from jerr.New()"}`,
		},
		{`simple with a jerr.New() error`,
			args{message: "Failed", wrapped: jerr.New("Hello from jerr.New()")},
			`{"message":"Failed","details":{"message":"Hello from jerr.New()"}}`,
		},
		{`empty with a jerr.Wrap() error`,
			args{message: "", wrapped: jerr.Wrap(errors.New("from wrap()'s details"), "Hello from jerr.Wrap()")},
			`{"message":"Hello from jerr.Wrap()","details":"from wrap()'s details"}`,
		},
		{`simple with a jerr.Wrap() error`,
			args{message: "Failed", wrapped: jerr.Wrap(errors.New("Something is wrong"), "Check the details, please")},
			`{"message":"Failed","details":{"message":"Check the details, please","details":"Something is wrong"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jerr.Wrap(tt.args.wrapped, tt.args.message).Error()

			if got != tt.want {
				t.Error("Got:", got, "Want:", tt.want)
			}
		})
	}
}
