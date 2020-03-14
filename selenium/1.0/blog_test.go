package blog_test

import (
	"flag"
	"io/ioutil"
	"log"
	"testing"

	"github.com/tebeka/selenium"
)

var (
	browserName = flag.String("browser", "chrome", "browser name")
	gridUrl     = flag.String("grid", "http://localhost:4444/wd/hub", "grid url")
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
	// open search page
	err := driver.Get("https://leileiluoluo.com/")
	if nil != err {
		t.Errorf("search page open error, err: %s", err)
	}

	// type keyword
	elem, err := driver.FindElement(selenium.ByID, "s")
	if nil != err {
		t.Errorf("find element error, err: %s", err)
	}
	err = elem.SendKeys("istio")
	if nil != err {
		t.Errorf("send keys error, err: %s", err)
	}

	// click search
	elem, err = driver.FindElement(selenium.ByID, "searchsubmit")
	if nil != err {
		t.Errorf("find element error, err: %s", err)
	}
	elem.Click()

	// assert
	elems, err := driver.FindElements(selenium.ByCSSSelector, "h3>a")
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
