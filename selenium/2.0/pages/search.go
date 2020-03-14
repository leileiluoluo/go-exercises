package pages

import (
	"errors"
	"fmt"

	"github.com/tebeka/selenium"
)

const (
	typeBy   = "#id"
	clickBy  = "#searchsubmit"
	resultBy = "h3>a"

	searchPage = "https://leileiluoluo.com/"
)

var drv selenium.WebDriver

type SearchPage struct {
}

// initializer
func NewSearchPage(driver selenium.WebDriver) *SearchPage {
	drv = driver
	return &SearchPage{}
}

// open search page
func (sp *SearchPage) openSearchPage() error {
	err := drv.Get(searchPage)
	if nil != err {
		return errors.New(fmt.Sprintf("search page open error, err: %s", err))
	}
	return nil
}

// type keyword
func (sp *SearchPage) typeKeyword(keyword string) error {
	elem, err := drv.FindElement(selenium.ByID, "s")
	if nil != err {
		return errors.New(fmt.Sprintf("find element error, err: %s", err))
	}

	err = elem.SendKeys(keyword)
	if nil != err {
		return errors.New(fmt.Sprintf("send keys error, err: %s", err))
	}
	return nil
}

// click search
func (sp *SearchPage) clickSearch() error {
	elem, err := drv.FindElement(selenium.ByID, "searchsubmit")
	if nil != err {
		return errors.New(fmt.Sprintf("find element error, err: %s", err))
	}
	elem.Click()
	return nil
}

// Search by keyword
// return count of search result
func (sp *SearchPage) Search(keyword string) (int, error) {
	// open search page
	err := sp.openSearchPage()
	if nil != err {
		return 0, err
	}

	// type keyword
	err = sp.typeKeyword(keyword)
	if nil != err {
		return 0, err
	}

	// click search
	err = sp.clickSearch()
	if nil != err {
		return 0, err
	}

	// search result
	elems, err := drv.FindElements(selenium.ByCSSSelector, "h3>a")
	if nil != err {
		return 0, errors.New(fmt.Sprintf("find element error, err: %s", err))
	}
	return len(elems), nil
}
