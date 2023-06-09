package main

import (
	in "csv-combiner/internal/company/adapter/in/csv"
	out "csv-combiner/internal/company/adapter/out/csv"
	"csv-combiner/internal/company/domain/entity"
	"csv-combiner/internal/company/domain/service"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		log.Fatalf("Usage: go run main.go names.csv descriptions.csv combined.csv")
	}

	namesFile := os.Args[1]
	descriptionsFile := os.Args[2]
	combinedFile := os.Args[3]

	names := parseNames(namesFile)
	descriptions := parseDescriptions(descriptionsFile)

	combineAndWrite(names, descriptions, combinedFile)
}

func combineAndWrite(names map[string]entity.Company, descriptions map[string]entity.Company, combinedFile string) {
	combinedWriter, err := out.NewCSVWriter(combinedFile)
	if err != nil {
		log.Fatalf("Failed to create combined.csv: %s", err)
	}
	defer combinedWriter.Close()

	companyService := service.NewService(combinedWriter)
	err = companyService.WriteCombined(names, descriptions)
	if err != nil {
		log.Fatalf("Failed to combineAndWrite combined.csv: %s", err)
	}
}

func parseDescriptions(descriptionsFilePath string) map[string]entity.Company {
	descriptionsReader, err := in.NewCSVReader(descriptionsFilePath)
	if err != nil {
		log.Fatalf("Failed to open descriptions file: %s", err)
	}
	defer descriptionsReader.Close()

	descriptions, err := descriptionsReader.ParseDescriptions()
	if err != nil {
		log.Fatalf("Failed to parse descriptions file: %s", err)
	}
	return descriptions
}

func parseNames(namesFilePath string) map[string]entity.Company {
	namesReader, err := in.NewCSVReader(namesFilePath)
	if err != nil {
		log.Fatalf("Failed to open names file: %s", err)
	}
	defer namesReader.Close()

	names, err := namesReader.ParseNames()
	if err != nil {
		log.Fatalf("Failed to parse names file: %s", err)
	}
	return names
}
