package logic

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

//HTTPPost post data to client
func HTTPPost(url string, verb string, postData []byte) error {
	fmt.Println("URL:>", url)

	req, err := http.NewRequest(verb, url, bytes.NewBuffer(postData))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYW1laWQiOiJzeW5jLXNlcnZpY2UiLCJpZCI6IjU5ODFhMWU0MWQ0MWM4NGNhZTkwNGZkMiIsInVuaXF1ZV9uYW1lIjoic3luYy1zZXJ2aWNlIiwic3ViIjoic3luYy1zZXJ2aWNlIiwiaXNzIjoiaHR0cDovL2xvY2FsaG9zdC8iLCJhdWQiOiJiOWRjNzEyYzk1MmI0YWFmYjQ4MWFiZWRlMGZlYzRkOCIsImV4cCI6OTk5OTk5OTk5OSwibmJmIjoxNDk3MTc4MjQ1LCJyb2xlIjpbInJvc3RhbSJdLCJhcHAiOiI1YTk1NjNmMzg0OWUwNjc3MTA1ZGY1MjkifQ.IZGgGrZjXnXHIPGF6w_nAKonQRqbqETyrFD743STaC8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	return nil
}
