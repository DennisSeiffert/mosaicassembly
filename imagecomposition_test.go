package mosaicassembly

import (
    "testing"
)

func Test_mapImageFilenameToPositions(t *testing.T) {
	assertableSet := map[string]map[Coordinate]bool{
		"imageA" : map[Coordinate]bool{ Coordinate{0,0} : false, Coordinate{5,0} : false },
		"imageB" : map[Coordinate]bool{ Coordinate{1,1} : false, Coordinate{5,1} : false },
	}
	resultSet := map[string]map[Coordinate]bool{}
	imagePositions := ImagePositions{
		ImagePosition{"0", "imageA"}, 
		ImagePosition{"5", "imageA"},
		ImagePosition{"7", "imageB"}, 
		ImagePosition{"11", "imageB"},
		}
	
	mapImageFilenameToPositions(resultSet, imagePositions, 6)
	
	if len(resultSet) != len(assertableSet) {
		t.Error("length is different in result set.")
	}
	for key, value := range resultSet {
		coordinates, ok := assertableSet[key]
		if !ok { t.Error("Cannot find key " + key + " in assertion set.") }
		for coordinate, _ := range coordinates{
			_, ok = value[coordinate]
			if !ok {
				t.Error("cannot find coordinate: %v",coordinate)
			}
		} 
	}
}

