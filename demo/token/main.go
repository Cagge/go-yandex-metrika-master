package main

import (
	"fmt"
	"log"

	metrika "github.com/xboston/go-yandex-metrika"
)

func main() {

	log.Println("Start")

	token := "y0_AgAAAABEjid9AAh73QAAAADQ6YnhO9Zibw3ZRmatyQ56ons_aQW4dp8"

	metrika := metrika.NewMetrikaFromToken(token)
	metrika.SetDebug(true)
	if err := metrika.Authorize(); err != nil {
		log.Fatal(err.Error())
	}
	counterList, err := metrika.GetCounterList()
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(counterList)
	for _, counter := range counterList.Counters {
		log.Println(counter.ID, counter.Name, counter.Site)
	}

	log.Println("Finish")
}
