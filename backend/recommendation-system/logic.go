package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

type Data struct {
	zipcode       int64   `json:"zipcode"`
	population    int64   `json:"population"`
	landSqMile    float64 `json:"land_sq_mile"`
	densitySqMile float64 `json:"density_sq_mile"`
}

func main() {

	csvFile, _ := os.Open("data/population-density-per-zipcode.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	dict := make(map[int64]Data) // create a map,k: on zip code, v: Data object
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(line)

		zipcode, err := strconv.ParseInt(line[0], 16, 64)
		population, err := strconv.ParseInt(line[1], 16, 64)
		landSqMile, err := strconv.ParseFloat(line[2], 32)
		densitySqMile, err := strconv.ParseFloat(line[3], 32)

		dataObj := Data{zipcode, population, landSqMile, densitySqMile}

		dict[zipcode] = dataObj

	}

	// print the contents of the map
	for key, value := range dict {
		fmt.Println("Key:", key, "Value:", value)
	}

}
