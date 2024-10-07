package main

import (
	"Lenia/organism"
	"bufio"
	"fmt"
	"image"
	"image/gif"
	"os"
	"time"
)

func main() {
	start := time.Now()
	billy := organism.NewOrganism(200)
	fmt.Println("Billy ready !")
	createGIF(billy, "billy")
	hector := organism.NewOrganism(200)
	fmt.Println("Hector ready !")
	createGIF(hector, "hector")
	franck := organism.NewOrganism(200)
	fmt.Println("franck ready !")
	createGIF(franck, "franck")
	edward := organism.NewOrganism(200)
	fmt.Println("Edward ready !")
	createGIF(edward, "edward")
	bob := organism.NewOrganism(200)
	fmt.Println("Bob ready !")
	createGIF(bob, "bob")
	fmt.Printf("Ellapsed %f", time.Now().Sub(start).Seconds())
}

func createGIF(o *organism.Organism, imageName string) {
	var images []*image.Paletted
	var delays []int

	for i := 0; i < 100; i++ {
		fmt.Printf("Year: %v\n", i)
		delays = append(delays, 12)
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
