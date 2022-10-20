package parser

import (
	"encoding/json"

	"github.com/googleapis/google-cloudevents-go/cloud/audit/v1"
)

// Payload received from pub/sub
type Payload struct {
	Data []byte `json:"data"`
}

// PubSub Data structure
type PubSubMessage struct {
	InsertID         string                 `json:"insertId"`
	LogName          string                 `json:"logName"`
	Payload          audit.ProtoPayload     `json:"protoPayload"`
	ReceiveTimestamp string                 `json:"receiveTimestamp"`
	Resource         map[string]interface{} `json:"resource"`
	Severity         string                 `json:"severity"`
	Timestamp        string                 `json:"timestamp"`
}

type ParseHandler interface {
	Parse(input interface{}) (interface{}, error)
}

// RawLogHandler implements Parser interface
// it parses raw log data to protoPayload object
type RawLogHandler struct{}

func (p *RawLogHandler) Parse(input Payload) (audit.ProtoPayload, error) {
	data := input.Data
	var psm PubSubMessage
	err := json.Unmarshal(data, &psm)
	if err != nil {
		return psm.Payload, err
	}
	return psm.Payload, nil
}

// TODO: factory pattern / option pattern to be implemented
