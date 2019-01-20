package main 

import (
	"fmt"
	"net/http"
	"github.com/techcomsecurities/rest"
)

func main() {
	url := "http://www.mocky.io/v2/5c42d74b320000d92f73281d"
	resp, err := rest.NewRequest().Get(url)
	if err != nil {

	}
}