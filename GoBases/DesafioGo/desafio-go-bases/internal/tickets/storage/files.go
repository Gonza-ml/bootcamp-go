package storage

import (
	"encoding/csv"
	"os"
)

func ReadCSV(fileName string) (data [][]string, err error) {
	file, err := os.Open(fileName)
	if err != nil {
		return
	}

	defer file.Close()

	// reader csv
	reader := csv.NewReader(file)

	data, err = reader.ReadAll()
	return

}
