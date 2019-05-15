# Memes

This repo contains the ULTIMATE™ toolchain to creating the Dankest Memes™.

For now it only contains a Spongebob text generator.
```go
package main

import (
	"fmt"	
	
	"go.ilie.io/memes/spongebob"
)

func main() {
    fmt.Println(spongebob.Text("Hello world!"))
}
```
Output will be:
```
> hELlO wOrlD!
   ```

You can also use it as a CLI:
```sh
go build cmd./spongebob.go -o spongebob && ./spongebob "hello world"

> hELlO wOrlD
```

