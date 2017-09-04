package apigw

import (
	"net/http"
	"fmt"
)

type Credential struct {
	UserName string `json:"UserName"`
	Password string `json:"Password"`
	TokenID string `json:TokenID`
}


type Apigw struct {
	URL string `json:"URL"`
}

func Getburst(myurl string, myloop int64) {
	fmt.Println("In api / launch GET api gateway burst")
	// loop on myloop on myurl
	fmt.Println("URL", myurl)
	fmt.Println("Number of Request", myloop)
	var i int64
	for i = 0; i < myloop; i++ {
		fmt.Println("Loop", i)
		req, err := http.NewRequest("GET", myurl, nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
		fmt.Println("Error Connection: ")
		}
		defer resp.Body.Close()
	}
	
}

func Postburst(myurl string, myloop int64) {
	fmt.Println("In api / launch POST api gateway burst")
	// loop on myloop on myurl
	fmt.Println("URL", myurl)
	fmt.Println("Number of Request", myloop)
	var i int64
	for i = 0; i < myloop; i++ {
		fmt.Println("Loop", i)
		req, err := http.NewRequest("POST", myurl, nil)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
		fmt.Println("Error Connection: ")
		}
		defer resp.Body.Close()
	}
	
}
