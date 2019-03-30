package common

import "testing"

func TestPage_GetTotalPages(t *testing.T) {
	page := Page{
		StartIndex: 0,
		TotalRows:  21,
		PageRows:   20,
	}
	total := page.GetTotalPages()
	t.Log("assert total page is ", total)
	if total != 2 {
		t.Fatal(total)
	}
}

func TestPage_CurrentPage(t *testing.T) {
	page := Page{
		StartIndex: 0,
		TotalRows:  21,
		PageRows:   20,
	}
	currentPage := page.CurrentPage()
	t.Log("assert currentPage page is 1 actual is ", currentPage)
	if currentPage != 1 {
		t.Fatal(currentPage)
	}

	page.SetCurrentPage(2)

	currentPage = page.CurrentPage()
	t.Log("assert currentPage page is 2 actual is ", currentPage)
	if currentPage != 2 {
		t.Fatal(currentPage)
	}

	t.Log("assert startIndex is 20 actual is ", page.StartIndex)
	if page.StartIndex != 20 {
		t.Fatal(page.StartIndex)
	}
}
