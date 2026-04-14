// Stats structure for metrics collection from storage
package models

// Stats: structure for statistics with metrics
type Stats struct {
	Hits   int64 // Good reads (key not expired and found)
	Misses int64 // Bad reads (key expired or not found)
	Keys   int
}

// HitRate: method just to get percent of successful storage calls lol
func (s Stats) HitRate() float64 {
	total := s.Hits + s.Misses
	if total == 0 {
		return 0
	}
	return float64(s.Hits) / float64(total) * 100
}
