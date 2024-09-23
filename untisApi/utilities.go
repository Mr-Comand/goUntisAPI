package untisApi

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Mr-Comand/goUntisAPI/structs"
)

type Client struct {
	structs.ApiConfig
	structs.AuthResponse
	Logger *logger
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
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(reqBody) + " Error: " + err.Error())
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
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(reqBody) + " Error: " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	c.Logger.Debug("Called\t" + method + "\t at: " + "\"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(reqBody))

	var rpcResp structs.RPCResponse

	err = json.NewDecoder(resp.Body).Decode(&rpcResp)
	if err != nil {
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(reqBody) + " Error: " + err.Error())
		return nil, err
	}

	if rpcResp.Error != nil {
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(reqBody) + " Error: " + rpcResp.Error.Error())
		return nil, rpcResp.Error
	}
	c.Logger.Debug("Successful\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(reqBody))
	return &rpcResp, nil
}

func NewClient(apiConfig structs.ApiConfig, logger *log.Logger, LogLevel LogLevel) *Client {
	c := Client{
		ApiConfig:    apiConfig,
		AuthResponse: structs.AuthResponse{},
		Logger:       newLogger(logger, LogLevel),
	}

	return &c
}
