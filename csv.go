package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statistical function
type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// Parses the contents of a CSV into a slice of floating point numbers.
func csv2float(r io.Reader, column int) ([]float64, error) {
	// Create the CSV reader used to read in data from CSV files.
	cr := csv.NewReader(r)

	// Adjust for a 0 based index
	column--

	// Read in all CSV data
	allData, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	// Variable to hold converted data
	var data []float64

	// Looping through all records
	for i, row := range allData {
		// file does not have that many columns
		if i == 0 {
			continue
		}
		// Checking number of columns in CSV file
		if len(row) <= column {
			return nil, fmt.Errorf("%w: File has only %d columns", ErrInvalidColumn, len(row))
		}
		// Try to convert the data to a float value
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}
	// Return the slice of float64 and nil error
	return data, nil
}
