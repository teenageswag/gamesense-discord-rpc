package log

import "fmt"

func Info(message string) {
	fmt.Println("[LOG]", message)
}
func Success(message string) {
	fmt.Println("[SUCCESS]", message)
}
func Warning(message string) {
	fmt.Println("[WARNING]", message)
}
func Error(message string) {
	fmt.Println("[ERROR]", message)
}
