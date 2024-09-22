package untisApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Mr-Comand/goUntisAPI/structs"
)

type Client struct {
	structs.ApiConfig
	structs.AuthResponse
}

func (c *Client) CallRPC(method string, params interface{}) (*structs.RPCResponse, error) {
	reqBody, err := json.Marshal(structs.RPCRequest{
		ID:      "1",
		Method:  method,
		Params:  params,
		JSONRPC: "2.0",
	})
	if err != nil {
		return nil, err
	}
	// Create a new request
	req, err := http.NewRequest("POST", "https://"+c.Server+"/WebUntis/jsonrpc.do?school="+c.School, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header
	req.Header.Set("Content-Type", "application/json")
	if c.SessionID != "" {
		// Add the SESSIONID cookie
		cookie := &http.Cookie{
			Name:  "JSESSIONID",
			Value: c.SessionID,
		}
		req.AddCookie(cookie)
	}
	// Create a new HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rpcResp structs.RPCResponse

	err = json.NewDecoder(resp.Body).Decode(&rpcResp)
	if err != nil {
		return nil, err
	}

	if rpcResp.Error != nil {
		return nil, fmt.Errorf("RPC error(%d): %s", rpcResp.Error.Code, rpcResp.Error.Message)
	}

	return &rpcResp, nil
}

func NewClient(apiConfig structs.ApiConfig) *Client {
	// Example usage

	c := Client{
		apiConfig,
		structs.AuthResponse{},
	}

	// err := c.Authenticate()
	// if err != nil {
	// 	fmt.Println("Error authenticating:", err)
	// 	return &Client{}
	// }

	return &c
}
