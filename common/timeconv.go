package common

import "time"

func TimeToString(t time.Time) string {
	return t.Format(TIME_FORMAT)
}
