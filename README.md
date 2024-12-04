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

It's power comes from the fact that it satisfies interfaces from the fs package, allowing for modularity and testability. Let's say you wanted to create a storage-agnostic and testable key-value store: 

```go

type Database struct {
    filesystem realfs.WritableFs
}

func (d *Database) Get(name string) ([]byte, error) {
    return d.filesystem.OpenFile(name)
}

func (d *Databse) Set(name string, content []byte) error {
    return d.filesystem.WriteFile(name, content, 0640)
}

func (d *Database) Delete(name string) error {
    return d.filesystem.Remove(name)
}

```

To instantiate a real Databse you can do:

```go

db := Database{realfs.NewWritable()} 

```

And to test is equally easy:

```go

func TestDatabase(t *testing.T) {

    


}

```