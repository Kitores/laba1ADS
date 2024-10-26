package readCSV

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Sort struct {
	Name      string
	ExecTimes []int64
	Sizes     []string
}

func ReadCSV(sortName string) Sort {
	file, err := os.Open(fmt.Sprintf("./internal/sorts/csvData/%sAverageData.csv", sortName))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var execTimes []int64
	var sizes []string
	for i, record := range records {
		if i == 0 { // Пропускаем заголовок
			continue
		}
		size := record[0]
		sizes = append(sizes, size)

		execTimeInt, err := strconv.Atoi(record[1]) // Преобразуем вреся в int
		if err != nil {
			log.Fatal(err)
		}
		execTimeInt64 := int64(execTimeInt)
		execTimes = append(execTimes, execTimeInt64)

	}
	sort := Sort{
		Name:      sortName,
		ExecTimes: execTimes,
		Sizes:     sizes,
	}

	fmt.Println(sort)
	return sort
}
