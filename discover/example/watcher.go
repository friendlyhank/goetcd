package main

import (
	"hank.com/goetcd/discover/client"
	"log"
	"time"
)

func main() {
	m, _ := client.NewEtcdV3Discovery("node/",[]string{"localhost:2379", "localhost:2381", "localhost:2383"})
	for{
		list,_ := m.GetServiceList()
		if len(list) == 0{
			continue
		}

		for _,endPoint:= range list{
			log.Printf("node:%v \n", endPoint)
		}
		time.Sleep(time.Second*5)
	}
}