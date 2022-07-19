/*
Copyright © 2022 Domingos Nunes mingosnunes94@gmail.com

*/
package utils

import (
	"fmt"

	"github.com/fatih/color"
)

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
	fmt.Println("ℹ️ " + text)
}

func PrintDebug(text string) {
	fmt.Println("[DEBUG]", text)
}
