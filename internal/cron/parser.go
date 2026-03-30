package cron

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var specialSchedules = map[string]string{
	"@yearly":   "0 0 1 1 *",
	"@annually": "0 0 1 1 *",
	"@monthly":  "0 0 1 * *",
	"@weekly":   "0 0 * * 0",
	"@daily":    "0 0 * * *",
	"@midnight": "0 0 * * *",
	"@hourly":   "0 * * * *",
}

func parseCrontab(path string) ([]*Entry, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var entries []*Entry
	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		e, err := parseLine(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum, err)
		}
		if e != nil {
			entries = append(entries, e)
		}
	}
	return entries, scanner.Err()
}

func parseLine(raw string) (*Entry, error) {
	line := raw
	if idx := strings.IndexByte(line, '#'); idx != -1 {
		line = line[:idx]
	}
	line = strings.TrimSpace(line)
	if line == "" {
		return nil, nil
	}
	if strings.HasPrefix(line, "@") {
		return parseSpecial(line, raw)
	}
	return parseStandard(line, raw)
}

func parseSpecial(line, raw string) (*Entry, error) {
	parts := strings.SplitN(line, " ", 2)
	keyword := strings.ToLower(parts[0])
	if len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return nil, fmt.Errorf("missing command after %q", keyword)
	}
	command := strings.TrimSpace(parts[1])

	if keyword == "@reboot" {
		return &Entry{Command: command, Reboot: true, Raw: raw}, nil
	}
	expr, ok := specialSchedules[keyword]
	if !ok {
		return nil, fmt.Errorf("unknown schedule keyword %q", keyword)
	}
	return parseStandard(expr+" "+command, raw)
}

func parseStandard(line, raw string) (*Entry, error) {
	fields := strings.Fields(line)
	if len(fields) < 6 {
		return nil, fmt.Errorf("expected at least 6 fields (5 time fields + command), got %d", len(fields))
	}

	minute, err := parseField(fields[0], 0, 59)
	if err != nil {
		return nil, fmt.Errorf("minute field: %w", err)
	}
	hour, err := parseField(fields[1], 0, 23)
	if err != nil {
		return nil, fmt.Errorf("hour field: %w", err)
	}
	dom, err := parseField(fields[2], 1, 31)
	if err != nil {
		return nil, fmt.Errorf("day-of-month field: %w", err)
	}
	month, err := parseField(fields[3], 1, 12)
	if err != nil {
		return nil, fmt.Errorf("month field: %w", err)
	}
	dow, err := parseField(fields[4], 0, 6)
	if err != nil {
		return nil, fmt.Errorf("day-of-week field: %w", err)
	}

	return &Entry{
		Minute:     minute,
		Hour:       hour,
		DayOfMonth: dom,
		Month:      month,
		DayOfWeek:  dow,
		Command:    strings.Join(fields[5:], " "),
		Raw:        raw,
	}, nil
}
