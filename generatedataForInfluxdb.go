package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"

	client "github.com/influxdata/influxdb/client/v2"
)

// this an example of time serie data set
// this data will be store in influxdb so that we can use it later
var animaltags = []string{"Data2 rex2;Rex2", "joseph2:jojo2", "entrance2:afstudeerstage2"}

const myDB = "MydataBase"

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086", // url that my data will be listening on
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}
	queryDB(c, "", "Create DATABASE "+myDB)
	//Create a bacthpoint object
	bp, err := client.NewBatchPoints(client.BatchPointsConfig{
		Database:  myDB,
		Precision: "s",
	})
	if err != nil {
		log.Fatal(err)
	}
	wg := sync.WaitGroup{}
	detectSignal := checkStopOsSignals(&wg)
	//ForLoop will run forever when a signal is not decteted
	//this loop generate randon data every 1 sec and add that to the bachtpoint
	for !(*detectSignal) {
		animaltag := animaltags[rand.Intn(len(animaltags))]
		split := strings.Split(animaltag, ";")
		tags := map[string]string{
			"animal_type2": split[0],
			"nickname2":    split[0],
		}
		fields := map[string]interface{}{
			"weight": rand.Intn(300) + 1, //generate random number for the weight
		}
		fmt.Println(animaltag, fields["weight"])
		pt, err := client.NewPoint("weightmeasures", tags, fields, time.Now())
		if err != nil {
			log.Println(err)
			continue
		}
		bp.AddPoint(pt)
		time.Sleep(1 * time.Second)
	}
	log.Println("Exit signal triggered, writting data ... ")
	if err := c.Write(bp); err != nil {
		log.Fatal(err)
	}
	wg.Wait()
	log.Println("Exiting program ...")
}

//this special in go. this function allow you to detect signal coming from your Operating systeem when your OS is shutdown
func checkStopOsSignals(wg *sync.WaitGroup) *bool {
	Signal := false
	go func(s *bool) {
		wg.Add(1)
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch
		log.Println("Exit signal recieved ... ")
		*s = true
		wg.Done()
	}(&Signal)
	return &Signal
}
func queryDB(c client.Client, database, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	response, err := c.Query(q)
	if err != nil {
		return res, err
	}
	if response.Error() != nil {
		return res, response.Error()
	}

	return response.Results, nil
}
