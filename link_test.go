package link

import "testing"

func TestPagination(t *testing.T) {
	cases := map[string]struct {
		page       int
		perPage    int
		totalCount int
		url        string
		expected   string
	}{
		"pattern 1": {1, 30, 100, "http://foo/bar", `<http://foo/bar?page=4&per_page=30>; rel="last", <http://foo/bar?page=2&per_page=30>; rel="next"`},
		"pattern 2": {2, 30, 100, "http://foo/bar", `<http://foo/bar?page=1&per_page=30>; rel="first", <http://foo/bar?page=4&per_page=30>; rel="last", <http://foo/bar?page=3&per_page=30>; rel="next", <http://foo/bar?page=1&per_page=30>; rel="prev"`},
		"pattern 3": {3, 30, 100, "http://foo/bar", `<http://foo/bar?page=1&per_page=30>; rel="first", <http://foo/bar?page=4&per_page=30>; rel="last", <http://foo/bar?page=4&per_page=30>; rel="next", <http://foo/bar?page=2&per_page=30>; rel="prev"`},
		"pattern 4": {4, 30, 100, "http://foo/bar", `<http://foo/bar?page=1&per_page=30>; rel="first", <http://foo/bar?page=3&per_page=30>; rel="prev"`},
		"pattern 5": {1, 30, 130, "http://foo/bar", `<http://foo/bar?page=5&per_page=30>; rel="last", <http://foo/bar?page=2&per_page=30>; rel="next"`},
		"pattern 6": {2, 30, 130, "http://foo/bar", `<http://foo/bar?page=1&per_page=30>; rel="first", <http://foo/bar?page=5&per_page=30>; rel="last", <http://foo/bar?page=3&per_page=30>; rel="next", <http://foo/bar?page=1&per_page=30>; rel="prev"`},
		"pattern 7": {3, 30, 130, "http://foo/bar", `<http://foo/bar?page=1&per_page=30>; rel="first", <http://foo/bar?page=5&per_page=30>; rel="last", <http://foo/bar?page=4&per_page=30>; rel="next", <http://foo/bar?page=2&per_page=30>; rel="prev"`},
		"pattern 8": {4, 30, 130, "http://foo/bar", `<http://foo/bar?page=1&per_page=30>; rel="first", <http://foo/bar?page=5&per_page=30>; rel="last", <http://foo/bar?page=5&per_page=30>; rel="next", <http://foo/bar?page=3&per_page=30>; rel="prev"`},
		"pattern 9": {5, 30, 130, "http://foo/bar", `<http://foo/bar?page=1&per_page=30>; rel="first", <http://foo/bar?page=4&per_page=30>; rel="prev"`},
	}

	for k, tc := range cases {
		if actual := Pagination(tc.page, tc.perPage, tc.totalCount, tc.url); actual != tc.expected {
			t.Errorf("%s: Pagination got %v, want %v", k, actual, tc.expected)
		}
	}
}

func TestGetLastPage(t *testing.T) {
	cases := map[string]struct {
		total    int
		perPage  int
		expected int
	}{
		"pattern 1": {50, 30, 2},
		"pattern 2": {50, 20, 3},
		"pattern 3": {50, 10, 5},
		"pattern 4": {100, 30, 4},
		"pattern 5": {100, 20, 5},
		"pattern 6": {100, 10, 10},
	}

	for k, tc := range cases {
		if actual := getLastPage(tc.total, tc.perPage); actual != tc.expected {
			t.Errorf("%s: getLastPage got %v, want %v", k, actual, tc.expected)
		}
	}
}
