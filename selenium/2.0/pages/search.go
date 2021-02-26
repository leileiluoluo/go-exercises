package pages

import (
	"errors"
	"fmt"
	"time"

	"github.com/tebeka/selenium"
)

const (
	blogURL                        = "https://leileiluoluo.com/"
	searchButtonIdSelector         = "searchOpen"
	keywordInputIdSelector         = "search-query"
	searchResultLoadingCssSelector = "#search-results #loadingDiv"
	searchResult                   = "#search-results .border-bottom"
)

var drv selenium.WebDriver

type SearchPage struct {
}

// initializer
func NewSearchPage(driver selenium.WebDriver) *SearchPage {
	drv = driver
	return &SearchPage{}
}

// open blog and click search button
func (sp *SearchPage) openBlogAndClickSearchButton() error {
	// open blog
	err := drv.Get(blogURL)
	if nil != err {
		return errors.New(fmt.Sprintf("search page open error, err: %s", err))
	}

	// click search button
	elem, err := drv.FindElement(selenium.ByID, searchButtonIdSelector)
	if nil != err {
		return errors.New(fmt.Sprintf("search button not found, err: %s", err))
	}
	return elem.Click()
}

// type keyword and enter
func (sp *SearchPage) typeKeyword(keyword string) error {
	elem, err := drv.FindElement(selenium.ByID, keywordInputIdSelector)
	if nil != err {
		return errors.New(fmt.Sprintf("keyword input element not found, err: %s", err))
	}
	return elem.SendKeys(keyword + "\n")
}

// wait until search result displayed
func (sp *SearchPage) waitUntilResultDisplayed() error {
	return drv.WaitWithTimeout(func(driver selenium.WebDriver) (bool, error) {
		elem, err := driver.FindElement(selenium.ByCSSSelector, searchResultLoadingCssSelector)
		if nil != err {
			return false, nil
		}
		visible, err := elem.IsDisplayed()
		return !visible, err
	}, 30*time.Second)
}

// Search by keyword
// return count of search result
func (sp *SearchPage) Search(keyword string) (int, error) {
	// open blog and click search button
	err := sp.openBlogAndClickSearchButton()
	if nil != err {
		return 0, err
	}

	// type keyword and enter
	err = sp.typeKeyword(keyword)
	if nil != err {
		return 0, err
	}

	// wait until search result displayed
	err = sp.waitUntilResultDisplayed()
	if nil != err {
		return 0, err
	}

	// return
	elems, err := drv.FindElements(selenium.ByCSSSelector, searchResult)
	if nil != err {
		return 0, errors.New(fmt.Sprintf("search element error, err: %s", err))
	}
	return len(elems), nil
}
