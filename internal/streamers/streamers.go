package streamers

import (
	"LycorisMonitor/internal/services"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func ReadFromFile(filename string) (*[]services.Streamer, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	var streamers []services.Streamer

	err = json.NewDecoder(bytes.NewReader(data)).Decode(&streamers)
	if err != nil {
		return nil, fmt.Errorf("failed to decode JSON from file %s: %w", filename, err)
	}

	return &streamers, nil
}

func WriteToFile(streamers *[]services.Streamer, filename string) error {
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	err = json.NewEncoder(writer).Encode(streamers)
	if err != nil {
		return fmt.Errorf("failed to encode JSON to file %s: %w", filename, err)
	}

	return nil
}

// Streamer и Preview структуры теперь находятся в internal/services/streamers.go
