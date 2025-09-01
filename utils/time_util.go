package utils

import "time"

func GetTime(tm time.Time) (string, int64) {
	_timeStr := tm.Format("2006-01-02 15:04:05")
	_timestamp := tm.Unix()
	return _timeStr, _timestamp
}

func TimestampToDatetimeStr(tm int64) string {
	_timeUnix := tm
	return time.Unix(_timeUnix, 0).Format("2006-01-02 15:04:05")
}

func TimestampToTime(tm int64) time.Time {
	_timeUnix := TimestampToDatetimeStr(tm)
	_formatTime, _ := time.Parse("2006-01-02 15:04:05", _timeUnix)
	return _formatTime
}

func DatetimeStrToTime(dateTime string) (tm time.Time) {
	_formatTime, _ := time.Parse("2006-01-02 15:04:05", dateTime)
	return _formatTime
}

func DatetimeStrToTimestamp(dateTime string) (tm int64) {
	_formatTime, _ := time.Parse("2006-01-02 15:04:05", dateTime)
	return _formatTime.Unix()
}
