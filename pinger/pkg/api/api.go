package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func SendData(ip, pingTime, lastSuccessfulPing string) error {
	payload := map[string]string{
		"ip_address":           ip,
		"ping_time":            pingTime,
		"last_successful_ping": lastSuccessfulPing,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("failed to marshal payload: %v", err)
	}

	resp, err := http.Post("http://back:8080/api/containers", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("backend returned status code: %d", resp.StatusCode)
	}

	return nil
}
