package stats

import (
	"fmt"
	"sort"
	"strings"
)

func NewFocalStats() *FocalStats {
	return &FocalStats{
		totalPhotos:    0,
		focalLengthMap: make(map[float64]int),
	}
}

type FocalStats struct {
	focalLengthMap map[float64]int
	totalPhotos    int
}

func (fs *FocalStats) AddClusteredFocalLength(length float64) {
	clusteredFocalLength := getClusteredFocalLength(length)
	fs.AddFocalLength(clusteredFocalLength)
}

func (fs *FocalStats) AddFocalLength(length float64) {
	fs.focalLengthMap[length]++
	fs.totalPhotos++
}

func (fs *FocalStats) GetMostUsedFocalLength() (float64, int) {
	var mostUsedLength float64
	maxCount := 0
	for length, count := range fs.focalLengthMap {
		if count > maxCount {
			mostUsedLength = length
			maxCount = count
		}
	}
	return mostUsedLength, maxCount
}

func (fs *FocalStats) TotalFotos() int {
	return fs.totalPhotos
}

func (fs *FocalStats) String() string {
	lengths := make([]float64, 0, len(fs.focalLengthMap))
	for length := range fs.focalLengthMap {
		lengths = append(lengths, length)
	}
	sort.Float64s(lengths)

	var sb strings.Builder
	for _, length := range lengths {
		sb.WriteString(fmt.Sprintf("%.2fmm: %d\n", length, fs.focalLengthMap[length]))
	}
	return sb.String()
}
