package main

import "github.com/firefly-zero/firefly-go/firefly"

var (
	imgBg      firefly.Image
	imgCat     firefly.Image
	imgDone    firefly.Image
	imgHolder  firefly.Image
	imgHolder2 firefly.Image
	imgPaws    firefly.Image
	imgRoll    firefly.Image
	imgStripe  firefly.Image
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
	// ...
}

func render() {
	firefly.DrawImage(imgBg, firefly.Point{})
	firefly.DrawImage(imgCat, firefly.Point{X: 166, Y: 39})
	firefly.DrawImage(imgHolder, firefly.Point{X: 63, Y: 68})

	stripe := imgStripe.Sub(firefly.Point{}, firefly.Size{W: 53, H: 240})
	firefly.DrawSubImage(stripe, firefly.Point{X: 93, Y: 71})

	roll := imgRoll.Sub(firefly.Point{}, firefly.Size{W: 84, H: 76})
	firefly.DrawSubImage(roll, firefly.Point{X: 70, Y: 31})

	firefly.DrawImage(imgHolder2, firefly.Point{X: 63, Y: 68})

	paws := imgPaws.Sub(firefly.Point{}, firefly.Size{W: 69, H: 82})
	firefly.DrawSubImage(paws, firefly.Point{X: 125, Y: 51})

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
