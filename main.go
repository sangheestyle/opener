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
	file, err := os.Open("config.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("There is no such file")
		} else {
			fmt.Println("Error opening file:", err)
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
	allocCtx, cancel := chromedp.NewExecAllocator(
		context.Background(),
		append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
		)...,
	)
	defer cancel()

	// Create a new browser context using the allocator context
	ctx, cancel := chromedp.NewContext(allocCtx)
	defer cancel()

	// Navigate to the website
	if err := chromedp.Run(ctx, chromedp.Navigate(`https://www.hometax.go.kr/`)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)

	// Click the login button
	if err := chromedp.Run(ctx, chromedp.Click(`#textbox81212912`)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)

	// Switch to the iframe
	if err := chromedp.Run(ctx, chromedp.WaitVisible(`#txppIframe`, chromedp.ByID), chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.Run(ctx, chromedp.Click(`#txppIframe`))
	})); err != nil {
		log.Fatal(err)
	}

	// Click the certificate button
	if err := chromedp.Run(ctx, chromedp.Click(`//*[@id="anchor22"]/span`)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)

	// Switch to another iframe
	if err := chromedp.Run(ctx, chromedp.WaitVisible(`#dscert`, chromedp.ByID), chromedp.ActionFunc(func(ctx context.Context) error {
		return chromedp.Run(ctx, chromedp.Click(`#dscert`))
	})); err != nil {
		log.Fatal(err)
	}

	// Select the certificate and click
	if err := chromedp.Run(ctx, chromedp.Click(`//*[@id="columntabledataTable"]/div[1]`), chromedp.Click(fmt.Sprintf(`//*[@title="%s"]`, certName))); err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)

	// Enter the password
	if err := chromedp.Run(ctx, chromedp.SendKeys(`#input_cert_pw`, certPasswd)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)

	// Click the confirmation button
	if err := chromedp.Run(ctx, chromedp.Click(`#btn_confirm_iframe`)); err != nil {
		log.Fatal(err)
	}
}
