package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// Open the file for reading
	file, e := os.Open("config.txt")
	if e != nil {
		if os.IsNotExist(e) {
			fmt.Println("There is no such file")
		} else {
			fmt.Println("Error opening file:", e)
		}
		return
	}
	defer file.Close()

	var certName string
	var certPasswd string

	// Create a new scanner
	scanner := bufio.NewScanner(file)

	// Read and trim the first line
	if scanner.Scan() {
		certName = strings.TrimSpace(scanner.Text())
	}

	// Read and trim the second line
	if scanner.Scan() {
		certPasswd = strings.TrimSpace(scanner.Text())
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("certName: ", certName)
	fmt.Println("certPasswd: ", certPasswd)

	// Set up allocator options to disable headless mode
	allocCtx, _ := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)...,
	)
	// defer cancel()

	// Create a new browser context using the allocator context
	ctx, _ := chromedp.NewContext(
		allocCtx,
		// chromedp.WithDebugf(log.Printf),
	)
	// defer cancel()

	if err := chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(`https://www.hometax.go.kr/`),
		chromedp.Sleep(3 * time.Second),
		chromedp.Click(`#textbox81212912`),
		chromedp.Sleep(2 * time.Second),

		chromedp.WaitVisible(`//*[@id="anchor22"]/span`),
		chromedp.Click(`//*[@id="anchor22"]/span`),

		chromedp.WaitVisible(`//*[@id="columntabledataTable"]/div[1]`),
		chromedp.Click(`//*[@id="columntabledataTable"]/div[1]`),
		chromedp.Click(fmt.Sprintf(`//*[@title="%s"]`, certName)),
		chromedp.WaitVisible(`//*[@id="input_cert_pw"]`),
		chromedp.SendKeys(`//*[@id="input_cert_pw"]`, certPasswd),
		chromedp.Sleep(1 * time.Second),

		chromedp.Click(`//*[@id="btn_confirm_iframe"]/span`),
	}); err != nil {
		log.Fatal(err)
	}

	log.Println(">> Success!! Do your work!!")
}
