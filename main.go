package main

import (
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"

)

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
		url:  "https://www.bestbuy.com/site/amd-ryzen-9-5900x-4th-gen-12-core-24-threads-unlocked-desktop-processor-without-cooler/6438942.p?skuId=6438942",
		name: "5900x",
	},
}

func isInStockBB(cpu LinkForCPU) {
	res, err := http.Get(cpu.url);
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()
	if res.StatusCode >= 300 {
		log.Fatal("Status code >= 300");
	}



}

func main() {
	const bb = ""
	isInStockBB(Links[0])

	//client.Get()
}
