package strings

import "strings"

// Customer String Replacer with dynamic placeholder and replacer
func StringDynamicReplacer(text string, replacers ...string) string {
	replacer := strings.NewReplacer(replacers...)
	return replacer.Replace(text)
}
