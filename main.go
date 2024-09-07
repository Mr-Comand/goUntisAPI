package main

import (
	"encoding/json"
	"fmt"

	"github.com/Mr-Comand/goUntisAPI/goUntisAPI/structs"
	"github.com/Mr-Comand/goUntisAPI/goUntisAPI/untisApi"
)

func main() {
	apiConfig := structs.ApiConfig{Server: "", User: "", Password: "", Useragent: "client", School: "FannyLGym"}
	c := untisApi.NewClient(apiConfig)
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
