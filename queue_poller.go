package main

import (
	"fmt"
	"github.com/hazelcast/hazelcast-go-client"
	"github.com/hazelcast/hazelcast-go-client/config/property"
)

func main() {

	clientConfig := hazelcast.NewConfig()

	clientConfig.SetClientName("queue_poller")

	clientConfig.NetworkConfig().AddAddress("127.0.0.1:5701")

	clientConfig.GroupConfig().SetName("devoxx")
        clientConfig.GroupConfig().SetPassword("")

	// Off by default
	clientConfig.SetProperty(property.StatisticsEnabled.Name(), "true")

	hazelcastClient, _ := hazelcast.NewClientWithConfig(clientConfig)

	fmt.Printf("%+v\n", hazelcastClient.Cluster().GetMembers());

        for _, member := range hazelcastClient.Cluster().GetMembers() {
   		fmt.Printf("%+v\n", member.Address());
	}

	disputesQueue, _ := hazelcastClient.GetQueue("disputes")

	fmt.Printf("===============\n")		
	fmt.Printf("Queue '%s'\n", disputesQueue.Name())
        done := false
        for !done {
		fmt.Printf("===============\n")		
		item, _ := disputesQueue.Take()
		fmt.Printf("Item '%s'\n", item)
        }
	fmt.Printf("===============\n")		

	hazelcastClient.Shutdown()       
}
