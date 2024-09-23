package untisApi

import (
	"encoding/json"

	"github.com/Mr-Comand/goUntisAPI/structs"
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

	c.Logger.Info("Authenticated successfully, session ID:", c.SessionID)

	return nil
}
func (c *Client) ContinueSession(SessionID string) error {
	c.SessionID = SessionID
	return c.Test()
}
func (c *Client) Logout() error {
	_, err := c.CallRPC("logout", struct{}{})
	if err == nil {
		c.Logger.Info("Logout successful, session ID:", c.SessionID)
		c.AuthResponse = structs.AuthResponse{}
	}
	return err
}
