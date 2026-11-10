import { defineConfig } from 'vitepress'

export default defineConfig({
  title: "pro-cron",
  description: "A stripped-down crond clone in Go that reads a crontab, runs matching jobs each minute, and logs to stdout.",
  base: '/pro-cron/',
  cleanUrls: true,
  lastUpdated: true,
  head: [
    ['link', { rel: 'icon', href: '/favicon.svg' }],
  ],
  themeConfig: {
    logo: { light: '/logo.svg', dark: '/logo.svg' },
    siteTitle: "pro-cron",
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Guide', link: '/guide' },
      { text: 'Getting started', link: '/getting-started' },
      {
        text: 'GitHub',
        link: 'https://github.com/mediuswareltd/pro-cron'
      }
    ],
    sidebar: [
      {
        text: 'Documentation',
        items: [
          { text: 'Overview', link: '/' },
          { text: 'Getting started', link: '/getting-started' },
          { text: 'Guide', link: '/guide' },
          { text: 'About', link: '/about' }
        ]
      }
    ],
    socialLinks: [
      { icon: 'github', link: 'https://github.com/mediuswareltd/pro-cron' }
    ],
    footer: {
      message: 'Built by Mediusware',
      copyright: 'Copyright © Mediusware Ltd'
    },
    editLink: {
      pattern: 'https://github.com/mediuswareltd/pro-cron/edit/main/docs/:path',
      text: 'Edit this page on GitHub'
    },
    search: {
      provider: 'local'
    }
  }
})
