# link

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
