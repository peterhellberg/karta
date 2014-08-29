package main

import (
	"errors"
	"flag"
	"image/jpeg"
	"image/png"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

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
		http.Handle("/", baseHandler(ctx, kartaHandler))
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

func kartaHandler(c *context, r *http.Request, w http.ResponseWriter) error {
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
		if strings.Contains(r.URL.Path, ".jpg") {
			w.Header().Add("Content-Type", "image/jpeg")

			quality := getInt(r, "q", 100, 0, 100)

			opts := &jpeg.Options{Quality: quality}
			err := jpeg.Encode(w, k.Image, opts)
			if err != nil {
				return err
			}

			c.logf("Served JPEG: quality=%v count=%v width=%v height=%v iterations=%v seed=%v",
				quality, count, width, height, iterations, seed)
		} else {
			w.Header().Add("Content-Type", "image/png")

			err := png.Encode(w, k.Image)
			if err != nil {
				return err
			}

			c.logf("Served PNG: count=%v width=%v height=%v iterations=%v seed=%v",
				count, width, height, iterations, seed)
		}

	} else {
		return errors.New("could not generate image")
	}

	return nil
}
