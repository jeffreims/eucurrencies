# eucurrencies
Small lib to fetch Euro foreign exchange reference rates


Fetch the lib

`go get github.com/jeffreims/eucurrencies`


```
package main

import "github.com/jeffreims/eucurrencies"
import "fmt"

func main() {

	var data eucurrencies.EUCurrencies
	err := data.Fetch()
	if err == nil {
		fmt.Println("NOK:", data.NOK)
	}
}
```