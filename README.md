# Ungo
Ungo and a package that has the function of unshorten urls.
# How to install
You should run the following commands:<br>
<code>go get github.com/ReiGelado/Ungo</code>
and
<code>go install github.com/ReiGelado/Ungo</code>
<br>There, you can now use :)
# Shorteners
The available shorteners are:<br>
<code> ad7biz ---> http://ad7.biz <p>
adfly ---> http://adf.ly <p>
adfocus ---> http://adfoc.us <p>
amankanlink --> http://kombatch.amankan.link <p>
coegin ---> http://coeg.in <p>
googl ---> http://goo.gl <p>
linkbucks ---> http://linkbucks.com/<p>
linkshrink ---> http://linkshrink.net/ <p>
linktl ---> http://link.tl <p>
shst ---> http://sh.st <p>
tco ---> http://t.co <p>
urlgogs ---> http://urlgo.gs <p></code>

# Example
A simple example of how to use the library :)

```go
package main

import (
	"fmt"
	"ungo"
)

func main() {
	url, err := ungo.Shorten(ungo.Config{Url: "http://adf.ly/tYjLr", Shortener:"adfly"})
	if err != nil {
		panic(err)
	}
	fmt.Println(url)
}

```