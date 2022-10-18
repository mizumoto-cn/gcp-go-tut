// Package logs/bq contains examples for handling Cloud Functions logs.
package bq

import (
	"context"
	"fmt"
	"log"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/cloudevents/sdk-go/v2/event"
)

func init() {
	functions.CloudEvent("HelloAuditLog", helloAuditLog)
}

// AuditLogEntry represents a LogEntry as described at
// https://cloud.google.com/logging/docs/reference/v2/rest/v2/LogEntry
type AuditLogEntry struct {
	ProtoPayload *AuditLogProtoPayload `json:"protoPayload"`
}

// AuditLogProtoPayload represents AuditLog within the LogEntry.protoPayload
// See https://cloud.google.com/logging/docs/reference/audit/auditlog/rest/Shared.Types/AuditLog
type AuditLogProtoPayload struct {
	MethodName         string                 `json:"methodName"`
	ResourceName       string                 `json:"resourceName"`
	AuthenticationInfo map[string]interface{} `json:"authenticationInfo"`
}

// helloAuditLog receives a CloudEvent containing an AuditLogEntry, and logs a few fields.
func helloAuditLog(ctx context.Context, e event.Event) error {
	// Print out details from the CloudEvent itself
	// See https://github.com/cloudevents/spec/blob/v1.0.1/spec.md#subject
	// for details on the Subject field
	log.Printf("Event Type: %s", e.Type())
	log.Printf("Subject: %s", e.Subject())

	// Decode the Cloud Audit Logging message embedded in the CloudEvent
	logentry := &AuditLogEntry{}
	if err := e.DataAs(logentry); err != nil {
		ferr := fmt.Errorf("event.DataAs: %w", err)
		log.Print(ferr)
		return ferr
	}
	// Print out some of the information contained in the Cloud Audit Logging event
	// See https://cloud.google.com/logging/docs/audit#audit_log_entry_structure
	// for a full description of available fields.
	log.Printf("API Method: %s", logentry.ProtoPayload.MethodName)
	log.Printf("Resource Name: %s", logentry.ProtoPayload.ResourceName)
	if v, ok := logentry.ProtoPayload.AuthenticationInfo["principalEmail"]; ok {
		log.Printf("Principal: %s", v)
	}
	return nil
}
