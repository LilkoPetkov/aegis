package pass_export

import (
	"aegis/internal/queries"

	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ExportPasswordsCsv(filePath string) {
	csvExport, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file: %s", err)
	}
	defer csvExport.Close()

	writer := csv.NewWriter(csvExport)
	defer writer.Flush()

	writeDataCsv(writer)
}

func writeDataCsv(writer *csv.Writer) {
	rows := queries.FetchAllUsers()
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatalf("Error getting columns: %s", err)
	}
	if err := writer.Write(columns); err != nil {
		log.Fatalf("Writer error: %s", err)
	}

	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))

	for i := range values {
		valuePtrs[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatalf("Error copying columns from rows: %s", err)
		}

		record := make([]string, len(columns))
		for i, val := range values {
			if val == nil {
				record[i] = ""
			} else {
				record[i] = fmt.Sprintf("%v", val)
			}
		}

		if err := writer.Write(record); err != nil {
			log.Fatalf("Error writing record to CSV: %s", err)
		}
	}
}
