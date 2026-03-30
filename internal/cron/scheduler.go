package cron

import (
	"log"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

// Scheduler loads a crontab file and runs matching jobs every minute.
type Scheduler struct {
	path    string
	entries []*Entry
	logger  *log.Logger
	mu      sync.RWMutex
	stop    chan struct{}
}

// NewScheduler creates a Scheduler that reads jobs from path.
func NewScheduler(path string, logger *log.Logger) *Scheduler {
	return &Scheduler{
		path:   path,
		logger: logger,
		stop:   make(chan struct{}),
	}
}

// Load reads and parses the crontab file for the first time.
func (s *Scheduler) Load() error {
	entries, err := parseCrontab(s.path)
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.entries = entries
	s.mu.Unlock()
	s.logger.Printf("loaded %d job(s)", len(entries))
	return nil
}

// Reload re-reads the crontab file without stopping the scheduler.
func (s *Scheduler) Reload() error {
	entries, err := parseCrontab(s.path)
	if err != nil {
		return err
	}
	s.mu.Lock()
	s.entries = entries
	s.mu.Unlock()
	s.logger.Printf("reloaded %d job(s)", len(entries))
	return nil
}

// Start begins the scheduling loop in a background goroutine.
func (s *Scheduler) Start() {
	go s.loop()
}

// Stop signals the scheduling loop to exit.
func (s *Scheduler) Stop() {
	close(s.stop)
}

func (s *Scheduler) loop() {
	// Fire @reboot jobs immediately on startup.
	s.mu.RLock()
	for _, e := range s.entries {
		if e.Reboot {
			go s.run(e)
		}
	}
	s.mu.RUnlock()

	// Sleep until the top of the next minute so ticks are always aligned.
	now := time.Now()
	next := now.Truncate(time.Minute).Add(time.Minute)
	select {
	case <-time.After(next.Sub(now)):
	case <-s.stop:
		return
	}

	s.tick(time.Now())

	ticker := time.NewTicker(time.Minute)
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			s.tick(t)
		case <-s.stop:
			return
		}
	}
}

func (s *Scheduler) tick(t time.Time) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, e := range s.entries {
		if !e.Reboot && e.Match(t) {
			go s.run(e)
		}
	}
}

func (s *Scheduler) run(e *Entry) {
	s.logger.Printf("start  | %s", e.Command)
	cmd := shellCommand(e.Command)
	out, err := cmd.CombinedOutput()
	if err != nil {
		s.logger.Printf("error  | %s: %v", e.Command, err)
	}
	if len(out) > 0 {
		for _, line := range strings.Split(strings.TrimRight(string(out), "\n\r"), "\n") {
			s.logger.Printf("output | %s", strings.TrimRight(line, "\r"))
		}
	}
	s.logger.Printf("done   | %s", e.Command)
}

// shellCommand wraps a raw command string in the platform shell.
func shellCommand(command string) *exec.Cmd {
	if runtime.GOOS == "windows" {
		return exec.Command("cmd", "/C", command)
	}
	return exec.Command("sh", "-c", command)
}
