/*
Package link provides a function to create the Link Header like pagination of GitHub API.
*/
package link

import (
	"fmt"
	"strings"
)

var (
	// DefaultPage is the default value of Page.
	DefaultPage = 1

	// DefaultPerPage is the default value of PerPage.
	DefaultPerPage = 30

	// DefaultTotalCount is the default value of TotalCount.
	DefaultTotalCount = 1
)

// Pagination returns a string for the Link Header.
func Pagination(page, perPage, totalCount int, url string) string {
	if page < 1 {
		page = DefaultPage
	}

	if perPage < 1 {
		perPage = DefaultPerPage
	}

	if totalCount < 1 {
		totalCount = DefaultTotalCount
	}

	lastPage := getLastPage(totalCount, perPage)

	var links []string
	for _, v := range []struct {
		rel  string
		page int
	}{
		{"first", 1},
		{"last", lastPage},
		{"next", page + 1},
		{"prev", page - 1},
	} {
		if v.page > lastPage || v.page < 1 || v.page == page {
			continue
		}
		links = append(links, fmt.Sprintf(`<%s?page=%d&per_page=%d>; rel="%s"`, url, v.page, perPage, v.rel))
	}

	return strings.Join(links, ", ")
}

func getLastPage(total, perPage int) int {
	if (total % perPage) != 0 {
		return (total / perPage) + 1
	}
	return (total / perPage)
}
