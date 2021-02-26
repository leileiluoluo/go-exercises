package blog_test

import (
	"flag"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/tebeka/selenium"
)

var (
	browserName                    = flag.String("browser", "chrome", "browser name")
	gridUrl                        = flag.String("grid", "http://localhost:4444/wd/hub", "grid url")
	blogURL                        = "https://leileiluoluo.com/"
	searchButtonIdSelector         = "searchOpen"
	keywordInputIdSelector         = "search-query"
	searchResultLoadingCssSelector = "#search-results #loadingDiv"
	searchResult                   = "#search-results .border-bottom"

	keyword = "istio"
)

var driver selenium.WebDriver

func setup() func() {
	// new remote driver
	caps := selenium.Capabilities{"browserName": *browserName}
	webDriver, err := selenium.NewRemote(caps, *gridUrl)
	if nil != err {
		panic(err)
	}
	driver = webDriver

	// teardown
	return func() {
		driver.Quit()
	}
}

func screenshot(filename string) {
	bytes, err := driver.Screenshot()
	if nil != err {
		log.Printf("take screenshot error, err: %s", err)
		return
	}

	err = ioutil.WriteFile(filename, bytes, 0666)
	if nil != err {
		log.Printf("save screenshot error, err: %s", err)
	}
}

func TestSearch(t *testing.T) {
	// open blog
	err := driver.Get(blogURL)
	if nil != err {
		t.Errorf("search page open error, err: %s", err)
	}

	// click search button
	elem, err := driver.FindElement(selenium.ByID, searchButtonIdSelector)
	if nil != err {
		t.Errorf("search button not found, err: %s", err)
	}
	elem.Click()

	// type keyword and enter
	elem, err = driver.FindElement(selenium.ByID, keywordInputIdSelector)
	if nil != err {
		t.Errorf("keyword input element not found, err: %s", err)
	}
	elem.SendKeys(keyword + "\n")

	// wait until search result displayed
	driver.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		elem, err = driver.FindElement(selenium.ByCSSSelector, searchResultLoadingCssSelector)
		if nil != err {
			return false, nil
		}
		visible, err := elem.IsDisplayed()
		return !visible, err
	}, 30*time.Second)

	// assert
	elems, err := driver.FindElements(selenium.ByCSSSelector, searchResult)
	if nil != err || len(elems) < 1 {
		t.Errorf("no search result, err: %s", err)
	}

	// save screenshot
	screenshot("search.png")
}

func TestMain(m *testing.M) {
	// parse flags
	flag.Parse()

	// setup / teardown
	teardown := setup()
	defer teardown()

	// run tests
	m.Run()
}
