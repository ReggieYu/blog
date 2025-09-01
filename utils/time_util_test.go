package utils

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeUtil(t *testing.T) {
	fmt.Println(GetTime(time.Now()))
	fmt.Println(TimestampToTime(1655124564))
	fmt.Println(DatetimeStrToTime("2025-09-01 23:39:00"))
	fmt.Println(DatetimeStrToTimestamp("2025-09-01 23:39:00"))
}
