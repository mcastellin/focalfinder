package processor

import (
	"fmt"

	"github.com/mcastellin/focalfinder/pkg/equiv"
	"github.com/mcastellin/focalfinder/pkg/exif"
	"github.com/mcastellin/focalfinder/pkg/stats"
)

func NewMetadataProcessor() *MetadataProcessor {
	return &MetadataProcessor{
		Stats: stats.NewFocalStats(),

		UnknownMakeModels: map[string]equiv.CameraCropFactor{},
	}
}

type MetadataProcessor struct {
	Stats *stats.FocalStats

	UnknownMakeModels map[string]equiv.CameraCropFactor
}

func (p *MetadataProcessor) ProcessImage(filePath string) error {
	x, err := exif.Decode(filePath)
	if err != nil {
		return err
	}

	var length float64
	var cameraMake, model string

	if length, err = x.ExifFocalLength(); err != nil {
		return nil
	}
	if cameraMake, err = x.ExifMake(); err != nil {
		return nil
	}
	if model, err = x.ExifModel(); err != nil {
		return nil
	}

	equivLength, err := equiv.Get35mmEquivalent(length, cameraMake, model)
	if err != nil {
		makeModel := fmt.Sprintf("%s-%s", cameraMake, model)
		p.UnknownMakeModels[makeModel] = equiv.CameraCropFactor{
			Make:  cameraMake,
			Model: model,
		}
		return nil
	}
	p.Stats.AddClusteredFocalLength(equivLength)
	return nil
}
