package main

import "github.com/firefly-zero/firefly-go/firefly"

const maxProgress = 8000

var (
	imgBg      firefly.Image
	imgCat     firefly.Image
	imgDone    firefly.Image
	imgHolder  firefly.Image
	imgHolder2 firefly.Image
	imgPaws    firefly.Image
	imgRoll    firefly.Image
	imgStripe  firefly.Image

	// The Y value of touchpad on the previous iteration.
	oldY *int

	// How much of toilet paper has been unrolled
	progress int

	// If the player made any progess on this update
	moving bool

	pawsFrame   int
	stripeFrame int
)

func init() {
	firefly.Boot = boot
	firefly.Update = update
	firefly.Render = render
}

func boot() {
	loadAssets()
}

func update() {
	moving = false
	pad, touched := firefly.ReadPad(firefly.Combined)
	if touched {
		if oldY != nil {
			diff := (pad.Y - *oldY) / 10
			if diff > 0 {
				progress += diff
				moving = true
			}
		}
		oldY = &pad.Y
	} else {
		oldY = nil
	}

	if moving {
		pawsFrame++
		if pawsFrame >= 8 {
			pawsFrame = 0
		}
		stripeFrame++
		if stripeFrame >= 8 {
			stripeFrame = 0
		}
	}
}

func render() {
	firefly.DrawImage(imgBg, firefly.Point{})
	firefly.DrawImage(imgCat, firefly.Point{X: 166, Y: 39})
	firefly.DrawImage(imgHolder, firefly.Point{X: 63, Y: 68})
	renderStripe()
	renderRoll()
	firefly.DrawImage(imgHolder2, firefly.Point{X: 63, Y: 68})
	renderPaws()
}

func renderStripe() {
	if progress >= maxProgress {
		return
	}
	x := 53 * (pawsFrame / 4)
	sub := imgStripe.Sub(firefly.Point{X: x}, firefly.Size{W: 53, H: 240})
	firefly.DrawSubImage(sub, firefly.Point{X: 93, Y: 71})
}

func renderRoll() {
	const width = 84
	x := width * (progress / (maxProgress / 4))
	if progress >= maxProgress {
		x = width * 4
	}
	sub := imgRoll.Sub(firefly.Point{X: x}, firefly.Size{W: width, H: 76})
	firefly.DrawSubImage(sub, firefly.Point{X: 70, Y: 31})
}

func renderPaws() {
	x := 69 * (pawsFrame / 4)
	sub := imgPaws.Sub(firefly.Point{X: x}, firefly.Size{W: 69, H: 82})
	firefly.DrawSubImage(sub, firefly.Point{X: 125, Y: 51})
}

func loadAssets() {
	imgBg = firefly.LoadFile("bg", nil).Must().Image()
	imgCat = firefly.LoadFile("cat", nil).Must().Image()
	imgDone = firefly.LoadFile("done", nil).Must().Image()
	imgHolder = firefly.LoadFile("holder", nil).Must().Image()
	imgHolder2 = firefly.LoadFile("holder2", nil).Must().Image()
	imgPaws = firefly.LoadFile("paws", nil).Must().Image()
	imgRoll = firefly.LoadFile("roll", nil).Must().Image()
	imgStripe = firefly.LoadFile("stripe", nil).Must().Image()
}
