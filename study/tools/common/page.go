package common

type Page struct {
	StartIndex int
	PageRows   int
	TotalRows  int
}

func (page *Page) CurrentPage() int {
	if page.TotalRows < 0 {
		return 1
	}
	return (page.StartIndex + page.PageRows) / page.PageRows
}

func (page *Page) SetCurrentPage(p int) {
	if page.PageRows <= 0 {
		page.StartIndex = 0
	} else {
		page.StartIndex = (p - 1) * page.PageRows
	}
}

func (page *Page) GetTotalPages() int {
	var totalPages int
	if page.PageRows <= 0 {
		totalPages = 1
	} else {
		totalPages = (page.TotalRows + page.PageRows - 1) / page.PageRows
	}
	if totalPages < 1 {
		return 1
	}
	return totalPages
}

//private int StartIndex;
//private int pageRows;
//private int totalRows;
//
//public EqlPage() {
//
//}
//
//public EqlPage(int StartIndex, int pageRows) {
//this.setStartIndex(StartIndex);
//this.setPageRows(pageRows);
//}
//
//public int getStartIndex() {
//return StartIndex;
//}
//
//public void setStartIndex(int StartIndex) {
//this.StartIndex = StartIndex;
//}
//
//public int getPageRows() {
//return pageRows;
//}
//
//public void setPageRows(int pageRows) {
//this.pageRows = pageRows;
//}
//
//public int getTotalRows() {
//return totalRows;
//}
//
//public void setTotalRows(int totalRows) {
//this.totalRows = totalRows;
//}
//
//public int getCurrentPage() {
//return getPageRows() <= 0 ? 1 : (getStartIndex() + getPageRows()) / getPageRows();
//}
//
//public void setCurrentPage(int currentPage) {
//this.StartIndex = getPageRows() <= 0 ? 0 : (currentPage - 1) * getPageRows();
//}
//
//public int getTotalPages() {
//int totalPages = getPageRows() <= 0 ? 1 : (getTotalRows() + getPageRows() - 1) / getPageRows();
//return totalPages < 1 ? 1 : totalPages;
//}
