package main

import (
	"fmt"
	"strings"

	"github.com/fancurson/links/links"
)

var exampleHTML = `
<html>
<body>
	<h1>Hello!</h1>
	<a href="/other-page">A link to another page</a>
	<a href="/other-page">A link to second page</a>
</body>
</html>
`

func main() {
	r := strings.NewReader(exampleHTML)
	data, err := links.Parse(r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", data)
}
