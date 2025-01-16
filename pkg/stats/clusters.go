package stats

// Focal length clusters for grouping similar focal lengths
var focalLengthClusters = []struct {
	Min    float64
	Max    float64
	Mapped float64
}{
	{0, 21, 16},      // Ultra-wide
	{21, 29, 24},     // Wide-angle
	{29, 45, 35},     // Standard wide
	{45, 64, 50},     // Standard prime
	{64, 90, 85},     // Portrait prime
	{90, 106, 100},   // Telephoto prime
	{106, 139, 135},  // Medium telephoto
	{139, 201, 150},  // Long telephoto
	{201, 301, 250},  // Super telephoto
	{301, 1000, 400}, // Extreme telephoto
}

// getClusteredFocalLength maps a focal length to its cluster.
func getClusteredFocalLength(focalLength float64) float64 {
	for _, cluster := range focalLengthClusters {
		if focalLength >= cluster.Min && focalLength <= cluster.Max {
			return cluster.Mapped
		}
	}
	return focalLength // Return original if no cluster matches
}
