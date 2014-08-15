package noise

import "bitbucket.org/s_l_teichmann/simplexnoise"

type Noiser interface {
	Noise2D() float64
}

type Noise struct {
	*simplexnoise.SimplexNoise
}

func New(seed int64) *Noise {
	return &Noise{simplexnoise.NewSimplexNoise(seed)}
}
