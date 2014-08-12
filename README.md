# Karta

Just playing around with [Voronoi diagrams](http://en.wikipedia.org/wiki/Voronoi_diagram).

![Lloyd Relaxation 0-16 iterations](http://assets.c7.se/viz/lloyd-relaxation.gif)

The goal is to create something similar to what is described in the article
[Polygonal Map Generation for Games](http://www-cs-students.stanford.edu/~amitp/game-programming/polygon-map-generation/)

## Installation

```bash
go get github.com/peterhellberg/karta
```

## Development

I spend most of my time in [iTerm 2](http://iterm2.com/) and this project is no different.

This is how I preview the map in a split pane:

```bash
rerun -p "**/*.go" -c -x -b -- "go run main.go -width 55 -height 55 -count 3 && aimg -w 55 karta.png"
```

## License

**The MIT License (MIT)**

Copyright (C) 2014 [Peter Hellberg](http://c7.se/)

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
