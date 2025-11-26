package number

import (
	"log"
	"strconv"
	"time"
)

func NumberGetFormattedRupiah(money int) string {
	stringifiedMoney := strconv.Itoa(money)

	return "Rp" + stringifiedMoney
}

func StringToUint(numberedStr string) uint {
	id, err := strconv.ParseUint(numberedStr, 10, 64)
	if err != nil {
		log.Printf("Gagal konversi string ke uint: %v", err)
		return 0
	}
	return uint(id)
}

func TimeToNumber(t time.Time) int64 {
	return t.UnixNano()
}

func NumberToDays(startNano, endNano int64) int64 {
	const dayInNano = int64(24 * time.Hour)
	return (endNano - startNano) / dayInNano
}

func HourToUnix(val int) int64 {
	return time.Now().Add(time.Hour * time.Duration(val)).Unix()
}

func StringToInt(val string) (*int, error) {
	res, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
