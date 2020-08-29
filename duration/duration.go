package duration

import (
	"bytes"
	"fmt"
)

// Convert func: to convert input int to duration string
func Convert(seconds int) string {
	var minutes, hours, days int
	if seconds > 59 {
		minutes, seconds = seconds/60, seconds%60
		if minutes > 59 {
			hours, minutes = minutes/60, minutes%60
			if hours > 23 {
				days, hours = hours/24, hours%24
			}
		}
	}

	var buffer bytes.Buffer
	if days > 0 {
		buffer.WriteString(fmt.Sprintf("%dd", days))
	}
	if hours > 0 {
		buffer.WriteString(fmt.Sprintf("%dh", hours))
	}
	if minutes > 0 {
		buffer.WriteString(fmt.Sprintf("%dm", minutes))
	}
	buffer.WriteString(fmt.Sprintf("%ds", seconds))
	return buffer.String()
}
