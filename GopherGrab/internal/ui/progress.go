package ui

import (
	"fmt"
	"strings"
)

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorRed    = "\033[31m"
	ColorCyan   = "\033[36m"
	ColorYellow = "\033[33m"
)

func FormatBytes(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}

func PrintDivider() {
	fmt.Println(strings.Repeat("-", 60))
}

func PrintHeader(workerCount, jobCount int) {
	fmt.Printf("%s GopherGrab started with  %d workers. Downloading  %d file...%s\n\n",
		ColorCyan, workerCount, jobCount, ColorReset)
}
