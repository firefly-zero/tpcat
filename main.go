package main

import "github.com/firefly-zero/firefly-go/firefly"

const maxProgress = 2000

var (
	imgBg      firefly.Image
	imgCat     firefly.Image
	imgDone    firefly.Image
	imgHolder  firefly.Image
	imgsPaws   []firefly.Image
	imgsRoll   []firefly.Image
	imgsStripe []firefly.Image

	// The Y value of touchpad on the previous iteration.
	oldY *int

	// How much of toilet paper has been unrolled
	progress int

	afterProgress int

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
			diff := (*oldY - pad.Y) / 10
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
		if progress >= maxProgress {
			afterProgress++
		}
		pawsFrame++
		if pawsFrame >= len(imgsPaws) {
			pawsFrame = 0
		}
		stripeFrame++
		if stripeFrame >= len(imgsStripe) {
			stripeFrame = 0
		}
	}
}

func render() {
	if afterProgress >= 4 {
		firefly.DrawImage(imgDone, firefly.Point{})
		return
	}
	firefly.DrawImage(imgBg, firefly.Point{})
	firefly.DrawImage(imgCat, firefly.Point{})
	firefly.DrawImage(imgHolder, firefly.Point{})
	renderStripe()
	renderRoll()
	renderPaws()
}

func renderStripe() {
	if progress >= maxProgress {
		return
	}
	img := imgsStripe[stripeFrame]
	firefly.DrawImage(img, firefly.Point{})
}

func renderRoll() {
	idx := progress * 4 / maxProgress
	if progress >= maxProgress {
		idx = 4
	}
	img := imgsRoll[idx]
	firefly.DrawImage(img, firefly.Point{})
}

func renderPaws() {
	img := imgsPaws[pawsFrame]
	firefly.DrawImage(img, firefly.Point{})
}

func loadAssets() {
	imgBg = firefly.LoadFile("bg", nil).Must().Image()
	imgCat = firefly.LoadFile("cat", nil).Must().Image()
	imgDone = firefly.LoadFile("done", nil).Must().Image()
	imgHolder = firefly.LoadFile("holder", nil).Must().Image()
	imgsPaws = []firefly.Image{
		firefly.LoadFile("paws1", nil).Must().Image(),
		firefly.LoadFile("paws2", nil).Must().Image(),
		firefly.LoadFile("paws3", nil).Must().Image(),
		firefly.LoadFile("paws4", nil).Must().Image(),
	}
	imgsRoll = []firefly.Image{
		firefly.LoadFile("roll1", nil).Must().Image(),
		firefly.LoadFile("roll2", nil).Must().Image(),
		firefly.LoadFile("roll3", nil).Must().Image(),
		firefly.LoadFile("roll4", nil).Must().Image(),
		firefly.LoadFile("roll5", nil).Must().Image(),
	}
	imgsStripe = []firefly.Image{
		firefly.LoadFile("stripe1", nil).Must().Image(),
		firefly.LoadFile("stripe2", nil).Must().Image(),
		firefly.LoadFile("stripe3", nil).Must().Image(),
	}
}
