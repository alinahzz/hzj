package util

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"time"
)

func Int64ToInt(i int64) int{

	str := strconv.FormatInt(i, 10)

	v, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return v

}

func RandomInt(max int) int{

	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))

	if err != nil {
		return 0
	}

	return Int64ToInt(n.Int64())

}

func TodayStartTime() int64{

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())

	return today.Unix()

}

func NextDayStartTime() int64{

	now := time.Now()
	today := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	tomorrow := today.AddDate(0, 0, 1)

	return tomorrow.Unix()
}


