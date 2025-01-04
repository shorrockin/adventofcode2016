package colors

import "fmt"

func Red(in string) string    { return Color(in, "\033[31m") }
func Blue(in string) string   { return Color(in, "\033[34m") }
func Green(in string) string  { return Color(in, "\033[32m") }
func Yellow(in string) string { return Color(in, "\033[33m") }

func Color(in, code string) string {
	return fmt.Sprintf("%s%s%s", code, in, "\033[0m")
}
