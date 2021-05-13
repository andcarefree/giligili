package cache

import (
	"fmt"
	"strconv"
)

const (
	DailyRankKey = "rank:daily"
)

func VideoClickKey(id uint) string {
	return fmt.Sprintf("click:video:%s",strconv.FormatUint(uint64(id),10))
}
 
