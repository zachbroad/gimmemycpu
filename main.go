package main

import (
	"gopkg.in/go-toast/toast.v1"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

import "fmt"

type LinkForCPU struct {
	url  string
	name string
}

var links = []LinkForCPU{
	{
		url:  "https://www.bestbuy.com/site/amd-ryzen-9-5900x-4th-gen-12-core-24-threads-unlocked-desktop-processor-without-cooler/6438942.p?skuId=6438942",
		name: "5900x",
	},
	{
		url:  "https://www.amazon.com/AMD-Ryzen-5900X-24-Thread-Processor/dp/B08164VTWH/ref=cm_cr_arp_d_product_top?ie=UTF8",
		name: "5900x",
	},
	{
		url:  "https://www.bhphotovideo.com/c/product/1598373-REG/amd_100_100000061wof_ryzen_9_5900x_3_7.html?SID=trd-us-9843011037592730000",
		name: "5900x",
	},
	{
		url:  "https://www.newegg.com/amd-ryzen-9-5900x/p/N82E16819113664",
		name: "5900x",
	},
	{
		url:  "https://www.bestbuy.com/site/amd-ryzen-9-5950x-4th-gen-16-core-32-threads-unlocked-desktop-processor-without-cooler/6438941.p?skuId=6438941",
		name: "5950x",
	},
}

var outOfStockStrings = []string{
	"currently unavailable",
	"out of stock",
	"sold out",
	"currently sold out",
	"notify when available",
}

func main() {
	var wg sync.WaitGroup

	for _, link := range links {
		wg.Add(1)
		go func(cpu LinkForCPU) {
			for {
				_ = checkLink(cpu)
				time.Sleep(5 * time.Second)
			}
			wg.Done()
		}(link)
	}
	wg.Wait()
}

func checkLink(cpu LinkForCPU) bool {
	req, _ := http.NewRequest("GET", cpu.url, nil)
	req.Header.Set("Accept", "text/html,application/xhtml+xml")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.103 Safari/537.36")

	url, err := url.Parse(cpu.url)
	if err != nil {
		log.Fatalf("Error parsing URL: %s", err)
	}

	client := &http.Client{}
	res, _ := client.Do(req)

	defer res.Body.Close()
	if res.StatusCode >= 300 {
		log.Fatal(res.StatusCode)
	}

	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	bodystr := string(bs)

	return checkIfInStock(bodystr, cpu, url)
}

func checkIfInStock(body string, cpu LinkForCPU, url *url.URL) bool {
	for _, str := range outOfStockStrings {
		if strings.Contains(
			strings.ToLower(body),
			strings.ToLower(str),
		) {
			fmt.Printf("%s not in stock at %s!\n", cpu.name, url.Host)
			return false
		}

	}

	inStockAlert(cpu, url)
	return true
}

func inStockAlert(cpu LinkForCPU, url *url.URL) {
	sendTextMessage()
	fmt.Printf("%s is in stock at %s!\n", cpu.name, url.Host)

	t := toast.Notification{
		AppID:               "GimmeMyCPU",
		Title:               "Stock Alert!",
		Message:             fmt.Sprintf("%s is in stock at %s!", cpu.name, url.Host),
		Icon:                "",
		ActivationType:      "",
		ActivationArguments: cpu.url,
		Actions:             nil,
		Audio:               toast.LoopingAlarm9,
		Loop:                true,
	}
	_ = t.Push()
}

func sendTextMessage() {
	println("TODO: Send text message")
}
