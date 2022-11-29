package noise

import "github.com/peterhellberg/gfx"

// Noise wraps simplexnoise.SimplexNoise
type Noise struct {
	*gfx.SimplexNoise
}

// New generates a new Noise instance
func New(seed int64) *Noise {
	return &Noise{gfx.NewSimplexNoise(seed)}
}
