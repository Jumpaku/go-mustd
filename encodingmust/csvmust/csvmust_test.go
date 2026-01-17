package csvmust_test

import (
	"bytes"
	"strings"
	"testing"

	"github.com/Jumpaku/go-mustd/encodingmust/csvmust"
	"github.com/Jumpaku/go-mustd/iomust"
)

func TestReader(t *testing.T) {
	t.Run("Read single record", func(t *testing.T) {
		csvData := "name,age\nAlice,30\nBob,25"
		r := iomust.ReaderOf(strings.NewReader(csvData))
		reader := csvmust.NewReader(r)

		record := reader.Read()
		if len(record) != 2 {
			t.Errorf("expected 2 fields, got %d", len(record))
		}
		if record[0] != "name" || record[1] != "age" {
			t.Errorf("expected [name age], got %v", record)
		}
	})

	t.Run("ReadAll", func(t *testing.T) {
		csvData := "name,age\nAlice,30\nBob,25"
		r := iomust.ReaderOf(strings.NewReader(csvData))
		reader := csvmust.NewReader(r)

		records := reader.ReadAll()
		if len(records) != 3 {
			t.Errorf("expected 3 records, got %d", len(records))
		}
	})

	t.Run("FieldPos", func(t *testing.T) {
		csvData := "name,age\nAlice,30"
		r := iomust.ReaderOf(strings.NewReader(csvData))
		reader := csvmust.NewReader(r)

		reader.Read()
		line, col := reader.FieldPos(0)
		if line == 0 && col == 0 {
			t.Error("FieldPos returned zero values")
		}
	})

	t.Run("InputOffset", func(t *testing.T) {
		csvData := "name,age\nAlice,30"
		r := iomust.ReaderOf(strings.NewReader(csvData))
		reader := csvmust.NewReader(r)

		reader.Read()
		offset := reader.InputOffset()
		if offset == 0 {
			t.Error("InputOffset returned zero")
		}
	})
}

func TestWriter(t *testing.T) {
	t.Run("Write single record", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := iomust.WriterOf(buf)
		writer := csvmust.NewWriter(w)

		writer.Write([]string{"name", "age"})
		writer.Flush()

		if buf.String() != "name,age\n" {
			t.Errorf("expected 'name,age\\n', got %s", buf.String())
		}
	})

	t.Run("WriteAll", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := iomust.WriterOf(buf)
		writer := csvmust.NewWriter(w)

		records := [][]string{
			{"name", "age"},
			{"Alice", "30"},
			{"Bob", "25"},
		}
		writer.WriteAll(records)
		writer.Flush()

		expected := "name,age\nAlice,30\nBob,25\n"
		if buf.String() != expected {
			t.Errorf("expected %q, got %q", expected, buf.String())
		}
	})

	t.Run("Flush", func(t *testing.T) {
		buf := &bytes.Buffer{}
		w := iomust.WriterOf(buf)
		writer := csvmust.NewWriter(w)

		writer.Write([]string{"test"})
		// Before flush, buffer might be empty
		writer.Flush()

		if buf.Len() == 0 {
			t.Error("Flush did not write data")
		}
	})
}
