package utils

import (
	"log"
	"time"
)

func MarkStartTime() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func MarkEndTime(startTime int64) {
	endTime := time.Now().UnixNano() / int64(time.Millisecond)
	log.Println("end time:", endTime)
	log.Println("execution time:", endTime-startTime)
}
