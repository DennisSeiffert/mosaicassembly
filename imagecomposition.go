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
		fmt.Printf("opening filename " +  imageName + "\n")		
		img, err := openImage(imageName)			
		if err != nil {
			fmt.Printf("Cannot open image %s", imageName)
		} else {			
			positionImage(coordinates, img, dst)
			save("out.png", dst)
		}				
	}   
}

func openImage(filename string) (image.Image, error) {
	imgfile, err := os.Open("/home/dennis/Desktop/img/" + filename)
	defer imgfile.Close()							
	if err != nil {
		return nil,err
	}
	img, _, _ := image.Decode(imgfile)
	return img,nil			
}

func positionImage(coordinates map[Coordinate]bool, img image.Image, dst *image.RGBA){
	for coordinate,_ := range coordinates {								
		draw.NearestNeighbor.Scale(dst, image.Rect(coordinate.x, coordinate.y, coordinate.x+5, coordinate.y+5), img, img.Bounds(), draw.Over, nil)	
	}
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
