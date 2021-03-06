package Services_GIS

import (
	// "bufio"
	. "Framework/Framework_Definitions"
	"fmt"
	"github.com/disintegration/imaging"
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

const ()

// ------------------------------------------- Datasets ------------------------------------------- //

type dataset struct {
	name    string
	indices map[string]int
	data    [][]string
}

// A range of data, half open. [lower, upper)
type dataRange struct {
	lower float64
	upper float64
}

// dataRange constructor
func NewRange(lower, upper float64) dataRange {
	if lower == -1 {
		lower = 0
	}
	if upper == -1 {
		upper = math.MaxFloat64
	}
	return dataRange{lower, upper}
}

// Format for data requests
type dataRequest struct {
	datasetName string
	requested   map[string][]dataRange
}

// // Specific to city dataset
// // https://simplemaps.com/data/world-cities
// var cityIndex = map[string]int{
// 	"name":       1,
// 	"longitude":  2,
// 	"latitude":   3,
// 	"population": 4,
// 	"country":    5,
// }

// ------------------------------------------- Public ------------------------------------------- //

// See GIS_Service_datasets.go for details about specific datasets. Put user provided dataset details in that file

type GISService struct {
	image        *image.RGBA
	yStep, xStep float64
	centerPixel  image.Point
	datasets     map[string]*dataset
}

// Initialize the service
func (s *GISService) Init() string {
	// Load blank map
	imgFile := loadImage("Data\\blank_world_map.jpg")
	s.initMap(&imgFile, image.Point{180, -60}, image.Point{-180, 90})
	s.image = image.NewRGBA(imgFile.Bounds())
	draw.Draw(s.image, imgFile.Bounds(), imgFile, imgFile.Bounds().Min, draw.Src)

	// Load datasets
	s.loadCityDataset()
	s.loadCountryDataset()

	return "GISService"
}

// Respond to events posted to the service
func (s *GISService) RunFunction(event Event, sendChannel chan Event) Event {
	returnEvent := NewEvent(NONE, "", "")

	switch eventType := event.Type; eventType {

	case GLOBAL_START:

	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED

	case ADD_DATA:

	case REQUEST_DATASET_NAMES:
		returnEvent.Type = DATASET_CHANGE
		returnEvent.Parameter = &s.datasets

	case GENERATE_MAP:
		s.generateMap(event.Parameter.(**[]dataRequest))
		saveImage(s.image, "Data\\Output.jpg", true)

	case CHECK_DATA_REQUEST:
		returnEvent.Type = IS_VALID_DATA_REQUEST
		returnEvent.Parameter = s.checkDataRequest(event.Parameter.(dataRequest))
	}
	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

// Functions that depend on (or change) service state go here

// Draws data from dataRequest
func (s *GISService) generateMap(requestsPointer **[]dataRequest) {

	for _, request := range **requestsPointer {
		set := s.datasets[request.datasetName]

		for i, _ := range set.data {

			for dataType, dataRange := range request.requested {

				floatData, _ := strconv.ParseFloat(set.data[i][set.indices[dataType]], 64)
				if isInRange(floatData, &dataRange) {
					lon, _ := strconv.ParseFloat(set.data[i][set.indices["longitude"]], 64)
					lat, _ := strconv.ParseFloat(set.data[i][set.indices["latitude"]], 64)
					location := s.mapToPixels(lat, lon)
					drawDot(s.image, location, (int)(floatData/(1e6)), true)
				}
			}
		}
	}
}

func (s *GISService) initMap(img *image.Image, topleft, bottomright image.Point) {
	bounds := (*img).Bounds()
	s.yStep = math.Abs((float64)(bounds.Max.Y-bounds.Min.Y) / (float64)(topleft.Y-bottomright.Y))
	s.xStep = math.Abs((float64)(bounds.Max.X-bounds.Min.X) / (float64)(topleft.X-bottomright.X))
	yCenterPixel := (int)((float64)(-topleft.Y) * s.yStep)
	xCenterPixel := (int)((float64)(topleft.X) * s.xStep)
	s.centerPixel = image.Point{xCenterPixel, yCenterPixel}
}

func (s *GISService) mapToPixels(longitude, latitude float64) image.Point {
	x := s.centerPixel.X - (int)(s.xStep*(float64)(longitude))
	y := s.centerPixel.Y + (int)(s.yStep*(float64)(latitude))
	return image.Point{x, y}
}

// Checks if all requests ask for datasets and datatypes that exist in the currently loaded datasets
func (s *GISService) checkDataRequest(request dataRequest) bool {

	// for _, request := range requests {
	// Check if datasetName exists
	var matchedSet *dataset
	for _, set := range s.datasets {
		if set.name == request.datasetName {
			matchedSet = set
			break
		}
	}
	if matchedSet == nil {
		fmt.Println(request.datasetName + " is not a valid dataset name.")
		return false
	}

	for keyRequested, _ := range request.requested {
		flag := false
		for keyDataSet, _ := range (*matchedSet).indices {
			if keyRequested == keyDataSet {
				flag = true
				break
			}
		}
		if flag == false {
			fmt.Println(keyRequested + " is not a valid data type in " + matchedSet.name + " dataset.")
			return false
		}
	}
	// }
	return true
}

// ------------------------------------------- Utility ------------------------------------------- //

// Purely functional routines go here
// None of these functions should depend on (or change) state

// Respond to errors
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Loads an image frome a fileName
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

// Saves image to location provided by fileName. rotate bool is a flag for rotating saved image 180 degrees
func saveImage(img image.Image, fileName string, rotate bool) {

	f, err := os.Create(fileName)
	checkError(err)
	if rotate {
		img = imaging.Rotate(img, 180.0, color.RGBA{0, 0, 0, 0})
	}
	err = jpeg.Encode(f, img, nil)
	checkError(err)
}

// Draws square dot centered on point in image. Can be outline or filled
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

// Checks if value is contained in any of a set of ranges
func isInRange(value float64, data *[]dataRange) bool {

	for i, _ := range *data {
		if value >= (*data)[i].lower && value < (*data)[i].upper {
			return true
		}
	}
	return false
}
