package Services_GIS

import (
	// "bufio"
	"encoding/csv"
	"fmt"
	// "github.com/disintegration/imaging"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"strconv"
	"strings"
)

// Lat/Long to pixels
var yStep, xStep float64
var centerPixels image.Point
var globalImg *image.RGBA

// ------------------------------------------- Main ------------------------------------------- //

func main() {
	imgFile := loadImage("blank_world_map.jpg")
	initMap(&imgFile, image.Point{180, -60}, image.Point{-180, 90})

	img := image.NewRGBA(imgFile.Bounds())
	draw.Draw(img, imgFile.Bounds(), imgFile, imgFile.Bounds().Min, draw.Src)

	globalImg = img
	csvFile, err := os.Open("cities.csv")
	checkError(err)
	cities, err := csv.NewReader(csvFile).ReadAll()
	checkError(err)
	var location image.Point
	for i, _ := range cities {
		lon, _ := strconv.ParseFloat(cities[i][2], 32)
		lat, _ := strconv.ParseFloat(cities[i][3], 32)
		location = mapToPixels(lat, lon)
		drawDot(img, location, 3, true)
	}
	fmt.Println(cities)
	saveImage(img, "OUTPUT.jpg")
}

func loadImage(fileName string) image.Image {
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	i := strings.Index(fileName, ".")
	if i == -1 {
		fmt.Println("Could not load image.")
		return nil
	}
	end := fileName[i+1:]

	var img image.Image
	switch end {
	case "png":
		img, err = png.Decode(file)
	case "jpg", "jpeg":
		img, err = jpeg.Decode(file)
	}
	checkError(err)
	return img
}

func saveImage(img image.Image, fileName string) {
	f, err := os.Create(fileName)
	checkError(err)
	// img = imaging.Rotate(img, 180.0, color.RGBA{0, 0, 0, 0})
	err = jpeg.Encode(f, img, nil)
	checkError(err)
}

func initMap(img *image.Image, topleft, bottomright image.Point) {
	bounds := (*img).Bounds()
	yStep = math.Abs((float64)(bounds.Max.Y-bounds.Min.Y) / (float64)(topleft.Y-bottomright.Y))
	xStep = math.Abs((float64)(bounds.Max.X-bounds.Min.X) / (float64)(topleft.X-bottomright.X))
	yCenterPixels := (int)((float64)(-topleft.Y) * yStep)
	xCenterPixels := (int)((float64)(topleft.X) * xStep)
	centerPixels = image.Point{xCenterPixels, yCenterPixels}
}

func mapToPixels(longitude, latitude float64) image.Point {
	x := centerPixels.X - (int)(xStep*(float64)(longitude))
	y := centerPixels.Y + (int)(yStep*(float64)(latitude))
	return image.Point{x, y}
}

func drawDot(img *image.RGBA, location image.Point, size int, filled bool) {
	if filled {
		for i := -size; i < size; i++ {
			for j := -size; j < size; j++ {
				img.Set(location.X+i, location.Y+j, color.RGBA{200, 50, 100, 255})
			}
		}
	} else {
		for i := -size; i < size; i++ {
			img.Set(location.X+i, location.Y+size, color.RGBA{200, 50, 100, 255})
			img.Set(location.X+i, location.Y-size, color.RGBA{200, 50, 100, 255})
			img.Set(location.X+size, location.Y+i, color.RGBA{200, 50, 100, 255})
			img.Set(location.X-size, location.Y-i, color.RGBA{200, 50, 100, 255})
		}
	}
}

// ------------------------------------------- Utility ------------------------------------------- //

// Respond to errors
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
