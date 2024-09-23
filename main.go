package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Mr-Comand/goUntisAPI/structs"
	"github.com/Mr-Comand/goUntisAPI/untisApi"
)

func main() {
	apiConfig := structs.ApiConfig{Server: "", User: "", Password: "", Useragent: "client", School: "FannyLGym"}
	c := untisApi.NewClient(apiConfig, log.Default(), untisApi.DEBUG)
	err := c.Authenticate()
	if err != nil {
		fmt.Println("Error authenticating:", err)
	}
	c.Test()
	data, err := c.GetClassesOfSchoolYear(1)
	dataJson, _ := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling person slice:", err)
		return
	}
	fmt.Println(string(dataJson), err)
	c.Logout()
}
