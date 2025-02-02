package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Simulasi penggunaan token curian untuk akses API
	token := "TOKEN_CURAN"

	client := &http.Client{}
	req, _ := http.NewRequest("GET", "http://target.com/api/user", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response:", string(body))
}