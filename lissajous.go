package main

//генерирует анимированное изображение из случайных фигур Лиссажу

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var pallete = []color.Color{color.White, color.Black}

const (
	whiteIdx = 0 // Первый цвет палитры
	blackIdx = 1 // Следующий цвет палитры
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // количество полный колебаний Х
		res     = 0.001 // угловое разрешение
		size    = 100   // канва изображения охватывает [size... + size]
		nframes = 64    // количество кадров анимации
		delay   = 8     // задержка между кадрами, единица = 10мс
	)
	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0 // относительная частота колебаний у
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // разность фаз
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIdx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
