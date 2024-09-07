package untisApi

import (
	"encoding/json"
	"fmt"

	"github.com/Mr-Comand/goUntisAPI/goUntisAPI/structs"
)

func (c *Client) Authenticate() error {
	params := structs.AuthParams{
		User:     c.User,
		Password: c.Password,
		Client:   c.Useragent,
	}
	rpcResp, err := c.CallRPC("authenticate", params)
	if err != nil {
		return err
	}
	var authResp structs.AuthResponse
	err = json.Unmarshal(rpcResp.Result, &authResp)
	if err != nil {
		return err
	}
	c.AuthResponse = authResp

	fmt.Println("Authenticated successfully, session ID:", c.SessionID)

	return nil
}
func (c *Client) Logout() error {
	rpcResp, err := c.CallRPC("logout", struct{}{})
	fmt.Println(rpcResp, err)
	return err
}
