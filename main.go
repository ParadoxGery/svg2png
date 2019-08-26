package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/rogpeppe/misc/svg"
)

func main() {
	inName := flag.String("in", "", "svg to read from")
	outName := flag.String("out", "", "name for the png")
	flag.Parse()

	if *inName == "" || *outName == "" {
		log.Fatal("the in and out parameter need to be set")
	}

	file, err := os.Open(*inName)
	if err != nil {
		log.Fatal("failed reading svg", err)
	}
	defer file.Close()

	size := image.Point{1000, 1000}
	des, err := svg.Render(file, size)

	out, err := os.Create(*outName)
	if err != nil {
		log.Fatal("failed creating png file", err)
	}

	err = png.Encode(out, des)
	if err != nil {
		log.Fatal("failed writing png to file", err)
	}

	log.Println("done")
}
