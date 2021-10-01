package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/twpayne/go-polyline"
)

func main() {
	//Prompts the log Path
	fmt.Print("Enter Log Path: ")
	var rutaLog string
	fmt.Scanln(&rutaLog)
	rutaLog = strings.ReplaceAll(rutaLog, "\\", "/")
	fmt.Print("\n")
	//coords stores geolocations values
	coords := make([][]float64, 0)

	file, err := os.Open(rutaLog)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	regex := regexp.MustCompile(`(?:latitude:\s)?(?P<lat>[\-\+]?[\d]*\.[\d]*),(?:\slongitude:\s)?(?P<long>[\-\+]?[\d]*\.[\d]*)`)
	count := 0
	for scanner.Scan() {
		logLine := scanner.Text()
		if regex.FindString(logLine) != "" {
			match := regex.FindStringSubmatch(logLine)
			//matching groups that are defined in regex
			lat, _ := strconv.ParseFloat(match[1], 64)
			long, _ := strconv.ParseFloat(match[2], 64)

			latlong := []float64{lat, long}
			coords = append(coords, latlong)
			count++
		}
	}

	fmt.Println(string(polyline.EncodeCoords(coords)))
	fmt.Println("\n", count, " coordinates")
	fmt.Println("Press the Enter Key to finish!")
	fmt.Scanln() // wait for Enter Key

}
