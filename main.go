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
	numPointsDelta := 0
	multiplierDelta := 0
	for !win.Closed() {
		win.Clear(colornames.White)

		if win.JustPressed(pixelgl.KeyEscape) {
			win.SetClosed(true)
		}

		numPointsDelta = 0
		multiplierDelta = 0

		// q/a = +/- 100 points
		if win.JustPressed(pixelgl.KeyQ) {
			numPointsDelta += 100
		}
		if win.JustPressed(pixelgl.KeyA) {
			numPointsDelta -= 100
		}
		// w/s = +/- 10  points
		if win.JustPressed(pixelgl.KeyW) {
			numPointsDelta += 10
		}
		if win.JustPressed(pixelgl.KeyS) {
			numPointsDelta -= 10
		}
		// e/d = +/- 1   points
		if win.JustPressed(pixelgl.KeyE) {
			numPointsDelta++
		}
		if win.JustPressed(pixelgl.KeyD) {
			numPointsDelta--
		}

		// u/j = +/- 100 multiplier
		if win.JustPressed(pixelgl.KeyU) {
			multiplierDelta += 100
		}
		if win.JustPressed(pixelgl.KeyJ) {
			multiplierDelta -= 100
		}
		// i/k = +/- 10  multiplier
		if win.JustPressed(pixelgl.KeyI) {
			multiplierDelta += 10
		}
		if win.JustPressed(pixelgl.KeyK) {
			multiplierDelta -= 10
		}
		// o/l = +/- 1   multiplier
		if win.JustPressed(pixelgl.KeyO) {
			multiplierDelta++
		}
		if win.JustPressed(pixelgl.KeyL) {
			multiplierDelta--
		}
		if numPointsDelta != 0 {
			numPoints += numPointsDelta
			if numPoints > 1000 {
				numPoints = 1000
			} else if numPoints < 0 {
				numPoints = 0
			}
			replot = true
		}
		if multiplierDelta != 0 {
			multiplier += multiplierDelta
			if multiplier > 1000 {
				multiplier = 1000
			} else if multiplier < 0 {
				multiplier = 0
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
