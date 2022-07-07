package utils

import (
	"encoding/json"
	"log"
	"strconv"
	"sync"

	"github.com/Muhammadjon226/first_service/models"
	"github.com/go-resty/resty/v2"
)

var response []models.Data

//GetPostsFromOpenAPIHandler ...
func GetPostsFromOpenAPIHandler() ([]models.Data, error) {
	wg := sync.WaitGroup{}

	wg.Add(5)

	go getFirstPart(&wg)
	go getSecondPart(&wg)
	go getThirdPart(&wg)
	go getFourthPart(&wg)
	go getFifthPart(&wg)

	wg.Wait()
	return response, nil
}
//getFirstPart ...
func getFirstPart(wg *sync.WaitGroup) []models.Data {
	defer wg.Done()

	client := resty.New()
	

	for i := 1; i <= 10; i++ {
		page := strconv.Itoa(i)
	
		url := "https://gorest.co.in/public/v1/posts?page=" + page
		data := models.Response{}
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
		if err != nil {
			log.Println("http request err: ", err)
		}
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			log.Println(page, " umarshall err: ", err)
		}
		response = append(response, data.Data...)
	}
	return response
}
//getSecondPart ...
func getSecondPart(wg *sync.WaitGroup) []models.Data {
	defer wg.Done()

	client := resty.New()
	

	for i := 11; i <= 20; i++ {
		page := strconv.Itoa(i)
	
		url := "https://gorest.co.in/public/v1/posts?page=" + page
		data := models.Response{}
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
		if err != nil {
			log.Println("http request err: ", err)
		}
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			log.Println(page, " umarshall err: ", err)
		}
		response = append(response, data.Data...)
	}
	return response
}

//getThirdPart ...
func getThirdPart(wg *sync.WaitGroup) []models.Data {
	defer wg.Done()

	client := resty.New()
	

	for i := 21; i <= 30; i++ {
		page := strconv.Itoa(i)
	
		url := "https://gorest.co.in/public/v1/posts?page=" + page
		data := models.Response{}
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
		if err != nil {
			log.Println("http request err: ", err)
		}
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			log.Println(page, " umarshall err: ", err)
		}
		response = append(response, data.Data...)
	}
	return response
}

//getFourthPart ...
func getFourthPart(wg *sync.WaitGroup) []models.Data {
	defer wg.Done()

	client := resty.New()
	

	for i := 31; i <= 40; i++ {
		page := strconv.Itoa(i)
	
		url := "https://gorest.co.in/public/v1/posts?page=" + page
		data := models.Response{}
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
		if err != nil {
			log.Println("http request err: ", err)
		}
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			log.Println(page, " umarshall err: ", err)
		}
		response = append(response, data.Data...)
	}
	return response
}

//getFifthPart ...
func getFifthPart(wg *sync.WaitGroup) []models.Data {
	defer wg.Done()

	client := resty.New()
	

	for i := 41; i <= 50; i++ {
		page := strconv.Itoa(i)
	
		url := "https://gorest.co.in/public/v1/posts?page=" + page
		data := models.Response{}
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			Get(url)
		if err != nil {
			log.Println("http request err: ", err)
		}
		err = json.Unmarshal(resp.Body(), &data)
		if err != nil {
			log.Println(page, " umarshall err: ", err)
		}
		response = append(response, data.Data...)
	}
	return response
}