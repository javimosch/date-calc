# date-calc
Calculate date offsets. Add/subtract days from a date. Written in Go.
```bash
curl -LO https://github.com/javimosch/date-calc/releases/latest/download/date-calc-linux-amd64
chmod +x date-calc-linux-amd64; sudo mv date-calc-linux-amd64 /usr/local/bin/date-calc
```
Usage: `date-calc 30` → date 30 days from now. `date-calc -7 2024-01-15` → 7 days before Jan 15

---

This tool is a plugin for [supercli](https://github.com/javimosch/supercli) — an AI-friendly, config-driven dynamic CLI platform.
