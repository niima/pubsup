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
	req.Header.Set("Authorization", "bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJuYW1laWQiOiIyNCIsInVuaXF1ZV9uYW1lIjoidW5rbm93bigyNCkiLCJzdWIiOiJ1bmtub3duKDI0KSIsImlzcyI6Imh0dHA6Ly9sb2NhbGhvc3QvIiwiYXVkIjoiYjlkYzcxMmM5NTJiNGFhZmI0ODFhYmVkZTBmZWM0ZDgiLCJleHAiOjk5MTUwNTY1OTQ1MiwibmJmIjoxNTA0NDQ5ODUyLCJpZCI6IjU5YWJmZWNhYjE2NDY4MWRkODM3OTM1NiIsInVzZXJuYW1lIjoibmltYSJ9.nH-2mVs4waZ3luq7LozZE8a1qvC-oK-ELCv_hFeDBT8")

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
