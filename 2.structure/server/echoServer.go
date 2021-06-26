package main

import (
	"GoPl/2.structure/mandelbrotFractal"
	"GoPl/2.structure/surface"
	"fmt"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/gif", gifHandler)
	http.HandleFunc("/surface", surfaceHandler)
	http.HandleFunc("/fractal", mandelbrotHandler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%v", r.Header)
}
func gifHandler(response http.ResponseWriter, request *http.Request) {
	cycles, err := strconv.Atoi(request.URL.Query().Get("cycles"))
	if err != nil || cycles > 100 {
		cycles = 5
	}
	Lissajous(response, cycles)
}

func surfaceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	fmt.Fprintf(w, surface.PrintSurface())
}

func mandelbrotHandler(w http.ResponseWriter, r *http.Request) {
	png.Encode(w, mandelbrotFractal.PrintFractal())
}
