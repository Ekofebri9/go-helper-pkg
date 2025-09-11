package time

import (
	"log/slog"
	"time"
)

var IDNLocation *time.Location

func Init() {
	var err error
	IDNLocation, err = time.LoadLocation(IDNZone)
	if err != nil {
		slog.Warn("failed to load location Asia/Jakarta, fallback to GMT+7", "error", err)
		IDNLocation = time.FixedZone(IDNZone, 7*3600) // Fallback to GMT+7
	}
}

func Now() time.Time {
	return time.Now().In(IDNLocation)
}

func ParseDate(dateStr string) (time.Time, error) {
	return time.ParseInLocation(DateFormat, dateStr, IDNLocation)
}

func ParseTime(timeStr string) (time.Time, error) {
	return time.ParseInLocation(TimeFormat, timeStr, IDNLocation)
}

func ParseDateTime(dateTimeStr string) (time.Time, error) {
	return time.ParseInLocation(DateTimeFormat, dateTimeStr, IDNLocation)
}

func ParseWithFormat(format, value string) (time.Time, error) {
	return time.ParseInLocation(format, value, IDNLocation)
}
