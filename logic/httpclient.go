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
	req.Header.Set("Authorization", "bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYW1laWQiOiJzeW5jLXNlcnZpY2UiLCJ1bmlxdWVfbmFtZSI6InN5bmMtc2VydmljZSIsInN1YiI6InN5bmMtc2VydmljZSIsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3QvIiwiYXVkIjoiYjlkYzcxMmM5NTJiNGFhZmI0ODFhYmVkZTBmZWM0ZDgiLCJleHAiOjk5OTk5OTk5OTksIm5iZiI6MTQ5NzE3ODI0NX0.2OVSsDT1kC5o6JaChZk_OZVUq_w-c-fB17sIP-kI92A")

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
