/*
This a test code to read database
In this case this code read a database name "MydataBase".
select * from weightmeasures where "animal_type"='joseph:jojo
*/
package main

import (
	"log"

	client "github.com/influxdata/influxdb/client/v2"
)

func main() {
	// Create a new HTTPClient
	c, err := client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://localhost:8086",
		Username: "",
		Password: "",
	})
	if err != nil {
		log.Fatal(err)
	}

	res, err := MyDBquery(c, "MydataBase", `select * from weightmeasures where "animal_type"='joseph:jojo'`)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range res {
		log.Println("message: ", v.Messages)
		for _, s := range v.Series {
			log.Println("series name: ", s.Name)
			log.Println("series colums:", s.Columns)
			log.Println("series values:", s.Values)
		}

	}
}

func MyDBquery(c client.Client, database, cmd string) (res []client.Result, err error) {
	q := client.Query{
		Command:  cmd,
		Database: database,
	}
	if response, err := c.Query(q); err == nil {
		if response.Error() != nil {
			return res, response.Error()
		}
		res = response.Results
	} else {
		return res, err
	}
	return res, nil
}
