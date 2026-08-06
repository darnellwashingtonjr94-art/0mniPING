package exporter

import (
	"encoding/json"
	"os"

	"github.com/username/omniping/pkg/telemetry"
)

type JSONExporter struct {
	filePath string
}

func NewJSONExporter(filePath string) *JSONExporter {
	return &JSONExporter{filePath: filePath}
}

func (j *JSONExporter) Export(results []telemetry.Result) error {
	file, err := os.OpenFile(j.filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	for _, res := range results {
		if err := encoder.Encode(res); err != nil {
			return err
		}
	}
	return nil
}
