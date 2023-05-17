package service

import (
	"csv-combiner/internal/company/domain/entity"
	"errors"
	"testing"
)

type MockWriter struct {
	WrittenData [][]string
	Flushed     bool
	WriteError  error
	FlushError  error
}

func (w *MockWriter) Write(data []string) error {
	if w.WriteError != nil {
		return w.WriteError
	}
	w.WrittenData = append(w.WrittenData, data)
	return nil
}

func (w *MockWriter) Flush() error {
	if w.FlushError != nil {
		return w.FlushError
	}
	w.Flushed = true
	return nil
}

func TestService_WriteCombined(t *testing.T) {

	testCases := []struct {
		Name         string
		Names        map[string]entity.Company
		Descriptions map[string]entity.Company
		ExpectedData [][]string
		WriteError   error
	}{
		{
			Name: "Successfully write combined data",
			Names: map[string]entity.Company{
				"1": {ID: "1", Name: "Company A"},
				"2": {ID: "2", Name: "Company B"},
				"3": {ID: "3", Name: "Company C"},
			},
			Descriptions: map[string]entity.Company{
				"1": {ID: "1", Description: "Description A"},
				"2": {ID: "2", Description: "Description B"},
			},
			ExpectedData: [][]string{
				{"CompanyID", "CompanyName", "CompanyDescription"},
				{"1", "Company A", "Description A"},
				{"2", "Company B", "Description B"},
				{"3", "Company C", ""},
			},
		},
		{
			Name: "Error when writing data",
			Names: map[string]entity.Company{
				"1": {ID: "1", Name: "Company A"},
			},
			Descriptions: map[string]entity.Company{
				"1": {ID: "1", Description: "Description A"},
			},
			WriteError: errors.New("write error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mockWriter := &MockWriter{
				WriteError: tc.WriteError,
			}
			service := NewService(mockWriter)

			err := service.WriteCombined(tc.Names, tc.Descriptions)
			if err != nil && tc.WriteError == nil {
				t.Errorf("unexpected error: %v", err)
			}

			if len(mockWriter.WrittenData) != len(tc.ExpectedData) {
				t.Errorf("unexpected number of write calls, got %d, want %d", len(mockWriter.WrittenData), len(tc.ExpectedData))
			}

			for i, expected := range tc.ExpectedData {
				if i >= len(mockWriter.WrittenData) {
					t.Errorf("expected data written at index %d, but no write call found", i)
					continue
				}
				actual := mockWriter.WrittenData[i]
				if !stringSlicesEqual(actual, expected) {
					t.Errorf("unexpected data written at index %d, got %v, want %v", i, actual, expected)
				}
			}

			if tc.WriteError == nil && !mockWriter.Flushed {
				t.Error("expected Flush to be called, but it was not")
			}

			if tc.WriteError != nil && mockWriter.Flushed {
				t.Error("expected Flush not to be called, but it was")
			}
		})
	}
}

func stringSlicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, val := range a {
		if val != b[i] {
			return false
		}
	}
	return true
}
