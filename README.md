# link

[![CircleCI](https://circleci.com/gh/tchssk/link.svg?style=shield&circle-token=3be34073d5a63f25fc6aff4e39487afead79df72)](https://circleci.com/gh/tchssk/link)

Creates the Link Header like pagination of GitHub API.

## Installation

```sh
$ go get github.com/tchssk/link
```

## Usage

```go
l := link.Pagination(page, perPage, totalCount, url)
http.ResponseWriter.Header().Set("Link", l)
```

## License

MIT License

## Author

Taichi Sasaki
