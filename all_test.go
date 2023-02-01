package spidate

import (
	"fmt"
	"github.com/SmartPrintsInk/prettyfie"
	"testing"
	"time"
)

func TestAll(t *testing.T) {
	fmt.Printf("Test %+v\n", prettyfie.Pretty(GetMonthData(time.February, MySQL)))
	fmt.Printf("Today %+v\n", prettyfie.Pretty(GetToday(Human)))
	fmt.Printf("Today Add Days %+v\n", prettyfie.Pretty(AddDaysToToday(30, Human)))
	fmt.Printf("Range %+v\n", prettyfie.Pretty(AddDaysToDate("2022-12-01", 45, MongoDB)))
}
