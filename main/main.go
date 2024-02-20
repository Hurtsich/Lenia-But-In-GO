package main

import (
	"Lenia/organism"
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"os"
)

func main() {
	billy := organism.NewOrganism(20)
	fmt.Println("Billy ready !")
	createGIF(billy, "billy")
}

func createGIF(o *organism.Organism, imageName string) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 100; i++ {
		fmt.Printf("Year: %v\n", i)
		delays = append(delays, 7)
		photo := o.Photo()
		images = append(images, photo)
		o.Breathe(10)
	}
	defer o.Die()

	f, err := os.Create("../data/" + imageName + ".gif")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	err = gif.EncodeAll(w, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	if err != nil {
		fmt.Println(err)
	}
}
