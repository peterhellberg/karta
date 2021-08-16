# Karta

Just playing around with [Voronoi diagrams](http://en.wikipedia.org/wiki/Voronoi_diagram).

[![GoDoc](https://godoc.org/github.com/peterhellberg/karta?status.svg)](https://godoc.org/github.com/peterhellberg/karta)

![Karta progress 012](https://assets.c7.se/skitch/karta_progress_012-20140823-155942.png)

The goal is to create something similar to what is described in the article
[Polygonal Map Generation for Games](http://www-cs-students.stanford.edu/~amitp/game-programming/polygon-map-generation/)

## Installation

```bash
go get github.com/peterhellberg/karta/cmd/karta
```

## Development

I spend most of my time in [iTerm 2](http://iterm2.com/) and this project is no different.

This is how I preview the map in a split pane:

```bash
rerun -p "**/*.go" -c -x -b -- \
"go run cmd/karta/main.go -width 55 -height 55 && aimg -w 55 karta.png"
```

![Karta development](https://assets.c7.se/skitch/karta_development_iterm_vim_aimg-20140812-204452.png)

## Usage

The project includes two binaries `karta` and `karta-server`, the latter serving both PNG and JPEG.

### Command line arguments

```
Usage of karta:
  -count=2048: The number of sites in the voronoi diagram
  -height=512: The height of the map in pixels
  -iterations=1: The number of iterations of Lloyd's algorithm to run (max 16)
  -output="karta.png": Output filename
  -seed=3: The starting seed for the map generator
  -show=false: Show generated map using Preview.app
  -width=512: The width of the map in pixels
```

### Query string parameters

  - **s** - size
  - **w** - width
  - **h** - height
  - **c** - count
  - **i** - iterations

## Progress

First i just plotted out random dots:

![Karta progress 001](https://assets.c7.se/skitch/karta_progress_001-20140812-223244.png)

Then drew a voroni diagram:

![Karta progress 002](https://assets.c7.se/skitch/karta_progress_002-20140812-223300.png)

Found the centroids:

![Karta progress 003](https://assets.c7.se/skitch/karta_progress_003-20140812-223327.png)

Ran the diagram through [Lloyd's algorithm](http://en.wikipedia.org/wiki/Lloyd%27s_algorithm):

![Karta progress 004](https://assets.c7.se/skitch/karta_progress_004-20140812-223344.png)

A few iterations later:

![Karta progress 005](https://assets.c7.se/skitch/karta_progress_005-20140812-223400.png)

Animating 0-16 iterations:

![Lloyd Relaxation 0-16 iterations](http://assets.c7.se/viz/lloyd-relaxation.gif)

Started coloring the map according to the distance from the center:

![Karta progress 006](https://assets.c7.se/skitch/karta_progress_006-20140812-223423.png)

Added some randomness:

![Karta progress 007](https://assets.c7.se/skitch/karta_progress_007-20140812-223203.png)

Added a different types of elevation:

![Karta progress 008](https://assets.c7.se/skitch/karta_progress_008-20140813-005713.png)

Removed the centroid markers:

![Karta progress 009](https://assets.c7.se/skitch/karta_progress_009-20140813-005845.png)

Started working on using [Simplex noise](http://en.wikipedia.org/wiki/Simplex_noise) for
island shape and elevation:

![Karta progress 010](https://assets.c7.se/skitch/karta_progress_010-20140813-010005.png)

Elevation based on noise + distance from center of map:

![Karta progress 011](https://assets.c7.se/skitch/karta_progress_011-20140816-013747.png)

Tweaked noise and yellow beaches:

![Karta progress 012](https://assets.c7.se/skitch/karta_progress_012-20140823-155942.png)

## License

**The MIT License (MIT)**

Copyright (C) 2014 [Peter Hellberg](https://c7.se/)

> Permission is hereby granted, free of charge, to any person obtaining
> a copy of this software and associated documentation files (the "Software"),
> to deal in the Software without restriction, including without limitation
> the rights to use, copy, modify, merge, publish, distribute, sublicense,
> and/or sell copies of the Software, and to permit persons to whom the
> Software is furnished to do so, subject to the following conditions:
>
> The above copyright notice and this permission notice shall be included
> in all copies or substantial portions of the Software.
>
> THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
> EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES
> OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
> IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
> DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
> TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
> OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
