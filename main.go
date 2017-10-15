package main

import (
	"fmt"
	"math"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

var (
	windowSize   = 1000.0
	imd          *imdraw.IMDraw
	points       []pixel.Vec
	lineSize     = 2.0
	circleRadius = 4.0

	numPoints     int
	multiplier    int
	maxPoints     = 1000
	maxMultiplier = 1000
)

func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Modular Multiplication Around A Circle by Christopher Silva",
		Bounds: pixel.R(0, 0, windowSize, windowSize),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	imd = imdraw.New(nil)
	imd.Color = colornames.Black
	imd.EndShape = imdraw.RoundEndShape

	numPoints = 10
	multiplier = 2

	plot()

	replot := false
	for !win.Closed() {
		win.Clear(colornames.White)

		// q/a = +/- 100 points
		if win.JustPressed(pixelgl.KeyQ) {
			if numPoints < maxPoints {
				numPoints += 100
			} else {
				numPoints = maxPoints
			}
			replot = true
		}
		if win.JustPressed(pixelgl.KeyA) {
			if numPoints >= 100 {
				numPoints -= 100
			} else {
				numPoints = 0
			}
			replot = true
		}
		// w/s = +/- 10  points
		if win.JustPressed(pixelgl.KeyW) {
			if numPoints < maxPoints {
				numPoints += 10
			} else {
				numPoints = maxPoints
			}
			replot = true
		}
		if win.JustPressed(pixelgl.KeyS) {
			if numPoints >= 10 {
				numPoints -= 10
			} else {
				numPoints = 0
			}
			replot = true
		}
		// e/d = +/- 1   points
		if win.JustPressed(pixelgl.KeyE) {
			if numPoints < maxPoints {
				numPoints++
			}
			replot = true
		}
		if win.JustPressed(pixelgl.KeyD) {
			if numPoints > 0 {
				numPoints--
			}
			replot = true
		}

		// u/j = +/- 100 multiplier
		if win.JustPressed(pixelgl.KeyU) {
			if multiplier < maxMultiplier {
				multiplier += 100
			} else {
				multiplier = maxMultiplier
			}
			replot = true
		}
		if win.JustPressed(pixelgl.KeyJ) {
			if multiplier >= 100 {
				multiplier -= 100
			} else {
				multiplier = 0
			}
			replot = true
		}
		// i/k = +/- 10  multiplier
		if win.JustPressed(pixelgl.KeyI) {
			if multiplier < maxMultiplier {
				multiplier += 10
			} else {
				multiplier = maxMultiplier
			}
			replot = true
		}
		if win.JustPressed(pixelgl.KeyK) {
			if multiplier >= 10 {
				multiplier -= 10
			} else {
				multiplier = 0
			}
			replot = true
		}
		// o/l = +/- 1   multiplier
		if win.JustPressed(pixelgl.KeyO) {
			if multiplier < maxMultiplier {
				multiplier++
			}
			replot = true
		}
		if win.JustPressed(pixelgl.KeyL) {
			if multiplier > 0 {
				multiplier--
			}
			replot = true
		}
		if replot {
			fmt.Printf("Points: %d Multiplier: %d\n", numPoints, multiplier)
			plot()
			replot = false
		}

		imd.Draw(win)

		win.Update()
	}
}

func plot() {
	imd.Clear()

	delta := (2 * math.Pi) / float64(numPoints)

	points = make([]pixel.Vec, numPoints)

	for i := 0; i < numPoints; i++ {
		points[i].X = convertCoord(math.Sin(delta * float64(i)))
		points[i].Y = convertCoord(math.Cos(delta * float64(i)))
		drawCircle(points[i])
	}

	for i := 0; i < numPoints; i++ {
		drawLine(points[i], points[i*multiplier%numPoints])
	}
}

func convertCoord(val float64) float64 {
	return (((val - -1) * (windowSize - 6 - 6)) / (1 - -1)) + 6
}

func drawCircle(center pixel.Vec) {
	imd.Push(center)
	imd.Circle(circleRadius, lineSize)
}

func drawLine(from, to pixel.Vec) {
	imd.Push(from, to)
	imd.Line(lineSize)
}

func main() {
	pixelgl.Run(run)
}
