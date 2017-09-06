package praytimes

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestGetTimes(t *testing.T) {
	times := GetTimes(time.Now(), []float64{43, -80}, -5, 0, "")

	fmt.Println("Prayer Times for today in Waterloo/Canada")
	for _, val := range []string{"Fajr", "Sunrise", "Dhuhr", "Asr", "Maghrib", "Isha", "Midnight"} {
		fmt.Printf("%v: %v\n", val, times[strings.ToLower(val)])
	}

	fmt.Println()
	times = GetTimes(time.Now(), []float64{-6.9034443, 107.5731164}, 7)
	fmt.Println("Prayer Times for today in Bandung/Indonesia")
	for _, val := range []string{"Fajr", "Sunrise", "Dhuhr", "Asr", "Maghrib", "Isha", "Midnight"} {
		fmt.Printf("%v: %v\n", val, times[strings.ToLower(val)])
	}

	fmt.Println()
}

func TestAdjust(t *testing.T) {
	setting := GetSetting()

	fmt.Println("Setting for", GetMethod(), "method")
	for _, val := range []string{"highLats", "midnight", "imsak", "fajr", "dhuhr", "asr", "maghrib", "isha"} {
		fmt.Printf("%v: %v | ", val, setting[val])
	}

	fmt.Print("\n")

	SetMethod("Egypt")
	setting = GetSetting()
	fmt.Println("Setting for", GetMethod(), "method")

	for _, val := range []string{"highLats", "midnight", "imsak", "fajr", "dhuhr", "asr", "maghrib", "isha"} {
		fmt.Printf("%v: %v | ", val, setting[val])
	}
	fmt.Print("\n")
}
