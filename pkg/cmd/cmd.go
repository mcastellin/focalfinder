package processor

import (
	"fmt"
	"os"
	"path/filepath"
)

func walkDirectory(dirPath string, processor *MetadataProcessor) error {
	return filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		ext := filepath.Ext(info.Name())
		if ext == ".jpg" || ext == ".jpeg" || ext == ".JPG" || ext == ".JPEG" {
			if err := processor.ProcessImage(path); err != nil {
				fmt.Printf("Error processing file %s: %v\n", path, err)
			}
		}
		return nil
	})
}

func ProcessImages(dirs []string, factors []string) {
	processor := NewMetadataProcessor()

	for _, dir := range dirs {
		err := walkDirectory(dir, processor)
		if err != nil {
			fmt.Printf("Error walking directory: %v\n", err)
			os.Exit(1)
		}
	}

	if processor.Stats.TotalFotos() == 0 {
		fmt.Println("No photos found with EXIF data in the specified directory.")
		os.Exit(0)
	}

	mostUsedLength, count := processor.Stats.GetMostUsedFocalLength()
	fmt.Printf("Total photos analyzed: %d\n", processor.Stats.TotalFotos())
	fmt.Printf("Most used focal length: %.2fmm (used in %d photos)\n", mostUsedLength, count)

	// Print full statistics
	fmt.Println("\nAll focal lengths and counts:")
	fmt.Println(processor.Stats)

	// Print missing metadata
	if len(processor.UnknownMakeModels) > 0 {
		fmt.Println("\nCould not compute crop factors for the following camera make and models:")
		for _, m := range processor.UnknownMakeModels {
			fmt.Printf("Make: %s, Model: %s\n", m.Make, m.Model)
		}
	}
}
