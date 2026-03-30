package cron

import "time"

// Entry is a single parsed crontab line.
type Entry struct {
	Minute     *Field
	Hour       *Field
	DayOfMonth *Field
	Month      *Field
	DayOfWeek  *Field
	Command    string
	Reboot     bool // true for @reboot entries
	Raw        string
}

// Match reports whether this entry should fire at time t.
func (e *Entry) Match(t time.Time) bool {
	return e.Minute.Match(t.Minute()) &&
		e.Hour.Match(t.Hour()) &&
		e.DayOfMonth.Match(t.Day()) &&
		e.Month.Match(int(t.Month())) &&
		e.DayOfWeek.Match(int(t.Weekday()))
}
