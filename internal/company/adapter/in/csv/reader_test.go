package csv

import (
	"csv-combiner/internal/company/domain"
	"encoding/csv"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestNewCSVReader(t *testing.T) {
	// Test data file path
	filename := filepath.Join("testdata", "names.csv")

	// Test NewCSVReader function
	reader, err := NewCSVReader(filename)
	if err != nil {
		t.Fatalf("NewCSVReader returned an error: %s", err)
	}
	defer reader.Close()

	// Perform assertions on the reader object
	if reader.reader == nil {
		t.Error("Reader should not be nil")
	}
	if reader.file == nil {
		t.Error("File should not be nil")
	}
}

func TestReader_ParseNames(t *testing.T) {
	// Test data file path
	filename := filepath.Join("testdata", "names.csv")

	// Open the test data file
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open test data file: %s", err)
	}
	defer file.Close()

	reader := &Reader{reader: csv.NewReader(file)}

	// Test ParseNames function
	companies, err := reader.ParseNames()
	if err != nil {
		t.Fatalf("ParseNames returned an error: %s", err)
	}

	// Perform assertions on the parsed companies
	expected := map[string]domain.Company{
		"1": {ID: "1", Name: "Company A"},
		"2": {ID: "2", Name: "Company B"},
	}
	if len(companies) != len(expected) {
		t.Errorf("Expected %d companies, but got %d", len(expected), len(companies))
	}
	for k, v := range companies {
		if !reflect.DeepEqual(v, expected[k]) {
			t.Errorf("Expected company %v to be %+v, but got %+v", k, expected[k], v)
		}
	}
}

func TestReader_ParseDescriptions(t *testing.T) {
	// Test data file path
	filename := filepath.Join("testdata", "descriptions.csv")

	// Open the test data file
	file, err := os.Open(filename)
	if err != nil {
		t.Fatalf("failed to open test data file: %s", err)
	}
	defer file.Close()

	reader := &Reader{reader: csv.NewReader(file)}

	// Test ParseDescriptions function
	companies, err := reader.ParseDescriptions()
	if err != nil {
		t.Fatalf("ParseDescriptions returned an error: %s", err)
	}

	// Perform assertions on the parsed companies
	expected := map[string]domain.Company{
		"1": {ID: "1", Description: "Description for Company A"},
		"2": {ID: "2", Description: "Description for Company B"},
	}
	if len(companies) != len(expected) {
		t.Errorf("Expected %d companies, but got %d", len(expected), len(companies))
	}
	for k, v := range companies {
		if !reflect.DeepEqual(v, expected[k]) {
			t.Errorf("Expected company %v to be %+v, but got %+v", k, expected[k], v)
		}
	}
}
