package selenium_test

import (
	"fmt"
	"testing"

	"sourcegraph.com/sourcegraph/go-selenium"
)

func TestSingle(test *testing.T) {
	var webDriver selenium.WebDriver
	var err error
	caps := selenium.Capabilities(map[string]interface{}{
		"browserName":    "firefox",
		"platformName":   "Windows 10",
		"browserVersion": "latest",
		"build":          "golangsuite",
		"name":           "Testcase01",
	})
	if webDriver, err = selenium.NewRemote(caps, "http://anishmehta:2fcF9J3BAFO7UoVjT3NAl9LbO5p1VgSpKXGq8XO2t0WUa4189E@hub.lambdatest.com/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("https://www.lambdatest.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	err = webDriver.Get("https://www.google.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	err = webDriver.Get("https://www.facebook.com/")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

}
