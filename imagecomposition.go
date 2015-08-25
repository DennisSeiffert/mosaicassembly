package mosaicassembly

import (
	"fmt"
	"golang.org/x/image/draw"
	"image"
	"image/png"
	"log"
	"math"
	"os"
	"strconv"
)

type Coordinate struct {
	x int
	y int
}

func mapImageFilenameToPositions(set map[string]map[Coordinate]bool, imagePos ImagePositions, 
	width int) {
	for _, pos := range imagePos {
		position, _ := strconv.Atoi(pos.Pos)
		x := position % width
		y := int(math.Floor(float64(position) / float64(width)))
		imageFilename := pos.Image;
		s, ok := set[imageFilename]
		if !ok {
			s = make(map[Coordinate]bool)
			set[imageFilename] = s
		}		
		_, ok = s[Coordinate{ x, y}]
		if !ok {
			set[imageFilename][Coordinate{ x, y}] = false
		}
	}
}

func assembleMosaic(imagePos ImagePositions) {
	width := 400
	dst := image.NewRGBA(image.Rect(0, 0, 400, 400))
	imagesAtCoordinatesMap := map[string]map[Coordinate]bool{}
	mapImageFilenameToPositions(imagesAtCoordinatesMap, imagePos, width)
	
	for imageName, coordinates := range imagesAtCoordinatesMap {
		func(){
		imgfile, err := os.Open("/home/dennis/Desktop/img/" + imageName)
		fmt.Printf("opened filename " +  imageName + "\n")
		defer imgfile.Close()
		if err != nil {
			fmt.Println("img.jpg file not found!")
		} else {
			img, _, _ := image.Decode(imgfile)
			for coordinate,_ := range coordinates {				
				fmt.Printf("%+v\n", coordinate)
				draw.NearestNeighbor.Scale(dst, image.Rect(coordinate.x, coordinate.y, coordinate.x+5, coordinate.y+5), img, img.Bounds(), draw.Over, nil)	
			}
		}
		}();
		
		save("out.png", dst)
	}

    save("out.png", dst)
}

func save(filename string, dst draw.Image) {
	fDst, err := os.Create(filename)

	if err != nil {
		log.Fatal(err)
	}

	defer fDst.Close()

	err = png.Encode(fDst, dst)
	if err != nil {
		log.Fatal(err)
	}
}
