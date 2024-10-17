package untisApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Mr-Comand/goUntisAPI/structs"
)

type Client struct {
	structs.ApiConfig
	structs.AuthResponse
	Logger *logger
	Censor bool
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
	fmt.Println("test")
	logBody := reqBody
	if c.Censor {
		if logParams, ok := params.(structs.AuthParams); ok {
			password := logParams.Password
			logParams.Password = strings.Repeat("*", len(password))
			logBody, err = json.Marshal(structs.RPCRequest{
				ID:      "1",
				Method:  method,
				Params:  logParams,
				JSONRPC: "2.0",
			})
		}
		if err != nil {
			c.Logger.Error("Failed to  marshal log body")
		}
	}
	// Create a new request
	req, err := http.NewRequest("POST", "https://"+c.Server+"/WebUntis/jsonrpc.do?school="+c.School, bytes.NewBuffer(reqBody))
	if err != nil {
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(logBody) + " Error: " + err.Error())
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
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(logBody) + " Error: " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	c.Logger.Debug("Called\t" + method + "\t at: " + "\"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(logBody))

	var rpcResp structs.RPCResponse

	err = json.NewDecoder(resp.Body).Decode(&rpcResp)
	if err != nil {
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(logBody) + " Error: " + err.Error())
		return nil, err
	}

	if rpcResp.Error != nil {
		c.Logger.Error("Failed\t" + method + "\t at: \"https://" + c.Server + "/WebUntis/jsonrpc.do?school=" + c.School + "\" with body: " + string(logBody) + " Error: " + rpcResp.Error.Error())
		return nil, rpcResp.Error
	}
	c.Logger.Debug(fmt.Sprintf(
		"Successful\t%s\t at: \"https://%s/WebUntis/jsonrpc.do?school=%s\" with body: %s ContentLength: %d bytes",
		method,
		c.Server,
		c.School,
		string(logBody),
		resp.ContentLength,
	))
	return &rpcResp, nil
}

func NewClient(apiConfig structs.ApiConfig, logger *log.Logger, LogLevel LogLevel, censor bool) *Client {
	c := Client{
		ApiConfig:    apiConfig,
		AuthResponse: structs.AuthResponse{},
		Logger:       newLogger(logger, LogLevel),
		Censor:       censor,
	}

	return &c
}
