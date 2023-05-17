package csv

import (
	"encoding/csv"
	"os"
)

type Writer struct {
	writer *csv.Writer
	file   *os.File
}

func NewCSVWriter(filename string) (*Writer, error) {
	file, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	writer := csv.NewWriter(file)

	return &Writer{
		writer: writer,
		file:   file,
	}, nil
}

func (w *Writer) Write(record []string) error {
	return w.writer.Write(record)
}

func (w *Writer) Flush() error {
	w.writer.Flush()
	return w.writer.Error()
}

func (w *Writer) Close() error {
	w.writer.Flush()
	if err := w.writer.Error(); err != nil {
		return err
	}

	return w.file.Close()
}
