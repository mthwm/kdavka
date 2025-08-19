package davka16

import (
	"fmt"
	"strings"
	"time"
)

// formatC formats a string (C type): left-aligned, right-padded with spaces.
func formatC(value string, length int) string {
	if len(value) > length {
		value = value[:length]
	}
	return value + strings.Repeat(" ", length-len(value))
}

// formatN formats an int (N type): right-aligned, left-padded with spaces.
// For optional: ends with "0" if unset (value=0).
func formatN(value int, length int) string {
	s := fmt.Sprintf("%d", value)
	if len(s) > length {
		s = s[len(s)-length:] // Truncate
	}
	return strings.Repeat(" ", length-len(s)) + s
}

// formatD formats a date: "DDMMYYYY".
func formatD(t time.Time) string {
	return t.Format("02012006")
}

// formatMoney formats a float ($ type): with decimal point, right-aligned with spaces.
// For optional: ends with "0.00" if unset (value=0).
func formatMoney(value float64, total int, decimals int) string {
	s := fmt.Sprintf("%.*f", decimals, value)
	parts := strings.Split(s, ".")
	if len(parts) == 1 {
		parts = append(parts, strings.Repeat("0", decimals))
	} else if len(parts[1]) < decimals {
		parts[1] += strings.Repeat("0", decimals-len(parts[1]))
	}
	full := parts[0] + "." + parts[1]
	return strings.Repeat(" ", total-len(full)) + full
}
