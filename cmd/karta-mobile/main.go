package main

import (
	"image"
	"log"
	"time"

	"github.com/peterhellberg/karta"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event"
	"golang.org/x/mobile/f32"
	"golang.org/x/mobile/sprite"
	"golang.org/x/mobile/sprite/clock"
	"golang.org/x/mobile/sprite/glsprite"
)

var (
	start     = time.Now()
	lastClock = clock.Time(-1)
	eng       = glsprite.Engine()
	scene     *sprite.Node
	kn        *sprite.Node
)

func main() {
	app.Run(app.Callbacks{
		Draw: func() {
			if scene == nil {
				loadScene(40)
			}

			now := clock.Time(time.Since(start) * 60 / time.Second)
			if now == lastClock {
				return
			}
			lastClock = now

			eng.Render(scene, now)
		},
		Touch: func(e event.Touch) {
			if e.Type != event.TouchStart {
				return
			}

			eng.SetSubTex(kn, loadTexture((int(e.Loc.X) * (int(e.Loc.Y/25) + 1))))
		},
	})
}

func loadScene(count int) {
	kn = &sprite.Node{}

	scene = &sprite.Node{}
	scene.AppendChild(kn)

	eng.Register(kn)
	eng.Register(scene)
	eng.SetSubTex(kn, loadTexture(count))
	eng.SetTransform(kn, f32.Affine{
		{212, 0, 0},
		{0, 212, 0},
	})
}

func loadTexture(count int) sprite.SubTex {
	t, err := eng.LoadTexture(karta.Image(512, 512, count, 4))
	if err != nil {
		log.Fatal(err)
	}

	return sprite.SubTex{t, image.Rect(0, 0, 512, 512)}
}
