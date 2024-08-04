package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	var newdt time.Time
	rd := bufio.NewReader(os.Stdin)
	timeStr, _ := rd.ReadString('\n')

	// Удаление последнего символа при помощи пакета "strings".
	timeStr = strings.TrimSuffix(timeStr, "\n")

	firstTime, _ := time.Parse("2006-01-02 15:04:00", timeStr)
	newdt = firstTime.Add(time.Hour * 24)
	if firstTime.Hour() <= 13 {
		fmt.Println(timeStr)
	} else {
		fmt.Print(newdt.Format("2006-01-02 15:04:05"))
	}
}
