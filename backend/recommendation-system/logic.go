package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"

	types "./types"
)

func main() {

	csvFile, _ := os.Open("data/population-density-per-zipcode.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))
	data := make(map[string]types.Data) // create a map,k: on zip code, v: Data object
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		zipcode := line[0]
		population, err := strconv.ParseInt(line[1], 16, 64)
		landSqMile, err := strconv.ParseFloat(line[2], 32)
		densitySqMile, err := strconv.ParseFloat(line[3], 32)

		dataObj := types.Data{zipcode, population, landSqMile, densitySqMile}

		if line[0] == "94404" {
			fmt.Println(dataObj)
			data[zipcode] = dataObj
		} else {
			data[zipcode] = dataObj
		}
	}

	// San Mateo zip codes:  94401, 94402, 94403, 94404.
	// Santa Clara zip codes: 94089, 95002, 95008, 95013, 95014, 95032, 95035, 95037, 95050, 95054, 95070, 95110, 95111, 95112, 95113, 95116, 95117, 95118, 95119, 95120, 95121, 95122, 95123, 95124, 95125, 95126, 95127, 95128, 95129, 95130, 95131, 95132, 95133, 95134, 95135, 95136, 95138, 95139, 95140, 95148.

	countyZipCodes := make(map[string][]string)

	countyZipCodes["San Mateo"] = []string{"94401", "94402", "94403", "94404"}
	countyZipCodes["San Jose"] = []string{"94089", "95002", "95008", "95013", "95014", "95032", "95035", "95037", "95050", "95054", "95070", "95110", "95111", "95112", "95113", "95116", "95117", "95118", "95119", "95120", "95121", "95122", "95123", "95124", "95125", "95126", "95127", "95128", "95129", "95130", "95131", "95132", "95133", "95134", "95135", "95136", "95138", "95139", "95140", "95148"}

	type CountyZip struct {
		zipcode string
		county  string
	}

	densityMap := make(map[CountyZip]float64) // create a map,k: on zip code, v: population denisty

	for key, value := range countyZipCodes {
		arr := value

		for i := 0; i < len(arr); i++ {
			zipcode := arr[i]                         // get zip code from list
			dataOjb := data[zipcode]                  // get Data object based on zipcode
			densityPerSqMile := dataOjb.DensitySqMile // get the density per sq mile from the object
			k := CountyZip{zipcode, key}
			densityMap[k] = densityPerSqMile // put density in the map for each zip code
		}
	}

	for key, value := range densityMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	// Need to sort based on value
	// Connect with frontend

	// create a map to store san mateo values
	densityMapSanMateo := make(map[float64]CountyZip)
	densityMapSanJose := make(map[float64]CountyZip)
	var keysSanMateo = make([]float64, 100)
	var keysSanJose = make([]float64, 100)

	// split into two hashmaps
	counterSanMateo := 0
	counterSanJose := 0
	for key, value := range densityMap {
		if key.county == "San Mateo" {
			// add to array containing san mateo zipcodes
			// reverse key value pair
			densityMapSanMateo[value] = key
			keysSanMateo[counterSanMateo] = value
			counterSanMateo++

		} else {
			// add to array containining san jose zipcodes
			// reverse key value pair
			densityMapSanJose[value] = key
			keysSanJose[counterSanJose] = value
			counterSanJose++
		}
	}
	sort.Float64s(keysSanMateo)
	sort.Float64s(keysSanJose)

	// find the zipcode with the highest density for both cities
	largestKeySM := keysSanMateo[len(keysSanMateo)-1]
	largestKeySJ := keysSanJose[len(keysSanJose)-1]

	largestZipCodeSJ := densityMapSanJose[largestKeySJ].zipcode
	largestZipCodeSM := densityMapSanMateo[largestKeySM].zipcode

	fmt.Println("ZipCode with highest density population in San Mateo: ", largestZipCodeSM)
	fmt.Println("ZipCode with highest density population in San Jose: ", largestZipCodeSJ)

	// pass these values to frontend rendering
}
