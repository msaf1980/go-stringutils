package stringutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemplate_InitTemplate(t *testing.T) {
	tests := []struct {
		format   string
		template Template
		wantErr  bool
	}{
		{format: "", template: Template{"", ""}, wantErr: false},
		{format: "test", template: Template{"test", "test"}, wantErr: false},
		{
			format: "test %{query} %{log.status}",
			template: Template{
				"test %{query} %{log.status}",
				"test ", templateParam{"query", "query"}, " ", templateParam{"log.status", "log", "status"},
			},
			wantErr: false,
		},
		{
			format: "%{query}%{log.status} ",
			template: Template{
				"%{query}%{log.status} ",
				templateParam{"query", "query"}, templateParam{"log.status", "log", "status"}, " ",
			},
			wantErr: false,
		},
		{format: "%query}", template: nil, wantErr: true},
		{format: "%{query", template: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			tpl, err := InitTemplate(tt.format)
			if (err != nil) != tt.wantErr {
				t.Errorf("InitTemplate() error = %v, wantErr %v", err, tt.wantErr)
			}
			assert.Equalf(t, tt.template, tpl, "InitTemplate() results mismatch")
		})
	}
}

func TestTemplate_Execute(t *testing.T) {
	params := map[string]interface{}{
		"query":  "URL",
		"status": "success",
		"time":   "now",
		"one": map[string]interface{}{
			"two_one": "2_1", "two_two": "2_2",
			"three": map[string]interface{}{"three_four": "3_4"},
		},
	}
	tests := []struct {
		format  string
		want    string
		wantErr bool
	}{
		// one element
		{format: "", want: "", wantErr: false},
		{format: "test", want: "test", wantErr: false},
		{format: "%{query}", want: "URL", wantErr: false},
		{format: "%{not_exist}", want: "", wantErr: true},
		//several elements
		{format: "%{query} %{status}", want: "URL success", wantErr: false},
		{format: "%{query} %{status} at %{time}", want: "URL success at now", wantErr: false},
		// deep params
		{format: "%{query} %{one.two_one} %{one.three.three_four}", want: "URL 2_1 3_4", wantErr: false},
		{format: "%{query} %{one.two_one} %{one.three.three_four.2}", want: "", wantErr: true},
		{format: "%{query} %{one.two_one} %{one.three}", want: "", wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.format, func(t *testing.T) {
			tpl, err := InitTemplate(tt.format)
			if err != nil {
				t.Fatalf("InitTemplate() error = %v", err)
			}
			got, err := tpl.Execute(params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Execute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Template.Execute() = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

func Benchmark_TemplateExecute(b *testing.B) {
	params := map[string]interface{}{
		"query":  "URL",
		"status": "success",
		"time":   "now",
		"one": map[string]interface{}{
			"two_one": "2_1", "two_two": "2_2",
			"three": map[string]interface{}{"three_four": "3_4"},
		},
	}
	format := "%{query} %{one.two_one} %{one.three.three_four}"
	tpl, err := InitTemplate(format)
	if err != nil {
		b.Fatalf("InitTemplate() error = %v", err)
	}
	for i := 0; i < b.N; i++ {
		got, err := tpl.Execute(params)
		_ = got
		if err != nil {
			b.Fatalf("Template.Execute() error = %v", err)
			return
		}
	}
}

func Benchmark_TemplateInitAndExecute(b *testing.B) {
	params := map[string]interface{}{
		"query":  "URL",
		"status": "success",
		"time":   "now",
		"one": map[string]interface{}{
			"two_one": "2_1", "two_two": "2_2",
			"three": map[string]interface{}{"three_four": "3_4"},
		},
	}
	format := "%{query} %{one.two_one} %{one.three.three_four}"
	for i := 0; i < b.N; i++ {
		tpl, err := InitTemplate(format)
		if err != nil {
			b.Fatalf("InitTemplate() error = %v", err)
		}
		got, err := tpl.Execute(params)
		_ = got
		if err != nil {
			b.Fatalf("Template.Execute() error = %v", err)
			return
		}
	}
}
