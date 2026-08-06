package exporter

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/username/omniping/pkg/telemetry"
)

type CSVExporter struct {
	filePath string
}

func NewCSVExporter(filePath string) *CSVExporter {
	return &CSVExporter{filePath: filePath}
}

func (c *CSVExporter) Export(results []telemetry.Result) error {
	file, err := os.OpenFile(c.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, res := range results {
		record := []string{
			res.Timestamp.Format(time.RFC3339),
			res.Target,
			res.Layer,
			res.Probe,
			string(res.Status),
			strconv.FormatInt(res.Latency.Nanoseconds(), 10),
			res.Error,
		}
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	return nil
}
