package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"

	"github.com/go-co-op/gocron"
)

type Nilai struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

func Acak() {
	rand.Seed(time.Now().UnixMilli())
	angkaRandom := rand.Intn(20)
	rand.Seed(time.Now().Unix())
	angkaRandom2 := rand.Intn(20)
	data := Nilai{
		Water: angkaRandom,
		Wind:  angkaRandom2,
	}

	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile("value.json", file, 0644)

	if data.Water < 5 {
		fmt.Printf("Water %d Aman \n", angkaRandom)
	} else if data.Water > 4 && data.Water < 9 {
		fmt.Printf("Water %d Siaga \n", angkaRandom)
	} else {
		fmt.Printf("Water %d Bahaya \n", angkaRandom)
	}

	if data.Wind < 5 {
		fmt.Printf("Wind %d Aman \n", angkaRandom2)
	} else if data.Wind > 4 && data.Wind < 16 {
		fmt.Printf("Wind %d Siaga \n", angkaRandom2)
	} else {
		fmt.Printf("Wind %d Bahaya \n", angkaRandom2)
	}
}

func Cron() {
	s := gocron.NewScheduler(time.Local)
	s.Every("3s").Do(func() {
		Acak()
	})
	s.StartBlocking()
}

func main() {

	Cron()
}
