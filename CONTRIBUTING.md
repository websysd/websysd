# How to contribute

Contributions are why open source projects really excel.

## Changing files in the `assets` directory

If you made changes to any of the files in the `assets` directory you will need to run `go-bindata assets/...` inside the `websysd` directory. This will regenerate the contents of file `bindata.go`.

If you do not have `go-bindata` installed yet you can install it with `github.com/jteeuwen/go-bindata`