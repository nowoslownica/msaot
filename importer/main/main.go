package main

import (
	"encoding/csv"
	"fmt"
	"github.com/eakarpov/msaot/importer"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
	"log"
	"os"
)

func getData() *sheets.ValueRange {
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := importer.GetClient(config)

	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	// Prints the names and majors of students in a sample spreadsheet:
	// https://docs.google.com/spreadsheets/d/1BxiMVs0XRA5nFMdKvBdBZjgmUUqptlbs74OgvE2upms/edit
	spreadsheetId := "1Yl8Kk3M1yKCVa7cAsRxTQaJG5Vdnp6yGy6ceFGQb4z8"
	readRange := "words!A2:X"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	return resp
}

func main() {
	_, err := os.OpenFile("dictionary.csv", 1,644)
	if err != nil {
		resp := getData()
		if len(resp.Values) == 0 {
			fmt.Println("No data found.")
		} else {
			file, err := os.Create("dictionary.csv")
			if err != nil {
				log.Fatalf("Unable to create file: %v", err)
			}
			defer file.Close()

			writer := csv.NewWriter(file)
			defer writer.Flush()

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
		}
	}
	f, err := os.OpenFile("dictionary.csv", 1,644)
	if err != nil {
		log.Fatalf("Unable to open file %v", err)
	}
	fmt.Println(f.Name())
}
