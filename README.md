# csfloat

A cron-based csfloat deal finder.

## Usage
Create a developer key in your [CSFloat profile](https://csfloat.com/profile) by navigating to the `Developers`
tab. 
Copy the generated API key into a file called `key` that lies in the same dir as this project


```bash
go build -o float
```

Then run it via


```bash
./float
```

### Examples

```bash
# filter all offers by max price 25$
./float -m 2500 
```

```bash
# filter all offers by max price 25$ (TOP 5 only)
./float -m 2500 --top 5
```

```bash
# all offers with a discount of at least 15% 
./float -d 15
```

```bash
# also show offers with stickers on it 
./float --stickers
```

Usage per CLI:
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