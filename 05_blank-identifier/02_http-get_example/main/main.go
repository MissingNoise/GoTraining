package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)
// gets the html
func main()  {
	res, err :=http.Get("http://www.mcleods.com")
	// check for error
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	// check for an error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}