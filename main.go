package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	//"strings"
	"encoding/csv"
	"os"
)

type Scraper struct {
	URL string
}

func (s *Scraper) ScrapePage(url string) {
	resp, _ := http.Get(url)
	bytes, _ := ioutil.ReadAll(resp.Body)
	s.URL = string(bytes)
	resp.Body.Close()
}

type URLcaller struct {
    Sc Scraper
}

func (u *URLcaller) CallByCSV(file string) {
	f, _ := os.Open(file)
	defer f.Close()

	data, _ := csv.NewReader(f).ReadAll()

	for _, record := range data {
		u.Sc.ScrapePage(record[0])
		fmt.Println(u.Sc.URL) // Add this line to print the HTML.
	}
}


func main() {
	call := URLcaller{Sc: Scraper{}}
	call.CallByCSV("urls.csv")
	fmt.Println(call.Sc.URL)
}
