package embedlivesample

import (
	"bytes"
	"os"
)

func saveLiveSample(folder string, filePath string, params *LiveSampleParams) error {
	err := os.MkdirAll(folder, os.ModePerm)
	if err != nil {
		return err
	}
	// Will be put into book in postprocessing
	var liveSampleBuffer bytes.Buffer
	err = tLiveSample.Execute(&liveSampleBuffer, params)
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, liveSampleBuffer.Bytes(), 0644)

}
