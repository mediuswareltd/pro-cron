# Guide

## What it is

A stripped-down crond clone in Go that reads a crontab, runs matching jobs each minute, and logs to stdout.

## Stack

`Go`, `cron`

## Capabilities

### Classic crontab

Familiar Unix cron schedule syntax.

### One binary

No C deps, no root required.

### Foreground runner

Simple ops model with stdout logging.

### SIGHUP reload

Reload crontab without restart on Unix.

## Source

Repository: [mediuswareltd/pro-cron](https://github.com/mediuswareltd/pro-cron)
