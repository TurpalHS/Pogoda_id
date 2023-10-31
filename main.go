package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Responce struct {
	Main Main
}

type Main struct {
	Temp     float64 `json:"temp"`
	Humidity int     `json:"humidity"`
}

func main() {
	fmt.Println("Дарова, Заебал")
	var id string
	fmt.Print("Введите индекс своего мухосранска, чтобы узнать погоду, например: 2013159 \n")
	fmt.Fscan(os.Stdin, &id)
	idCity := id

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?id=%s&appid=4f594c4a2a73a44aebdf63f35837b419", idCity)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Неправильный индекс, гандон:", err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var pogoda Responce

	err = json.Unmarshal(body, &pogoda)
	if err != nil {
		fmt.Println(err)
	}

	tempCels := pogoda.Main.Temp
	vlagaProcent := pogoda.Main.Humidity

	fmt.Printf("Температура в городе: %.2f °C \n", tempCels)
	fmt.Printf("Влажность погоды: %d %% \n", vlagaProcent)

}
