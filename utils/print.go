package utils

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
)

func PrintSuccess(format string, args ...any) {
	fmt.Println(aurora.BrightGreen(fmt.Sprintf(format, args...)))
}

func PrintWarn(format string, args ...any) {
	fmt.Println(aurora.BrightYellow(fmt.Sprintf(format, args...)))
}

func PrintError(format string, args ...any) {
	fmt.Println(aurora.BrightRed(fmt.Sprintf(format, args...)))
}

func PrintInfo(format string, args ...any) {
	fmt.Println(aurora.BrightBlue(fmt.Sprintf(format, args...)))
}
