package csv

import (
	"csv-combiner/internal/company/domain/entity"
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

type Reader struct {
	reader *csv.Reader
	file   *os.File
}

func NewCSVReader(filename string) (*Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	reader.Comma = ','

	return &Reader{
		reader: reader,
		file:   file,
	}, nil
}

func (r *Reader) ParseNames() (map[string]entity.Company, error) {
	// Skip header row
	_, err := r.reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %s", err)
	}

	names := make(map[string]entity.Company)

	for {
		record, err := r.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read record: %s", err)
		}

		companyID := record[0]
		companyName := record[1]

		company := entity.Company{
			ID:   companyID,
			Name: companyName,
		}

		names[companyID] = company
	}

	return names, nil
}

func (r *Reader) ParseDescriptions() (map[string]entity.Company, error) {
	// Skip header row
	_, err := r.reader.Read()
	if err != nil {
		return nil, fmt.Errorf("failed to read header: %s", err)
	}

	descriptions := make(map[string]entity.Company)

	for {
		record, err := r.reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("failed to read record: %s", err)
		}

		companyID := record[0]
		companyDescription := record[1]

		company := entity.Company{
			ID:          companyID,
			Description: companyDescription,
		}

		descriptions[companyID] = company
	}

	return descriptions, nil
}

func (r *Reader) Close() error {
	return r.file.Close()
}
