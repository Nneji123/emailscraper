package emailscraper

import (
    "context"
    "fmt"
    "time"

    "github.com/go-rod/rod"
    "github.com/go-rod/rod/lib/launcher"
    "github.com/gocolly/colly/v2"
)

func initiateScrapingFromRod(response *colly.Response, timeout int) error {
    ctx := context.Background()

    // Launch Rod browser with options
    browser := rod.New().ControlURL(launcher.New().Headless(true).MustLaunch())

    // Handle timeout
    if timeout > 0 {
        ctx, _ = context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
    }

    page := browser.MustPage(response.Request.URL.String())

    // Get the HTML content of the page
    html, err := page.HTML()
    if err != nil {
        return fmt.Errorf("error getting HTML: %w", err)
    }

    response.Body = []byte(html)

    return nil
}
