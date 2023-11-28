package main

import (
	"fmt"
	"net/http"
	"log"
	"sync"
	"image"
	"image/gif"
	"image/color"
	"io"
	"math"
	"math/rand"
)

var mu sync.Mutex
var count int

var pallete = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main()  {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request){ 
		lissajous(w)
	})
	log.Fatal(http.ListenAndServe("localhost:8080",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL Path: %q\n", r.URL.Path)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Heaher: Key:%s Value:%s\n",k,v)
	}


	// parse query string
	if err := r.ParseForm(); err != nil {
		fmt.Println("error parsing query string")
	}

	for k, v := range r.Form {
		fmt.Fprintf(w, "FORM Key: %s, Value: %s\n", k, v)
	}

	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "Remore Addr: %s\n", r.RemoteAddr)
	mu.Lock()
	count++
	mu.Unlock()
}

func counter (w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "counter: %d\n", count)
	mu.Unlock()
}

func lissajous(out io.Writer) {

	const (
		cycles = 5
		res = 0.001
		size = 100
		nframes = 64
		delay = 8 // delay of frames in 10ms units
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0

	for i:=0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2 * size + 1, 2 * size + 1)
		img := image.NewPaletted(rect, pallete)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
			blackIndex)
		}
		phase+=0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}