# Google Agent Development Kit API Client (Go)

[![Go Reference](https://pkg.go.dev/badge/github.com/goudtho/adk-client-go.svg)](https://pkg.go.dev/github.com/goudtho/adk-client-go)

This repository contains a Go client for the REST API of the [Google Agent Development Kit](https://github.com/google/adk-python).

## Description

The [Google Agent Development Kit (ADK)](https://github.com/google/adk-python) provides tools and infrastructure for building, evaluating, and deploying AI agents. This Go client allows developers to interact with the ADK's REST API programmatically to manage resources such as applications, sessions, evaluations, and artifacts.

## Installation

To use this client in your Go project, you can install it using `go get`:

```bash
go get github.com/goudtho/adk-client-go
```

## Usage

Here's a basic example of how to create a client and list available applications:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	client "github.com/goudtho/adk-client-go"
)

func main() {
	// Replace with the actual server URL for the ADK API
	serverURL := "http://localhost:8000" // Example URL, replace with the real one

	// Create a new client instance
	c, err := client.NewClientWithResponses(serverURL)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: List available applications
	ctx := context.Background()
	resp, err := c.ListAppsWithResponse(ctx) //
	if err != nil {
		log.Fatalf("Failed to list apps: %v", err)
	}

	if resp.StatusCode() != 200 {
		log.Fatalf("Error listing apps: Status %s, Body: %s", resp.Status(), string(resp.Body))
	}

	if resp.JSON200 != nil {
		fmt.Println("Available Applications:")
		for _, appName := range *resp.JSON200 {
			fmt.Printf("- %s\n", appName)
		}
	} else {
		fmt.Println("No applications found or error parsing response.")
	}

	appName := "your_app_name"
	userID := "your_user_id"
	sessionID := "your_session_id" // Optional, can let server generate

	createSessionBody := client.CreateSessionWithIdJSONRequestBody{} // Add initial state if needed

	sessionResp, err := c.CreateSessionWithIdWithResponse(ctx, appName, userID, sessionID, createSessionBody)
	if err != nil {
		log.Fatalf("Failed to create session: %v", err)
	}
	if sessionResp.StatusCode() == 200 && sessionResp.JSON200 != nil {
		fmt.Printf("Created/Got Session ID: %s\n", sessionResp.JSON200.Id)
	} else {
		log.Printf("Error creating session: Status %s, Body: %s", sessionResp.Status(), string(sessionResp.Body))
	}

```

## Client Methods
* **`ListEvalSets(ctx context.Context, appName string, ...)`**
    * Retrieves a list of evaluation set IDs for a specific application.

* **`CreateEvalSet(ctx context.Context, appName string, evalSetId string, ...)`**
    * Creates a new evaluation set with a specified ID for a given application.

* **`AddSessionToEvalSet(ctx context.Context, appName string, evalSetId string, body AddSessionToEvalSetJSONRequestBody, ...)`**
    * Adds an existing session to a specified evaluation set.

* **`ListEvalsInEvalSet(ctx context.Context, appName string, evalSetId string, ...)`**
    * Retrieves a list of evaluation IDs within a specific evaluation set.

* **`RunEval(ctx context.Context, appName string, evalSetId string, body RunEvalJSONRequestBody, ...)`**
    * Executes evaluations based on the specified metrics for an evaluation set.

* **`ListSessions(ctx context.Context, appName string, userId string, ...)`**
    * Retrieves a list of sessions associated with a specific user and application.

* **`CreateSession(ctx context.Context, appName string, userId string, body CreateSessionJSONRequestBody, ...)`**
    * Creates a new session for a user and application, letting the server generate the session ID.

* **`DeleteSession(ctx context.Context, appName string, userId string, sessionId string, ...)`**
    * Deletes a specific session identified by its ID.

* **`GetSession(ctx context.Context, appName string, userId string, sessionId string, ...)`**
    * Retrieves the details of a specific session, including its events and state.

* **`CreateSessionWithId(ctx context.Context, appName string, userId string, sessionId string, body CreateSessionWithIdJSONRequestBody, ...)`**
    * Creates a new session with a client-specified ID or updates an existing one.

* **`ListArtifactNames(ctx context.Context, appName string, userId string, sessionId string, ...)`**
    * Retrieves a list of artifact names stored within a specific session.

* **`DeleteArtifact(ctx context.Context, appName string, userId string, sessionId string, artifactName string, ...)`**
    * Deletes a specific artifact (and all its versions) within a session.

* **`LoadArtifact(ctx context.Context, appName string, userId string, sessionId string, artifactName string, params *LoadArtifactParams, ...)`**
    * Loads the content of an artifact from a session. Can optionally specify a version via parameters.

* **`ListArtifactVersions(ctx context.Context, appName string, userId string, sessionId string, artifactName string, ...)`**
    * Retrieves a list of available version numbers for a specific artifact within a session.

* **`LoadArtifactVersion(ctx context.Context, appName string, userId string, sessionId string, artifactName string, versionId int, ...)`**
    * Loads the content of a specific version of an artifact from a session.

* **`GetEventGraph(ctx context.Context, appName string, userId string, sessionId string, eventId string, ...)`**
    * Retrieves the graph structure related to a specific event within a session.

* **`GetTraceDict(ctx context.Context, eventId string, ...)`**
    * Retrieves detailed trace information (as a dictionary/map) for a specific event ID, primarily for debugging.

* **`DevUi(ctx context.Context, ...)`**
    * Serves the ADK Developer UI static files or entry point.

* **`ListApps(ctx context.Context, ...)`**
    * Retrieves a list of all registered application names known to the ADK server.

* **`Run(ctx context.Context, body RunJSONRequestBody, ...)`**
    * Executes an agent run based on the request body (containing app name, user ID, session ID, and new message) and returns the resulting events in a single response.

* **`RunSse(ctx context.Context, body RunSseJSONRequestBody, ...)`**
    * Executes an agent run similar to `Run`, but streams the resulting events back using Server-Sent Events (SSE).

## Contributing

Contributions are welcome\! Please feel free to submit pull requests or open issues on the [GitHub repository](https://github.com/goudtho/adk-client-go).

## License

MIT License
