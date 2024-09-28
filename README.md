# CSFloat CLI

A cron-based CSFloat deal finder CLI.

## Usage
Create a developer key in your [CSFloat profile](https://csfloat.com/profile) by navigating to the `Developers`
tab. 
Copy the generated API key into a file called `key` that lies in the same dir as this project.
Then run 
```bash
go build -o float
```

to build the binary. Then run it via

```bash
./float
```
This expects the `key` file with the CSFloat API key in the same dir.
Alternatively, you can store the API key in any file you want and reference it via

```bash
./float --keyfile path/to/the/keyfile
```

The CLI has some configuration options you might consider.


### Examples

```bash
# filter all offers by max price 25$
./float -m 2500 
```

```bash
# filter all offers in range 20$ - 25$
./float -m 2500 -n 2000 
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

```bash
# also show auctions
./float -a
```

```bash
# filter all offers in range 50% - 95$ with stickers, include auctions and run in cron mode
./float -m 9500 -n 5000 -a -s --cron
```

All Options:
```
Usage:
  float [flags]

Flags:
  -a, --auctions         Also check auctions
  -c, --category int     Item category - [0: Any, 1: Normal, 2: Stattrak, 3: Souvenir] (default 1)
      --cron             Enable cron mode
  -d, --discount float   Min discount percentage (default 5)
  -h, --help             help for float
  -f, --keyfile string   The location of your API key file
  -k, --keyword string   The keyword. e.g a Skin Name like 'Asiimov' or 'Dragon Lore'
  -m, --max int          Max price in cents
  -n, --min int          Min price in cents
  -s, --stickers         Show stickers? (Default off)
  -t, --top int          Top List (default 10)
```

> `--max / --min` is based on cents, e.g `--max 125` corresponds to max price = 1.25$ 
> while `--discount` is a percentage, e.g `--discount 12` corresponds to 12%