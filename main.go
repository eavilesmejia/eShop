package main

import (
	"fmt"
	"net/http"
	"bufio"
	"os"
	"time"
)


//Urls to check
var urls = []string{
	"https://www.amazon.com/s/field-keywords=",
}


type HttpResponse struct {
	url      string
	response *http.Response
	err      error
}

func AsyncHttpGets(urls []string) chan *HttpResponse {
	ch := make(chan *HttpResponse, len(urls)) // buffered

	for _, url := range urls {
		//Go routine call
		go func(url string) {
			fmt.Printf("Fetching %s \n", url)
			resp, err := http.Get(url)
			resp.Body.Close()
			ch <- &HttpResponse{url, resp, err}
		}(url)
	}


	return ch

}

func ProcessResponse(){


	//for {
	//	select {
	//	case response := <-ch:
	//		fmt.Printf("%s was fetched\n", response.url)
	//		responses = append(responses, response)
	//		if len(responses) == len(urls) {
	//			return responses
	//		}
	//	case <-time.After(50 * time.Millisecond):
	//		fmt.Printf(".")
	//	}
	//}
}

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Type your search: ")
	text, _ := reader.ReadString('\n')

	results := AsyncHttpGets(urls)
	for _, result := range results {
		fmt.Printf("%s status: %s\n", result.url,
			result.response.Status)
	}
}
