package noise

import "bitbucket.org/s_l_teichmann/simplexnoise"

// Noise wraps simplexnoise.SimplexNoise
type Noise struct {
	*simplexnoise.SimplexNoise
}

// New generates a new Noise instance
func New(seed int64) *Noise {
	return &Noise{simplexnoise.NewSimplexNoise(seed)}
}
