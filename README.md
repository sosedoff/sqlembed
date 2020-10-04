# sqlembed

A tiny tool to embed `*.sql` files into a Go package

## Example

When you have a directory with SQL files (queries, etc) like this one:

```
create_users.sql
delete_users.sql
foo.sql
bar.sql
```

Running `sqlembed -path=./queries -package=queries` will produce the following output:

```
package queries

const (
  // file: ./queries/create_users.sql
  CreateUsers = `...`

  // file: ./queries/delete_users.sql
  DeleteUsers = `...`

  // and so on
)
```
