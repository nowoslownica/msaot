package main

import (
	"encoding/csv"
	"fmt"
	"github.com/eakarpov/msaot/db"
	"github.com/eakarpov/msaot/importer"
	"io"
	"log"
	"os"
)

func main() {
	_, err := os.Open("dictionary.csv")
	if err != nil {
		resp := importer.GetData()
		if len(resp.Values) == 0 {
			fmt.Println("No data found.")
		} else {
			file, err := os.Create("dictionary.csv")
			if err != nil {
				log.Fatalf("Unable to create file: %v", err)
			}

			writer := csv.NewWriter(file)

			for _, row := range resp.Values {
				str := make([]string, 0)
				for _, elem := range row {
					str = append(str, elem.(string))
				}
				err := writer.Write(str)
				if err != nil {
					log.Fatalf("Unable to write to file: %v", err)
				}
			}
			writer.Flush()
			err = file.Close()
			if err != nil {
				log.Fatalf("Unable to close the file: %v", err)
			}
		}
	}
	f, err := os.Open("dictionary.csv")
	if err != nil {
		log.Fatalf("Unable to open file %v", err)
	}
	defer f.Close()
	mainDB := db.GetByName("dictionary.db")
	err = mainDB.InitDB(false)
	if err != nil {
		log.Fatalf("Unable to init db: %v", err)
	}
	err = mainDB.Init()
	if err != nil {
		log.Fatalf("Unable to open db: %v", err)
	}
	csvr := csv.NewReader(f)
	csvr.FieldsPerRecord = -1
	err = importer.LoadGPositions(mainDB)
	if err != nil {
		log.Fatalf("Unable to load gPositions: %v", err)
	}
	index := 0
	for {
		record, err := csvr.Read()
		if err == io.EOF {
			break
		}
		if index > 100 {
			break
		}
		if err != nil {
			log.Fatalf("Error while reading file")
		}
		importer.ImportRecord(mainDB, record)
		index++
	}
}

