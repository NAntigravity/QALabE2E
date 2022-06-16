package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

//Require running application
func main() {
	totalCountBeforeTest, err := GetTotalCount()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Total count before test: %d", totalCountBeforeTest)
	err = IncrementTotalCount()
	if err != nil {
		log.Fatalln(err)
	}
	totalCountAfterTest, err := GetTotalCount()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Total count after test: %d", totalCountAfterTest)
	if totalCountAfterTest <= totalCountBeforeTest {
		log.Fatalln("Test count not incremented after sending request. You are a bad programmer. Try to find another job")
	}
	log.Printf("You are a good boy! Your code are beautiful!")
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
