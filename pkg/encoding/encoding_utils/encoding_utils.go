package encoding_utils

import (
    "strconv"
    "time"
)

func Substring(str string, pos int, length int) (res string) {
	runes := []rune(str)
	return string(runes[pos:length])
}

func TrimLeftChars(s string, n int) string {
    m := 0
    for i := range s {
        if m >= n {
            return s[i:]
        }
        m++
    }
    return s[:0]
}

func HexDate(str string) (d time.Time, err error) {
    if(str == "FFFF") {
        return time.Time{}, nil
    }
    n, err := strconv.ParseUint(str, 16, 32)
    date := time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC)
    date = date.AddDate(0,0,int(n))
    return date,err

}