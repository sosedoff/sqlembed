# sqlembed

A tiny tool to embed `*.sql` files into a Go package

## Installation

With Go:

```bash
go get -u github.com/sosedoff/sqlembed
```

## Usage

See a list of available flags with:

```bash
$ sqlembed -help

Usage of sqlembed:
  -package string
    	Output package name (default "queries")
  -path string
    	Path to directory containing SQL files
  -v	Show current version
```

## Example

When you have a directory with SQL query files like this one:

```
create_users.sql
delete_users.sql
foo.sql
bar.sql
```

Running `sqlembed -path=./queries -package=queries` will produce the following output:

```golang
package queries

const (
  // file: ./queries/create_users.sql
  CreateUsers = `...`

  // file: ./queries/delete_users.sql
  DeleteUsers = `...`

  // and so on
)
```

You can pipe the output into a file of your choice:

```bash
$ sqlembed -path=./queries -package=queries > queries/queries.go
```

The output file `queries/queries.go` is not very readable so it's a good idea to
add this file into your `.gitignore`.

## License

MIT
