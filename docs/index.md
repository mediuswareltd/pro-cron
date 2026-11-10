---
layout: home

hero:
  name: "pro-cron"
  text: "Minimal cron daemon in Go"
  tagline: "A stripped-down crond clone in Go that reads a crontab, runs matching jobs each minute, and logs to stdout."
  image:
    src: /hero.png
    alt: "pro-cron"
  actions:
  - theme: brand
    text: Get started
    link: /getting-started
  - theme: alt
    text: View on GitHub
    link: https://github.com/mediuswareltd/pro-cron

features:
  - title: "Classic crontab"
    details: "Familiar Unix cron schedule syntax."
  - title: "One binary"
    details: "No C deps, no root required."
  - title: "Foreground runner"
    details: "Simple ops model with stdout logging."
  - title: "SIGHUP reload"
    details: "Reload crontab without restart on Unix."
---
