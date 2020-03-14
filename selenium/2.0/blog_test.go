package blog_test

import (
	"flag"
	"testing"

	"github.com/olzhy/test/pages"
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

func TestSearch(t *testing.T) {
	sp := pages.NewSearchPage(driver)
	count, err := sp.Search("istio")
	if nil != err || count < 1 {
		t.Errorf("search error, count: %d, err: %s", count, err)
	}
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
