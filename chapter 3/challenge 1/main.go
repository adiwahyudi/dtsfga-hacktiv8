package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	for true {
		rand_water := rand.Intn(100-1) + 1
		rand_wind := rand.Intn(100-1) + 1

		data := map[string]interface{}{
			"water": rand_water,
			"wind":  rand_wind,
		}

		reqJson, err := json.Marshal(data)
		client := &http.Client{}

		req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(reqJson))

		req.Header.Set("Content-type", "application/json")
		if err != nil {
			log.Fatal(err)
		}

		res, eer := client.Do(req)
		if eer != nil {
			log.Fatalln(err)
		}

		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println(string(body))
		checkResult(rand_water, rand_wind)
		time.Sleep(15 * time.Second)
	}

}

func checkResult(water int, wind int) {
	status_water := ""
	status_wind := ""

	if water <= 5 {
		status_water = "aman"
	} else if water >= 6 && water <= 8 {
		status_water = "siaga"
	} else {
		status_water = "bahaya"
	}

	if wind <= 6 {
		status_wind = "aman"
	} else if wind >= 7 && wind <= 15 {
		status_wind = "siaga"
	} else {
		status_wind = "bahaya"
	}

	fmt.Printf("status water: %s\n", status_water)
	fmt.Printf("status wind: %s\n", status_wind)
}
