package midnight

import "time"

func GetMidNight() time.Time {
	n := time.Now()
	midnight := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, n.Location())
	return midnight
}
