# pro-cron

[![Docs](https://img.shields.io/badge/docs-GitHub%20Pages-0f766e)](https://mediuswareltd.github.io/pro-cron/)

<p align="center">
  <img src="arts/logo.png" width="250" height="250" alt="pro-cron logo" />
</p>

A minimal cron daemon written in Go — a stripped-down clone of `crond` that does exactly what you need and nothing more.

It reads a crontab file, wakes up at the top of every minute, and fires off any commands whose schedule matches the current time. No databases, no C dependencies, no root required.

---

## How it works

1. You write a crontab file (same syntax as classic Unix cron).
2. You start `pro-cron` pointing at that file.
3. It runs in the foreground, fires jobs at the right time, and logs everything to stdout.
4. Send `SIGHUP` (Linux/macOS) to reload the crontab without restarting.

---

## Build

You need Go 1.21 or newer.

```bash
go build -o pro-cron ./cmd/pro-cron
```

That's it — one binary, zero external dependencies.

---

## Usage

```
pro-cron [-f /path/to/crontab]
```

| Flag | Default (Linux/macOS) | Default (Windows) | Description |
|------|-----------------------|-------------------|-------------|
| `-f` | `~/.pro-cron` | `%USERPROFILE%\pro-cron.tab` | Path to your crontab file |

### Quick start

```bash
# Build
go build -o pro-cron ./cmd/pro-cron

# Copy the example crontab and edit it
cp pro-cron.proc ~/.pro-cron

# Start the daemon
./pro-cron
```

### Run in the background (Linux/macOS)

```bash
nohup ./pro-cron &> /var/log/pro-cron.log &
```

### Run on Windows

```powershell
Start-Process -FilePath ".\pro-cron.exe" -WindowStyle Hidden
```

---

## Crontab format

The format is the classic five-field cron syntax:

```
┌──────────── minute       (0–59)
│  ┌─────────── hour         (0–23)
│  │  ┌────────── day of month (1–31)
│  │  │  ┌───────── month        (1–12)
│  │  │  │  ┌──────── day of week  (0–6, 0 = Sunday)
│  │  │  │  │
*  *  *  *  *  command to run
```

Each field accepts:

| Syntax | Meaning | Example |
|--------|---------|---------|
| `*` | every value | `*` = every minute |
| `n` | exact value | `5` = at minute 5 |
| `a-b` | range | `1-5` = Monday through Friday |
| `*/n` | every n steps | `*/15` = every 15 minutes |
| `a-b/n` | range with step | `0-30/5` = 0, 5, 10, …, 30 |
| `a,b,c` | list | `1,3,5` = minutes 1, 3, and 5 |

Fields can be combined: `0,30 9-17 * * 1-5` means "at :00 and :30, during 9 AM–5 PM, weekdays."

### Special keywords

These are shorthand for common schedules:

| Keyword | Equivalent |
|---------|------------|
| `@reboot` | Once, when the daemon starts |
| `@hourly` | `0 * * * *` |
| `@daily` / `@midnight` | `0 0 * * *` |
| `@weekly` | `0 0 * * 0` |
| `@monthly` | `0 0 1 * *` |
| `@yearly` / `@annually` | `0 0 1 1 *` |

### Comments and blank lines

Lines starting with `#` are comments. Blank lines are ignored.

```
# This is a comment
* * * * * echo "every minute"
```

---

## Example crontab

```cron
# Ping a health-check endpoint every 5 minutes
*/5 * * * * curl -s https://example.com/health

# Back up the database every night at 2:30 AM
30 2 * * * /usr/local/bin/backup.sh

# Send a weekly digest every Monday at 9 AM
0 9 * * 1 /usr/local/bin/weekly-report.sh

# Run once on daemon startup
@reboot /usr/local/bin/init.sh
```

---

## Reloading the crontab

Edit your crontab file, then:

```bash
# Linux / macOS — send SIGHUP to the running daemon
kill -HUP $(pgrep pro-cron)
```

On Windows there is no SIGHUP. Just stop and restart the daemon after editing the file.

---

## Logging

All output goes to **stdout** in a simple, readable format:

```
[pro-cron] 2026/03/18 09:00:00 daemon started (pid 12345), crontab: /home/user/.pro-cron
[pro-cron] 2026/03/18 09:00:00 loaded 4 job(s)
[pro-cron] 2026/03/18 09:00:01 start  | /usr/local/bin/backup.sh
[pro-cron] 2026/03/18 09:00:03 output | Backup complete: 42 MB
[pro-cron] 2026/03/18 09:00:03 done   | /usr/local/bin/backup.sh
```

To write logs to a file, redirect stdout when starting:

```bash
./pro-cron >> /var/log/pro-cron.log 2>&1
```

---

## What it does NOT do

A few things the classic `crond` handles that `pro-cron` intentionally skips to stay minimal:

- **No mail delivery** — job output is logged, not emailed.
- **No per-user crontabs** — one crontab file, one daemon.
- **No environment variable definitions** inside the crontab (`VAR=value` lines are not supported).
- **No second-level precision** — minimum interval is one minute.

If you need any of these, consider using the system `crond` or a more full-featured scheduler.

---

## Signals

| Signal | Effect |
|--------|--------|
| `SIGTERM` / `SIGINT` / `Ctrl-C` | Graceful shutdown |
| `SIGHUP` | Reload crontab file (Linux/macOS only) |

---

## License

MIT

## Documentation

Full docs site (VitePress): **https://mediuswareltd.github.io/pro-cron/**

Local preview:

```bash
cd docs
npm install
npm run docs:dev
```
