package exif

import (
	"errors"
	"os"

	_exif "github.com/rwcarlsen/goexif/exif"
)

func Decode(filePath string) (*ImageExif, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	x, err := _exif.Decode(file)
	if err != nil {
		return nil, errors.New("could not decode EXIF data for file")
	}
	return &ImageExif{x: x}, nil
}

type ImageExif struct {
	x *_exif.Exif
}

func (ie *ImageExif) ExifMake() (string, error) {
	makeTag, err := ie.x.Get(_exif.Make)
	if err != nil {
		return "", errors.New("unable to retrieve camera make")
	}
	cameraMake, err := makeTag.StringVal()
	if err != nil {
		return "", errors.New("invalid camera make")
	}
	return cameraMake, nil
}

func (ie *ImageExif) ExifModel() (string, error) {
	modelTag, err := ie.x.Get(_exif.Model)
	if err != nil {
		return "", errors.New("unable to retrieve camera model")
	}
	model, err := modelTag.StringVal()
	if err != nil {
		return "", errors.New("invalid camera model")
	}
	return model, nil
}

func (ie *ImageExif) ExifFocalLength() (float64, error) {
	focalLength, err := ie.x.Get(_exif.FocalLength)
	if err != nil {
		return 0, err // Skip files without focal length data
	}

	num, denom, err := focalLength.Rat2(0)
	if err != nil {
		return 0, err
	}
	if denom == 0 {
		return 0, err
	}
	return float64(num) / float64(denom), nil
}
