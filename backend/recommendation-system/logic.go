package recommendation

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type Data struct {
	zipcode       int     `json:"zipcode"`
	population    int     `json:"population"`
	landSqMile    float32 `json:"land_sq_mile"`
	densitySqMile float32 `json:"density_sq_mile"`
}

func Reader() {

	csvFile, _ := os.Open("data/population-density-per-zipcode.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var lines []Data

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}

	// Output:
	// [first_name last_name username]
	// [Rob Pike rob]
	// [Ken Thompson ken]
	// [Robert Griesemer gri]
}
