package equiv

import (
	"errors"
	"strings"
)

type CameraCropFactor struct {
	Make       string
	Model      string
	CropFactor float64
}

var cameraCropFactors = []CameraCropFactor{
	{"Fujifilm", "GFX", 0.79},                 // Medium format
	{"Fujifilm", "X-T", 1.5},                  // APS-C
	{"Canon", "EOS R", 1.0},                   // Full-frame
	{"Nikon", "Z6", 1.0},                      // Full-frame
	{"NIKON CORPORATION", "NIKON D850", 1.0},  // Full-frame
	{"NIKON CORPORATION", "NIKON D7500", 1.5}, // APS-C
	{"NIKON CORPORATION", "NIKON D7200", 1.5}, // APS-C
	{"NIKON CORPORATION", "NIKON D5600", 1.5}, // APS-C
	{"NIKON CORPORATION", "NIKON D5100", 1.5}, // APS-C
	{"NIKON CORPORATION", "NIKON D3100", 1.5}, // APS-C
	{"NIKON CORPORATION", "NIKON D90", 1.5},   // APS-C
	{"Canon", "Canon EOS 70D", 1.5},           // APS-C
	{"Canon", "Canon EOS 600D", 1.5},          // APS-C
	{"Sony", "Alpha A7", 1.0},                 // Full-frame
	{"Sony", "Alpha A6000", 1.5},              // APS-C
	{"Leica", "M10", 1.0},                     // Full-frame
	{"Pentax", "645Z", 0.79},                  // Medium format
}

// Get35mmEquivalent determines the crop factor based on the camera make and model.
func Get35mmEquivalent(focalLength float64, cameraMake, model string) (float64, error) {
	cameraMake = strings.ToLower(cameraMake)
	model = strings.ToLower(model)

	for _, camera := range cameraCropFactors {
		if strings.Contains(cameraMake, strings.ToLower(camera.Make)) && strings.Contains(model, strings.ToLower(camera.Model)) {
			return focalLength * camera.CropFactor, nil
		}
	}

	return 0, errors.New("unknown crop factor for specified make and model")
}
