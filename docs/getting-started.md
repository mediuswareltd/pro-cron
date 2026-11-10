# Getting started

Run **pro-cron** on your machine.

## 1. Build

You need Go 1.21 or newer.

```bash
go build -o pro-cron ./cmd/pro-cron
```

That's it — one binary, zero external dependencies.

---

## 2. Usage

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

## 3. Quick start

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

## 4. Run

```bash
nohup ./pro-cron &> /var/log/pro-cron.log &
```

### Run on Windows

```powershell
Start-Process -FilePath ".\pro-cron.exe" -WindowStyle Hidden
```

---

## Next

- Full detail: [Guide](/guide)
- Source: [github.com/mediuswareltd/pro-cron](https://github.com/mediuswareltd/pro-cron)
