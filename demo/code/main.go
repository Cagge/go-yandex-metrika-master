package main

import (
	"log"

	metrika "github.com/xboston/go-yandex-metrika"
)

func main() {

	log.Println("Start")

	//code := "code"
	clientID := "6b98228df9894e148a3edbc6cbb08a6f"
	clientSecret := "19a7a17fc6f44f579658f60b6f147a44"
	username := "kiryanov.da"
	//password := "74539hsv61"

	//metrika := metrika.NewMetrika(clientID, clientSecret, username, password)
	metrika.SetDebug(true)

	if err := metrika.Authorize(); err != nil {
		log.Fatal(err.Error())
	}

	counterList, _ := metrika.GetCounterList()

	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
