package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)

// backup 1.go
func main() {
	antalScan := 0
	var EnablePrint = false
	var TTime time.Time
	var x byte
	x = ';'
	checkx := 0
	contentFile, err := ioutil.ReadFile("ExampleOfData.csv")
	if err != nil {
	}
	csvstring := string(contentFile)

	ScanCSV := func(data []byte, atEOF bool) (advance int, token []byte, err error) {

		commaidx := bytes.IndexByte(data, x)

		if commaidx > 0 {
			buffer2 := data[:commaidx]
			return commaidx + 1, bytes.TrimSpace(buffer2), nil
		}

		if atEOF {

			if len(data) > 0 {
				return len(data), bytes.TrimSpace(data), nil
			}
		}

		return 0, nil, nil
	}

	scanner := bufio.NewScanner(strings.NewReader(csvstring))
	scanner.Split(ScanCSV)

	for scanner.Scan() {
		//flip x between '\n' en ';' according to a value
		antalScan++
		if x == ';' {
			checkx++
			if checkx == 3 {
				EnablePrint = true
				x = '\n'
				str := scanner.Text()
				t, err := time.Parse(time.RFC3339, str)
				if err != nil {
					fmt.Println(err)
				}
				TTime = t
			}
		} else if checkx <= 4 {
			EnablePrint = false
			checkx = 0
			x = ';'
		}

		if EnablePrint {
			fmt.Print(checkx)
			fmt.Print("-----------")
			fmt.Println(TTime)
		} else {
			fmt.Println(scanner.Text())
		}
	}

}
