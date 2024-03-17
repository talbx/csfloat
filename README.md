# csfloat

A cron-based csfloat deal finder.

## Usage

```bash
go build -o float
```

```
Usage:
  float [flags]

Flags:
  -c, --category int        Item category - default normal (1) (default 1)
  -d, --discount float      Min Discount (default 5)
  -v, --discountValue int   Min Discount Value (cents) (default 10)
  -g, --gun string          Gun
  -h, --help                help for float
  -m, --max int             Max price
  -s, --stickers            Show stickers?
  -t, --top int             Top List (default 10)
```

> `--max` is based on cents, e.g `--max 125` corresponds to max price = 1.25$ 