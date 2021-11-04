package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
)

type Contact struct {
	Name   string
	Number string
}

func ReadContacts(csvFile string) ([]*Contact, error) {

	f, err := os.Open(csvFile)
	if err != nil {
		fmt.Printf("unable to read csv file. Error: %s", err.Error())
		return nil, errors.New("unable to read csv file")
	}

	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'

	csvRecords := make([]*Contact, 0)

	// read header
	_, err = csvReader.Read()
	if err != nil {
		fmt.Printf("unable to read header. Error: %s", err.Error())
		return nil, errors.New("unable to read header")
	}

	for {
		record, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("error reading csv record. Error: %s\n", err.Error())
			return nil, errors.New("error reading csv record")
		}

		csvRecord := &Contact{
			Name:   record[0],
			Number: record[1],
		}

		csvRecords = append(csvRecords, csvRecord)
	}

	return csvRecords, nil
}
