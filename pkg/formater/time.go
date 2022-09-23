package formater

import "time"

func TimeStampUnixTo(timestamp int64) string {
	return time.Unix(timestamp, 0).Format(time.RFC3339)
}
