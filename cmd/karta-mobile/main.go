package main

import (
	"image"
	"log"
	"time"

	"github.com/peterhellberg/karta"

	"golang.org/x/mobile/app"
	"golang.org/x/mobile/event/lifecycle"
	"golang.org/x/mobile/event/paint"
	"golang.org/x/mobile/exp/f32"
	"golang.org/x/mobile/exp/sprite"
	"golang.org/x/mobile/exp/sprite/clock"
	"golang.org/x/mobile/exp/sprite/glsprite"
)

var (
	start     = time.Now()
	lastClock = clock.Time(-1)
	eng       = glsprite.Engine(nil)
	scene     *sprite.Node
	kn        *sprite.Node
)

func main() {
	app.Main(func(a app.App) {
		for e := range a.Events() {
			switch e := a.Filter(e).(type) {
			case lifecycle.Event:
				// ...
			case paint.Event:
				log.Print("Call OpenGL here.")

				if scene == nil {
					loadScene(40)
				}

				now := clock.Time(time.Since(start) * 60 / time.Second)
				if now == lastClock {
					return
				}
				lastClock = now

				eng.Render(scene, now)

				a.Publish()
			}
		}
	})

	app.Main(
	//app.Callbacks{
	//Draw: func() {
	//	if scene == nil {
	//		loadScene(40)
	//	}

	//	now := clock.Time(time.Since(start) * 60 / time.Second)
	//	if now == lastClock {
	//		return
	//	}
	//	lastClock = now

	//	eng.Render(scene, now)
	//},
	//Touch: func(e touch.Event) {
	//	if e.Type != touch.TypeBegin {
	//		return
	//	}

	//	eng.SetSubTex(kn, loadTexture((int(e.Loc.X) * (int(e.Loc.Y/25) + 1))))
	//},
	//}
	)
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
