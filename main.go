package main

import (
	"crypto/tls"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/http"
	"os"
	"sync"
	"time"
)

type TopURL struct {
	GlobalRanking string
	DomainRanking string
	Domain        string
	Country       string
}

func DownloadMillionDomains() error {
	if _, err := os.Stat("majestic_million.csv"); errors.Is(err, os.ErrNotExist) {
		log.Println("Downloading million domain file...")

		url := "https://downloads.majestic.com/majestic_million.csv"
		out, err := os.Create("majestic_million.csv")
		if err != nil {
			return err
		}
		defer out.Close()

		resp, err := http.Get(url)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			return fmt.Errorf("bad status: %s", resp.Status)
		}

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			return err
		}
		log.Println("Download finalized!")
	} else {
		log.Println("Domain CSV file already exists!")
	}

	return nil
}

func CreateUrlList(urlList *os.File) []TopURL {
	csvReader := csv.NewReader(urlList)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var urls []TopURL
	for _, line := range data {
		if line[3] == "br" {
			url := TopURL{
				GlobalRanking: line[0],
				DomainRanking: line[1],
				Domain:        line[2],
				Country:       line[3],
			}
			urls = append(urls, url)
		}
	}

	return urls
}

func CheckSSL(server, port, ranking string) {
	conn, err := tls.DialWithDialer(&net.Dialer{Timeout: 180 * time.Second}, "tcp", "www."+server+":"+port, nil)
	if err != nil {
		log.Printf("Server %s ranking %s não suporta certificado SSL", server, ranking)
		return
	}
	defer conn.Close()

	expiresAt := conn.ConnectionState().PeerCertificates[0].NotAfter
	currentTime := time.Now()
	diff := expiresAt.Sub(currentTime)
	log.Printf("Tempo restante para o server %s expirar: %1.f dia(s). Ranking %s", server, math.Round(diff.Hours()/24), ranking)
}

func main() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	err = DownloadMillionDomains()
	if err != nil {
		log.Fatal(err)
	}

	domainFile, err := os.Open("majestic_million.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer domainFile.Close()

	urls := CreateUrlList(domainFile)

	var wg sync.WaitGroup
	for i := 0; i < len(urls); i++ {
		wg.Add(1)
		go func(url TopURL) {
			CheckSSL(url.Domain, "443", url.DomainRanking)
			defer wg.Done()
		}(urls[i])
	}
	wg.Wait()
}
