package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

//Require running application
func main() {

}

func GetTotalCount() (int, error) {
	resp, err := http.Get("http://localhost:8080/stat/total")
	if err != nil {
		return 0, err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	//Convert the body to type string
	totalCount, err := strconv.Atoi(string(body))
	if err != nil {
		return 0, err
	}
	return totalCount, err
}

func IncrementTotalCount() error {
	form := url.Values{}
	form.Add("userName", "Toby")
	resp, err := http.PostForm("http://localhost:8080/visit", form)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf("Status code: %d ", resp.StatusCode)
	}
	return nil
}
