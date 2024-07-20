# go-real-fs

A real filesystem that behaves normally and implements [fs.FS](https://pkg.go.dev/io/fs#FS) and wraps [os.DirFS](https://pkg.go.dev/os#DirFS).

Using `fs.FS` to represent filesystems has many advantages, but it can be clunky to work with for reasons [outlined in this rejected proposal](https://github.com/golang/go/issues/47803).

go-real-fs implements this behaviour like so:

```go
package main

import (
    "github.com/sean9999/go-real-fs"
)

func main(){

    realFs := realfs.New()
    data, _ := realFs.ReadFile("/etc/passwd")
    fmt.Println("here is the contents of /etc/passwd")
    os.Stdout.Write(data)

}
```

go-real-fs simply rewrites paths from context-dependant to root-relative:

```go
func (rfs realFS) correctPath(relativePath string) (string, error) {

	fullPath, err := filepath.Abs(relativePath)
	if err != nil {
		return "", err
	}
	return filepath.Rel("/", fullPath)

}
```