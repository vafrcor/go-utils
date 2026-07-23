package filesystem

import (
	"bytes"
	"fmt"
	"mime/multipart"
	"slices"
	"strings"
)

func IsValidUploadedFileMimeType(file *multipart.FileHeader) bool {
	valid := false
	if slices.Contains(
		[]string{
			"text/csv", "text/plain", "application/vnd.ms-excel",
		}, file.Header.Get("Content-Type"),
	) {
		valid = true
	}

	return valid
}

func ConvertHeaderAndRowToCsv(headers []string, rows [][]interface{}) *bytes.Buffer {
	// Escape headers if they contain commas
	for i, header := range headers {
		if strings.Contains(header, ",") {
			headers[i] = fmt.Sprintf("\"%s\"", header)
		}
	}
	strOut := strings.Join(headers, ",")

	for _, row := range rows {
		strOut += "\r\n"
		var rowStr []string
		for _, r := range row {
			// Convert to string
			val := fmt.Sprintf("%v", r)

			// If value contains comma or newline, wrap in quotes
			if strings.Contains(val, ",") || strings.Contains(
				val, "\r",
			) || strings.Contains(
				val, "\n",
			) || strings.Contains(val, "\"") {
				// Escape any existing quotes by doubling them
				val = strings.ReplaceAll(val, "\"", "\"\"")
				val = fmt.Sprintf("\"%s\"", val)
			}

			rowStr = append(rowStr, val)
		}
		strOut += strings.Join(rowStr, ",")
	}
	return bytes.NewBuffer([]byte(strOut))
}
