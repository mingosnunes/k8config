package utils

import "github.com/fatih/color"

var PrintYellow *color.Color = color.New(color.FgHiYellow)
var PrintRed *color.Color = color.New(color.FgHiRed)
var PrintGreen *color.Color = color.New(color.FgHiGreen)

func PrintSuccess(text string) {
	PrintGreen.Println("✅ " + text)
}

func PrintWaring(text string) {
	println("⚠️ " + text)
}

func PrintError(text string) {
	PrintRed.Println("❌ " + text)
}

func PrintInfo(text string) {
	println("ℹ️ " + text)
}
