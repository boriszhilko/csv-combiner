package service

import (
	"csv-combiner/internal/company/domain/entity"
	"fmt"
)

type Writer interface {
	Write([]string) error
	Flush() error
}

type Service struct {
	writer Writer
}

func NewService(writer Writer) *Service {
	return &Service{
		writer: writer,
	}
}

func (s *Service) WriteCombined(names map[string]entity.Company, descriptions map[string]entity.Company) error {
	companies := s.combine(names, descriptions)
	err := s.write(companies)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) combine(names map[string]entity.Company, descriptions map[string]entity.Company) []entity.Company {
	var combined []entity.Company

	for k, name := range names {
		company := entity.Company{
			ID:   k,
			Name: name.Name,
		}

		desc, ok := descriptions[k]
		if !ok {
			fmt.Println("No description for company ", k)
		} else {
			company.Description = desc.Description
		}

		combined = append(combined, company)
	}

	return combined
}

func (s *Service) write(companies []entity.Company) error {
	// Write headers to combined.csv
	headers := []string{"CompanyID", "CompanyName", "CompanyDescription"}
	err := s.writer.Write(headers)
	if err != nil {
		return fmt.Errorf("failed to write headers: %s", err)
	}

	// Write data to combined.csv
	for _, company := range companies {
		err = s.writer.Write([]string{company.ID, company.Name, company.Description})
		if err != nil {
			return fmt.Errorf("failed to write record: %s", err)
		}
	}

	err = s.writer.Flush()
	if err != nil {
		return fmt.Errorf("failed to flush writer: %s", err)
	}

	return nil
}
