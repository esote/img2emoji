package main

import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"github.com/esote/img2emoji"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("usage: img2emoji file")
		os.Exit(1)
	}

	var (
		img image.Image
		err error
	)
	if os.Args[1] == "-" {
		img, err = openStdin()
	} else {
		img, err = open(os.Args[1])
	}
	if err != nil {
		log.Fatal(err)
	}

	out, err := img2emoji.Convert(img, img2emoji.DefaultMapping)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(out)
}

func open(in string) (image.Image, error) {
	file, err := os.Open(in)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	return img, err
}

func openStdin() (image.Image, error) {
	img, _, err := image.Decode(os.Stdin)
	return img, err
}
