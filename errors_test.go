package jerr_test

import (
	"errors"
	"testing"

	"github.com/erdaltsksn/jerr"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		message string
		output  string
	}{
		{`empty`,
			"",
			"{}",
		},
		{`simple`,
			"Hello World",
			`{"message":"Hello World"}`,
		},
		{`i18n`,
			"ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә",
			`{"message":"ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә"}`,
		},
		{`html`,
			`<p class='title'>Paragraph<hr /></p>`,
			`{"message":"<p class='title'>Paragraph<hr /></p>"}`,
		},
		{`html with double quote`,
			`<div class="title"></div>`,
			`{"message":"<div class=\"title\"></div>"}`,
		},
		{`newline`,
			// DO NOT remove the new line in this string literal
			`New
Line`,
			`{"message":"New\nLine"}`,
		},
		{`newline with \n`,
			`New\nLine`,
			`{"message":"New\nLine"}`,
		},
		{`tab`,
			`json	error`,
			`{"message":"json\terror"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jerr.New(tt.message).Error()
			want := tt.output

			if got != want {
				t.Error("Got:", got, ",", "Want:", want)
			}
		})
	}
}

func TestWrap(t *testing.T) {
	tests := []struct {
		name    string
		message string
		wrapped error
		output  string
	}{
		{`empty`,
			"", errors.New("Hello from wrapped"),
			`{"message":"Hello from wrapped"}`,
		},
		{`simple`,
			"Failed", errors.New("You have entered a wrong number"),
			`{"message":"Failed","details":"You have entered a wrong number"}`,
		},
		{`i18n`,
			"ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә", errors.New("abc"),
			`{"message":"ı ğ ü ş i ö ç ä I Ğ Ü Ş İ Ö Ç â ê Ä η ή ί ώ w Ω Ә","details":"abc"}`,
		},
		{`html`,
			`<p class='title'>Paragraph<hr /></p>`, errors.New("abc"),
			`{"message":"<p class='title'>Paragraph<hr /></p>","details":"abc"}`,
		},
		{`html with double quote`,
			`<div class="title"></div>`, errors.New("abc"),
			`{"message":"<div class=\"title\"></div>","details":"abc"}`,
		},
		{`newline`,
			// DO NOT remove the new line in this string literal
			`New
Line`, errors.New("abc"),
			`{"message":"New\nLine","details":"abc"}`,
		},
		{`newline with \n`,
			`New\nLine`, errors.New("abc"),
			`{"message":"New\nLine","details":"abc"}`,
		},
		{`tab`,
			`json	error`, errors.New("abc"),
			`{"message":"json\terror","details":"abc"}`,
		},
		{`empty with a jerr.New() error`,
			"", jerr.New("Hello from jerr.New()"),
			`{"message":"Hello from jerr.New()"}`,
		},
		{`simple with a jerr.New() error`,
			"Failed", jerr.New("Hello from jerr.New()"),
			`{"message":"Failed","details":{"message":"Hello from jerr.New()"}}`,
		},
		{`empty with a jerr.Wrap() error`,
			"", jerr.Wrap(errors.New("from wrap()'s details"), "Hello from jerr.Wrap()"),
			`{"message":"Hello from jerr.Wrap()","details":"from wrap()'s details"}`,
		},
		{`simple with a jerr.Wrap() error`,
			"Failed", jerr.Wrap(errors.New("Something is wrong"), "Check the details, please"),
			`{"message":"Failed","details":{"message":"Check the details, please","details":"Something is wrong"}}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jerr.Wrap(tt.wrapped, tt.message).Error()
			want := tt.output

			if got != want {
				t.Error("Got:", got, ",", "Want:", want)
			}
		})
	}
}
