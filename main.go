package main

import "github.com/firefly-zero/firefly-go/firefly"

const maxProgress = 8000

var (
	imgBg      firefly.Image
	imgCat     firefly.Image
	imgDone    firefly.Image
	imgHolder  firefly.Image
	imgsPaws   [4]firefly.Image
	imgsRoll   [5]firefly.Image
	imgsStripe [3]firefly.Image

	// The Y value of touchpad on the previous iteration.
	oldY *int

	// How much of toilet paper has been unrolled
	progress int

	afterProgress int
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
	pad, touched := firefly.ReadPad(firefly.Combined)
	if touched {
		if oldY != nil {
			diff := (*oldY - pad.Y) / 10
			if diff > 0 {
				updateDiff(diff)
			}
		}
		oldY = &pad.Y
	} else {
		oldY = nil
	}
}

func updateDiff(diff int) {
	progress += diff
	if progress >= maxProgress {
		afterProgress++
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
	idx := (progress / 40) % len(imgsStripe)
	img := imgsStripe[idx]
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
	idx := (progress / 40) % len(imgsPaws)
	img := imgsPaws[idx]
	firefly.DrawImage(img, firefly.Point{})
}

func loadAssets() {
	imgBg = loadImage("bg")
	imgCat = loadImage("cat")
	imgDone = loadImage("done")
	imgHolder = loadImage("holder")
	imgsPaws = [...]firefly.Image{
		loadImage("paws1"),
		loadImage("paws2"),
		loadImage("paws3"),
		loadImage("paws4"),
	}
	imgsRoll = [...]firefly.Image{
		loadImage("roll1"),
		loadImage("roll2"),
		loadImage("roll3"),
		loadImage("roll4"),
		loadImage("roll5"),
	}
	imgsStripe = [...]firefly.Image{
		loadImage("stripe1"),
		loadImage("stripe2"),
		loadImage("stripe3"),
	}
}

func loadImage(name string) firefly.Image {
	return firefly.LoadFile(name, nil).Must().Image()
}
