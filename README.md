Installing

```go get -u github.com/kenjione/importer```

Getting Started

```golang
package main

import "github.com/kenjione/importer"

func main() {
  i := importer.NewImporter(&importer.Config{
    DatabaseHost:     "hostname",
    DatabaseUser:     "user",
    DatabasePassword: "password",
  })

  defer i.Close()

  stats := i.Parse("path/to/file")

  fmt.Println("Accepted:", stats.Accepted)
  fmt.Println("Invalid:", stats.Invalid)
  fmt.Println("Not saved:", stats.NotSaved)
  fmt.Println("Duration:", stats.Duration)

  lbytes, err := i.FindByIP("some.ip.address.here")

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println("Location: ", lbytes)
}

```