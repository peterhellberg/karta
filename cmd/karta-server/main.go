package main

import (
	"errors"
	"flag"
	"html/template"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/peterhellberg/karta"
)

const defaultPort = "8100"

func main() {
	// Parse the command line flags
	flag.Parse()

	// Setup the logger
	l := log.New(os.Stdout, "", 0)

	// Get the port to bind the server to
	p := getEnv("PORT", defaultPort)

	// Start the server
	err := ListenAndServe(l, p, func(ctx *context) {
		http.Handle("/", baseHandler(ctx, kartaIndexHandler))
		http.Handle("/map.jpg", baseHandler(ctx, kartaJPEGHandler))
		http.Handle("/map.png", baseHandler(ctx, kartaPNGHandler))
	})

	if err != nil {
		panic(err)
	}
}

func getEnv(key, fallback string) string {
	v := os.Getenv(key)
	if v != "" {
		return v
	}
	return fallback
}

func getInt(r *http.Request, s string, def, min, max int) int {
	val, err := strconv.Atoi(r.URL.Query().Get(s))
	if err != nil || val < min || val > max {
		return def
	}

	return val
}

var index = template.Must(template.New("index").Parse(`<!DOCTYPE html>
<html>
	<head>
		<title>Karta</title>
		<link rel="stylesheet" href="//cdn.jsdelivr.net/g/pure@0.5.0(base-min.css+grids-min.css+forms-min.css)">
		<style>
			body {
				margin: 0;
				border-top: 8px solid #F20765;
				background: #fefefe;
				color: #191A1A;
				font: 16px/22px "HelveticaNeue-Light","HelveticaNeue",Helvetica,sans-serif;
			}

			main {
				padding: 2em;
			}

			h1 {
				font-size: 32px;
				line-height: 38px;
				margin-bottom: 22px;
			}

			h2 {
				color: #555;
				font-size: 26px;
				line-height: 28px;
				margin-bottom: 11px;
			}
		</style>
	</head>
	<body>
		<main>
			<h1>Karta</h1>

			<form action="/map.png" method="GET" class="pure-form pure-form-stacked">
				<fieldset>
					<legend>Generate</legend>
					<input id="w" name="w" size="3" placeholder="width">
					<input id="h" name="h" size="3" placeholder="height">
					<input id="c" name="c" size="3" placeholder="count">
					<input id="s" name="s" size="3" placeholder="seed">
					<button type="submit" class="pure-button pure-button-primary">Generate</button>
				</fieldset>
			</form>

			<h2>Examples</h2>
			<img src="/map.png?w=256&h=256&s=24&c=128"  width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=256"  width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=512"  width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=512"  width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=768"  width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=1024" width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=1152" width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=1280" width="256" height="256" />
			<img src="/map.png?w=256&h=256&s=24&c=2048" width="256" height="256" />
		</main>
	</body>
</html>
`))

func kartaIndexHandler(c *context, r *http.Request, w http.ResponseWriter) error {
	if r.Method != "GET" {
		return errors.New("no such handler")
	}

	return index.Execute(w, struct{}{})
}

func kartaJPEGHandler(c *context, r *http.Request, w http.ResponseWriter) error {
	if r.Method != "GET" {
		return errors.New("no such handler")
	}

	seed := getInt(r, "s", 3, 0, math.MaxUint32)
	width := getInt(r, "w", 512, 10, 5120)
	height := getInt(r, "h", 512, 10, 5120)
	count := getInt(r, "c", 2048, 3, 10000)
	iterations := getInt(r, "i", 1, 0, 16)

	// Seed the random number generator
	rand.Seed(int64(seed))

	// Create a new karta
	k := karta.New(width, height, count, iterations)

	if k.Generate() == nil {
		w.Header().Add("Content-Type", "image/jpeg")

		quality := getInt(r, "q", 100, 0, 100)

		opts := &jpeg.Options{Quality: quality}
		err := jpeg.Encode(w, k.Image, opts)
		if err != nil {
			return err
		}

		c.logf("Served JPG count=%v width=%v height=%v iterations=%v seed=%v quality=%v",
			count, width, height, iterations, seed, quality)
	} else {
		return errors.New("could not generate image")
	}

	return nil
}

func kartaPNGHandler(c *context, r *http.Request, w http.ResponseWriter) error {
	if r.Method != "GET" {
		return errors.New("no such handler")
	}

	seed := getInt(r, "s", 3, 0, math.MaxUint32)
	width := getInt(r, "w", 512, 10, 5120)
	height := getInt(r, "h", 512, 10, 5120)
	count := getInt(r, "c", 2048, 3, 10000)
	iterations := getInt(r, "i", 1, 0, 16)

	// Seed the random number generator
	rand.Seed(int64(seed))

	// Create a new karta
	k := karta.New(width, height, count, iterations)

	if k.Generate() == nil {
		w.Header().Add("Content-Type", "image/png")

		err := png.Encode(w, k.Image)
		if err != nil {
			return err
		}

		c.logf("Served PNG count=%v width=%v height=%v iterations=%v seed=%v",
			count, width, height, iterations, seed)
	} else {
		return errors.New("could not generate image")
	}

	return nil
}
