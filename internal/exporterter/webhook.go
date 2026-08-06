package exporter

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/darnellwashingjr-94/omniping/pkg/telemetry"
)

type WebhookExporter struct {
	url    string
	client *http.Client
}

func NewWebhookExporter(url string) *WebhookExporter {
	return &WebhookExporter{
		url:    url,
		client: &http.Client{Timeout: 5 * time.Second},
	}
}

func (w *WebhookExporter) Export(results []telemetry.Result) error {
	failed := telemetry.FilterFailed(results)
	if len(failed) == 0 {
		return nil
	}

	payload, err := json.Marshal(failed)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", w.url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := w.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
