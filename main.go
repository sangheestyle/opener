package main

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func main() {
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
	if err := chromedp.Run(ctx, chromedp.Click(`//*[@id="columntabledataTable"]/div[1]`), chromedp.Click(`//*[@title="인증서 명"]`)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)

	// Enter the password
	passwd := "패스워드 입력"
	if err := chromedp.Run(ctx, chromedp.SendKeys(`#input_cert_pw`, passwd)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(3 * time.Second)

	// Click the confirmation button
	if err := chromedp.Run(ctx, chromedp.Click(`#btn_confirm_iframe`)); err != nil {
		log.Fatal(err)
	}
}
