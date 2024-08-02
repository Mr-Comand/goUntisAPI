package structs

import (
	"encoding/json"
	"fmt"
)

// RPCRequest represents a JSON-RPC request.
type RPCRequest struct {
	ID      string      `json:"id"`
	Method  string      `json:"method"`
	Params  interface{} `json:"params"`
	JSONRPC string      `json:"jsonrpc"`
}

// RPCResponse represents a JSON-RPC response.
type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      string          `json:"id"`
	Result  json.RawMessage `json:"result"`
	Error   *RPCError       `json:"error,omitempty"`
}

// RPCError represents an error in a JSON-RPC response.
type RPCError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *RPCError) Error() string {
	return fmt.Sprintf("RPCError: Code=%d, Message=%s", e.Code, e.Message)
}
