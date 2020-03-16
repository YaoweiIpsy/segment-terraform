package main

import (
	"encoding/json"
	"log"
	segmentApi "segment-terraform/segment/client"
)

func main() {
	//plugin.Serve(&plugin.ServeOpts{
	//	ProviderFunc: func() terraform.ResourceProvider {
	//		return segment.Provider()
	//	},
	//})

	client := segmentApi.NewClient("4K1I2qQs9Kfy06bqXRV-TsKa9hU8Numwf5OjLvfrnBs.1qN8nrmLJWHRYEjVLMKVHV03P_3zLHhs4WtrkZJLIZk", "ipsy-yaowei")
	client.DeleteDestination("javascript", "repeater")
	data, err := client.CreateDestination("javascript", "repeater", true, []segmentApi.DestinationConfig{
		segmentApi.DestinationConfig{
			Name:        "workspaces/ipsy-yaowei/sources/javascript/destinations/repeater/config/repeatKeys",
			DisplayName: "Write keys",
			Value:       []string{"asfasdfasfasf", "adfasdfasasdfasf", "adfasasddffdafasf", "-------"},
			Type:        "list",
		},
	}) //.ListDestinations("javascript")
	log.Println(err)
	out, _ := json.Marshal(data)
	log.Println(string(out))

	//token := ""
	//for {
	//	data, err := c.ListSources(token) //c.CreateSource(fmt.Sprintf("wwuu-%02d",i), "catalog/sources/javascript")
	//	if err != nil {
	//		log.Fatal(err)
	//		println(err)
	//		return
	//	}
	//	if data.Sources == nil {
	//		break
	//	}
	//	for _,source := range data.Sources {
	//		fmt.Println("-----", source.Name)
	//		err = c.DeleteSource(source.Name)
	//		if err != nil {
	//			log.Println(err)
	//		}
	//	}
	//	////println(string(json.Marshal(data)))
	//	//out, err := json.Marshal(data)
	//	//if err != nil {
	//	//	println(err)
	//	//} else {
	//	//	println(string(out))
	//	//}
	//	//token = data.NextPageToken
	//	//if token == "" {
	//	//	break
	//	//}
	//}
	////for i := 100; i < 5000; i++ {
	//
	////}
}
