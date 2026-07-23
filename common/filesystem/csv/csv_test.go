package csv

import (
	"mime/multipart"
	"net/textproto"
	"reflect"
	"testing"
)

func TestIsValidUploadedFileMimeType(t *testing.T) {
	tests := []struct {
		name        string
		contentType string
		want        bool
	}{
		{name: "CSV", contentType: "text/csv", want: true},
		{name: "plain text", contentType: "text/plain", want: true},
		{name: "Microsoft Excel", contentType: "application/vnd.ms-excel", want: true},
		{name: "unsupported", contentType: "application/json", want: false},
		{name: "missing", contentType: "", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			file := &multipart.FileHeader{
				Header: textproto.MIMEHeader{},
			}
			if tt.contentType != "" {
				file.Header.Set("Content-Type", tt.contentType)
			}

			if got := IsValidUploadedFileMimeType(file); got != tt.want {
				t.Errorf("IsValidUploadedFileMimeType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertHeaderAndRowToCsv(t *testing.T) {
	tests := []struct {
		name        string
		headers     []string
		rows        [][]interface{}
		want        string
		wantHeaders []string
	}{
		{
			name:        "empty input",
			want:        "",
			wantHeaders: nil,
		},
		{
			name:        "headers only",
			headers:     []string{"name", "age"},
			want:        "name,age",
			wantHeaders: []string{"name", "age"},
		},
		{
			name:        "basic row with mixed values",
			headers:     []string{"name", "age", "active"},
			rows:        [][]interface{}{{"Alice", 30, true}},
			want:        "name,age,active\r\nAlice,30,true",
			wantHeaders: []string{"name", "age", "active"},
		},
		{
			name:        "header containing comma",
			headers:     []string{"last, first", "age"},
			rows:        [][]interface{}{{"Doe, Jane", 28}},
			want:        "\"last, first\",age\r\n\"Doe, Jane\",28",
			wantHeaders: []string{"\"last, first\"", "age"},
		},
		{
			name:    "row values requiring escaping",
			headers: []string{"comma", "newline", "carriage return", "quote"},
			rows: [][]interface{}{{
				"one,two",
				"line one\nline two",
				"left\rright",
				`say "hello"`,
			}},
			want: "comma,newline,carriage return,quote\r\n" +
				"\"one,two\",\"line one\nline two\",\"left\rright\",\"say \"\"hello\"\"\"",
			wantHeaders: []string{"comma", "newline", "carriage return", "quote"},
		},
		{
			name:        "multiple rows",
			headers:     []string{"id", "name"},
			rows:        [][]interface{}{{1, "Alice"}, {2, "Bob"}},
			want:        "id,name\r\n1,Alice\r\n2,Bob",
			wantHeaders: []string{"id", "name"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConvertHeaderAndRowToCsv(tt.headers, tt.rows).String()
			if got != tt.want {
				t.Errorf("ConvertHeaderAndRowToCsv() = %q, want %q", got, tt.want)
			}
			if !reflect.DeepEqual(tt.headers, tt.wantHeaders) {
				t.Errorf("headers after ConvertHeaderAndRowToCsv() = %#v, want %#v", tt.headers, tt.wantHeaders)
			}
		})
	}
}
