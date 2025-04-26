package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/oapi-codegen/runtime"
)

// Defines values for APIKeyIn.
const (
	Cookie APIKeyIn = "cookie"
	Header APIKeyIn = "header"
	Query  APIKeyIn = "query"
)

// Defines values for AuthCredentialTypes.
const (
	AuthCredentialTypesApiKey         AuthCredentialTypes = "apiKey"
	AuthCredentialTypesHttp           AuthCredentialTypes = "http"
	AuthCredentialTypesOauth2         AuthCredentialTypes = "oauth2"
	AuthCredentialTypesOpenIdConnect  AuthCredentialTypes = "openIdConnect"
	AuthCredentialTypesServiceAccount AuthCredentialTypes = "serviceAccount"
)

// Defines values for EvalStatus.
const (
	N1 EvalStatus = 1
	N2 EvalStatus = 2
	N3 EvalStatus = 3
)

// Defines values for Language.
const (
	LANGUAGEUNSPECIFIED Language = "LANGUAGE_UNSPECIFIED"
	PYTHON              Language = "PYTHON"
)

// Defines values for Outcome.
const (
	OUTCOMEDEADLINEEXCEEDED Outcome = "OUTCOME_DEADLINE_EXCEEDED"
	OUTCOMEFAILED           Outcome = "OUTCOME_FAILED"
	OUTCOMEOK               Outcome = "OUTCOME_OK"
	OUTCOMEUNSPECIFIED      Outcome = "OUTCOME_UNSPECIFIED"
)

// Defines values for SecuritySchemeType.
const (
	SecuritySchemeTypeApiKey        SecuritySchemeType = "apiKey"
	SecuritySchemeTypeHttp          SecuritySchemeType = "http"
	SecuritySchemeTypeOauth2        SecuritySchemeType = "oauth2"
	SecuritySchemeTypeOpenIdConnect SecuritySchemeType = "openIdConnect"
)

// APIKey defines model for APIKey.
type APIKey struct {
	Description          *APIKey_Description    `json:"description,omitempty"`
	In                   APIKeyIn               `json:"in"`
	Name                 string                 `json:"name"`
	Type                 *SecuritySchemeType    `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// APIKeyDescription0 defines model for .
type APIKeyDescription0 = string

// APIKeyDescription1 defines model for .
type APIKeyDescription1 = string

// APIKey_Description defines model for APIKey.Description.
type APIKey_Description struct {
	union json.RawMessage
}

// APIKeyIn defines model for APIKeyIn.
type APIKeyIn string

// AddSessionToEvalSetRequest defines model for AddSessionToEvalSetRequest.
type AddSessionToEvalSetRequest struct {
	EvalId    string `json:"eval_id"`
	SessionId string `json:"session_id"`
	UserId    string `json:"user_id"`
}

// AgentRunRequest defines model for AgentRunRequest.
type AgentRunRequest struct {
	AppName string `json:"app_name"`

	// NewMessage Contains the multi-part content of a message.
	NewMessage ContentInput `json:"new_message"`
	SessionId  string       `json:"session_id"`
	Streaming  *bool        `json:"streaming,omitempty"`
	UserId     string       `json:"user_id"`
}

// AuthConfig The auth config sent by tool asking client to collect auth credentials and
//
// adk and client will help to fill in the response
type AuthConfig struct {
	AuthScheme AuthConfig_AuthScheme `json:"auth_scheme"`

	// ExchangedAuthCredential Data class representing an authentication credential.
	//
	// To exchange for the actual credential, please use
	// CredentialExchanger.exchange_credential().
	//
	// Examples: API Key Auth
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.API_KEY,
	//     api_key="1234",
	// )
	//
	// Example: HTTP Auth
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.HTTP,
	//     http=HttpAuth(
	//         scheme="basic",
	//         credentials=HttpCredentials(username="user", password="password"),
	//     ),
	// )
	//
	// Example: OAuth2 Bearer Token in HTTP Header
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.HTTP,
	//     http=HttpAuth(
	//         scheme="bearer",
	//         credentials=HttpCredentials(token="eyAkaknabna...."),
	//     ),
	// )
	//
	// Example: OAuth2 Auth with Authorization Code Flow
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.OAUTH2,
	//     oauth2=OAuth2Auth(
	//         client_id="1234",
	//         client_secret="secret",
	//     ),
	// )
	//
	// Example: OpenID Connect Auth
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.OPEN_ID_CONNECT,
	//     oauth2=OAuth2Auth(
	//         client_id="1234",
	//         client_secret="secret",
	//         redirect_uri="https://example.com",
	//         scopes=["scope1", "scope2"],
	//     ),
	// )
	//
	// Example: Auth with resource reference
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.API_KEY,
	//     resource_ref="projects/1234/locations/us-central1/resources/resource1",
	// )
	ExchangedAuthCredential *AuthCredential `json:"exchanged_auth_credential,omitempty"`

	// RawAuthCredential Data class representing an authentication credential.
	//
	// To exchange for the actual credential, please use
	// CredentialExchanger.exchange_credential().
	//
	// Examples: API Key Auth
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.API_KEY,
	//     api_key="1234",
	// )
	//
	// Example: HTTP Auth
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.HTTP,
	//     http=HttpAuth(
	//         scheme="basic",
	//         credentials=HttpCredentials(username="user", password="password"),
	//     ),
	// )
	//
	// Example: OAuth2 Bearer Token in HTTP Header
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.HTTP,
	//     http=HttpAuth(
	//         scheme="bearer",
	//         credentials=HttpCredentials(token="eyAkaknabna...."),
	//     ),
	// )
	//
	// Example: OAuth2 Auth with Authorization Code Flow
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.OAUTH2,
	//     oauth2=OAuth2Auth(
	//         client_id="1234",
	//         client_secret="secret",
	//     ),
	// )
	//
	// Example: OpenID Connect Auth
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.OPEN_ID_CONNECT,
	//     oauth2=OAuth2Auth(
	//         client_id="1234",
	//         client_secret="secret",
	//         redirect_uri="https://example.com",
	//         scopes=["scope1", "scope2"],
	//     ),
	// )
	//
	// Example: Auth with resource reference
	// AuthCredential(
	//     auth_type=AuthCredentialTypes.API_KEY,
	//     resource_ref="projects/1234/locations/us-central1/resources/resource1",
	// )
	RawAuthCredential *AuthCredential `json:"raw_auth_credential,omitempty"`
}

// AuthConfig_AuthScheme defines model for AuthConfig.AuthScheme.
type AuthConfig_AuthScheme struct {
	union json.RawMessage
}

// AuthCredential Data class representing an authentication credential.
//
// To exchange for the actual credential, please use
// CredentialExchanger.exchange_credential().
//
// Examples: API Key Auth
// AuthCredential(
//
//	auth_type=AuthCredentialTypes.API_KEY,
//	api_key="1234",
//
// )
//
// Example: HTTP Auth
// AuthCredential(
//
//	auth_type=AuthCredentialTypes.HTTP,
//	http=HttpAuth(
//	    scheme="basic",
//	    credentials=HttpCredentials(username="user", password="password"),
//	),
//
// )
//
// Example: OAuth2 Bearer Token in HTTP Header
// AuthCredential(
//
//	auth_type=AuthCredentialTypes.HTTP,
//	http=HttpAuth(
//	    scheme="bearer",
//	    credentials=HttpCredentials(token="eyAkaknabna...."),
//	),
//
// )
//
// Example: OAuth2 Auth with Authorization Code Flow
// AuthCredential(
//
//	auth_type=AuthCredentialTypes.OAUTH2,
//	oauth2=OAuth2Auth(
//	    client_id="1234",
//	    client_secret="secret",
//	),
//
// )
//
// Example: OpenID Connect Auth
// AuthCredential(
//
//	auth_type=AuthCredentialTypes.OPEN_ID_CONNECT,
//	oauth2=OAuth2Auth(
//	    client_id="1234",
//	    client_secret="secret",
//	    redirect_uri="https://example.com",
//	    scopes=["scope1", "scope2"],
//	),
//
// )
//
// Example: Auth with resource reference
// AuthCredential(
//
//	auth_type=AuthCredentialTypes.API_KEY,
//	resource_ref="projects/1234/locations/us-central1/resources/resource1",
//
// )
type AuthCredential struct {
	ApiKey *AuthCredential_ApiKey `json:"api_key,omitempty"`

	// AuthType Represents the type of authentication credential.
	AuthType             AuthCredentialTypes            `json:"auth_type"`
	Http                 *AuthCredential_Http           `json:"http,omitempty"`
	Oauth2               *AuthCredential_Oauth2         `json:"oauth2,omitempty"`
	ResourceRef          *AuthCredential_ResourceRef    `json:"resource_ref,omitempty"`
	ServiceAccount       *AuthCredential_ServiceAccount `json:"service_account,omitempty"`
	AdditionalProperties map[string]interface{}         `json:"-"`
}

// AuthCredentialApiKey0 defines model for .
type AuthCredentialApiKey0 = string

// AuthCredentialApiKey1 defines model for .
type AuthCredentialApiKey1 = string

// AuthCredential_ApiKey defines model for AuthCredential.ApiKey.
type AuthCredential_ApiKey struct {
	union json.RawMessage
}

// AuthCredentialHttp1 defines model for .
type AuthCredentialHttp1 = string

// AuthCredential_Http defines model for AuthCredential.Http.
type AuthCredential_Http struct {
	union json.RawMessage
}

// AuthCredentialOauth21 defines model for .
type AuthCredentialOauth21 = string

// AuthCredential_Oauth2 defines model for AuthCredential.Oauth2.
type AuthCredential_Oauth2 struct {
	union json.RawMessage
}

// AuthCredentialResourceRef0 defines model for .
type AuthCredentialResourceRef0 = string

// AuthCredentialResourceRef1 defines model for .
type AuthCredentialResourceRef1 = string

// AuthCredential_ResourceRef defines model for AuthCredential.ResourceRef.
type AuthCredential_ResourceRef struct {
	union json.RawMessage
}

// AuthCredentialServiceAccount1 defines model for .
type AuthCredentialServiceAccount1 = string

// AuthCredential_ServiceAccount defines model for AuthCredential.ServiceAccount.
type AuthCredential_ServiceAccount struct {
	union json.RawMessage
}

// AuthCredentialTypes Represents the type of authentication credential.
type AuthCredentialTypes string

// Blob Content blob.
type Blob struct {
	// Data Required. Raw bytes.
	Data *Blob_Data `json:"data,omitempty"`

	// MimeType Required. The IANA standard MIME type of the source data.
	MimeType *Blob_MimeType `json:"mimeType,omitempty"`
}

// BlobData0 defines model for .
type BlobData0 = string

// BlobData1 defines model for .
type BlobData1 = string

// Blob_Data Required. Raw bytes.
type Blob_Data struct {
	union json.RawMessage
}

// BlobMimeType0 defines model for .
type BlobMimeType0 = string

// BlobMimeType1 defines model for .
type BlobMimeType1 = string

// Blob_MimeType Required. The IANA standard MIME type of the source data.
type Blob_MimeType struct {
	union json.RawMessage
}

// CodeExecutionResult Result of executing the [ExecutableCode].
//
// Always follows a `part` containing the [ExecutableCode].
type CodeExecutionResult struct {
	// Outcome Required. Outcome of the code execution.
	Outcome *CodeExecutionResult_Outcome `json:"outcome,omitempty"`

	// Output Optional. Contains stdout when code execution is successful, stderr or other description otherwise.
	Output *CodeExecutionResult_Output `json:"output,omitempty"`
}

// CodeExecutionResultOutcome1 defines model for .
type CodeExecutionResultOutcome1 = string

// CodeExecutionResult_Outcome Required. Outcome of the code execution.
type CodeExecutionResult_Outcome struct {
	union json.RawMessage
}

// CodeExecutionResultOutput0 defines model for .
type CodeExecutionResultOutput0 = string

// CodeExecutionResultOutput1 defines model for .
type CodeExecutionResultOutput1 = string

// CodeExecutionResult_Output Optional. Contains stdout when code execution is successful, stderr or other description otherwise.
type CodeExecutionResult_Output struct {
	union json.RawMessage
}

// ContentInput Contains the multi-part content of a message.
type ContentInput struct {
	// Parts List of parts that constitute a single message. Each part may have
	//       a different IANA MIME type.
	Parts *ContentInput_Parts `json:"parts,omitempty"`

	// Role Optional. The producer of the content. Must be either 'user' or
	//       'model'. Useful to set for multi-turn conversations, otherwise can be
	//       empty. If role is not specified, SDK will determine the role.
	Role *ContentInput_Role `json:"role,omitempty"`
}

// ContentInputParts0 defines model for .
type ContentInputParts0 = []PartInput

// ContentInputParts1 defines model for .
type ContentInputParts1 = string

// ContentInput_Parts List of parts that constitute a single message. Each part may have
//
//	a different IANA MIME type.
type ContentInput_Parts struct {
	union json.RawMessage
}

// ContentInputRole0 defines model for .
type ContentInputRole0 = string

// ContentInputRole1 defines model for .
type ContentInputRole1 = string

// ContentInput_Role Optional. The producer of the content. Must be either 'user' or
//
//	'model'. Useful to set for multi-turn conversations, otherwise can be
//	empty. If role is not specified, SDK will determine the role.
type ContentInput_Role struct {
	union json.RawMessage
}

// ContentOutput Contains the multi-part content of a message.
type ContentOutput struct {
	// Parts List of parts that constitute a single message. Each part may have
	//       a different IANA MIME type.
	Parts *ContentOutput_Parts `json:"parts,omitempty"`

	// Role Optional. The producer of the content. Must be either 'user' or
	//       'model'. Useful to set for multi-turn conversations, otherwise can be
	//       empty. If role is not specified, SDK will determine the role.
	Role *ContentOutput_Role `json:"role,omitempty"`
}

// ContentOutputParts0 defines model for .
type ContentOutputParts0 = []PartOutput

// ContentOutputParts1 defines model for .
type ContentOutputParts1 = string

// ContentOutput_Parts List of parts that constitute a single message. Each part may have
//
//	a different IANA MIME type.
type ContentOutput_Parts struct {
	union json.RawMessage
}

// ContentOutputRole0 defines model for .
type ContentOutputRole0 = string

// ContentOutputRole1 defines model for .
type ContentOutputRole1 = string

// ContentOutput_Role Optional. The producer of the content. Must be either 'user' or
//
//	'model'. Useful to set for multi-turn conversations, otherwise can be
//	empty. If role is not specified, SDK will determine the role.
type ContentOutput_Role struct {
	union json.RawMessage
}

// EvalMetric defines model for EvalMetric.
type EvalMetric struct {
	MetricName string  `json:"metric_name"`
	Threshold  float32 `json:"threshold"`
}

// EvalStatus defines model for EvalStatus.
type EvalStatus int

// Event Represents an event in a conversation between agents and users.
//
// It is used to store the content of the conversation, as well as the actions
// taken by the agents like function calls, etc.
//
// Attributes:
//
//	invocation_id: The invocation ID of the event.
//	author: "user" or the name of the agent, indicating who appended the event
//	  to the session.
//	actions: The actions taken by the agent.
//	long_running_tool_ids: The ids of the long running function calls.
//	branch: The branch of the event.
//	id: The unique identifier of the event.
//	timestamp: The timestamp of the event.
//	is_final_response: Whether the event is the final response of the agent.
//	get_function_calls: Returns the function calls in the event.
type Event struct {
	// Actions Represents the actions attached to an event.
	Actions            *EventActions             `json:"actions,omitempty"`
	Author             string                    `json:"author"`
	Branch             *Event_Branch             `json:"branch,omitempty"`
	Content            *Event_Content            `json:"content,omitempty"`
	ErrorCode          *Event_ErrorCode          `json:"error_code,omitempty"`
	ErrorMessage       *Event_ErrorMessage       `json:"error_message,omitempty"`
	GroundingMetadata  *Event_GroundingMetadata  `json:"grounding_metadata,omitempty"`
	Id                 *string                   `json:"id,omitempty"`
	Interrupted        *Event_Interrupted        `json:"interrupted,omitempty"`
	InvocationId       *string                   `json:"invocation_id,omitempty"`
	LongRunningToolIds *Event_LongRunningToolIds `json:"long_running_tool_ids,omitempty"`
	Partial            *Event_Partial            `json:"partial,omitempty"`
	Timestamp          *float32                  `json:"timestamp,omitempty"`
	TurnComplete       *Event_TurnComplete       `json:"turn_complete,omitempty"`
}

// EventBranch0 defines model for .
type EventBranch0 = string

// EventBranch1 defines model for .
type EventBranch1 = string

// Event_Branch defines model for Event.Branch.
type Event_Branch struct {
	union json.RawMessage
}

// EventContent1 defines model for .
type EventContent1 = string

// Event_Content defines model for Event.Content.
type Event_Content struct {
	union json.RawMessage
}

// EventErrorCode0 defines model for .
type EventErrorCode0 = string

// EventErrorCode1 defines model for .
type EventErrorCode1 = string

// Event_ErrorCode defines model for Event.ErrorCode.
type Event_ErrorCode struct {
	union json.RawMessage
}

// EventErrorMessage0 defines model for .
type EventErrorMessage0 = string

// EventErrorMessage1 defines model for .
type EventErrorMessage1 = string

// Event_ErrorMessage defines model for Event.ErrorMessage.
type Event_ErrorMessage struct {
	union json.RawMessage
}

// EventGroundingMetadata1 defines model for .
type EventGroundingMetadata1 = string

// Event_GroundingMetadata defines model for Event.GroundingMetadata.
type Event_GroundingMetadata struct {
	union json.RawMessage
}

// EventInterrupted0 defines model for .
type EventInterrupted0 = bool

// EventInterrupted1 defines model for .
type EventInterrupted1 = string

// Event_Interrupted defines model for Event.Interrupted.
type Event_Interrupted struct {
	union json.RawMessage
}

// EventLongRunningToolIds0 defines model for .
type EventLongRunningToolIds0 = []string

// EventLongRunningToolIds1 defines model for .
type EventLongRunningToolIds1 = string

// Event_LongRunningToolIds defines model for Event.LongRunningToolIds.
type Event_LongRunningToolIds struct {
	union json.RawMessage
}

// EventPartial0 defines model for .
type EventPartial0 = bool

// EventPartial1 defines model for .
type EventPartial1 = string

// Event_Partial defines model for Event.Partial.
type Event_Partial struct {
	union json.RawMessage
}

// EventTurnComplete0 defines model for .
type EventTurnComplete0 = bool

// EventTurnComplete1 defines model for .
type EventTurnComplete1 = string

// Event_TurnComplete defines model for Event.TurnComplete.
type Event_TurnComplete struct {
	union json.RawMessage
}

// EventActions Represents the actions attached to an event.
type EventActions struct {
	ArtifactDelta        *map[string]int                 `json:"artifact_delta,omitempty"`
	Escalate             *EventActions_Escalate          `json:"escalate,omitempty"`
	RequestedAuthConfigs *map[string]AuthConfig          `json:"requested_auth_configs,omitempty"`
	SkipSummarization    *EventActions_SkipSummarization `json:"skip_summarization,omitempty"`
	StateDelta           *map[string]interface{}         `json:"state_delta,omitempty"`
	TransferToAgent      *EventActions_TransferToAgent   `json:"transfer_to_agent,omitempty"`
}

// EventActionsEscalate0 defines model for .
type EventActionsEscalate0 = bool

// EventActionsEscalate1 defines model for .
type EventActionsEscalate1 = string

// EventActions_Escalate defines model for EventActions.Escalate.
type EventActions_Escalate struct {
	union json.RawMessage
}

// EventActionsSkipSummarization0 defines model for .
type EventActionsSkipSummarization0 = bool

// EventActionsSkipSummarization1 defines model for .
type EventActionsSkipSummarization1 = string

// EventActions_SkipSummarization defines model for EventActions.SkipSummarization.
type EventActions_SkipSummarization struct {
	union json.RawMessage
}

// EventActionsTransferToAgent0 defines model for .
type EventActionsTransferToAgent0 = string

// EventActionsTransferToAgent1 defines model for .
type EventActionsTransferToAgent1 = string

// EventActions_TransferToAgent defines model for EventActions.TransferToAgent.
type EventActions_TransferToAgent struct {
	union json.RawMessage
}

// ExecutableCode Code generated by the model that is meant to be executed, and the result returned to the model.
//
// Generated when using the [FunctionDeclaration] tool and
// [FunctionCallingConfig] mode is set to [Mode.CODE].
type ExecutableCode struct {
	// Code Required. The code to be executed.
	Code *ExecutableCode_Code `json:"code,omitempty"`

	// Language Required. Programming language of the `code`.
	Language *ExecutableCode_Language `json:"language,omitempty"`
}

// ExecutableCodeCode0 defines model for .
type ExecutableCodeCode0 = string

// ExecutableCodeCode1 defines model for .
type ExecutableCodeCode1 = string

// ExecutableCode_Code Required. The code to be executed.
type ExecutableCode_Code struct {
	union json.RawMessage
}

// ExecutableCodeLanguage1 defines model for .
type ExecutableCodeLanguage1 = string

// ExecutableCode_Language Required. Programming language of the `code`.
type ExecutableCode_Language struct {
	union json.RawMessage
}

// FileData URI based data.
type FileData struct {
	// FileUri Required. URI.
	FileUri *FileData_FileUri `json:"fileUri,omitempty"`

	// MimeType Required. The IANA standard MIME type of the source data.
	MimeType *FileData_MimeType `json:"mimeType,omitempty"`
}

// FileDataFileUri0 defines model for .
type FileDataFileUri0 = string

// FileDataFileUri1 defines model for .
type FileDataFileUri1 = string

// FileData_FileUri Required. URI.
type FileData_FileUri struct {
	union json.RawMessage
}

// FileDataMimeType0 defines model for .
type FileDataMimeType0 = string

// FileDataMimeType1 defines model for .
type FileDataMimeType1 = string

// FileData_MimeType Required. The IANA standard MIME type of the source data.
type FileData_MimeType struct {
	union json.RawMessage
}

// FunctionCall A function call.
type FunctionCall struct {
	// Args Optional. Required. The function parameters and values in JSON object format. See [FunctionDeclaration.parameters] for parameter details.
	Args *FunctionCall_Args `json:"args,omitempty"`

	// Id The unique id of the function call. If populated, the client to execute the
	//    `function_call` and return the response with the matching `id`.
	Id *FunctionCall_Id `json:"id,omitempty"`

	// Name Required. The name of the function to call. Matches [FunctionDeclaration.name].
	Name *FunctionCall_Name `json:"name,omitempty"`
}

// FunctionCallArgs0 defines model for .
type FunctionCallArgs0 map[string]interface{}

// FunctionCallArgs1 defines model for .
type FunctionCallArgs1 = string

// FunctionCall_Args Optional. Required. The function parameters and values in JSON object format. See [FunctionDeclaration.parameters] for parameter details.
type FunctionCall_Args struct {
	union json.RawMessage
}

// FunctionCallId0 defines model for .
type FunctionCallId0 = string

// FunctionCallId1 defines model for .
type FunctionCallId1 = string

// FunctionCall_Id The unique id of the function call. If populated, the client to execute the
//
//	`function_call` and return the response with the matching `id`.
type FunctionCall_Id struct {
	union json.RawMessage
}

// FunctionCallName0 defines model for .
type FunctionCallName0 = string

// FunctionCallName1 defines model for .
type FunctionCallName1 = string

// FunctionCall_Name Required. The name of the function to call. Matches [FunctionDeclaration.name].
type FunctionCall_Name struct {
	union json.RawMessage
}

// FunctionResponse A function response.
type FunctionResponse struct {
	// Id The id of the function call this response is for. Populated by the client
	//    to match the corresponding function call `id`.
	Id *FunctionResponse_Id `json:"id,omitempty"`

	// Name Required. The name of the function to call. Matches [FunctionDeclaration.name] and [FunctionCall.name].
	Name *FunctionResponse_Name `json:"name,omitempty"`

	// Response Required. The function response in JSON object format. Use "output" key to specify function output and "error" key to specify error details (if any). If "output" and "error" keys are not specified, then whole "response" is treated as function output.
	Response *FunctionResponse_Response `json:"response,omitempty"`
}

// FunctionResponseId0 defines model for .
type FunctionResponseId0 = string

// FunctionResponseId1 defines model for .
type FunctionResponseId1 = string

// FunctionResponse_Id The id of the function call this response is for. Populated by the client
//
//	to match the corresponding function call `id`.
type FunctionResponse_Id struct {
	union json.RawMessage
}

// FunctionResponseName0 defines model for .
type FunctionResponseName0 = string

// FunctionResponseName1 defines model for .
type FunctionResponseName1 = string

// FunctionResponse_Name Required. The name of the function to call. Matches [FunctionDeclaration.name] and [FunctionCall.name].
type FunctionResponse_Name struct {
	union json.RawMessage
}

// FunctionResponseResponse0 defines model for .
type FunctionResponseResponse0 map[string]interface{}

// FunctionResponseResponse1 defines model for .
type FunctionResponseResponse1 = string

// FunctionResponse_Response Required. The function response in JSON object format. Use "output" key to specify function output and "error" key to specify error details (if any). If "output" and "error" keys are not specified, then whole "response" is treated as function output.
type FunctionResponse_Response struct {
	union json.RawMessage
}

// GroundingChunk Grounding chunk.
type GroundingChunk struct {
	// RetrievedContext Grounding chunk from context retrieved by the retrieval tools.
	RetrievedContext *GroundingChunk_RetrievedContext `json:"retrievedContext,omitempty"`

	// Web Grounding chunk from the web.
	Web *GroundingChunk_Web `json:"web,omitempty"`
}

// GroundingChunkRetrievedContext1 defines model for .
type GroundingChunkRetrievedContext1 = string

// GroundingChunk_RetrievedContext Grounding chunk from context retrieved by the retrieval tools.
type GroundingChunk_RetrievedContext struct {
	union json.RawMessage
}

// GroundingChunkWeb1 defines model for .
type GroundingChunkWeb1 = string

// GroundingChunk_Web Grounding chunk from the web.
type GroundingChunk_Web struct {
	union json.RawMessage
}

// GroundingChunkRetrievedContext Chunk from context retrieved by the retrieval tools.
type GroundingChunkRetrievedContext struct {
	// Text Text of the attribution.
	Text *GroundingChunkRetrievedContext_Text `json:"text,omitempty"`

	// Title Title of the attribution.
	Title *GroundingChunkRetrievedContext_Title `json:"title,omitempty"`

	// Uri URI reference of the attribution.
	Uri *GroundingChunkRetrievedContext_Uri `json:"uri,omitempty"`
}

// GroundingChunkRetrievedContextText0 defines model for .
type GroundingChunkRetrievedContextText0 = string

// GroundingChunkRetrievedContextText1 defines model for .
type GroundingChunkRetrievedContextText1 = string

// GroundingChunkRetrievedContext_Text Text of the attribution.
type GroundingChunkRetrievedContext_Text struct {
	union json.RawMessage
}

// GroundingChunkRetrievedContextTitle0 defines model for .
type GroundingChunkRetrievedContextTitle0 = string

// GroundingChunkRetrievedContextTitle1 defines model for .
type GroundingChunkRetrievedContextTitle1 = string

// GroundingChunkRetrievedContext_Title Title of the attribution.
type GroundingChunkRetrievedContext_Title struct {
	union json.RawMessage
}

// GroundingChunkRetrievedContextUri0 defines model for .
type GroundingChunkRetrievedContextUri0 = string

// GroundingChunkRetrievedContextUri1 defines model for .
type GroundingChunkRetrievedContextUri1 = string

// GroundingChunkRetrievedContext_Uri URI reference of the attribution.
type GroundingChunkRetrievedContext_Uri struct {
	union json.RawMessage
}

// GroundingChunkWeb Chunk from the web.
type GroundingChunkWeb struct {
	// Domain Domain of the (original) URI.
	Domain *GroundingChunkWeb_Domain `json:"domain,omitempty"`

	// Title Title of the chunk.
	Title *GroundingChunkWeb_Title `json:"title,omitempty"`

	// Uri URI reference of the chunk.
	Uri *GroundingChunkWeb_Uri `json:"uri,omitempty"`
}

// GroundingChunkWebDomain0 defines model for .
type GroundingChunkWebDomain0 = string

// GroundingChunkWebDomain1 defines model for .
type GroundingChunkWebDomain1 = string

// GroundingChunkWeb_Domain Domain of the (original) URI.
type GroundingChunkWeb_Domain struct {
	union json.RawMessage
}

// GroundingChunkWebTitle0 defines model for .
type GroundingChunkWebTitle0 = string

// GroundingChunkWebTitle1 defines model for .
type GroundingChunkWebTitle1 = string

// GroundingChunkWeb_Title Title of the chunk.
type GroundingChunkWeb_Title struct {
	union json.RawMessage
}

// GroundingChunkWebUri0 defines model for .
type GroundingChunkWebUri0 = string

// GroundingChunkWebUri1 defines model for .
type GroundingChunkWebUri1 = string

// GroundingChunkWeb_Uri URI reference of the chunk.
type GroundingChunkWeb_Uri struct {
	union json.RawMessage
}

// GroundingMetadata Metadata returned to client when grounding is enabled.
type GroundingMetadata struct {
	// GroundingChunks List of supporting references retrieved from specified grounding source.
	GroundingChunks *GroundingMetadata_GroundingChunks `json:"groundingChunks,omitempty"`

	// GroundingSupports Optional. List of grounding support.
	GroundingSupports *GroundingMetadata_GroundingSupports `json:"groundingSupports,omitempty"`

	// RetrievalMetadata Optional. Output only. Retrieval metadata.
	RetrievalMetadata *GroundingMetadata_RetrievalMetadata `json:"retrievalMetadata,omitempty"`

	// RetrievalQueries Optional. Queries executed by the retrieval tools.
	RetrievalQueries *GroundingMetadata_RetrievalQueries `json:"retrievalQueries,omitempty"`

	// SearchEntryPoint Optional. Google search entry for the following-up web searches.
	SearchEntryPoint *GroundingMetadata_SearchEntryPoint `json:"searchEntryPoint,omitempty"`

	// WebSearchQueries Optional. Web search queries for the following-up web search.
	WebSearchQueries *GroundingMetadata_WebSearchQueries `json:"webSearchQueries,omitempty"`
}

// GroundingMetadataGroundingChunks0 defines model for .
type GroundingMetadataGroundingChunks0 = []GroundingChunk

// GroundingMetadataGroundingChunks1 defines model for .
type GroundingMetadataGroundingChunks1 = string

// GroundingMetadata_GroundingChunks List of supporting references retrieved from specified grounding source.
type GroundingMetadata_GroundingChunks struct {
	union json.RawMessage
}

// GroundingMetadataGroundingSupports0 defines model for .
type GroundingMetadataGroundingSupports0 = []GroundingSupport

// GroundingMetadataGroundingSupports1 defines model for .
type GroundingMetadataGroundingSupports1 = string

// GroundingMetadata_GroundingSupports Optional. List of grounding support.
type GroundingMetadata_GroundingSupports struct {
	union json.RawMessage
}

// GroundingMetadataRetrievalMetadata1 defines model for .
type GroundingMetadataRetrievalMetadata1 = string

// GroundingMetadata_RetrievalMetadata Optional. Output only. Retrieval metadata.
type GroundingMetadata_RetrievalMetadata struct {
	union json.RawMessage
}

// GroundingMetadataRetrievalQueries0 defines model for .
type GroundingMetadataRetrievalQueries0 = []string

// GroundingMetadataRetrievalQueries1 defines model for .
type GroundingMetadataRetrievalQueries1 = string

// GroundingMetadata_RetrievalQueries Optional. Queries executed by the retrieval tools.
type GroundingMetadata_RetrievalQueries struct {
	union json.RawMessage
}

// GroundingMetadataSearchEntryPoint1 defines model for .
type GroundingMetadataSearchEntryPoint1 = string

// GroundingMetadata_SearchEntryPoint Optional. Google search entry for the following-up web searches.
type GroundingMetadata_SearchEntryPoint struct {
	union json.RawMessage
}

// GroundingMetadataWebSearchQueries0 defines model for .
type GroundingMetadataWebSearchQueries0 = []string

// GroundingMetadataWebSearchQueries1 defines model for .
type GroundingMetadataWebSearchQueries1 = string

// GroundingMetadata_WebSearchQueries Optional. Web search queries for the following-up web search.
type GroundingMetadata_WebSearchQueries struct {
	union json.RawMessage
}

// GroundingSupport Grounding support.
type GroundingSupport struct {
	// ConfidenceScores Confidence score of the support references. Ranges from 0 to 1. 1 is the most confident. This list must have the same size as the grounding_chunk_indices.
	ConfidenceScores *GroundingSupport_ConfidenceScores `json:"confidenceScores,omitempty"`

	// GroundingChunkIndices A list of indices (into 'grounding_chunk') specifying the citations associated with the claim. For instance [1,3,4] means that grounding_chunk[1], grounding_chunk[3], grounding_chunk[4] are the retrieved content attributed to the claim.
	GroundingChunkIndices *GroundingSupport_GroundingChunkIndices `json:"groundingChunkIndices,omitempty"`

	// Segment Segment of the content this support belongs to.
	Segment *GroundingSupport_Segment `json:"segment,omitempty"`
}

// GroundingSupportConfidenceScores0 defines model for .
type GroundingSupportConfidenceScores0 = []float32

// GroundingSupportConfidenceScores1 defines model for .
type GroundingSupportConfidenceScores1 = string

// GroundingSupport_ConfidenceScores Confidence score of the support references. Ranges from 0 to 1. 1 is the most confident. This list must have the same size as the grounding_chunk_indices.
type GroundingSupport_ConfidenceScores struct {
	union json.RawMessage
}

// GroundingSupportGroundingChunkIndices0 defines model for .
type GroundingSupportGroundingChunkIndices0 = []int

// GroundingSupportGroundingChunkIndices1 defines model for .
type GroundingSupportGroundingChunkIndices1 = string

// GroundingSupport_GroundingChunkIndices A list of indices (into 'grounding_chunk') specifying the citations associated with the claim. For instance [1,3,4] means that grounding_chunk[1], grounding_chunk[3], grounding_chunk[4] are the retrieved content attributed to the claim.
type GroundingSupport_GroundingChunkIndices struct {
	union json.RawMessage
}

// GroundingSupportSegment1 defines model for .
type GroundingSupportSegment1 = string

// GroundingSupport_Segment Segment of the content this support belongs to.
type GroundingSupport_Segment struct {
	union json.RawMessage
}

// HTTPBase defines model for HTTPBase.
type HTTPBase struct {
	Description          *HTTPBase_Description  `json:"description,omitempty"`
	Scheme               string                 `json:"scheme"`
	Type                 *SecuritySchemeType    `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// HTTPBaseDescription0 defines model for .
type HTTPBaseDescription0 = string

// HTTPBaseDescription1 defines model for .
type HTTPBaseDescription1 = string

// HTTPBase_Description defines model for HTTPBase.Description.
type HTTPBase_Description struct {
	union json.RawMessage
}

// HTTPBearer defines model for HTTPBearer.
type HTTPBearer struct {
	BearerFormat         *HTTPBearer_BearerFormat `json:"bearerFormat,omitempty"`
	Description          *HTTPBearer_Description  `json:"description,omitempty"`
	Scheme               *string                  `json:"scheme,omitempty"`
	Type                 *SecuritySchemeType      `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}   `json:"-"`
}

// HTTPBearerBearerFormat0 defines model for .
type HTTPBearerBearerFormat0 = string

// HTTPBearerBearerFormat1 defines model for .
type HTTPBearerBearerFormat1 = string

// HTTPBearer_BearerFormat defines model for HTTPBearer.BearerFormat.
type HTTPBearer_BearerFormat struct {
	union json.RawMessage
}

// HTTPBearerDescription0 defines model for .
type HTTPBearerDescription0 = string

// HTTPBearerDescription1 defines model for .
type HTTPBearerDescription1 = string

// HTTPBearer_Description defines model for HTTPBearer.Description.
type HTTPBearer_Description struct {
	union json.RawMessage
}

// HTTPValidationError defines model for HTTPValidationError.
type HTTPValidationError struct {
	Detail *[]ValidationError `json:"detail,omitempty"`
}

// HttpAuth The credentials and metadata for HTTP authentication.
type HttpAuth struct {
	// Credentials Represents the secret token value for HTTP authentication, like user name, password, oauth token, etc.
	Credentials          HttpCredentials        `json:"credentials"`
	Scheme               string                 `json:"scheme"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// HttpCredentials Represents the secret token value for HTTP authentication, like user name, password, oauth token, etc.
type HttpCredentials struct {
	Password             *HttpCredentials_Password `json:"password,omitempty"`
	Token                *HttpCredentials_Token    `json:"token,omitempty"`
	Username             *HttpCredentials_Username `json:"username,omitempty"`
	AdditionalProperties map[string]interface{}    `json:"-"`
}

// HttpCredentialsPassword0 defines model for .
type HttpCredentialsPassword0 = string

// HttpCredentialsPassword1 defines model for .
type HttpCredentialsPassword1 = string

// HttpCredentials_Password defines model for HttpCredentials.Password.
type HttpCredentials_Password struct {
	union json.RawMessage
}

// HttpCredentialsToken0 defines model for .
type HttpCredentialsToken0 = string

// HttpCredentialsToken1 defines model for .
type HttpCredentialsToken1 = string

// HttpCredentials_Token defines model for HttpCredentials.Token.
type HttpCredentials_Token struct {
	union json.RawMessage
}

// HttpCredentialsUsername0 defines model for .
type HttpCredentialsUsername0 = string

// HttpCredentialsUsername1 defines model for .
type HttpCredentialsUsername1 = string

// HttpCredentials_Username defines model for HttpCredentials.Username.
type HttpCredentials_Username struct {
	union json.RawMessage
}

// Language Required. Programming language of the `code`.
type Language string

// OAuth2 defines model for OAuth2.
type OAuth2 struct {
	Description          *OAuth2_Description    `json:"description,omitempty"`
	Flows                OAuthFlows             `json:"flows"`
	Type                 *SecuritySchemeType    `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"-"`
}

// OAuth2Description0 defines model for .
type OAuth2Description0 = string

// OAuth2Description1 defines model for .
type OAuth2Description1 = string

// OAuth2_Description defines model for OAuth2.Description.
type OAuth2_Description struct {
	union json.RawMessage
}

// OAuth2Auth Represents credential value and its metadata for a OAuth2 credential.
type OAuth2Auth struct {
	AuthCode             *OAuth2Auth_AuthCode        `json:"auth_code,omitempty"`
	AuthResponseUri      *OAuth2Auth_AuthResponseUri `json:"auth_response_uri,omitempty"`
	AuthUri              *OAuth2Auth_AuthUri         `json:"auth_uri,omitempty"`
	ClientId             *OAuth2Auth_ClientId        `json:"client_id,omitempty"`
	ClientSecret         *OAuth2Auth_ClientSecret    `json:"client_secret,omitempty"`
	RedirectUri          *OAuth2Auth_RedirectUri     `json:"redirect_uri,omitempty"`
	State                *OAuth2Auth_State           `json:"state,omitempty"`
	Token                *OAuth2Auth_Token           `json:"token,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"-"`
}

// OAuth2AuthAuthCode0 defines model for .
type OAuth2AuthAuthCode0 = string

// OAuth2AuthAuthCode1 defines model for .
type OAuth2AuthAuthCode1 = string

// OAuth2Auth_AuthCode defines model for OAuth2Auth.AuthCode.
type OAuth2Auth_AuthCode struct {
	union json.RawMessage
}

// OAuth2AuthAuthResponseUri0 defines model for .
type OAuth2AuthAuthResponseUri0 = string

// OAuth2AuthAuthResponseUri1 defines model for .
type OAuth2AuthAuthResponseUri1 = string

// OAuth2Auth_AuthResponseUri defines model for OAuth2Auth.AuthResponseUri.
type OAuth2Auth_AuthResponseUri struct {
	union json.RawMessage
}

// OAuth2AuthAuthUri0 defines model for .
type OAuth2AuthAuthUri0 = string

// OAuth2AuthAuthUri1 defines model for .
type OAuth2AuthAuthUri1 = string

// OAuth2Auth_AuthUri defines model for OAuth2Auth.AuthUri.
type OAuth2Auth_AuthUri struct {
	union json.RawMessage
}

// OAuth2AuthClientId0 defines model for .
type OAuth2AuthClientId0 = string

// OAuth2AuthClientId1 defines model for .
type OAuth2AuthClientId1 = string

// OAuth2Auth_ClientId defines model for OAuth2Auth.ClientId.
type OAuth2Auth_ClientId struct {
	union json.RawMessage
}

// OAuth2AuthClientSecret0 defines model for .
type OAuth2AuthClientSecret0 = string

// OAuth2AuthClientSecret1 defines model for .
type OAuth2AuthClientSecret1 = string

// OAuth2Auth_ClientSecret defines model for OAuth2Auth.ClientSecret.
type OAuth2Auth_ClientSecret struct {
	union json.RawMessage
}

// OAuth2AuthRedirectUri0 defines model for .
type OAuth2AuthRedirectUri0 = string

// OAuth2AuthRedirectUri1 defines model for .
type OAuth2AuthRedirectUri1 = string

// OAuth2Auth_RedirectUri defines model for OAuth2Auth.RedirectUri.
type OAuth2Auth_RedirectUri struct {
	union json.RawMessage
}

// OAuth2AuthState0 defines model for .
type OAuth2AuthState0 = string

// OAuth2AuthState1 defines model for .
type OAuth2AuthState1 = string

// OAuth2Auth_State defines model for OAuth2Auth.State.
type OAuth2Auth_State struct {
	union json.RawMessage
}

// OAuth2AuthToken0 defines model for .
type OAuth2AuthToken0 map[string]interface{}

// OAuth2AuthToken1 defines model for .
type OAuth2AuthToken1 = string

// OAuth2Auth_Token defines model for OAuth2Auth.Token.
type OAuth2Auth_Token struct {
	union json.RawMessage
}

// OAuthFlowAuthorizationCode defines model for OAuthFlowAuthorizationCode.
type OAuthFlowAuthorizationCode struct {
	AuthorizationUrl     string                                 `json:"authorizationUrl"`
	RefreshUrl           *OAuthFlowAuthorizationCode_RefreshUrl `json:"refreshUrl,omitempty"`
	Scopes               *map[string]string                     `json:"scopes,omitempty"`
	TokenUrl             string                                 `json:"tokenUrl"`
	AdditionalProperties map[string]interface{}                 `json:"-"`
}

// OAuthFlowAuthorizationCodeRefreshUrl0 defines model for .
type OAuthFlowAuthorizationCodeRefreshUrl0 = string

// OAuthFlowAuthorizationCodeRefreshUrl1 defines model for .
type OAuthFlowAuthorizationCodeRefreshUrl1 = string

// OAuthFlowAuthorizationCode_RefreshUrl defines model for OAuthFlowAuthorizationCode.RefreshUrl.
type OAuthFlowAuthorizationCode_RefreshUrl struct {
	union json.RawMessage
}

// OAuthFlowClientCredentials defines model for OAuthFlowClientCredentials.
type OAuthFlowClientCredentials struct {
	RefreshUrl           *OAuthFlowClientCredentials_RefreshUrl `json:"refreshUrl,omitempty"`
	Scopes               *map[string]string                     `json:"scopes,omitempty"`
	TokenUrl             string                                 `json:"tokenUrl"`
	AdditionalProperties map[string]interface{}                 `json:"-"`
}

// OAuthFlowClientCredentialsRefreshUrl0 defines model for .
type OAuthFlowClientCredentialsRefreshUrl0 = string

// OAuthFlowClientCredentialsRefreshUrl1 defines model for .
type OAuthFlowClientCredentialsRefreshUrl1 = string

// OAuthFlowClientCredentials_RefreshUrl defines model for OAuthFlowClientCredentials.RefreshUrl.
type OAuthFlowClientCredentials_RefreshUrl struct {
	union json.RawMessage
}

// OAuthFlowImplicit defines model for OAuthFlowImplicit.
type OAuthFlowImplicit struct {
	AuthorizationUrl     string                        `json:"authorizationUrl"`
	RefreshUrl           *OAuthFlowImplicit_RefreshUrl `json:"refreshUrl,omitempty"`
	Scopes               *map[string]string            `json:"scopes,omitempty"`
	AdditionalProperties map[string]interface{}        `json:"-"`
}

// OAuthFlowImplicitRefreshUrl0 defines model for .
type OAuthFlowImplicitRefreshUrl0 = string

// OAuthFlowImplicitRefreshUrl1 defines model for .
type OAuthFlowImplicitRefreshUrl1 = string

// OAuthFlowImplicit_RefreshUrl defines model for OAuthFlowImplicit.RefreshUrl.
type OAuthFlowImplicit_RefreshUrl struct {
	union json.RawMessage
}

// OAuthFlowPassword defines model for OAuthFlowPassword.
type OAuthFlowPassword struct {
	RefreshUrl           *OAuthFlowPassword_RefreshUrl `json:"refreshUrl,omitempty"`
	Scopes               *map[string]string            `json:"scopes,omitempty"`
	TokenUrl             string                        `json:"tokenUrl"`
	AdditionalProperties map[string]interface{}        `json:"-"`
}

// OAuthFlowPasswordRefreshUrl0 defines model for .
type OAuthFlowPasswordRefreshUrl0 = string

// OAuthFlowPasswordRefreshUrl1 defines model for .
type OAuthFlowPasswordRefreshUrl1 = string

// OAuthFlowPassword_RefreshUrl defines model for OAuthFlowPassword.RefreshUrl.
type OAuthFlowPassword_RefreshUrl struct {
	union json.RawMessage
}

// OAuthFlows defines model for OAuthFlows.
type OAuthFlows struct {
	AuthorizationCode    *OAuthFlows_AuthorizationCode `json:"authorizationCode,omitempty"`
	ClientCredentials    *OAuthFlows_ClientCredentials `json:"clientCredentials,omitempty"`
	Implicit             *OAuthFlows_Implicit          `json:"implicit,omitempty"`
	Password             *OAuthFlows_Password          `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}        `json:"-"`
}

// OAuthFlowsAuthorizationCode1 defines model for .
type OAuthFlowsAuthorizationCode1 = string

// OAuthFlows_AuthorizationCode defines model for OAuthFlows.AuthorizationCode.
type OAuthFlows_AuthorizationCode struct {
	union json.RawMessage
}

// OAuthFlowsClientCredentials1 defines model for .
type OAuthFlowsClientCredentials1 = string

// OAuthFlows_ClientCredentials defines model for OAuthFlows.ClientCredentials.
type OAuthFlows_ClientCredentials struct {
	union json.RawMessage
}

// OAuthFlowsImplicit1 defines model for .
type OAuthFlowsImplicit1 = string

// OAuthFlows_Implicit defines model for OAuthFlows.Implicit.
type OAuthFlows_Implicit struct {
	union json.RawMessage
}

// OAuthFlowsPassword1 defines model for .
type OAuthFlowsPassword1 = string

// OAuthFlows_Password defines model for OAuthFlows.Password.
type OAuthFlows_Password struct {
	union json.RawMessage
}

// OpenIdConnect defines model for OpenIdConnect.
type OpenIdConnect struct {
	Description          *OpenIdConnect_Description `json:"description,omitempty"`
	OpenIdConnectUrl     string                     `json:"openIdConnectUrl"`
	Type                 *SecuritySchemeType        `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}     `json:"-"`
}

// OpenIdConnectDescription0 defines model for .
type OpenIdConnectDescription0 = string

// OpenIdConnectDescription1 defines model for .
type OpenIdConnectDescription1 = string

// OpenIdConnect_Description defines model for OpenIdConnect.Description.
type OpenIdConnect_Description struct {
	union json.RawMessage
}

// OpenIdConnectWithConfig defines model for OpenIdConnectWithConfig.
type OpenIdConnectWithConfig struct {
	AuthorizationEndpoint             string                                                     `json:"authorization_endpoint"`
	Description                       *OpenIdConnectWithConfig_Description                       `json:"description,omitempty"`
	GrantTypesSupported               *OpenIdConnectWithConfig_GrantTypesSupported               `json:"grant_types_supported,omitempty"`
	RevocationEndpoint                *OpenIdConnectWithConfig_RevocationEndpoint                `json:"revocation_endpoint,omitempty"`
	Scopes                            *OpenIdConnectWithConfig_Scopes                            `json:"scopes,omitempty"`
	TokenEndpoint                     string                                                     `json:"token_endpoint"`
	TokenEndpointAuthMethodsSupported *OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported `json:"token_endpoint_auth_methods_supported,omitempty"`
	Type                              *SecuritySchemeType                                        `json:"type,omitempty"`
	UserinfoEndpoint                  *OpenIdConnectWithConfig_UserinfoEndpoint                  `json:"userinfo_endpoint,omitempty"`
	AdditionalProperties              map[string]interface{}                                     `json:"-"`
}

// OpenIdConnectWithConfigDescription0 defines model for .
type OpenIdConnectWithConfigDescription0 = string

// OpenIdConnectWithConfigDescription1 defines model for .
type OpenIdConnectWithConfigDescription1 = string

// OpenIdConnectWithConfig_Description defines model for OpenIdConnectWithConfig.Description.
type OpenIdConnectWithConfig_Description struct {
	union json.RawMessage
}

// OpenIdConnectWithConfigGrantTypesSupported0 defines model for .
type OpenIdConnectWithConfigGrantTypesSupported0 = []string

// OpenIdConnectWithConfigGrantTypesSupported1 defines model for .
type OpenIdConnectWithConfigGrantTypesSupported1 = string

// OpenIdConnectWithConfig_GrantTypesSupported defines model for OpenIdConnectWithConfig.GrantTypesSupported.
type OpenIdConnectWithConfig_GrantTypesSupported struct {
	union json.RawMessage
}

// OpenIdConnectWithConfigRevocationEndpoint0 defines model for .
type OpenIdConnectWithConfigRevocationEndpoint0 = string

// OpenIdConnectWithConfigRevocationEndpoint1 defines model for .
type OpenIdConnectWithConfigRevocationEndpoint1 = string

// OpenIdConnectWithConfig_RevocationEndpoint defines model for OpenIdConnectWithConfig.RevocationEndpoint.
type OpenIdConnectWithConfig_RevocationEndpoint struct {
	union json.RawMessage
}

// OpenIdConnectWithConfigScopes0 defines model for .
type OpenIdConnectWithConfigScopes0 = []string

// OpenIdConnectWithConfigScopes1 defines model for .
type OpenIdConnectWithConfigScopes1 = string

// OpenIdConnectWithConfig_Scopes defines model for OpenIdConnectWithConfig.Scopes.
type OpenIdConnectWithConfig_Scopes struct {
	union json.RawMessage
}

// OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0 defines model for .
type OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0 = []string

// OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1 defines model for .
type OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1 = string

// OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported defines model for OpenIdConnectWithConfig.TokenEndpointAuthMethodsSupported.
type OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported struct {
	union json.RawMessage
}

// OpenIdConnectWithConfigUserinfoEndpoint0 defines model for .
type OpenIdConnectWithConfigUserinfoEndpoint0 = string

// OpenIdConnectWithConfigUserinfoEndpoint1 defines model for .
type OpenIdConnectWithConfigUserinfoEndpoint1 = string

// OpenIdConnectWithConfig_UserinfoEndpoint defines model for OpenIdConnectWithConfig.UserinfoEndpoint.
type OpenIdConnectWithConfig_UserinfoEndpoint struct {
	union json.RawMessage
}

// Outcome Required. Outcome of the code execution.
type Outcome string

// PartInput A datatype containing media content.
//
// Exactly one field within a Part should be set, representing the specific type
// of content being conveyed. Using multiple fields within the same `Part`
// instance is considered invalid.
type PartInput struct {
	// CodeExecutionResult Optional. Result of executing the [ExecutableCode].
	CodeExecutionResult *PartInput_CodeExecutionResult `json:"codeExecutionResult,omitempty"`

	// ExecutableCode Optional. Code generated by the model that is meant to be executed.
	ExecutableCode *PartInput_ExecutableCode `json:"executableCode,omitempty"`

	// FileData Optional. URI based data.
	FileData *PartInput_FileData `json:"fileData,omitempty"`

	// FunctionCall Optional. A predicted [FunctionCall] returned from the model that contains a string representing the [FunctionDeclaration.name] with the parameters and their values.
	FunctionCall *PartInput_FunctionCall `json:"functionCall,omitempty"`

	// FunctionResponse Optional. The result output of a [FunctionCall] that contains a string representing the [FunctionDeclaration.name] and a structured JSON object containing any output from the function call. It is used as context to the model.
	FunctionResponse *PartInput_FunctionResponse `json:"functionResponse,omitempty"`

	// InlineData Optional. Inlined bytes data.
	InlineData *PartInput_InlineData `json:"inlineData,omitempty"`

	// Text Optional. Text part (can be code).
	Text *PartInput_Text `json:"text,omitempty"`

	// Thought Indicates if the part is thought from the model.
	Thought *PartInput_Thought `json:"thought,omitempty"`

	// VideoMetadata Metadata for a given video.
	VideoMetadata *PartInput_VideoMetadata `json:"videoMetadata,omitempty"`
}

// PartInputCodeExecutionResult1 defines model for .
type PartInputCodeExecutionResult1 = string

// PartInput_CodeExecutionResult Optional. Result of executing the [ExecutableCode].
type PartInput_CodeExecutionResult struct {
	union json.RawMessage
}

// PartInputExecutableCode1 defines model for .
type PartInputExecutableCode1 = string

// PartInput_ExecutableCode Optional. Code generated by the model that is meant to be executed.
type PartInput_ExecutableCode struct {
	union json.RawMessage
}

// PartInputFileData1 defines model for .
type PartInputFileData1 = string

// PartInput_FileData Optional. URI based data.
type PartInput_FileData struct {
	union json.RawMessage
}

// PartInputFunctionCall1 defines model for .
type PartInputFunctionCall1 = string

// PartInput_FunctionCall Optional. A predicted [FunctionCall] returned from the model that contains a string representing the [FunctionDeclaration.name] with the parameters and their values.
type PartInput_FunctionCall struct {
	union json.RawMessage
}

// PartInputFunctionResponse1 defines model for .
type PartInputFunctionResponse1 = string

// PartInput_FunctionResponse Optional. The result output of a [FunctionCall] that contains a string representing the [FunctionDeclaration.name] and a structured JSON object containing any output from the function call. It is used as context to the model.
type PartInput_FunctionResponse struct {
	union json.RawMessage
}

// PartInputInlineData1 defines model for .
type PartInputInlineData1 = string

// PartInput_InlineData Optional. Inlined bytes data.
type PartInput_InlineData struct {
	union json.RawMessage
}

// PartInputText0 defines model for .
type PartInputText0 = string

// PartInputText1 defines model for .
type PartInputText1 = string

// PartInput_Text Optional. Text part (can be code).
type PartInput_Text struct {
	union json.RawMessage
}

// PartInputThought0 defines model for .
type PartInputThought0 = bool

// PartInputThought1 defines model for .
type PartInputThought1 = string

// PartInput_Thought Indicates if the part is thought from the model.
type PartInput_Thought struct {
	union json.RawMessage
}

// PartInputVideoMetadata1 defines model for .
type PartInputVideoMetadata1 = string

// PartInput_VideoMetadata Metadata for a given video.
type PartInput_VideoMetadata struct {
	union json.RawMessage
}

// PartOutput A datatype containing media content.
//
// Exactly one field within a Part should be set, representing the specific type
// of content being conveyed. Using multiple fields within the same `Part`
// instance is considered invalid.
type PartOutput struct {
	// CodeExecutionResult Optional. Result of executing the [ExecutableCode].
	CodeExecutionResult *PartOutput_CodeExecutionResult `json:"codeExecutionResult,omitempty"`

	// ExecutableCode Optional. Code generated by the model that is meant to be executed.
	ExecutableCode *PartOutput_ExecutableCode `json:"executableCode,omitempty"`

	// FileData Optional. URI based data.
	FileData *PartOutput_FileData `json:"fileData,omitempty"`

	// FunctionCall Optional. A predicted [FunctionCall] returned from the model that contains a string representing the [FunctionDeclaration.name] with the parameters and their values.
	FunctionCall *PartOutput_FunctionCall `json:"functionCall,omitempty"`

	// FunctionResponse Optional. The result output of a [FunctionCall] that contains a string representing the [FunctionDeclaration.name] and a structured JSON object containing any output from the function call. It is used as context to the model.
	FunctionResponse *PartOutput_FunctionResponse `json:"functionResponse,omitempty"`

	// InlineData Optional. Inlined bytes data.
	InlineData *PartOutput_InlineData `json:"inlineData,omitempty"`

	// Text Optional. Text part (can be code).
	Text *PartOutput_Text `json:"text,omitempty"`

	// Thought Indicates if the part is thought from the model.
	Thought *PartOutput_Thought `json:"thought,omitempty"`

	// VideoMetadata Metadata for a given video.
	VideoMetadata *PartOutput_VideoMetadata `json:"videoMetadata,omitempty"`
}

// PartOutputCodeExecutionResult1 defines model for .
type PartOutputCodeExecutionResult1 = string

// PartOutput_CodeExecutionResult Optional. Result of executing the [ExecutableCode].
type PartOutput_CodeExecutionResult struct {
	union json.RawMessage
}

// PartOutputExecutableCode1 defines model for .
type PartOutputExecutableCode1 = string

// PartOutput_ExecutableCode Optional. Code generated by the model that is meant to be executed.
type PartOutput_ExecutableCode struct {
	union json.RawMessage
}

// PartOutputFileData1 defines model for .
type PartOutputFileData1 = string

// PartOutput_FileData Optional. URI based data.
type PartOutput_FileData struct {
	union json.RawMessage
}

// PartOutputFunctionCall1 defines model for .
type PartOutputFunctionCall1 = string

// PartOutput_FunctionCall Optional. A predicted [FunctionCall] returned from the model that contains a string representing the [FunctionDeclaration.name] with the parameters and their values.
type PartOutput_FunctionCall struct {
	union json.RawMessage
}

// PartOutputFunctionResponse1 defines model for .
type PartOutputFunctionResponse1 = string

// PartOutput_FunctionResponse Optional. The result output of a [FunctionCall] that contains a string representing the [FunctionDeclaration.name] and a structured JSON object containing any output from the function call. It is used as context to the model.
type PartOutput_FunctionResponse struct {
	union json.RawMessage
}

// PartOutputInlineData1 defines model for .
type PartOutputInlineData1 = string

// PartOutput_InlineData Optional. Inlined bytes data.
type PartOutput_InlineData struct {
	union json.RawMessage
}

// PartOutputText0 defines model for .
type PartOutputText0 = string

// PartOutputText1 defines model for .
type PartOutputText1 = string

// PartOutput_Text Optional. Text part (can be code).
type PartOutput_Text struct {
	union json.RawMessage
}

// PartOutputThought0 defines model for .
type PartOutputThought0 = bool

// PartOutputThought1 defines model for .
type PartOutputThought1 = string

// PartOutput_Thought Indicates if the part is thought from the model.
type PartOutput_Thought struct {
	union json.RawMessage
}

// PartOutputVideoMetadata1 defines model for .
type PartOutputVideoMetadata1 = string

// PartOutput_VideoMetadata Metadata for a given video.
type PartOutput_VideoMetadata struct {
	union json.RawMessage
}

// RetrievalMetadata Metadata related to retrieval in the grounding flow.
type RetrievalMetadata struct {
	// GoogleSearchDynamicRetrievalScore Optional. Score indicating how likely information from Google Search could help answer the prompt. The score is in the range `[0, 1]`, where 0 is the least likely and 1 is the most likely. This score is only populated when Google Search grounding and dynamic retrieval is enabled. It will be compared to the threshold to determine whether to trigger Google Search.
	GoogleSearchDynamicRetrievalScore *RetrievalMetadata_GoogleSearchDynamicRetrievalScore `json:"googleSearchDynamicRetrievalScore,omitempty"`
}

// RetrievalMetadataGoogleSearchDynamicRetrievalScore0 defines model for .
type RetrievalMetadataGoogleSearchDynamicRetrievalScore0 = float32

// RetrievalMetadataGoogleSearchDynamicRetrievalScore1 defines model for .
type RetrievalMetadataGoogleSearchDynamicRetrievalScore1 = string

// RetrievalMetadata_GoogleSearchDynamicRetrievalScore Optional. Score indicating how likely information from Google Search could help answer the prompt. The score is in the range `[0, 1]`, where 0 is the least likely and 1 is the most likely. This score is only populated when Google Search grounding and dynamic retrieval is enabled. It will be compared to the threshold to determine whether to trigger Google Search.
type RetrievalMetadata_GoogleSearchDynamicRetrievalScore struct {
	union json.RawMessage
}

// RunEvalRequest defines model for RunEvalRequest.
type RunEvalRequest struct {
	EvalIds     []string     `json:"eval_ids"`
	EvalMetrics []EvalMetric `json:"eval_metrics"`
}

// RunEvalResult defines model for RunEvalResult.
type RunEvalResult struct {
	EvalId            string          `json:"eval_id"`
	EvalMetricResults [][]interface{} `json:"eval_metric_results"`
	EvalSetId         string          `json:"eval_set_id"`
	FinalEvalStatus   EvalStatus      `json:"final_eval_status"`
	SessionId         string          `json:"session_id"`
}

// SearchEntryPoint Google search entry point.
type SearchEntryPoint struct {
	// RenderedContent Optional. Web content snippet that can be embedded in a web page or an app webview.
	RenderedContent *SearchEntryPoint_RenderedContent `json:"renderedContent,omitempty"`

	// SdkBlob Optional. Base64 encoded JSON representing array of tuple.
	SdkBlob *SearchEntryPoint_SdkBlob `json:"sdkBlob,omitempty"`
}

// SearchEntryPointRenderedContent0 defines model for .
type SearchEntryPointRenderedContent0 = string

// SearchEntryPointRenderedContent1 defines model for .
type SearchEntryPointRenderedContent1 = string

// SearchEntryPoint_RenderedContent Optional. Web content snippet that can be embedded in a web page or an app webview.
type SearchEntryPoint_RenderedContent struct {
	union json.RawMessage
}

// SearchEntryPointSdkBlob0 defines model for .
type SearchEntryPointSdkBlob0 = string

// SearchEntryPointSdkBlob1 defines model for .
type SearchEntryPointSdkBlob1 = string

// SearchEntryPoint_SdkBlob Optional. Base64 encoded JSON representing array of tuple.
type SearchEntryPoint_SdkBlob struct {
	union json.RawMessage
}

// SecuritySchemeType defines model for SecuritySchemeType.
type SecuritySchemeType string

// Segment Segment of the content.
type Segment struct {
	// EndIndex Output only. End index in the given Part, measured in bytes. Offset from the start of the Part, exclusive, starting at zero.
	EndIndex *Segment_EndIndex `json:"endIndex,omitempty"`

	// PartIndex Output only. The index of a Part object within its parent Content object.
	PartIndex *Segment_PartIndex `json:"partIndex,omitempty"`

	// StartIndex Output only. Start index in the given Part, measured in bytes. Offset from the start of the Part, inclusive, starting at zero.
	StartIndex *Segment_StartIndex `json:"startIndex,omitempty"`

	// Text Output only. The text corresponding to the segment from the response.
	Text *Segment_Text `json:"text,omitempty"`
}

// SegmentEndIndex0 defines model for .
type SegmentEndIndex0 = int

// SegmentEndIndex1 defines model for .
type SegmentEndIndex1 = string

// Segment_EndIndex Output only. End index in the given Part, measured in bytes. Offset from the start of the Part, exclusive, starting at zero.
type Segment_EndIndex struct {
	union json.RawMessage
}

// SegmentPartIndex0 defines model for .
type SegmentPartIndex0 = int

// SegmentPartIndex1 defines model for .
type SegmentPartIndex1 = string

// Segment_PartIndex Output only. The index of a Part object within its parent Content object.
type Segment_PartIndex struct {
	union json.RawMessage
}

// SegmentStartIndex0 defines model for .
type SegmentStartIndex0 = int

// SegmentStartIndex1 defines model for .
type SegmentStartIndex1 = string

// Segment_StartIndex Output only. Start index in the given Part, measured in bytes. Offset from the start of the Part, inclusive, starting at zero.
type Segment_StartIndex struct {
	union json.RawMessage
}

// SegmentText0 defines model for .
type SegmentText0 = string

// SegmentText1 defines model for .
type SegmentText1 = string

// Segment_Text Output only. The text corresponding to the segment from the response.
type Segment_Text struct {
	union json.RawMessage
}

// ServiceAccount Represents Google Service Account configuration.
type ServiceAccount struct {
	Scopes                   []string                                 `json:"scopes"`
	ServiceAccountCredential *ServiceAccount_ServiceAccountCredential `json:"service_account_credential,omitempty"`
	UseDefaultCredential     *ServiceAccount_UseDefaultCredential     `json:"use_default_credential,omitempty"`
	AdditionalProperties     map[string]interface{}                   `json:"-"`
}

// ServiceAccountServiceAccountCredential1 defines model for .
type ServiceAccountServiceAccountCredential1 = string

// ServiceAccount_ServiceAccountCredential defines model for ServiceAccount.ServiceAccountCredential.
type ServiceAccount_ServiceAccountCredential struct {
	union json.RawMessage
}

// ServiceAccountUseDefaultCredential0 defines model for .
type ServiceAccountUseDefaultCredential0 = bool

// ServiceAccountUseDefaultCredential1 defines model for .
type ServiceAccountUseDefaultCredential1 = string

// ServiceAccount_UseDefaultCredential defines model for ServiceAccount.UseDefaultCredential.
type ServiceAccount_UseDefaultCredential struct {
	union json.RawMessage
}

// ServiceAccountCredential Represents Google Service Account configuration.
//
// Attributes:
//
//	type: The type should be "service_account".
//	project_id: The project ID.
//	private_key_id: The ID of the private key.
//	private_key: The private key.
//	client_email: The client email.
//	client_id: The client ID.
//	auth_uri: The authorization URI.
//	token_uri: The token URI.
//	auth_provider_x509_cert_url: URL for auth provider's X.509 cert.
//	client_x509_cert_url: URL for the client's X.509 cert.
//	universe_domain: The universe domain.
//
// Example:
//
//	config = ServiceAccountCredential(
//	    type_="service_account",
//	    project_id="your_project_id",
//	    private_key_id="your_private_key_id",
//	    private_key="-----BEGIN PRIVATE KEY-----...",
//	    client_email="...@....iam.gserviceaccount.com",
//	    client_id="your_client_id",
//	    auth_uri="https://accounts.google.com/o/oauth2/auth",
//	    token_uri="https://oauth2.googleapis.com/token",
//	    auth_provider_x509_cert_url="https://www.googleapis.com/oauth2/v1/certs",
//	    client_x509_cert_url="https://www.googleapis.com/robot/v1/metadata/x509/...",
//	    universe_domain="googleapis.com"
//	)
//
//
//	config = ServiceAccountConfig.model_construct(**{
//	    ...service account config dict
//	})
type ServiceAccountCredential struct {
	AuthProviderX509CertUrl string                 `json:"auth_provider_x509_cert_url"`
	AuthUri                 string                 `json:"auth_uri"`
	ClientEmail             string                 `json:"client_email"`
	ClientId                string                 `json:"client_id"`
	ClientX509CertUrl       string                 `json:"client_x509_cert_url"`
	PrivateKey              string                 `json:"private_key"`
	PrivateKeyId            string                 `json:"private_key_id"`
	ProjectId               string                 `json:"project_id"`
	TokenUri                string                 `json:"token_uri"`
	Type                    *string                `json:"type,omitempty"`
	UniverseDomain          string                 `json:"universe_domain"`
	AdditionalProperties    map[string]interface{} `json:"-"`
}

// Session Represents a series of interactions between a user and agents.
//
// Attributes:
//
//	id: The unique identifier of the session.
//	app_name: The name of the app.
//	user_id: The id of the user.
//	state: The state of the session.
//	events: The events of the session, e.g. user input, model response, function
//	  call/response, etc.
//	last_update_time: The last update time of the session.
type Session struct {
	AppName        string                  `json:"app_name"`
	Events         *[]Event                `json:"events,omitempty"`
	Id             string                  `json:"id"`
	LastUpdateTime *float32                `json:"last_update_time,omitempty"`
	State          *map[string]interface{} `json:"state,omitempty"`
	UserId         string                  `json:"user_id"`
}

// ValidationError defines model for ValidationError.
type ValidationError struct {
	Loc  []ValidationError_Loc_Item `json:"loc"`
	Msg  string                     `json:"msg"`
	Type string                     `json:"type"`
}

// ValidationErrorLoc0 defines model for .
type ValidationErrorLoc0 = string

// ValidationErrorLoc1 defines model for .
type ValidationErrorLoc1 = int

// ValidationError_Loc_Item defines model for ValidationError.loc.Item.
type ValidationError_Loc_Item struct {
	union json.RawMessage
}

// VideoMetadata Metadata describes the input video content.
type VideoMetadata struct {
	// EndOffset Optional. The end offset of the video.
	EndOffset *VideoMetadata_EndOffset `json:"endOffset,omitempty"`

	// StartOffset Optional. The start offset of the video.
	StartOffset *VideoMetadata_StartOffset `json:"startOffset,omitempty"`
}

// VideoMetadataEndOffset0 defines model for .
type VideoMetadataEndOffset0 = string

// VideoMetadataEndOffset1 defines model for .
type VideoMetadataEndOffset1 = string

// VideoMetadata_EndOffset Optional. The end offset of the video.
type VideoMetadata_EndOffset struct {
	union json.RawMessage
}

// VideoMetadataStartOffset0 defines model for .
type VideoMetadataStartOffset0 = string

// VideoMetadataStartOffset1 defines model for .
type VideoMetadataStartOffset1 = string

// VideoMetadata_StartOffset Optional. The start offset of the video.
type VideoMetadata_StartOffset struct {
	union json.RawMessage
}

// CreateSessionJSONBody defines parameters for CreateSession.
type CreateSessionJSONBody struct {
	union json.RawMessage
}

// CreateSessionJSONBody0 defines parameters for CreateSession.
type CreateSessionJSONBody0 map[string]interface{}

// CreateSessionJSONBody1 defines parameters for CreateSession.
type CreateSessionJSONBody1 = string

// CreateSessionWithIdJSONBody defines parameters for CreateSessionWithId.
type CreateSessionWithIdJSONBody struct {
	union json.RawMessage
}

// CreateSessionWithIdJSONBody0 defines parameters for CreateSessionWithId.
type CreateSessionWithIdJSONBody0 map[string]interface{}

// CreateSessionWithIdJSONBody1 defines parameters for CreateSessionWithId.
type CreateSessionWithIdJSONBody1 = string

// LoadArtifactParams defines parameters for LoadArtifact.
type LoadArtifactParams struct {
	Version *struct {
		union json.RawMessage
	} `form:"version,omitempty" json:"version,omitempty"`
}

// LoadArtifactParamsVersion0 defines parameters for LoadArtifact.
type LoadArtifactParamsVersion0 = int

// LoadArtifactParamsVersion1 defines parameters for LoadArtifact.
type LoadArtifactParamsVersion1 = string

// AddSessionToEvalSetJSONRequestBody defines body for AddSessionToEvalSet for application/json ContentType.
type AddSessionToEvalSetJSONRequestBody = AddSessionToEvalSetRequest

// RunEvalJSONRequestBody defines body for RunEval for application/json ContentType.
type RunEvalJSONRequestBody = RunEvalRequest

// CreateSessionJSONRequestBody defines body for CreateSession for application/json ContentType.
type CreateSessionJSONRequestBody CreateSessionJSONBody

// CreateSessionWithIdJSONRequestBody defines body for CreateSessionWithId for application/json ContentType.
type CreateSessionWithIdJSONRequestBody CreateSessionWithIdJSONBody

// RunJSONRequestBody defines body for Run for application/json ContentType.
type RunJSONRequestBody = AgentRunRequest

// RunSseJSONRequestBody defines body for RunSse for application/json ContentType.
type RunSseJSONRequestBody = AgentRunRequest

// Getter for additional properties for APIKey. Returns the specified
// element and whether it was found
func (a APIKey) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for APIKey
func (a *APIKey) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for APIKey to handle AdditionalProperties
func (a *APIKey) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["in"]; found {
		err = json.Unmarshal(raw, &a.In)
		if err != nil {
			return fmt.Errorf("error reading 'in': %w", err)
		}
		delete(object, "in")
	}

	if raw, found := object["name"]; found {
		err = json.Unmarshal(raw, &a.Name)
		if err != nil {
			return fmt.Errorf("error reading 'name': %w", err)
		}
		delete(object, "name")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for APIKey to handle AdditionalProperties
func (a APIKey) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'description': %w", err)
		}
	}

	object["in"], err = json.Marshal(a.In)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'in': %w", err)
	}

	object["name"], err = json.Marshal(a.Name)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'name': %w", err)
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for AuthCredential. Returns the specified
// element and whether it was found
func (a AuthCredential) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for AuthCredential
func (a *AuthCredential) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for AuthCredential to handle AdditionalProperties
func (a *AuthCredential) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["api_key"]; found {
		err = json.Unmarshal(raw, &a.ApiKey)
		if err != nil {
			return fmt.Errorf("error reading 'api_key': %w", err)
		}
		delete(object, "api_key")
	}

	if raw, found := object["auth_type"]; found {
		err = json.Unmarshal(raw, &a.AuthType)
		if err != nil {
			return fmt.Errorf("error reading 'auth_type': %w", err)
		}
		delete(object, "auth_type")
	}

	if raw, found := object["http"]; found {
		err = json.Unmarshal(raw, &a.Http)
		if err != nil {
			return fmt.Errorf("error reading 'http': %w", err)
		}
		delete(object, "http")
	}

	if raw, found := object["oauth2"]; found {
		err = json.Unmarshal(raw, &a.Oauth2)
		if err != nil {
			return fmt.Errorf("error reading 'oauth2': %w", err)
		}
		delete(object, "oauth2")
	}

	if raw, found := object["resource_ref"]; found {
		err = json.Unmarshal(raw, &a.ResourceRef)
		if err != nil {
			return fmt.Errorf("error reading 'resource_ref': %w", err)
		}
		delete(object, "resource_ref")
	}

	if raw, found := object["service_account"]; found {
		err = json.Unmarshal(raw, &a.ServiceAccount)
		if err != nil {
			return fmt.Errorf("error reading 'service_account': %w", err)
		}
		delete(object, "service_account")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for AuthCredential to handle AdditionalProperties
func (a AuthCredential) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.ApiKey != nil {
		object["api_key"], err = json.Marshal(a.ApiKey)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'api_key': %w", err)
		}
	}

	object["auth_type"], err = json.Marshal(a.AuthType)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'auth_type': %w", err)
	}

	if a.Http != nil {
		object["http"], err = json.Marshal(a.Http)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'http': %w", err)
		}
	}

	if a.Oauth2 != nil {
		object["oauth2"], err = json.Marshal(a.Oauth2)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'oauth2': %w", err)
		}
	}

	if a.ResourceRef != nil {
		object["resource_ref"], err = json.Marshal(a.ResourceRef)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'resource_ref': %w", err)
		}
	}

	if a.ServiceAccount != nil {
		object["service_account"], err = json.Marshal(a.ServiceAccount)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'service_account': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for HTTPBase. Returns the specified
// element and whether it was found
func (a HTTPBase) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for HTTPBase
func (a *HTTPBase) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for HTTPBase to handle AdditionalProperties
func (a *HTTPBase) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["scheme"]; found {
		err = json.Unmarshal(raw, &a.Scheme)
		if err != nil {
			return fmt.Errorf("error reading 'scheme': %w", err)
		}
		delete(object, "scheme")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for HTTPBase to handle AdditionalProperties
func (a HTTPBase) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'description': %w", err)
		}
	}

	object["scheme"], err = json.Marshal(a.Scheme)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'scheme': %w", err)
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for HTTPBearer. Returns the specified
// element and whether it was found
func (a HTTPBearer) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for HTTPBearer
func (a *HTTPBearer) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for HTTPBearer to handle AdditionalProperties
func (a *HTTPBearer) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["bearerFormat"]; found {
		err = json.Unmarshal(raw, &a.BearerFormat)
		if err != nil {
			return fmt.Errorf("error reading 'bearerFormat': %w", err)
		}
		delete(object, "bearerFormat")
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["scheme"]; found {
		err = json.Unmarshal(raw, &a.Scheme)
		if err != nil {
			return fmt.Errorf("error reading 'scheme': %w", err)
		}
		delete(object, "scheme")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for HTTPBearer to handle AdditionalProperties
func (a HTTPBearer) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.BearerFormat != nil {
		object["bearerFormat"], err = json.Marshal(a.BearerFormat)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'bearerFormat': %w", err)
		}
	}

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'description': %w", err)
		}
	}

	if a.Scheme != nil {
		object["scheme"], err = json.Marshal(a.Scheme)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'scheme': %w", err)
		}
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for HttpAuth. Returns the specified
// element and whether it was found
func (a HttpAuth) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for HttpAuth
func (a *HttpAuth) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for HttpAuth to handle AdditionalProperties
func (a *HttpAuth) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["credentials"]; found {
		err = json.Unmarshal(raw, &a.Credentials)
		if err != nil {
			return fmt.Errorf("error reading 'credentials': %w", err)
		}
		delete(object, "credentials")
	}

	if raw, found := object["scheme"]; found {
		err = json.Unmarshal(raw, &a.Scheme)
		if err != nil {
			return fmt.Errorf("error reading 'scheme': %w", err)
		}
		delete(object, "scheme")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for HttpAuth to handle AdditionalProperties
func (a HttpAuth) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["credentials"], err = json.Marshal(a.Credentials)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'credentials': %w", err)
	}

	object["scheme"], err = json.Marshal(a.Scheme)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'scheme': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for HttpCredentials. Returns the specified
// element and whether it was found
func (a HttpCredentials) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for HttpCredentials
func (a *HttpCredentials) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for HttpCredentials to handle AdditionalProperties
func (a *HttpCredentials) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["password"]; found {
		err = json.Unmarshal(raw, &a.Password)
		if err != nil {
			return fmt.Errorf("error reading 'password': %w", err)
		}
		delete(object, "password")
	}

	if raw, found := object["token"]; found {
		err = json.Unmarshal(raw, &a.Token)
		if err != nil {
			return fmt.Errorf("error reading 'token': %w", err)
		}
		delete(object, "token")
	}

	if raw, found := object["username"]; found {
		err = json.Unmarshal(raw, &a.Username)
		if err != nil {
			return fmt.Errorf("error reading 'username': %w", err)
		}
		delete(object, "username")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for HttpCredentials to handle AdditionalProperties
func (a HttpCredentials) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Password != nil {
		object["password"], err = json.Marshal(a.Password)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'password': %w", err)
		}
	}

	if a.Token != nil {
		object["token"], err = json.Marshal(a.Token)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'token': %w", err)
		}
	}

	if a.Username != nil {
		object["username"], err = json.Marshal(a.Username)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'username': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OAuth2. Returns the specified
// element and whether it was found
func (a OAuth2) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OAuth2
func (a *OAuth2) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OAuth2 to handle AdditionalProperties
func (a *OAuth2) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["flows"]; found {
		err = json.Unmarshal(raw, &a.Flows)
		if err != nil {
			return fmt.Errorf("error reading 'flows': %w", err)
		}
		delete(object, "flows")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OAuth2 to handle AdditionalProperties
func (a OAuth2) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'description': %w", err)
		}
	}

	object["flows"], err = json.Marshal(a.Flows)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'flows': %w", err)
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OAuth2Auth. Returns the specified
// element and whether it was found
func (a OAuth2Auth) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OAuth2Auth
func (a *OAuth2Auth) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OAuth2Auth to handle AdditionalProperties
func (a *OAuth2Auth) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["auth_code"]; found {
		err = json.Unmarshal(raw, &a.AuthCode)
		if err != nil {
			return fmt.Errorf("error reading 'auth_code': %w", err)
		}
		delete(object, "auth_code")
	}

	if raw, found := object["auth_response_uri"]; found {
		err = json.Unmarshal(raw, &a.AuthResponseUri)
		if err != nil {
			return fmt.Errorf("error reading 'auth_response_uri': %w", err)
		}
		delete(object, "auth_response_uri")
	}

	if raw, found := object["auth_uri"]; found {
		err = json.Unmarshal(raw, &a.AuthUri)
		if err != nil {
			return fmt.Errorf("error reading 'auth_uri': %w", err)
		}
		delete(object, "auth_uri")
	}

	if raw, found := object["client_id"]; found {
		err = json.Unmarshal(raw, &a.ClientId)
		if err != nil {
			return fmt.Errorf("error reading 'client_id': %w", err)
		}
		delete(object, "client_id")
	}

	if raw, found := object["client_secret"]; found {
		err = json.Unmarshal(raw, &a.ClientSecret)
		if err != nil {
			return fmt.Errorf("error reading 'client_secret': %w", err)
		}
		delete(object, "client_secret")
	}

	if raw, found := object["redirect_uri"]; found {
		err = json.Unmarshal(raw, &a.RedirectUri)
		if err != nil {
			return fmt.Errorf("error reading 'redirect_uri': %w", err)
		}
		delete(object, "redirect_uri")
	}

	if raw, found := object["state"]; found {
		err = json.Unmarshal(raw, &a.State)
		if err != nil {
			return fmt.Errorf("error reading 'state': %w", err)
		}
		delete(object, "state")
	}

	if raw, found := object["token"]; found {
		err = json.Unmarshal(raw, &a.Token)
		if err != nil {
			return fmt.Errorf("error reading 'token': %w", err)
		}
		delete(object, "token")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OAuth2Auth to handle AdditionalProperties
func (a OAuth2Auth) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.AuthCode != nil {
		object["auth_code"], err = json.Marshal(a.AuthCode)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'auth_code': %w", err)
		}
	}

	if a.AuthResponseUri != nil {
		object["auth_response_uri"], err = json.Marshal(a.AuthResponseUri)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'auth_response_uri': %w", err)
		}
	}

	if a.AuthUri != nil {
		object["auth_uri"], err = json.Marshal(a.AuthUri)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'auth_uri': %w", err)
		}
	}

	if a.ClientId != nil {
		object["client_id"], err = json.Marshal(a.ClientId)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'client_id': %w", err)
		}
	}

	if a.ClientSecret != nil {
		object["client_secret"], err = json.Marshal(a.ClientSecret)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'client_secret': %w", err)
		}
	}

	if a.RedirectUri != nil {
		object["redirect_uri"], err = json.Marshal(a.RedirectUri)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'redirect_uri': %w", err)
		}
	}

	if a.State != nil {
		object["state"], err = json.Marshal(a.State)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'state': %w", err)
		}
	}

	if a.Token != nil {
		object["token"], err = json.Marshal(a.Token)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'token': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OAuthFlowAuthorizationCode. Returns the specified
// element and whether it was found
func (a OAuthFlowAuthorizationCode) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OAuthFlowAuthorizationCode
func (a *OAuthFlowAuthorizationCode) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OAuthFlowAuthorizationCode to handle AdditionalProperties
func (a *OAuthFlowAuthorizationCode) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["authorizationUrl"]; found {
		err = json.Unmarshal(raw, &a.AuthorizationUrl)
		if err != nil {
			return fmt.Errorf("error reading 'authorizationUrl': %w", err)
		}
		delete(object, "authorizationUrl")
	}

	if raw, found := object["refreshUrl"]; found {
		err = json.Unmarshal(raw, &a.RefreshUrl)
		if err != nil {
			return fmt.Errorf("error reading 'refreshUrl': %w", err)
		}
		delete(object, "refreshUrl")
	}

	if raw, found := object["scopes"]; found {
		err = json.Unmarshal(raw, &a.Scopes)
		if err != nil {
			return fmt.Errorf("error reading 'scopes': %w", err)
		}
		delete(object, "scopes")
	}

	if raw, found := object["tokenUrl"]; found {
		err = json.Unmarshal(raw, &a.TokenUrl)
		if err != nil {
			return fmt.Errorf("error reading 'tokenUrl': %w", err)
		}
		delete(object, "tokenUrl")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OAuthFlowAuthorizationCode to handle AdditionalProperties
func (a OAuthFlowAuthorizationCode) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["authorizationUrl"], err = json.Marshal(a.AuthorizationUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'authorizationUrl': %w", err)
	}

	if a.RefreshUrl != nil {
		object["refreshUrl"], err = json.Marshal(a.RefreshUrl)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'refreshUrl': %w", err)
		}
	}

	if a.Scopes != nil {
		object["scopes"], err = json.Marshal(a.Scopes)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'scopes': %w", err)
		}
	}

	object["tokenUrl"], err = json.Marshal(a.TokenUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'tokenUrl': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OAuthFlowClientCredentials. Returns the specified
// element and whether it was found
func (a OAuthFlowClientCredentials) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OAuthFlowClientCredentials
func (a *OAuthFlowClientCredentials) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OAuthFlowClientCredentials to handle AdditionalProperties
func (a *OAuthFlowClientCredentials) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["refreshUrl"]; found {
		err = json.Unmarshal(raw, &a.RefreshUrl)
		if err != nil {
			return fmt.Errorf("error reading 'refreshUrl': %w", err)
		}
		delete(object, "refreshUrl")
	}

	if raw, found := object["scopes"]; found {
		err = json.Unmarshal(raw, &a.Scopes)
		if err != nil {
			return fmt.Errorf("error reading 'scopes': %w", err)
		}
		delete(object, "scopes")
	}

	if raw, found := object["tokenUrl"]; found {
		err = json.Unmarshal(raw, &a.TokenUrl)
		if err != nil {
			return fmt.Errorf("error reading 'tokenUrl': %w", err)
		}
		delete(object, "tokenUrl")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OAuthFlowClientCredentials to handle AdditionalProperties
func (a OAuthFlowClientCredentials) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.RefreshUrl != nil {
		object["refreshUrl"], err = json.Marshal(a.RefreshUrl)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'refreshUrl': %w", err)
		}
	}

	if a.Scopes != nil {
		object["scopes"], err = json.Marshal(a.Scopes)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'scopes': %w", err)
		}
	}

	object["tokenUrl"], err = json.Marshal(a.TokenUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'tokenUrl': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OAuthFlowImplicit. Returns the specified
// element and whether it was found
func (a OAuthFlowImplicit) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OAuthFlowImplicit
func (a *OAuthFlowImplicit) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OAuthFlowImplicit to handle AdditionalProperties
func (a *OAuthFlowImplicit) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["authorizationUrl"]; found {
		err = json.Unmarshal(raw, &a.AuthorizationUrl)
		if err != nil {
			return fmt.Errorf("error reading 'authorizationUrl': %w", err)
		}
		delete(object, "authorizationUrl")
	}

	if raw, found := object["refreshUrl"]; found {
		err = json.Unmarshal(raw, &a.RefreshUrl)
		if err != nil {
			return fmt.Errorf("error reading 'refreshUrl': %w", err)
		}
		delete(object, "refreshUrl")
	}

	if raw, found := object["scopes"]; found {
		err = json.Unmarshal(raw, &a.Scopes)
		if err != nil {
			return fmt.Errorf("error reading 'scopes': %w", err)
		}
		delete(object, "scopes")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OAuthFlowImplicit to handle AdditionalProperties
func (a OAuthFlowImplicit) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["authorizationUrl"], err = json.Marshal(a.AuthorizationUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'authorizationUrl': %w", err)
	}

	if a.RefreshUrl != nil {
		object["refreshUrl"], err = json.Marshal(a.RefreshUrl)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'refreshUrl': %w", err)
		}
	}

	if a.Scopes != nil {
		object["scopes"], err = json.Marshal(a.Scopes)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'scopes': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OAuthFlowPassword. Returns the specified
// element and whether it was found
func (a OAuthFlowPassword) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OAuthFlowPassword
func (a *OAuthFlowPassword) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OAuthFlowPassword to handle AdditionalProperties
func (a *OAuthFlowPassword) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["refreshUrl"]; found {
		err = json.Unmarshal(raw, &a.RefreshUrl)
		if err != nil {
			return fmt.Errorf("error reading 'refreshUrl': %w", err)
		}
		delete(object, "refreshUrl")
	}

	if raw, found := object["scopes"]; found {
		err = json.Unmarshal(raw, &a.Scopes)
		if err != nil {
			return fmt.Errorf("error reading 'scopes': %w", err)
		}
		delete(object, "scopes")
	}

	if raw, found := object["tokenUrl"]; found {
		err = json.Unmarshal(raw, &a.TokenUrl)
		if err != nil {
			return fmt.Errorf("error reading 'tokenUrl': %w", err)
		}
		delete(object, "tokenUrl")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OAuthFlowPassword to handle AdditionalProperties
func (a OAuthFlowPassword) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.RefreshUrl != nil {
		object["refreshUrl"], err = json.Marshal(a.RefreshUrl)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'refreshUrl': %w", err)
		}
	}

	if a.Scopes != nil {
		object["scopes"], err = json.Marshal(a.Scopes)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'scopes': %w", err)
		}
	}

	object["tokenUrl"], err = json.Marshal(a.TokenUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'tokenUrl': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OAuthFlows. Returns the specified
// element and whether it was found
func (a OAuthFlows) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OAuthFlows
func (a *OAuthFlows) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OAuthFlows to handle AdditionalProperties
func (a *OAuthFlows) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["authorizationCode"]; found {
		err = json.Unmarshal(raw, &a.AuthorizationCode)
		if err != nil {
			return fmt.Errorf("error reading 'authorizationCode': %w", err)
		}
		delete(object, "authorizationCode")
	}

	if raw, found := object["clientCredentials"]; found {
		err = json.Unmarshal(raw, &a.ClientCredentials)
		if err != nil {
			return fmt.Errorf("error reading 'clientCredentials': %w", err)
		}
		delete(object, "clientCredentials")
	}

	if raw, found := object["implicit"]; found {
		err = json.Unmarshal(raw, &a.Implicit)
		if err != nil {
			return fmt.Errorf("error reading 'implicit': %w", err)
		}
		delete(object, "implicit")
	}

	if raw, found := object["password"]; found {
		err = json.Unmarshal(raw, &a.Password)
		if err != nil {
			return fmt.Errorf("error reading 'password': %w", err)
		}
		delete(object, "password")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OAuthFlows to handle AdditionalProperties
func (a OAuthFlows) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.AuthorizationCode != nil {
		object["authorizationCode"], err = json.Marshal(a.AuthorizationCode)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'authorizationCode': %w", err)
		}
	}

	if a.ClientCredentials != nil {
		object["clientCredentials"], err = json.Marshal(a.ClientCredentials)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'clientCredentials': %w", err)
		}
	}

	if a.Implicit != nil {
		object["implicit"], err = json.Marshal(a.Implicit)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'implicit': %w", err)
		}
	}

	if a.Password != nil {
		object["password"], err = json.Marshal(a.Password)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'password': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OpenIdConnect. Returns the specified
// element and whether it was found
func (a OpenIdConnect) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenIdConnect
func (a *OpenIdConnect) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenIdConnect to handle AdditionalProperties
func (a *OpenIdConnect) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["openIdConnectUrl"]; found {
		err = json.Unmarshal(raw, &a.OpenIdConnectUrl)
		if err != nil {
			return fmt.Errorf("error reading 'openIdConnectUrl': %w", err)
		}
		delete(object, "openIdConnectUrl")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OpenIdConnect to handle AdditionalProperties
func (a OpenIdConnect) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'description': %w", err)
		}
	}

	object["openIdConnectUrl"], err = json.Marshal(a.OpenIdConnectUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'openIdConnectUrl': %w", err)
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for OpenIdConnectWithConfig. Returns the specified
// element and whether it was found
func (a OpenIdConnectWithConfig) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for OpenIdConnectWithConfig
func (a *OpenIdConnectWithConfig) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for OpenIdConnectWithConfig to handle AdditionalProperties
func (a *OpenIdConnectWithConfig) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["authorization_endpoint"]; found {
		err = json.Unmarshal(raw, &a.AuthorizationEndpoint)
		if err != nil {
			return fmt.Errorf("error reading 'authorization_endpoint': %w", err)
		}
		delete(object, "authorization_endpoint")
	}

	if raw, found := object["description"]; found {
		err = json.Unmarshal(raw, &a.Description)
		if err != nil {
			return fmt.Errorf("error reading 'description': %w", err)
		}
		delete(object, "description")
	}

	if raw, found := object["grant_types_supported"]; found {
		err = json.Unmarshal(raw, &a.GrantTypesSupported)
		if err != nil {
			return fmt.Errorf("error reading 'grant_types_supported': %w", err)
		}
		delete(object, "grant_types_supported")
	}

	if raw, found := object["revocation_endpoint"]; found {
		err = json.Unmarshal(raw, &a.RevocationEndpoint)
		if err != nil {
			return fmt.Errorf("error reading 'revocation_endpoint': %w", err)
		}
		delete(object, "revocation_endpoint")
	}

	if raw, found := object["scopes"]; found {
		err = json.Unmarshal(raw, &a.Scopes)
		if err != nil {
			return fmt.Errorf("error reading 'scopes': %w", err)
		}
		delete(object, "scopes")
	}

	if raw, found := object["token_endpoint"]; found {
		err = json.Unmarshal(raw, &a.TokenEndpoint)
		if err != nil {
			return fmt.Errorf("error reading 'token_endpoint': %w", err)
		}
		delete(object, "token_endpoint")
	}

	if raw, found := object["token_endpoint_auth_methods_supported"]; found {
		err = json.Unmarshal(raw, &a.TokenEndpointAuthMethodsSupported)
		if err != nil {
			return fmt.Errorf("error reading 'token_endpoint_auth_methods_supported': %w", err)
		}
		delete(object, "token_endpoint_auth_methods_supported")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if raw, found := object["userinfo_endpoint"]; found {
		err = json.Unmarshal(raw, &a.UserinfoEndpoint)
		if err != nil {
			return fmt.Errorf("error reading 'userinfo_endpoint': %w", err)
		}
		delete(object, "userinfo_endpoint")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for OpenIdConnectWithConfig to handle AdditionalProperties
func (a OpenIdConnectWithConfig) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["authorization_endpoint"], err = json.Marshal(a.AuthorizationEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'authorization_endpoint': %w", err)
	}

	if a.Description != nil {
		object["description"], err = json.Marshal(a.Description)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'description': %w", err)
		}
	}

	if a.GrantTypesSupported != nil {
		object["grant_types_supported"], err = json.Marshal(a.GrantTypesSupported)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'grant_types_supported': %w", err)
		}
	}

	if a.RevocationEndpoint != nil {
		object["revocation_endpoint"], err = json.Marshal(a.RevocationEndpoint)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'revocation_endpoint': %w", err)
		}
	}

	if a.Scopes != nil {
		object["scopes"], err = json.Marshal(a.Scopes)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'scopes': %w", err)
		}
	}

	object["token_endpoint"], err = json.Marshal(a.TokenEndpoint)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'token_endpoint': %w", err)
	}

	if a.TokenEndpointAuthMethodsSupported != nil {
		object["token_endpoint_auth_methods_supported"], err = json.Marshal(a.TokenEndpointAuthMethodsSupported)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'token_endpoint_auth_methods_supported': %w", err)
		}
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	if a.UserinfoEndpoint != nil {
		object["userinfo_endpoint"], err = json.Marshal(a.UserinfoEndpoint)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'userinfo_endpoint': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ServiceAccount. Returns the specified
// element and whether it was found
func (a ServiceAccount) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ServiceAccount
func (a *ServiceAccount) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ServiceAccount to handle AdditionalProperties
func (a *ServiceAccount) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["scopes"]; found {
		err = json.Unmarshal(raw, &a.Scopes)
		if err != nil {
			return fmt.Errorf("error reading 'scopes': %w", err)
		}
		delete(object, "scopes")
	}

	if raw, found := object["service_account_credential"]; found {
		err = json.Unmarshal(raw, &a.ServiceAccountCredential)
		if err != nil {
			return fmt.Errorf("error reading 'service_account_credential': %w", err)
		}
		delete(object, "service_account_credential")
	}

	if raw, found := object["use_default_credential"]; found {
		err = json.Unmarshal(raw, &a.UseDefaultCredential)
		if err != nil {
			return fmt.Errorf("error reading 'use_default_credential': %w", err)
		}
		delete(object, "use_default_credential")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ServiceAccount to handle AdditionalProperties
func (a ServiceAccount) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["scopes"], err = json.Marshal(a.Scopes)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'scopes': %w", err)
	}

	if a.ServiceAccountCredential != nil {
		object["service_account_credential"], err = json.Marshal(a.ServiceAccountCredential)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'service_account_credential': %w", err)
		}
	}

	if a.UseDefaultCredential != nil {
		object["use_default_credential"], err = json.Marshal(a.UseDefaultCredential)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'use_default_credential': %w", err)
		}
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// Getter for additional properties for ServiceAccountCredential. Returns the specified
// element and whether it was found
func (a ServiceAccountCredential) Get(fieldName string) (value interface{}, found bool) {
	if a.AdditionalProperties != nil {
		value, found = a.AdditionalProperties[fieldName]
	}
	return
}

// Setter for additional properties for ServiceAccountCredential
func (a *ServiceAccountCredential) Set(fieldName string, value interface{}) {
	if a.AdditionalProperties == nil {
		a.AdditionalProperties = make(map[string]interface{})
	}
	a.AdditionalProperties[fieldName] = value
}

// Override default JSON handling for ServiceAccountCredential to handle AdditionalProperties
func (a *ServiceAccountCredential) UnmarshalJSON(b []byte) error {
	object := make(map[string]json.RawMessage)
	err := json.Unmarshal(b, &object)
	if err != nil {
		return err
	}

	if raw, found := object["auth_provider_x509_cert_url"]; found {
		err = json.Unmarshal(raw, &a.AuthProviderX509CertUrl)
		if err != nil {
			return fmt.Errorf("error reading 'auth_provider_x509_cert_url': %w", err)
		}
		delete(object, "auth_provider_x509_cert_url")
	}

	if raw, found := object["auth_uri"]; found {
		err = json.Unmarshal(raw, &a.AuthUri)
		if err != nil {
			return fmt.Errorf("error reading 'auth_uri': %w", err)
		}
		delete(object, "auth_uri")
	}

	if raw, found := object["client_email"]; found {
		err = json.Unmarshal(raw, &a.ClientEmail)
		if err != nil {
			return fmt.Errorf("error reading 'client_email': %w", err)
		}
		delete(object, "client_email")
	}

	if raw, found := object["client_id"]; found {
		err = json.Unmarshal(raw, &a.ClientId)
		if err != nil {
			return fmt.Errorf("error reading 'client_id': %w", err)
		}
		delete(object, "client_id")
	}

	if raw, found := object["client_x509_cert_url"]; found {
		err = json.Unmarshal(raw, &a.ClientX509CertUrl)
		if err != nil {
			return fmt.Errorf("error reading 'client_x509_cert_url': %w", err)
		}
		delete(object, "client_x509_cert_url")
	}

	if raw, found := object["private_key"]; found {
		err = json.Unmarshal(raw, &a.PrivateKey)
		if err != nil {
			return fmt.Errorf("error reading 'private_key': %w", err)
		}
		delete(object, "private_key")
	}

	if raw, found := object["private_key_id"]; found {
		err = json.Unmarshal(raw, &a.PrivateKeyId)
		if err != nil {
			return fmt.Errorf("error reading 'private_key_id': %w", err)
		}
		delete(object, "private_key_id")
	}

	if raw, found := object["project_id"]; found {
		err = json.Unmarshal(raw, &a.ProjectId)
		if err != nil {
			return fmt.Errorf("error reading 'project_id': %w", err)
		}
		delete(object, "project_id")
	}

	if raw, found := object["token_uri"]; found {
		err = json.Unmarshal(raw, &a.TokenUri)
		if err != nil {
			return fmt.Errorf("error reading 'token_uri': %w", err)
		}
		delete(object, "token_uri")
	}

	if raw, found := object["type"]; found {
		err = json.Unmarshal(raw, &a.Type)
		if err != nil {
			return fmt.Errorf("error reading 'type': %w", err)
		}
		delete(object, "type")
	}

	if raw, found := object["universe_domain"]; found {
		err = json.Unmarshal(raw, &a.UniverseDomain)
		if err != nil {
			return fmt.Errorf("error reading 'universe_domain': %w", err)
		}
		delete(object, "universe_domain")
	}

	if len(object) != 0 {
		a.AdditionalProperties = make(map[string]interface{})
		for fieldName, fieldBuf := range object {
			var fieldVal interface{}
			err := json.Unmarshal(fieldBuf, &fieldVal)
			if err != nil {
				return fmt.Errorf("error unmarshaling field %s: %w", fieldName, err)
			}
			a.AdditionalProperties[fieldName] = fieldVal
		}
	}
	return nil
}

// Override default JSON handling for ServiceAccountCredential to handle AdditionalProperties
func (a ServiceAccountCredential) MarshalJSON() ([]byte, error) {
	var err error
	object := make(map[string]json.RawMessage)

	object["auth_provider_x509_cert_url"], err = json.Marshal(a.AuthProviderX509CertUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'auth_provider_x509_cert_url': %w", err)
	}

	object["auth_uri"], err = json.Marshal(a.AuthUri)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'auth_uri': %w", err)
	}

	object["client_email"], err = json.Marshal(a.ClientEmail)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'client_email': %w", err)
	}

	object["client_id"], err = json.Marshal(a.ClientId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'client_id': %w", err)
	}

	object["client_x509_cert_url"], err = json.Marshal(a.ClientX509CertUrl)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'client_x509_cert_url': %w", err)
	}

	object["private_key"], err = json.Marshal(a.PrivateKey)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'private_key': %w", err)
	}

	object["private_key_id"], err = json.Marshal(a.PrivateKeyId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'private_key_id': %w", err)
	}

	object["project_id"], err = json.Marshal(a.ProjectId)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'project_id': %w", err)
	}

	object["token_uri"], err = json.Marshal(a.TokenUri)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'token_uri': %w", err)
	}

	if a.Type != nil {
		object["type"], err = json.Marshal(a.Type)
		if err != nil {
			return nil, fmt.Errorf("error marshaling 'type': %w", err)
		}
	}

	object["universe_domain"], err = json.Marshal(a.UniverseDomain)
	if err != nil {
		return nil, fmt.Errorf("error marshaling 'universe_domain': %w", err)
	}

	for fieldName, field := range a.AdditionalProperties {
		object[fieldName], err = json.Marshal(field)
		if err != nil {
			return nil, fmt.Errorf("error marshaling '%s': %w", fieldName, err)
		}
	}
	return json.Marshal(object)
}

// AsAPIKeyDescription0 returns the union data inside the APIKey_Description as a APIKeyDescription0
func (t APIKey_Description) AsAPIKeyDescription0() (APIKeyDescription0, error) {
	var body APIKeyDescription0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAPIKeyDescription0 overwrites any union data inside the APIKey_Description as the provided APIKeyDescription0
func (t *APIKey_Description) FromAPIKeyDescription0(v APIKeyDescription0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAPIKeyDescription0 performs a merge with any union data inside the APIKey_Description, using the provided APIKeyDescription0
func (t *APIKey_Description) MergeAPIKeyDescription0(v APIKeyDescription0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAPIKeyDescription1 returns the union data inside the APIKey_Description as a APIKeyDescription1
func (t APIKey_Description) AsAPIKeyDescription1() (APIKeyDescription1, error) {
	var body APIKeyDescription1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAPIKeyDescription1 overwrites any union data inside the APIKey_Description as the provided APIKeyDescription1
func (t *APIKey_Description) FromAPIKeyDescription1(v APIKeyDescription1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAPIKeyDescription1 performs a merge with any union data inside the APIKey_Description, using the provided APIKeyDescription1
func (t *APIKey_Description) MergeAPIKeyDescription1(v APIKeyDescription1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t APIKey_Description) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *APIKey_Description) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsAPIKey returns the union data inside the AuthConfig_AuthScheme as a APIKey
func (t AuthConfig_AuthScheme) AsAPIKey() (APIKey, error) {
	var body APIKey
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAPIKey overwrites any union data inside the AuthConfig_AuthScheme as the provided APIKey
func (t *AuthConfig_AuthScheme) FromAPIKey(v APIKey) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAPIKey performs a merge with any union data inside the AuthConfig_AuthScheme, using the provided APIKey
func (t *AuthConfig_AuthScheme) MergeAPIKey(v APIKey) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHTTPBase returns the union data inside the AuthConfig_AuthScheme as a HTTPBase
func (t AuthConfig_AuthScheme) AsHTTPBase() (HTTPBase, error) {
	var body HTTPBase
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBase overwrites any union data inside the AuthConfig_AuthScheme as the provided HTTPBase
func (t *AuthConfig_AuthScheme) FromHTTPBase(v HTTPBase) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBase performs a merge with any union data inside the AuthConfig_AuthScheme, using the provided HTTPBase
func (t *AuthConfig_AuthScheme) MergeHTTPBase(v HTTPBase) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2 returns the union data inside the AuthConfig_AuthScheme as a OAuth2
func (t AuthConfig_AuthScheme) AsOAuth2() (OAuth2, error) {
	var body OAuth2
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2 overwrites any union data inside the AuthConfig_AuthScheme as the provided OAuth2
func (t *AuthConfig_AuthScheme) FromOAuth2(v OAuth2) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2 performs a merge with any union data inside the AuthConfig_AuthScheme, using the provided OAuth2
func (t *AuthConfig_AuthScheme) MergeOAuth2(v OAuth2) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnect returns the union data inside the AuthConfig_AuthScheme as a OpenIdConnect
func (t AuthConfig_AuthScheme) AsOpenIdConnect() (OpenIdConnect, error) {
	var body OpenIdConnect
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnect overwrites any union data inside the AuthConfig_AuthScheme as the provided OpenIdConnect
func (t *AuthConfig_AuthScheme) FromOpenIdConnect(v OpenIdConnect) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnect performs a merge with any union data inside the AuthConfig_AuthScheme, using the provided OpenIdConnect
func (t *AuthConfig_AuthScheme) MergeOpenIdConnect(v OpenIdConnect) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHTTPBearer returns the union data inside the AuthConfig_AuthScheme as a HTTPBearer
func (t AuthConfig_AuthScheme) AsHTTPBearer() (HTTPBearer, error) {
	var body HTTPBearer
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBearer overwrites any union data inside the AuthConfig_AuthScheme as the provided HTTPBearer
func (t *AuthConfig_AuthScheme) FromHTTPBearer(v HTTPBearer) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBearer performs a merge with any union data inside the AuthConfig_AuthScheme, using the provided HTTPBearer
func (t *AuthConfig_AuthScheme) MergeHTTPBearer(v HTTPBearer) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectWithConfig returns the union data inside the AuthConfig_AuthScheme as a OpenIdConnectWithConfig
func (t AuthConfig_AuthScheme) AsOpenIdConnectWithConfig() (OpenIdConnectWithConfig, error) {
	var body OpenIdConnectWithConfig
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfig overwrites any union data inside the AuthConfig_AuthScheme as the provided OpenIdConnectWithConfig
func (t *AuthConfig_AuthScheme) FromOpenIdConnectWithConfig(v OpenIdConnectWithConfig) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfig performs a merge with any union data inside the AuthConfig_AuthScheme, using the provided OpenIdConnectWithConfig
func (t *AuthConfig_AuthScheme) MergeOpenIdConnectWithConfig(v OpenIdConnectWithConfig) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AuthConfig_AuthScheme) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AuthConfig_AuthScheme) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsAuthCredentialApiKey0 returns the union data inside the AuthCredential_ApiKey as a AuthCredentialApiKey0
func (t AuthCredential_ApiKey) AsAuthCredentialApiKey0() (AuthCredentialApiKey0, error) {
	var body AuthCredentialApiKey0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthCredentialApiKey0 overwrites any union data inside the AuthCredential_ApiKey as the provided AuthCredentialApiKey0
func (t *AuthCredential_ApiKey) FromAuthCredentialApiKey0(v AuthCredentialApiKey0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthCredentialApiKey0 performs a merge with any union data inside the AuthCredential_ApiKey, using the provided AuthCredentialApiKey0
func (t *AuthCredential_ApiKey) MergeAuthCredentialApiKey0(v AuthCredentialApiKey0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAuthCredentialApiKey1 returns the union data inside the AuthCredential_ApiKey as a AuthCredentialApiKey1
func (t AuthCredential_ApiKey) AsAuthCredentialApiKey1() (AuthCredentialApiKey1, error) {
	var body AuthCredentialApiKey1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthCredentialApiKey1 overwrites any union data inside the AuthCredential_ApiKey as the provided AuthCredentialApiKey1
func (t *AuthCredential_ApiKey) FromAuthCredentialApiKey1(v AuthCredentialApiKey1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthCredentialApiKey1 performs a merge with any union data inside the AuthCredential_ApiKey, using the provided AuthCredentialApiKey1
func (t *AuthCredential_ApiKey) MergeAuthCredentialApiKey1(v AuthCredentialApiKey1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AuthCredential_ApiKey) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AuthCredential_ApiKey) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsHttpAuth returns the union data inside the AuthCredential_Http as a HttpAuth
func (t AuthCredential_Http) AsHttpAuth() (HttpAuth, error) {
	var body HttpAuth
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHttpAuth overwrites any union data inside the AuthCredential_Http as the provided HttpAuth
func (t *AuthCredential_Http) FromHttpAuth(v HttpAuth) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHttpAuth performs a merge with any union data inside the AuthCredential_Http, using the provided HttpAuth
func (t *AuthCredential_Http) MergeHttpAuth(v HttpAuth) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAuthCredentialHttp1 returns the union data inside the AuthCredential_Http as a AuthCredentialHttp1
func (t AuthCredential_Http) AsAuthCredentialHttp1() (AuthCredentialHttp1, error) {
	var body AuthCredentialHttp1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthCredentialHttp1 overwrites any union data inside the AuthCredential_Http as the provided AuthCredentialHttp1
func (t *AuthCredential_Http) FromAuthCredentialHttp1(v AuthCredentialHttp1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthCredentialHttp1 performs a merge with any union data inside the AuthCredential_Http, using the provided AuthCredentialHttp1
func (t *AuthCredential_Http) MergeAuthCredentialHttp1(v AuthCredentialHttp1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AuthCredential_Http) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AuthCredential_Http) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2Auth returns the union data inside the AuthCredential_Oauth2 as a OAuth2Auth
func (t AuthCredential_Oauth2) AsOAuth2Auth() (OAuth2Auth, error) {
	var body OAuth2Auth
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2Auth overwrites any union data inside the AuthCredential_Oauth2 as the provided OAuth2Auth
func (t *AuthCredential_Oauth2) FromOAuth2Auth(v OAuth2Auth) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2Auth performs a merge with any union data inside the AuthCredential_Oauth2, using the provided OAuth2Auth
func (t *AuthCredential_Oauth2) MergeOAuth2Auth(v OAuth2Auth) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAuthCredentialOauth21 returns the union data inside the AuthCredential_Oauth2 as a AuthCredentialOauth21
func (t AuthCredential_Oauth2) AsAuthCredentialOauth21() (AuthCredentialOauth21, error) {
	var body AuthCredentialOauth21
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthCredentialOauth21 overwrites any union data inside the AuthCredential_Oauth2 as the provided AuthCredentialOauth21
func (t *AuthCredential_Oauth2) FromAuthCredentialOauth21(v AuthCredentialOauth21) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthCredentialOauth21 performs a merge with any union data inside the AuthCredential_Oauth2, using the provided AuthCredentialOauth21
func (t *AuthCredential_Oauth2) MergeAuthCredentialOauth21(v AuthCredentialOauth21) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AuthCredential_Oauth2) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AuthCredential_Oauth2) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsAuthCredentialResourceRef0 returns the union data inside the AuthCredential_ResourceRef as a AuthCredentialResourceRef0
func (t AuthCredential_ResourceRef) AsAuthCredentialResourceRef0() (AuthCredentialResourceRef0, error) {
	var body AuthCredentialResourceRef0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthCredentialResourceRef0 overwrites any union data inside the AuthCredential_ResourceRef as the provided AuthCredentialResourceRef0
func (t *AuthCredential_ResourceRef) FromAuthCredentialResourceRef0(v AuthCredentialResourceRef0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthCredentialResourceRef0 performs a merge with any union data inside the AuthCredential_ResourceRef, using the provided AuthCredentialResourceRef0
func (t *AuthCredential_ResourceRef) MergeAuthCredentialResourceRef0(v AuthCredentialResourceRef0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAuthCredentialResourceRef1 returns the union data inside the AuthCredential_ResourceRef as a AuthCredentialResourceRef1
func (t AuthCredential_ResourceRef) AsAuthCredentialResourceRef1() (AuthCredentialResourceRef1, error) {
	var body AuthCredentialResourceRef1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthCredentialResourceRef1 overwrites any union data inside the AuthCredential_ResourceRef as the provided AuthCredentialResourceRef1
func (t *AuthCredential_ResourceRef) FromAuthCredentialResourceRef1(v AuthCredentialResourceRef1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthCredentialResourceRef1 performs a merge with any union data inside the AuthCredential_ResourceRef, using the provided AuthCredentialResourceRef1
func (t *AuthCredential_ResourceRef) MergeAuthCredentialResourceRef1(v AuthCredentialResourceRef1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AuthCredential_ResourceRef) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AuthCredential_ResourceRef) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsServiceAccount returns the union data inside the AuthCredential_ServiceAccount as a ServiceAccount
func (t AuthCredential_ServiceAccount) AsServiceAccount() (ServiceAccount, error) {
	var body ServiceAccount
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromServiceAccount overwrites any union data inside the AuthCredential_ServiceAccount as the provided ServiceAccount
func (t *AuthCredential_ServiceAccount) FromServiceAccount(v ServiceAccount) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeServiceAccount performs a merge with any union data inside the AuthCredential_ServiceAccount, using the provided ServiceAccount
func (t *AuthCredential_ServiceAccount) MergeServiceAccount(v ServiceAccount) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsAuthCredentialServiceAccount1 returns the union data inside the AuthCredential_ServiceAccount as a AuthCredentialServiceAccount1
func (t AuthCredential_ServiceAccount) AsAuthCredentialServiceAccount1() (AuthCredentialServiceAccount1, error) {
	var body AuthCredentialServiceAccount1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromAuthCredentialServiceAccount1 overwrites any union data inside the AuthCredential_ServiceAccount as the provided AuthCredentialServiceAccount1
func (t *AuthCredential_ServiceAccount) FromAuthCredentialServiceAccount1(v AuthCredentialServiceAccount1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeAuthCredentialServiceAccount1 performs a merge with any union data inside the AuthCredential_ServiceAccount, using the provided AuthCredentialServiceAccount1
func (t *AuthCredential_ServiceAccount) MergeAuthCredentialServiceAccount1(v AuthCredentialServiceAccount1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t AuthCredential_ServiceAccount) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *AuthCredential_ServiceAccount) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBlobData0 returns the union data inside the Blob_Data as a BlobData0
func (t Blob_Data) AsBlobData0() (BlobData0, error) {
	var body BlobData0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBlobData0 overwrites any union data inside the Blob_Data as the provided BlobData0
func (t *Blob_Data) FromBlobData0(v BlobData0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBlobData0 performs a merge with any union data inside the Blob_Data, using the provided BlobData0
func (t *Blob_Data) MergeBlobData0(v BlobData0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsBlobData1 returns the union data inside the Blob_Data as a BlobData1
func (t Blob_Data) AsBlobData1() (BlobData1, error) {
	var body BlobData1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBlobData1 overwrites any union data inside the Blob_Data as the provided BlobData1
func (t *Blob_Data) FromBlobData1(v BlobData1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBlobData1 performs a merge with any union data inside the Blob_Data, using the provided BlobData1
func (t *Blob_Data) MergeBlobData1(v BlobData1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Blob_Data) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Blob_Data) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBlobMimeType0 returns the union data inside the Blob_MimeType as a BlobMimeType0
func (t Blob_MimeType) AsBlobMimeType0() (BlobMimeType0, error) {
	var body BlobMimeType0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBlobMimeType0 overwrites any union data inside the Blob_MimeType as the provided BlobMimeType0
func (t *Blob_MimeType) FromBlobMimeType0(v BlobMimeType0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBlobMimeType0 performs a merge with any union data inside the Blob_MimeType, using the provided BlobMimeType0
func (t *Blob_MimeType) MergeBlobMimeType0(v BlobMimeType0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsBlobMimeType1 returns the union data inside the Blob_MimeType as a BlobMimeType1
func (t Blob_MimeType) AsBlobMimeType1() (BlobMimeType1, error) {
	var body BlobMimeType1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBlobMimeType1 overwrites any union data inside the Blob_MimeType as the provided BlobMimeType1
func (t *Blob_MimeType) FromBlobMimeType1(v BlobMimeType1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBlobMimeType1 performs a merge with any union data inside the Blob_MimeType, using the provided BlobMimeType1
func (t *Blob_MimeType) MergeBlobMimeType1(v BlobMimeType1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Blob_MimeType) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Blob_MimeType) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOutcome returns the union data inside the CodeExecutionResult_Outcome as a Outcome
func (t CodeExecutionResult_Outcome) AsOutcome() (Outcome, error) {
	var body Outcome
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOutcome overwrites any union data inside the CodeExecutionResult_Outcome as the provided Outcome
func (t *CodeExecutionResult_Outcome) FromOutcome(v Outcome) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOutcome performs a merge with any union data inside the CodeExecutionResult_Outcome, using the provided Outcome
func (t *CodeExecutionResult_Outcome) MergeOutcome(v Outcome) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsCodeExecutionResultOutcome1 returns the union data inside the CodeExecutionResult_Outcome as a CodeExecutionResultOutcome1
func (t CodeExecutionResult_Outcome) AsCodeExecutionResultOutcome1() (CodeExecutionResultOutcome1, error) {
	var body CodeExecutionResultOutcome1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCodeExecutionResultOutcome1 overwrites any union data inside the CodeExecutionResult_Outcome as the provided CodeExecutionResultOutcome1
func (t *CodeExecutionResult_Outcome) FromCodeExecutionResultOutcome1(v CodeExecutionResultOutcome1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCodeExecutionResultOutcome1 performs a merge with any union data inside the CodeExecutionResult_Outcome, using the provided CodeExecutionResultOutcome1
func (t *CodeExecutionResult_Outcome) MergeCodeExecutionResultOutcome1(v CodeExecutionResultOutcome1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t CodeExecutionResult_Outcome) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CodeExecutionResult_Outcome) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCodeExecutionResultOutput0 returns the union data inside the CodeExecutionResult_Output as a CodeExecutionResultOutput0
func (t CodeExecutionResult_Output) AsCodeExecutionResultOutput0() (CodeExecutionResultOutput0, error) {
	var body CodeExecutionResultOutput0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCodeExecutionResultOutput0 overwrites any union data inside the CodeExecutionResult_Output as the provided CodeExecutionResultOutput0
func (t *CodeExecutionResult_Output) FromCodeExecutionResultOutput0(v CodeExecutionResultOutput0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCodeExecutionResultOutput0 performs a merge with any union data inside the CodeExecutionResult_Output, using the provided CodeExecutionResultOutput0
func (t *CodeExecutionResult_Output) MergeCodeExecutionResultOutput0(v CodeExecutionResultOutput0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsCodeExecutionResultOutput1 returns the union data inside the CodeExecutionResult_Output as a CodeExecutionResultOutput1
func (t CodeExecutionResult_Output) AsCodeExecutionResultOutput1() (CodeExecutionResultOutput1, error) {
	var body CodeExecutionResultOutput1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCodeExecutionResultOutput1 overwrites any union data inside the CodeExecutionResult_Output as the provided CodeExecutionResultOutput1
func (t *CodeExecutionResult_Output) FromCodeExecutionResultOutput1(v CodeExecutionResultOutput1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCodeExecutionResultOutput1 performs a merge with any union data inside the CodeExecutionResult_Output, using the provided CodeExecutionResultOutput1
func (t *CodeExecutionResult_Output) MergeCodeExecutionResultOutput1(v CodeExecutionResultOutput1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t CodeExecutionResult_Output) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *CodeExecutionResult_Output) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentInputParts0 returns the union data inside the ContentInput_Parts as a ContentInputParts0
func (t ContentInput_Parts) AsContentInputParts0() (ContentInputParts0, error) {
	var body ContentInputParts0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentInputParts0 overwrites any union data inside the ContentInput_Parts as the provided ContentInputParts0
func (t *ContentInput_Parts) FromContentInputParts0(v ContentInputParts0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentInputParts0 performs a merge with any union data inside the ContentInput_Parts, using the provided ContentInputParts0
func (t *ContentInput_Parts) MergeContentInputParts0(v ContentInputParts0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentInputParts1 returns the union data inside the ContentInput_Parts as a ContentInputParts1
func (t ContentInput_Parts) AsContentInputParts1() (ContentInputParts1, error) {
	var body ContentInputParts1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentInputParts1 overwrites any union data inside the ContentInput_Parts as the provided ContentInputParts1
func (t *ContentInput_Parts) FromContentInputParts1(v ContentInputParts1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentInputParts1 performs a merge with any union data inside the ContentInput_Parts, using the provided ContentInputParts1
func (t *ContentInput_Parts) MergeContentInputParts1(v ContentInputParts1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ContentInput_Parts) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ContentInput_Parts) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentInputRole0 returns the union data inside the ContentInput_Role as a ContentInputRole0
func (t ContentInput_Role) AsContentInputRole0() (ContentInputRole0, error) {
	var body ContentInputRole0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentInputRole0 overwrites any union data inside the ContentInput_Role as the provided ContentInputRole0
func (t *ContentInput_Role) FromContentInputRole0(v ContentInputRole0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentInputRole0 performs a merge with any union data inside the ContentInput_Role, using the provided ContentInputRole0
func (t *ContentInput_Role) MergeContentInputRole0(v ContentInputRole0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentInputRole1 returns the union data inside the ContentInput_Role as a ContentInputRole1
func (t ContentInput_Role) AsContentInputRole1() (ContentInputRole1, error) {
	var body ContentInputRole1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentInputRole1 overwrites any union data inside the ContentInput_Role as the provided ContentInputRole1
func (t *ContentInput_Role) FromContentInputRole1(v ContentInputRole1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentInputRole1 performs a merge with any union data inside the ContentInput_Role, using the provided ContentInputRole1
func (t *ContentInput_Role) MergeContentInputRole1(v ContentInputRole1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ContentInput_Role) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ContentInput_Role) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentOutputParts0 returns the union data inside the ContentOutput_Parts as a ContentOutputParts0
func (t ContentOutput_Parts) AsContentOutputParts0() (ContentOutputParts0, error) {
	var body ContentOutputParts0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentOutputParts0 overwrites any union data inside the ContentOutput_Parts as the provided ContentOutputParts0
func (t *ContentOutput_Parts) FromContentOutputParts0(v ContentOutputParts0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentOutputParts0 performs a merge with any union data inside the ContentOutput_Parts, using the provided ContentOutputParts0
func (t *ContentOutput_Parts) MergeContentOutputParts0(v ContentOutputParts0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentOutputParts1 returns the union data inside the ContentOutput_Parts as a ContentOutputParts1
func (t ContentOutput_Parts) AsContentOutputParts1() (ContentOutputParts1, error) {
	var body ContentOutputParts1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentOutputParts1 overwrites any union data inside the ContentOutput_Parts as the provided ContentOutputParts1
func (t *ContentOutput_Parts) FromContentOutputParts1(v ContentOutputParts1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentOutputParts1 performs a merge with any union data inside the ContentOutput_Parts, using the provided ContentOutputParts1
func (t *ContentOutput_Parts) MergeContentOutputParts1(v ContentOutputParts1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ContentOutput_Parts) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ContentOutput_Parts) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentOutputRole0 returns the union data inside the ContentOutput_Role as a ContentOutputRole0
func (t ContentOutput_Role) AsContentOutputRole0() (ContentOutputRole0, error) {
	var body ContentOutputRole0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentOutputRole0 overwrites any union data inside the ContentOutput_Role as the provided ContentOutputRole0
func (t *ContentOutput_Role) FromContentOutputRole0(v ContentOutputRole0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentOutputRole0 performs a merge with any union data inside the ContentOutput_Role, using the provided ContentOutputRole0
func (t *ContentOutput_Role) MergeContentOutputRole0(v ContentOutputRole0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsContentOutputRole1 returns the union data inside the ContentOutput_Role as a ContentOutputRole1
func (t ContentOutput_Role) AsContentOutputRole1() (ContentOutputRole1, error) {
	var body ContentOutputRole1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentOutputRole1 overwrites any union data inside the ContentOutput_Role as the provided ContentOutputRole1
func (t *ContentOutput_Role) FromContentOutputRole1(v ContentOutputRole1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentOutputRole1 performs a merge with any union data inside the ContentOutput_Role, using the provided ContentOutputRole1
func (t *ContentOutput_Role) MergeContentOutputRole1(v ContentOutputRole1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ContentOutput_Role) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ContentOutput_Role) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventBranch0 returns the union data inside the Event_Branch as a EventBranch0
func (t Event_Branch) AsEventBranch0() (EventBranch0, error) {
	var body EventBranch0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventBranch0 overwrites any union data inside the Event_Branch as the provided EventBranch0
func (t *Event_Branch) FromEventBranch0(v EventBranch0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventBranch0 performs a merge with any union data inside the Event_Branch, using the provided EventBranch0
func (t *Event_Branch) MergeEventBranch0(v EventBranch0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventBranch1 returns the union data inside the Event_Branch as a EventBranch1
func (t Event_Branch) AsEventBranch1() (EventBranch1, error) {
	var body EventBranch1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventBranch1 overwrites any union data inside the Event_Branch as the provided EventBranch1
func (t *Event_Branch) FromEventBranch1(v EventBranch1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventBranch1 performs a merge with any union data inside the Event_Branch, using the provided EventBranch1
func (t *Event_Branch) MergeEventBranch1(v EventBranch1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_Branch) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_Branch) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsContentOutput returns the union data inside the Event_Content as a ContentOutput
func (t Event_Content) AsContentOutput() (ContentOutput, error) {
	var body ContentOutput
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromContentOutput overwrites any union data inside the Event_Content as the provided ContentOutput
func (t *Event_Content) FromContentOutput(v ContentOutput) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeContentOutput performs a merge with any union data inside the Event_Content, using the provided ContentOutput
func (t *Event_Content) MergeContentOutput(v ContentOutput) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventContent1 returns the union data inside the Event_Content as a EventContent1
func (t Event_Content) AsEventContent1() (EventContent1, error) {
	var body EventContent1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventContent1 overwrites any union data inside the Event_Content as the provided EventContent1
func (t *Event_Content) FromEventContent1(v EventContent1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventContent1 performs a merge with any union data inside the Event_Content, using the provided EventContent1
func (t *Event_Content) MergeEventContent1(v EventContent1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_Content) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_Content) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventErrorCode0 returns the union data inside the Event_ErrorCode as a EventErrorCode0
func (t Event_ErrorCode) AsEventErrorCode0() (EventErrorCode0, error) {
	var body EventErrorCode0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventErrorCode0 overwrites any union data inside the Event_ErrorCode as the provided EventErrorCode0
func (t *Event_ErrorCode) FromEventErrorCode0(v EventErrorCode0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventErrorCode0 performs a merge with any union data inside the Event_ErrorCode, using the provided EventErrorCode0
func (t *Event_ErrorCode) MergeEventErrorCode0(v EventErrorCode0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventErrorCode1 returns the union data inside the Event_ErrorCode as a EventErrorCode1
func (t Event_ErrorCode) AsEventErrorCode1() (EventErrorCode1, error) {
	var body EventErrorCode1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventErrorCode1 overwrites any union data inside the Event_ErrorCode as the provided EventErrorCode1
func (t *Event_ErrorCode) FromEventErrorCode1(v EventErrorCode1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventErrorCode1 performs a merge with any union data inside the Event_ErrorCode, using the provided EventErrorCode1
func (t *Event_ErrorCode) MergeEventErrorCode1(v EventErrorCode1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_ErrorCode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_ErrorCode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventErrorMessage0 returns the union data inside the Event_ErrorMessage as a EventErrorMessage0
func (t Event_ErrorMessage) AsEventErrorMessage0() (EventErrorMessage0, error) {
	var body EventErrorMessage0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventErrorMessage0 overwrites any union data inside the Event_ErrorMessage as the provided EventErrorMessage0
func (t *Event_ErrorMessage) FromEventErrorMessage0(v EventErrorMessage0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventErrorMessage0 performs a merge with any union data inside the Event_ErrorMessage, using the provided EventErrorMessage0
func (t *Event_ErrorMessage) MergeEventErrorMessage0(v EventErrorMessage0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventErrorMessage1 returns the union data inside the Event_ErrorMessage as a EventErrorMessage1
func (t Event_ErrorMessage) AsEventErrorMessage1() (EventErrorMessage1, error) {
	var body EventErrorMessage1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventErrorMessage1 overwrites any union data inside the Event_ErrorMessage as the provided EventErrorMessage1
func (t *Event_ErrorMessage) FromEventErrorMessage1(v EventErrorMessage1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventErrorMessage1 performs a merge with any union data inside the Event_ErrorMessage, using the provided EventErrorMessage1
func (t *Event_ErrorMessage) MergeEventErrorMessage1(v EventErrorMessage1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_ErrorMessage) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_ErrorMessage) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingMetadata returns the union data inside the Event_GroundingMetadata as a GroundingMetadata
func (t Event_GroundingMetadata) AsGroundingMetadata() (GroundingMetadata, error) {
	var body GroundingMetadata
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadata overwrites any union data inside the Event_GroundingMetadata as the provided GroundingMetadata
func (t *Event_GroundingMetadata) FromGroundingMetadata(v GroundingMetadata) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadata performs a merge with any union data inside the Event_GroundingMetadata, using the provided GroundingMetadata
func (t *Event_GroundingMetadata) MergeGroundingMetadata(v GroundingMetadata) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventGroundingMetadata1 returns the union data inside the Event_GroundingMetadata as a EventGroundingMetadata1
func (t Event_GroundingMetadata) AsEventGroundingMetadata1() (EventGroundingMetadata1, error) {
	var body EventGroundingMetadata1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventGroundingMetadata1 overwrites any union data inside the Event_GroundingMetadata as the provided EventGroundingMetadata1
func (t *Event_GroundingMetadata) FromEventGroundingMetadata1(v EventGroundingMetadata1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventGroundingMetadata1 performs a merge with any union data inside the Event_GroundingMetadata, using the provided EventGroundingMetadata1
func (t *Event_GroundingMetadata) MergeEventGroundingMetadata1(v EventGroundingMetadata1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_GroundingMetadata) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_GroundingMetadata) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventInterrupted0 returns the union data inside the Event_Interrupted as a EventInterrupted0
func (t Event_Interrupted) AsEventInterrupted0() (EventInterrupted0, error) {
	var body EventInterrupted0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventInterrupted0 overwrites any union data inside the Event_Interrupted as the provided EventInterrupted0
func (t *Event_Interrupted) FromEventInterrupted0(v EventInterrupted0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventInterrupted0 performs a merge with any union data inside the Event_Interrupted, using the provided EventInterrupted0
func (t *Event_Interrupted) MergeEventInterrupted0(v EventInterrupted0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventInterrupted1 returns the union data inside the Event_Interrupted as a EventInterrupted1
func (t Event_Interrupted) AsEventInterrupted1() (EventInterrupted1, error) {
	var body EventInterrupted1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventInterrupted1 overwrites any union data inside the Event_Interrupted as the provided EventInterrupted1
func (t *Event_Interrupted) FromEventInterrupted1(v EventInterrupted1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventInterrupted1 performs a merge with any union data inside the Event_Interrupted, using the provided EventInterrupted1
func (t *Event_Interrupted) MergeEventInterrupted1(v EventInterrupted1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_Interrupted) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_Interrupted) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventLongRunningToolIds0 returns the union data inside the Event_LongRunningToolIds as a EventLongRunningToolIds0
func (t Event_LongRunningToolIds) AsEventLongRunningToolIds0() (EventLongRunningToolIds0, error) {
	var body EventLongRunningToolIds0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventLongRunningToolIds0 overwrites any union data inside the Event_LongRunningToolIds as the provided EventLongRunningToolIds0
func (t *Event_LongRunningToolIds) FromEventLongRunningToolIds0(v EventLongRunningToolIds0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventLongRunningToolIds0 performs a merge with any union data inside the Event_LongRunningToolIds, using the provided EventLongRunningToolIds0
func (t *Event_LongRunningToolIds) MergeEventLongRunningToolIds0(v EventLongRunningToolIds0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventLongRunningToolIds1 returns the union data inside the Event_LongRunningToolIds as a EventLongRunningToolIds1
func (t Event_LongRunningToolIds) AsEventLongRunningToolIds1() (EventLongRunningToolIds1, error) {
	var body EventLongRunningToolIds1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventLongRunningToolIds1 overwrites any union data inside the Event_LongRunningToolIds as the provided EventLongRunningToolIds1
func (t *Event_LongRunningToolIds) FromEventLongRunningToolIds1(v EventLongRunningToolIds1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventLongRunningToolIds1 performs a merge with any union data inside the Event_LongRunningToolIds, using the provided EventLongRunningToolIds1
func (t *Event_LongRunningToolIds) MergeEventLongRunningToolIds1(v EventLongRunningToolIds1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_LongRunningToolIds) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_LongRunningToolIds) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventPartial0 returns the union data inside the Event_Partial as a EventPartial0
func (t Event_Partial) AsEventPartial0() (EventPartial0, error) {
	var body EventPartial0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventPartial0 overwrites any union data inside the Event_Partial as the provided EventPartial0
func (t *Event_Partial) FromEventPartial0(v EventPartial0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventPartial0 performs a merge with any union data inside the Event_Partial, using the provided EventPartial0
func (t *Event_Partial) MergeEventPartial0(v EventPartial0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventPartial1 returns the union data inside the Event_Partial as a EventPartial1
func (t Event_Partial) AsEventPartial1() (EventPartial1, error) {
	var body EventPartial1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventPartial1 overwrites any union data inside the Event_Partial as the provided EventPartial1
func (t *Event_Partial) FromEventPartial1(v EventPartial1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventPartial1 performs a merge with any union data inside the Event_Partial, using the provided EventPartial1
func (t *Event_Partial) MergeEventPartial1(v EventPartial1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_Partial) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_Partial) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventTurnComplete0 returns the union data inside the Event_TurnComplete as a EventTurnComplete0
func (t Event_TurnComplete) AsEventTurnComplete0() (EventTurnComplete0, error) {
	var body EventTurnComplete0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventTurnComplete0 overwrites any union data inside the Event_TurnComplete as the provided EventTurnComplete0
func (t *Event_TurnComplete) FromEventTurnComplete0(v EventTurnComplete0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventTurnComplete0 performs a merge with any union data inside the Event_TurnComplete, using the provided EventTurnComplete0
func (t *Event_TurnComplete) MergeEventTurnComplete0(v EventTurnComplete0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventTurnComplete1 returns the union data inside the Event_TurnComplete as a EventTurnComplete1
func (t Event_TurnComplete) AsEventTurnComplete1() (EventTurnComplete1, error) {
	var body EventTurnComplete1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventTurnComplete1 overwrites any union data inside the Event_TurnComplete as the provided EventTurnComplete1
func (t *Event_TurnComplete) FromEventTurnComplete1(v EventTurnComplete1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventTurnComplete1 performs a merge with any union data inside the Event_TurnComplete, using the provided EventTurnComplete1
func (t *Event_TurnComplete) MergeEventTurnComplete1(v EventTurnComplete1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Event_TurnComplete) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Event_TurnComplete) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventActionsEscalate0 returns the union data inside the EventActions_Escalate as a EventActionsEscalate0
func (t EventActions_Escalate) AsEventActionsEscalate0() (EventActionsEscalate0, error) {
	var body EventActionsEscalate0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventActionsEscalate0 overwrites any union data inside the EventActions_Escalate as the provided EventActionsEscalate0
func (t *EventActions_Escalate) FromEventActionsEscalate0(v EventActionsEscalate0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventActionsEscalate0 performs a merge with any union data inside the EventActions_Escalate, using the provided EventActionsEscalate0
func (t *EventActions_Escalate) MergeEventActionsEscalate0(v EventActionsEscalate0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventActionsEscalate1 returns the union data inside the EventActions_Escalate as a EventActionsEscalate1
func (t EventActions_Escalate) AsEventActionsEscalate1() (EventActionsEscalate1, error) {
	var body EventActionsEscalate1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventActionsEscalate1 overwrites any union data inside the EventActions_Escalate as the provided EventActionsEscalate1
func (t *EventActions_Escalate) FromEventActionsEscalate1(v EventActionsEscalate1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventActionsEscalate1 performs a merge with any union data inside the EventActions_Escalate, using the provided EventActionsEscalate1
func (t *EventActions_Escalate) MergeEventActionsEscalate1(v EventActionsEscalate1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t EventActions_Escalate) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *EventActions_Escalate) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventActionsSkipSummarization0 returns the union data inside the EventActions_SkipSummarization as a EventActionsSkipSummarization0
func (t EventActions_SkipSummarization) AsEventActionsSkipSummarization0() (EventActionsSkipSummarization0, error) {
	var body EventActionsSkipSummarization0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventActionsSkipSummarization0 overwrites any union data inside the EventActions_SkipSummarization as the provided EventActionsSkipSummarization0
func (t *EventActions_SkipSummarization) FromEventActionsSkipSummarization0(v EventActionsSkipSummarization0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventActionsSkipSummarization0 performs a merge with any union data inside the EventActions_SkipSummarization, using the provided EventActionsSkipSummarization0
func (t *EventActions_SkipSummarization) MergeEventActionsSkipSummarization0(v EventActionsSkipSummarization0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventActionsSkipSummarization1 returns the union data inside the EventActions_SkipSummarization as a EventActionsSkipSummarization1
func (t EventActions_SkipSummarization) AsEventActionsSkipSummarization1() (EventActionsSkipSummarization1, error) {
	var body EventActionsSkipSummarization1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventActionsSkipSummarization1 overwrites any union data inside the EventActions_SkipSummarization as the provided EventActionsSkipSummarization1
func (t *EventActions_SkipSummarization) FromEventActionsSkipSummarization1(v EventActionsSkipSummarization1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventActionsSkipSummarization1 performs a merge with any union data inside the EventActions_SkipSummarization, using the provided EventActionsSkipSummarization1
func (t *EventActions_SkipSummarization) MergeEventActionsSkipSummarization1(v EventActionsSkipSummarization1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t EventActions_SkipSummarization) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *EventActions_SkipSummarization) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsEventActionsTransferToAgent0 returns the union data inside the EventActions_TransferToAgent as a EventActionsTransferToAgent0
func (t EventActions_TransferToAgent) AsEventActionsTransferToAgent0() (EventActionsTransferToAgent0, error) {
	var body EventActionsTransferToAgent0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventActionsTransferToAgent0 overwrites any union data inside the EventActions_TransferToAgent as the provided EventActionsTransferToAgent0
func (t *EventActions_TransferToAgent) FromEventActionsTransferToAgent0(v EventActionsTransferToAgent0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventActionsTransferToAgent0 performs a merge with any union data inside the EventActions_TransferToAgent, using the provided EventActionsTransferToAgent0
func (t *EventActions_TransferToAgent) MergeEventActionsTransferToAgent0(v EventActionsTransferToAgent0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsEventActionsTransferToAgent1 returns the union data inside the EventActions_TransferToAgent as a EventActionsTransferToAgent1
func (t EventActions_TransferToAgent) AsEventActionsTransferToAgent1() (EventActionsTransferToAgent1, error) {
	var body EventActionsTransferToAgent1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromEventActionsTransferToAgent1 overwrites any union data inside the EventActions_TransferToAgent as the provided EventActionsTransferToAgent1
func (t *EventActions_TransferToAgent) FromEventActionsTransferToAgent1(v EventActionsTransferToAgent1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeEventActionsTransferToAgent1 performs a merge with any union data inside the EventActions_TransferToAgent, using the provided EventActionsTransferToAgent1
func (t *EventActions_TransferToAgent) MergeEventActionsTransferToAgent1(v EventActionsTransferToAgent1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t EventActions_TransferToAgent) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *EventActions_TransferToAgent) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsExecutableCodeCode0 returns the union data inside the ExecutableCode_Code as a ExecutableCodeCode0
func (t ExecutableCode_Code) AsExecutableCodeCode0() (ExecutableCodeCode0, error) {
	var body ExecutableCodeCode0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromExecutableCodeCode0 overwrites any union data inside the ExecutableCode_Code as the provided ExecutableCodeCode0
func (t *ExecutableCode_Code) FromExecutableCodeCode0(v ExecutableCodeCode0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeExecutableCodeCode0 performs a merge with any union data inside the ExecutableCode_Code, using the provided ExecutableCodeCode0
func (t *ExecutableCode_Code) MergeExecutableCodeCode0(v ExecutableCodeCode0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsExecutableCodeCode1 returns the union data inside the ExecutableCode_Code as a ExecutableCodeCode1
func (t ExecutableCode_Code) AsExecutableCodeCode1() (ExecutableCodeCode1, error) {
	var body ExecutableCodeCode1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromExecutableCodeCode1 overwrites any union data inside the ExecutableCode_Code as the provided ExecutableCodeCode1
func (t *ExecutableCode_Code) FromExecutableCodeCode1(v ExecutableCodeCode1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeExecutableCodeCode1 performs a merge with any union data inside the ExecutableCode_Code, using the provided ExecutableCodeCode1
func (t *ExecutableCode_Code) MergeExecutableCodeCode1(v ExecutableCodeCode1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ExecutableCode_Code) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ExecutableCode_Code) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsLanguage returns the union data inside the ExecutableCode_Language as a Language
func (t ExecutableCode_Language) AsLanguage() (Language, error) {
	var body Language
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromLanguage overwrites any union data inside the ExecutableCode_Language as the provided Language
func (t *ExecutableCode_Language) FromLanguage(v Language) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeLanguage performs a merge with any union data inside the ExecutableCode_Language, using the provided Language
func (t *ExecutableCode_Language) MergeLanguage(v Language) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsExecutableCodeLanguage1 returns the union data inside the ExecutableCode_Language as a ExecutableCodeLanguage1
func (t ExecutableCode_Language) AsExecutableCodeLanguage1() (ExecutableCodeLanguage1, error) {
	var body ExecutableCodeLanguage1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromExecutableCodeLanguage1 overwrites any union data inside the ExecutableCode_Language as the provided ExecutableCodeLanguage1
func (t *ExecutableCode_Language) FromExecutableCodeLanguage1(v ExecutableCodeLanguage1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeExecutableCodeLanguage1 performs a merge with any union data inside the ExecutableCode_Language, using the provided ExecutableCodeLanguage1
func (t *ExecutableCode_Language) MergeExecutableCodeLanguage1(v ExecutableCodeLanguage1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ExecutableCode_Language) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ExecutableCode_Language) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFileDataFileUri0 returns the union data inside the FileData_FileUri as a FileDataFileUri0
func (t FileData_FileUri) AsFileDataFileUri0() (FileDataFileUri0, error) {
	var body FileDataFileUri0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileDataFileUri0 overwrites any union data inside the FileData_FileUri as the provided FileDataFileUri0
func (t *FileData_FileUri) FromFileDataFileUri0(v FileDataFileUri0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileDataFileUri0 performs a merge with any union data inside the FileData_FileUri, using the provided FileDataFileUri0
func (t *FileData_FileUri) MergeFileDataFileUri0(v FileDataFileUri0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFileDataFileUri1 returns the union data inside the FileData_FileUri as a FileDataFileUri1
func (t FileData_FileUri) AsFileDataFileUri1() (FileDataFileUri1, error) {
	var body FileDataFileUri1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileDataFileUri1 overwrites any union data inside the FileData_FileUri as the provided FileDataFileUri1
func (t *FileData_FileUri) FromFileDataFileUri1(v FileDataFileUri1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileDataFileUri1 performs a merge with any union data inside the FileData_FileUri, using the provided FileDataFileUri1
func (t *FileData_FileUri) MergeFileDataFileUri1(v FileDataFileUri1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FileData_FileUri) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FileData_FileUri) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFileDataMimeType0 returns the union data inside the FileData_MimeType as a FileDataMimeType0
func (t FileData_MimeType) AsFileDataMimeType0() (FileDataMimeType0, error) {
	var body FileDataMimeType0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileDataMimeType0 overwrites any union data inside the FileData_MimeType as the provided FileDataMimeType0
func (t *FileData_MimeType) FromFileDataMimeType0(v FileDataMimeType0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileDataMimeType0 performs a merge with any union data inside the FileData_MimeType, using the provided FileDataMimeType0
func (t *FileData_MimeType) MergeFileDataMimeType0(v FileDataMimeType0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFileDataMimeType1 returns the union data inside the FileData_MimeType as a FileDataMimeType1
func (t FileData_MimeType) AsFileDataMimeType1() (FileDataMimeType1, error) {
	var body FileDataMimeType1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileDataMimeType1 overwrites any union data inside the FileData_MimeType as the provided FileDataMimeType1
func (t *FileData_MimeType) FromFileDataMimeType1(v FileDataMimeType1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileDataMimeType1 performs a merge with any union data inside the FileData_MimeType, using the provided FileDataMimeType1
func (t *FileData_MimeType) MergeFileDataMimeType1(v FileDataMimeType1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FileData_MimeType) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FileData_MimeType) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionCallArgs0 returns the union data inside the FunctionCall_Args as a FunctionCallArgs0
func (t FunctionCall_Args) AsFunctionCallArgs0() (FunctionCallArgs0, error) {
	var body FunctionCallArgs0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCallArgs0 overwrites any union data inside the FunctionCall_Args as the provided FunctionCallArgs0
func (t *FunctionCall_Args) FromFunctionCallArgs0(v FunctionCallArgs0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCallArgs0 performs a merge with any union data inside the FunctionCall_Args, using the provided FunctionCallArgs0
func (t *FunctionCall_Args) MergeFunctionCallArgs0(v FunctionCallArgs0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFunctionCallArgs1 returns the union data inside the FunctionCall_Args as a FunctionCallArgs1
func (t FunctionCall_Args) AsFunctionCallArgs1() (FunctionCallArgs1, error) {
	var body FunctionCallArgs1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCallArgs1 overwrites any union data inside the FunctionCall_Args as the provided FunctionCallArgs1
func (t *FunctionCall_Args) FromFunctionCallArgs1(v FunctionCallArgs1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCallArgs1 performs a merge with any union data inside the FunctionCall_Args, using the provided FunctionCallArgs1
func (t *FunctionCall_Args) MergeFunctionCallArgs1(v FunctionCallArgs1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FunctionCall_Args) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FunctionCall_Args) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionCallId0 returns the union data inside the FunctionCall_Id as a FunctionCallId0
func (t FunctionCall_Id) AsFunctionCallId0() (FunctionCallId0, error) {
	var body FunctionCallId0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCallId0 overwrites any union data inside the FunctionCall_Id as the provided FunctionCallId0
func (t *FunctionCall_Id) FromFunctionCallId0(v FunctionCallId0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCallId0 performs a merge with any union data inside the FunctionCall_Id, using the provided FunctionCallId0
func (t *FunctionCall_Id) MergeFunctionCallId0(v FunctionCallId0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFunctionCallId1 returns the union data inside the FunctionCall_Id as a FunctionCallId1
func (t FunctionCall_Id) AsFunctionCallId1() (FunctionCallId1, error) {
	var body FunctionCallId1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCallId1 overwrites any union data inside the FunctionCall_Id as the provided FunctionCallId1
func (t *FunctionCall_Id) FromFunctionCallId1(v FunctionCallId1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCallId1 performs a merge with any union data inside the FunctionCall_Id, using the provided FunctionCallId1
func (t *FunctionCall_Id) MergeFunctionCallId1(v FunctionCallId1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FunctionCall_Id) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FunctionCall_Id) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionCallName0 returns the union data inside the FunctionCall_Name as a FunctionCallName0
func (t FunctionCall_Name) AsFunctionCallName0() (FunctionCallName0, error) {
	var body FunctionCallName0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCallName0 overwrites any union data inside the FunctionCall_Name as the provided FunctionCallName0
func (t *FunctionCall_Name) FromFunctionCallName0(v FunctionCallName0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCallName0 performs a merge with any union data inside the FunctionCall_Name, using the provided FunctionCallName0
func (t *FunctionCall_Name) MergeFunctionCallName0(v FunctionCallName0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFunctionCallName1 returns the union data inside the FunctionCall_Name as a FunctionCallName1
func (t FunctionCall_Name) AsFunctionCallName1() (FunctionCallName1, error) {
	var body FunctionCallName1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCallName1 overwrites any union data inside the FunctionCall_Name as the provided FunctionCallName1
func (t *FunctionCall_Name) FromFunctionCallName1(v FunctionCallName1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCallName1 performs a merge with any union data inside the FunctionCall_Name, using the provided FunctionCallName1
func (t *FunctionCall_Name) MergeFunctionCallName1(v FunctionCallName1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FunctionCall_Name) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FunctionCall_Name) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionResponseId0 returns the union data inside the FunctionResponse_Id as a FunctionResponseId0
func (t FunctionResponse_Id) AsFunctionResponseId0() (FunctionResponseId0, error) {
	var body FunctionResponseId0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponseId0 overwrites any union data inside the FunctionResponse_Id as the provided FunctionResponseId0
func (t *FunctionResponse_Id) FromFunctionResponseId0(v FunctionResponseId0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponseId0 performs a merge with any union data inside the FunctionResponse_Id, using the provided FunctionResponseId0
func (t *FunctionResponse_Id) MergeFunctionResponseId0(v FunctionResponseId0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFunctionResponseId1 returns the union data inside the FunctionResponse_Id as a FunctionResponseId1
func (t FunctionResponse_Id) AsFunctionResponseId1() (FunctionResponseId1, error) {
	var body FunctionResponseId1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponseId1 overwrites any union data inside the FunctionResponse_Id as the provided FunctionResponseId1
func (t *FunctionResponse_Id) FromFunctionResponseId1(v FunctionResponseId1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponseId1 performs a merge with any union data inside the FunctionResponse_Id, using the provided FunctionResponseId1
func (t *FunctionResponse_Id) MergeFunctionResponseId1(v FunctionResponseId1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FunctionResponse_Id) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FunctionResponse_Id) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionResponseName0 returns the union data inside the FunctionResponse_Name as a FunctionResponseName0
func (t FunctionResponse_Name) AsFunctionResponseName0() (FunctionResponseName0, error) {
	var body FunctionResponseName0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponseName0 overwrites any union data inside the FunctionResponse_Name as the provided FunctionResponseName0
func (t *FunctionResponse_Name) FromFunctionResponseName0(v FunctionResponseName0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponseName0 performs a merge with any union data inside the FunctionResponse_Name, using the provided FunctionResponseName0
func (t *FunctionResponse_Name) MergeFunctionResponseName0(v FunctionResponseName0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFunctionResponseName1 returns the union data inside the FunctionResponse_Name as a FunctionResponseName1
func (t FunctionResponse_Name) AsFunctionResponseName1() (FunctionResponseName1, error) {
	var body FunctionResponseName1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponseName1 overwrites any union data inside the FunctionResponse_Name as the provided FunctionResponseName1
func (t *FunctionResponse_Name) FromFunctionResponseName1(v FunctionResponseName1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponseName1 performs a merge with any union data inside the FunctionResponse_Name, using the provided FunctionResponseName1
func (t *FunctionResponse_Name) MergeFunctionResponseName1(v FunctionResponseName1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FunctionResponse_Name) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FunctionResponse_Name) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionResponseResponse0 returns the union data inside the FunctionResponse_Response as a FunctionResponseResponse0
func (t FunctionResponse_Response) AsFunctionResponseResponse0() (FunctionResponseResponse0, error) {
	var body FunctionResponseResponse0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponseResponse0 overwrites any union data inside the FunctionResponse_Response as the provided FunctionResponseResponse0
func (t *FunctionResponse_Response) FromFunctionResponseResponse0(v FunctionResponseResponse0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponseResponse0 performs a merge with any union data inside the FunctionResponse_Response, using the provided FunctionResponseResponse0
func (t *FunctionResponse_Response) MergeFunctionResponseResponse0(v FunctionResponseResponse0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsFunctionResponseResponse1 returns the union data inside the FunctionResponse_Response as a FunctionResponseResponse1
func (t FunctionResponse_Response) AsFunctionResponseResponse1() (FunctionResponseResponse1, error) {
	var body FunctionResponseResponse1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponseResponse1 overwrites any union data inside the FunctionResponse_Response as the provided FunctionResponseResponse1
func (t *FunctionResponse_Response) FromFunctionResponseResponse1(v FunctionResponseResponse1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponseResponse1 performs a merge with any union data inside the FunctionResponse_Response, using the provided FunctionResponseResponse1
func (t *FunctionResponse_Response) MergeFunctionResponseResponse1(v FunctionResponseResponse1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t FunctionResponse_Response) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *FunctionResponse_Response) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkRetrievedContext returns the union data inside the GroundingChunk_RetrievedContext as a GroundingChunkRetrievedContext
func (t GroundingChunk_RetrievedContext) AsGroundingChunkRetrievedContext() (GroundingChunkRetrievedContext, error) {
	var body GroundingChunkRetrievedContext
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContext overwrites any union data inside the GroundingChunk_RetrievedContext as the provided GroundingChunkRetrievedContext
func (t *GroundingChunk_RetrievedContext) FromGroundingChunkRetrievedContext(v GroundingChunkRetrievedContext) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContext performs a merge with any union data inside the GroundingChunk_RetrievedContext, using the provided GroundingChunkRetrievedContext
func (t *GroundingChunk_RetrievedContext) MergeGroundingChunkRetrievedContext(v GroundingChunkRetrievedContext) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkRetrievedContext1 returns the union data inside the GroundingChunk_RetrievedContext as a GroundingChunkRetrievedContext1
func (t GroundingChunk_RetrievedContext) AsGroundingChunkRetrievedContext1() (GroundingChunkRetrievedContext1, error) {
	var body GroundingChunkRetrievedContext1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContext1 overwrites any union data inside the GroundingChunk_RetrievedContext as the provided GroundingChunkRetrievedContext1
func (t *GroundingChunk_RetrievedContext) FromGroundingChunkRetrievedContext1(v GroundingChunkRetrievedContext1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContext1 performs a merge with any union data inside the GroundingChunk_RetrievedContext, using the provided GroundingChunkRetrievedContext1
func (t *GroundingChunk_RetrievedContext) MergeGroundingChunkRetrievedContext1(v GroundingChunkRetrievedContext1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunk_RetrievedContext) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunk_RetrievedContext) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkWeb returns the union data inside the GroundingChunk_Web as a GroundingChunkWeb
func (t GroundingChunk_Web) AsGroundingChunkWeb() (GroundingChunkWeb, error) {
	var body GroundingChunkWeb
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWeb overwrites any union data inside the GroundingChunk_Web as the provided GroundingChunkWeb
func (t *GroundingChunk_Web) FromGroundingChunkWeb(v GroundingChunkWeb) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWeb performs a merge with any union data inside the GroundingChunk_Web, using the provided GroundingChunkWeb
func (t *GroundingChunk_Web) MergeGroundingChunkWeb(v GroundingChunkWeb) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkWeb1 returns the union data inside the GroundingChunk_Web as a GroundingChunkWeb1
func (t GroundingChunk_Web) AsGroundingChunkWeb1() (GroundingChunkWeb1, error) {
	var body GroundingChunkWeb1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWeb1 overwrites any union data inside the GroundingChunk_Web as the provided GroundingChunkWeb1
func (t *GroundingChunk_Web) FromGroundingChunkWeb1(v GroundingChunkWeb1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWeb1 performs a merge with any union data inside the GroundingChunk_Web, using the provided GroundingChunkWeb1
func (t *GroundingChunk_Web) MergeGroundingChunkWeb1(v GroundingChunkWeb1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunk_Web) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunk_Web) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkRetrievedContextText0 returns the union data inside the GroundingChunkRetrievedContext_Text as a GroundingChunkRetrievedContextText0
func (t GroundingChunkRetrievedContext_Text) AsGroundingChunkRetrievedContextText0() (GroundingChunkRetrievedContextText0, error) {
	var body GroundingChunkRetrievedContextText0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContextText0 overwrites any union data inside the GroundingChunkRetrievedContext_Text as the provided GroundingChunkRetrievedContextText0
func (t *GroundingChunkRetrievedContext_Text) FromGroundingChunkRetrievedContextText0(v GroundingChunkRetrievedContextText0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContextText0 performs a merge with any union data inside the GroundingChunkRetrievedContext_Text, using the provided GroundingChunkRetrievedContextText0
func (t *GroundingChunkRetrievedContext_Text) MergeGroundingChunkRetrievedContextText0(v GroundingChunkRetrievedContextText0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkRetrievedContextText1 returns the union data inside the GroundingChunkRetrievedContext_Text as a GroundingChunkRetrievedContextText1
func (t GroundingChunkRetrievedContext_Text) AsGroundingChunkRetrievedContextText1() (GroundingChunkRetrievedContextText1, error) {
	var body GroundingChunkRetrievedContextText1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContextText1 overwrites any union data inside the GroundingChunkRetrievedContext_Text as the provided GroundingChunkRetrievedContextText1
func (t *GroundingChunkRetrievedContext_Text) FromGroundingChunkRetrievedContextText1(v GroundingChunkRetrievedContextText1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContextText1 performs a merge with any union data inside the GroundingChunkRetrievedContext_Text, using the provided GroundingChunkRetrievedContextText1
func (t *GroundingChunkRetrievedContext_Text) MergeGroundingChunkRetrievedContextText1(v GroundingChunkRetrievedContextText1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunkRetrievedContext_Text) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunkRetrievedContext_Text) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkRetrievedContextTitle0 returns the union data inside the GroundingChunkRetrievedContext_Title as a GroundingChunkRetrievedContextTitle0
func (t GroundingChunkRetrievedContext_Title) AsGroundingChunkRetrievedContextTitle0() (GroundingChunkRetrievedContextTitle0, error) {
	var body GroundingChunkRetrievedContextTitle0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContextTitle0 overwrites any union data inside the GroundingChunkRetrievedContext_Title as the provided GroundingChunkRetrievedContextTitle0
func (t *GroundingChunkRetrievedContext_Title) FromGroundingChunkRetrievedContextTitle0(v GroundingChunkRetrievedContextTitle0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContextTitle0 performs a merge with any union data inside the GroundingChunkRetrievedContext_Title, using the provided GroundingChunkRetrievedContextTitle0
func (t *GroundingChunkRetrievedContext_Title) MergeGroundingChunkRetrievedContextTitle0(v GroundingChunkRetrievedContextTitle0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkRetrievedContextTitle1 returns the union data inside the GroundingChunkRetrievedContext_Title as a GroundingChunkRetrievedContextTitle1
func (t GroundingChunkRetrievedContext_Title) AsGroundingChunkRetrievedContextTitle1() (GroundingChunkRetrievedContextTitle1, error) {
	var body GroundingChunkRetrievedContextTitle1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContextTitle1 overwrites any union data inside the GroundingChunkRetrievedContext_Title as the provided GroundingChunkRetrievedContextTitle1
func (t *GroundingChunkRetrievedContext_Title) FromGroundingChunkRetrievedContextTitle1(v GroundingChunkRetrievedContextTitle1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContextTitle1 performs a merge with any union data inside the GroundingChunkRetrievedContext_Title, using the provided GroundingChunkRetrievedContextTitle1
func (t *GroundingChunkRetrievedContext_Title) MergeGroundingChunkRetrievedContextTitle1(v GroundingChunkRetrievedContextTitle1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunkRetrievedContext_Title) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunkRetrievedContext_Title) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkRetrievedContextUri0 returns the union data inside the GroundingChunkRetrievedContext_Uri as a GroundingChunkRetrievedContextUri0
func (t GroundingChunkRetrievedContext_Uri) AsGroundingChunkRetrievedContextUri0() (GroundingChunkRetrievedContextUri0, error) {
	var body GroundingChunkRetrievedContextUri0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContextUri0 overwrites any union data inside the GroundingChunkRetrievedContext_Uri as the provided GroundingChunkRetrievedContextUri0
func (t *GroundingChunkRetrievedContext_Uri) FromGroundingChunkRetrievedContextUri0(v GroundingChunkRetrievedContextUri0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContextUri0 performs a merge with any union data inside the GroundingChunkRetrievedContext_Uri, using the provided GroundingChunkRetrievedContextUri0
func (t *GroundingChunkRetrievedContext_Uri) MergeGroundingChunkRetrievedContextUri0(v GroundingChunkRetrievedContextUri0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkRetrievedContextUri1 returns the union data inside the GroundingChunkRetrievedContext_Uri as a GroundingChunkRetrievedContextUri1
func (t GroundingChunkRetrievedContext_Uri) AsGroundingChunkRetrievedContextUri1() (GroundingChunkRetrievedContextUri1, error) {
	var body GroundingChunkRetrievedContextUri1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkRetrievedContextUri1 overwrites any union data inside the GroundingChunkRetrievedContext_Uri as the provided GroundingChunkRetrievedContextUri1
func (t *GroundingChunkRetrievedContext_Uri) FromGroundingChunkRetrievedContextUri1(v GroundingChunkRetrievedContextUri1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkRetrievedContextUri1 performs a merge with any union data inside the GroundingChunkRetrievedContext_Uri, using the provided GroundingChunkRetrievedContextUri1
func (t *GroundingChunkRetrievedContext_Uri) MergeGroundingChunkRetrievedContextUri1(v GroundingChunkRetrievedContextUri1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunkRetrievedContext_Uri) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunkRetrievedContext_Uri) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkWebDomain0 returns the union data inside the GroundingChunkWeb_Domain as a GroundingChunkWebDomain0
func (t GroundingChunkWeb_Domain) AsGroundingChunkWebDomain0() (GroundingChunkWebDomain0, error) {
	var body GroundingChunkWebDomain0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWebDomain0 overwrites any union data inside the GroundingChunkWeb_Domain as the provided GroundingChunkWebDomain0
func (t *GroundingChunkWeb_Domain) FromGroundingChunkWebDomain0(v GroundingChunkWebDomain0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWebDomain0 performs a merge with any union data inside the GroundingChunkWeb_Domain, using the provided GroundingChunkWebDomain0
func (t *GroundingChunkWeb_Domain) MergeGroundingChunkWebDomain0(v GroundingChunkWebDomain0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkWebDomain1 returns the union data inside the GroundingChunkWeb_Domain as a GroundingChunkWebDomain1
func (t GroundingChunkWeb_Domain) AsGroundingChunkWebDomain1() (GroundingChunkWebDomain1, error) {
	var body GroundingChunkWebDomain1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWebDomain1 overwrites any union data inside the GroundingChunkWeb_Domain as the provided GroundingChunkWebDomain1
func (t *GroundingChunkWeb_Domain) FromGroundingChunkWebDomain1(v GroundingChunkWebDomain1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWebDomain1 performs a merge with any union data inside the GroundingChunkWeb_Domain, using the provided GroundingChunkWebDomain1
func (t *GroundingChunkWeb_Domain) MergeGroundingChunkWebDomain1(v GroundingChunkWebDomain1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunkWeb_Domain) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunkWeb_Domain) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkWebTitle0 returns the union data inside the GroundingChunkWeb_Title as a GroundingChunkWebTitle0
func (t GroundingChunkWeb_Title) AsGroundingChunkWebTitle0() (GroundingChunkWebTitle0, error) {
	var body GroundingChunkWebTitle0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWebTitle0 overwrites any union data inside the GroundingChunkWeb_Title as the provided GroundingChunkWebTitle0
func (t *GroundingChunkWeb_Title) FromGroundingChunkWebTitle0(v GroundingChunkWebTitle0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWebTitle0 performs a merge with any union data inside the GroundingChunkWeb_Title, using the provided GroundingChunkWebTitle0
func (t *GroundingChunkWeb_Title) MergeGroundingChunkWebTitle0(v GroundingChunkWebTitle0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkWebTitle1 returns the union data inside the GroundingChunkWeb_Title as a GroundingChunkWebTitle1
func (t GroundingChunkWeb_Title) AsGroundingChunkWebTitle1() (GroundingChunkWebTitle1, error) {
	var body GroundingChunkWebTitle1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWebTitle1 overwrites any union data inside the GroundingChunkWeb_Title as the provided GroundingChunkWebTitle1
func (t *GroundingChunkWeb_Title) FromGroundingChunkWebTitle1(v GroundingChunkWebTitle1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWebTitle1 performs a merge with any union data inside the GroundingChunkWeb_Title, using the provided GroundingChunkWebTitle1
func (t *GroundingChunkWeb_Title) MergeGroundingChunkWebTitle1(v GroundingChunkWebTitle1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunkWeb_Title) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunkWeb_Title) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingChunkWebUri0 returns the union data inside the GroundingChunkWeb_Uri as a GroundingChunkWebUri0
func (t GroundingChunkWeb_Uri) AsGroundingChunkWebUri0() (GroundingChunkWebUri0, error) {
	var body GroundingChunkWebUri0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWebUri0 overwrites any union data inside the GroundingChunkWeb_Uri as the provided GroundingChunkWebUri0
func (t *GroundingChunkWeb_Uri) FromGroundingChunkWebUri0(v GroundingChunkWebUri0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWebUri0 performs a merge with any union data inside the GroundingChunkWeb_Uri, using the provided GroundingChunkWebUri0
func (t *GroundingChunkWeb_Uri) MergeGroundingChunkWebUri0(v GroundingChunkWebUri0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingChunkWebUri1 returns the union data inside the GroundingChunkWeb_Uri as a GroundingChunkWebUri1
func (t GroundingChunkWeb_Uri) AsGroundingChunkWebUri1() (GroundingChunkWebUri1, error) {
	var body GroundingChunkWebUri1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingChunkWebUri1 overwrites any union data inside the GroundingChunkWeb_Uri as the provided GroundingChunkWebUri1
func (t *GroundingChunkWeb_Uri) FromGroundingChunkWebUri1(v GroundingChunkWebUri1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingChunkWebUri1 performs a merge with any union data inside the GroundingChunkWeb_Uri, using the provided GroundingChunkWebUri1
func (t *GroundingChunkWeb_Uri) MergeGroundingChunkWebUri1(v GroundingChunkWebUri1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingChunkWeb_Uri) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingChunkWeb_Uri) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingMetadataGroundingChunks0 returns the union data inside the GroundingMetadata_GroundingChunks as a GroundingMetadataGroundingChunks0
func (t GroundingMetadata_GroundingChunks) AsGroundingMetadataGroundingChunks0() (GroundingMetadataGroundingChunks0, error) {
	var body GroundingMetadataGroundingChunks0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataGroundingChunks0 overwrites any union data inside the GroundingMetadata_GroundingChunks as the provided GroundingMetadataGroundingChunks0
func (t *GroundingMetadata_GroundingChunks) FromGroundingMetadataGroundingChunks0(v GroundingMetadataGroundingChunks0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataGroundingChunks0 performs a merge with any union data inside the GroundingMetadata_GroundingChunks, using the provided GroundingMetadataGroundingChunks0
func (t *GroundingMetadata_GroundingChunks) MergeGroundingMetadataGroundingChunks0(v GroundingMetadataGroundingChunks0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingMetadataGroundingChunks1 returns the union data inside the GroundingMetadata_GroundingChunks as a GroundingMetadataGroundingChunks1
func (t GroundingMetadata_GroundingChunks) AsGroundingMetadataGroundingChunks1() (GroundingMetadataGroundingChunks1, error) {
	var body GroundingMetadataGroundingChunks1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataGroundingChunks1 overwrites any union data inside the GroundingMetadata_GroundingChunks as the provided GroundingMetadataGroundingChunks1
func (t *GroundingMetadata_GroundingChunks) FromGroundingMetadataGroundingChunks1(v GroundingMetadataGroundingChunks1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataGroundingChunks1 performs a merge with any union data inside the GroundingMetadata_GroundingChunks, using the provided GroundingMetadataGroundingChunks1
func (t *GroundingMetadata_GroundingChunks) MergeGroundingMetadataGroundingChunks1(v GroundingMetadataGroundingChunks1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingMetadata_GroundingChunks) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingMetadata_GroundingChunks) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingMetadataGroundingSupports0 returns the union data inside the GroundingMetadata_GroundingSupports as a GroundingMetadataGroundingSupports0
func (t GroundingMetadata_GroundingSupports) AsGroundingMetadataGroundingSupports0() (GroundingMetadataGroundingSupports0, error) {
	var body GroundingMetadataGroundingSupports0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataGroundingSupports0 overwrites any union data inside the GroundingMetadata_GroundingSupports as the provided GroundingMetadataGroundingSupports0
func (t *GroundingMetadata_GroundingSupports) FromGroundingMetadataGroundingSupports0(v GroundingMetadataGroundingSupports0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataGroundingSupports0 performs a merge with any union data inside the GroundingMetadata_GroundingSupports, using the provided GroundingMetadataGroundingSupports0
func (t *GroundingMetadata_GroundingSupports) MergeGroundingMetadataGroundingSupports0(v GroundingMetadataGroundingSupports0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingMetadataGroundingSupports1 returns the union data inside the GroundingMetadata_GroundingSupports as a GroundingMetadataGroundingSupports1
func (t GroundingMetadata_GroundingSupports) AsGroundingMetadataGroundingSupports1() (GroundingMetadataGroundingSupports1, error) {
	var body GroundingMetadataGroundingSupports1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataGroundingSupports1 overwrites any union data inside the GroundingMetadata_GroundingSupports as the provided GroundingMetadataGroundingSupports1
func (t *GroundingMetadata_GroundingSupports) FromGroundingMetadataGroundingSupports1(v GroundingMetadataGroundingSupports1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataGroundingSupports1 performs a merge with any union data inside the GroundingMetadata_GroundingSupports, using the provided GroundingMetadataGroundingSupports1
func (t *GroundingMetadata_GroundingSupports) MergeGroundingMetadataGroundingSupports1(v GroundingMetadataGroundingSupports1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingMetadata_GroundingSupports) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingMetadata_GroundingSupports) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsRetrievalMetadata returns the union data inside the GroundingMetadata_RetrievalMetadata as a RetrievalMetadata
func (t GroundingMetadata_RetrievalMetadata) AsRetrievalMetadata() (RetrievalMetadata, error) {
	var body RetrievalMetadata
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRetrievalMetadata overwrites any union data inside the GroundingMetadata_RetrievalMetadata as the provided RetrievalMetadata
func (t *GroundingMetadata_RetrievalMetadata) FromRetrievalMetadata(v RetrievalMetadata) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRetrievalMetadata performs a merge with any union data inside the GroundingMetadata_RetrievalMetadata, using the provided RetrievalMetadata
func (t *GroundingMetadata_RetrievalMetadata) MergeRetrievalMetadata(v RetrievalMetadata) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingMetadataRetrievalMetadata1 returns the union data inside the GroundingMetadata_RetrievalMetadata as a GroundingMetadataRetrievalMetadata1
func (t GroundingMetadata_RetrievalMetadata) AsGroundingMetadataRetrievalMetadata1() (GroundingMetadataRetrievalMetadata1, error) {
	var body GroundingMetadataRetrievalMetadata1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataRetrievalMetadata1 overwrites any union data inside the GroundingMetadata_RetrievalMetadata as the provided GroundingMetadataRetrievalMetadata1
func (t *GroundingMetadata_RetrievalMetadata) FromGroundingMetadataRetrievalMetadata1(v GroundingMetadataRetrievalMetadata1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataRetrievalMetadata1 performs a merge with any union data inside the GroundingMetadata_RetrievalMetadata, using the provided GroundingMetadataRetrievalMetadata1
func (t *GroundingMetadata_RetrievalMetadata) MergeGroundingMetadataRetrievalMetadata1(v GroundingMetadataRetrievalMetadata1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingMetadata_RetrievalMetadata) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingMetadata_RetrievalMetadata) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingMetadataRetrievalQueries0 returns the union data inside the GroundingMetadata_RetrievalQueries as a GroundingMetadataRetrievalQueries0
func (t GroundingMetadata_RetrievalQueries) AsGroundingMetadataRetrievalQueries0() (GroundingMetadataRetrievalQueries0, error) {
	var body GroundingMetadataRetrievalQueries0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataRetrievalQueries0 overwrites any union data inside the GroundingMetadata_RetrievalQueries as the provided GroundingMetadataRetrievalQueries0
func (t *GroundingMetadata_RetrievalQueries) FromGroundingMetadataRetrievalQueries0(v GroundingMetadataRetrievalQueries0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataRetrievalQueries0 performs a merge with any union data inside the GroundingMetadata_RetrievalQueries, using the provided GroundingMetadataRetrievalQueries0
func (t *GroundingMetadata_RetrievalQueries) MergeGroundingMetadataRetrievalQueries0(v GroundingMetadataRetrievalQueries0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingMetadataRetrievalQueries1 returns the union data inside the GroundingMetadata_RetrievalQueries as a GroundingMetadataRetrievalQueries1
func (t GroundingMetadata_RetrievalQueries) AsGroundingMetadataRetrievalQueries1() (GroundingMetadataRetrievalQueries1, error) {
	var body GroundingMetadataRetrievalQueries1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataRetrievalQueries1 overwrites any union data inside the GroundingMetadata_RetrievalQueries as the provided GroundingMetadataRetrievalQueries1
func (t *GroundingMetadata_RetrievalQueries) FromGroundingMetadataRetrievalQueries1(v GroundingMetadataRetrievalQueries1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataRetrievalQueries1 performs a merge with any union data inside the GroundingMetadata_RetrievalQueries, using the provided GroundingMetadataRetrievalQueries1
func (t *GroundingMetadata_RetrievalQueries) MergeGroundingMetadataRetrievalQueries1(v GroundingMetadataRetrievalQueries1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingMetadata_RetrievalQueries) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingMetadata_RetrievalQueries) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSearchEntryPoint returns the union data inside the GroundingMetadata_SearchEntryPoint as a SearchEntryPoint
func (t GroundingMetadata_SearchEntryPoint) AsSearchEntryPoint() (SearchEntryPoint, error) {
	var body SearchEntryPoint
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSearchEntryPoint overwrites any union data inside the GroundingMetadata_SearchEntryPoint as the provided SearchEntryPoint
func (t *GroundingMetadata_SearchEntryPoint) FromSearchEntryPoint(v SearchEntryPoint) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSearchEntryPoint performs a merge with any union data inside the GroundingMetadata_SearchEntryPoint, using the provided SearchEntryPoint
func (t *GroundingMetadata_SearchEntryPoint) MergeSearchEntryPoint(v SearchEntryPoint) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingMetadataSearchEntryPoint1 returns the union data inside the GroundingMetadata_SearchEntryPoint as a GroundingMetadataSearchEntryPoint1
func (t GroundingMetadata_SearchEntryPoint) AsGroundingMetadataSearchEntryPoint1() (GroundingMetadataSearchEntryPoint1, error) {
	var body GroundingMetadataSearchEntryPoint1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataSearchEntryPoint1 overwrites any union data inside the GroundingMetadata_SearchEntryPoint as the provided GroundingMetadataSearchEntryPoint1
func (t *GroundingMetadata_SearchEntryPoint) FromGroundingMetadataSearchEntryPoint1(v GroundingMetadataSearchEntryPoint1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataSearchEntryPoint1 performs a merge with any union data inside the GroundingMetadata_SearchEntryPoint, using the provided GroundingMetadataSearchEntryPoint1
func (t *GroundingMetadata_SearchEntryPoint) MergeGroundingMetadataSearchEntryPoint1(v GroundingMetadataSearchEntryPoint1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingMetadata_SearchEntryPoint) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingMetadata_SearchEntryPoint) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingMetadataWebSearchQueries0 returns the union data inside the GroundingMetadata_WebSearchQueries as a GroundingMetadataWebSearchQueries0
func (t GroundingMetadata_WebSearchQueries) AsGroundingMetadataWebSearchQueries0() (GroundingMetadataWebSearchQueries0, error) {
	var body GroundingMetadataWebSearchQueries0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataWebSearchQueries0 overwrites any union data inside the GroundingMetadata_WebSearchQueries as the provided GroundingMetadataWebSearchQueries0
func (t *GroundingMetadata_WebSearchQueries) FromGroundingMetadataWebSearchQueries0(v GroundingMetadataWebSearchQueries0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataWebSearchQueries0 performs a merge with any union data inside the GroundingMetadata_WebSearchQueries, using the provided GroundingMetadataWebSearchQueries0
func (t *GroundingMetadata_WebSearchQueries) MergeGroundingMetadataWebSearchQueries0(v GroundingMetadataWebSearchQueries0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingMetadataWebSearchQueries1 returns the union data inside the GroundingMetadata_WebSearchQueries as a GroundingMetadataWebSearchQueries1
func (t GroundingMetadata_WebSearchQueries) AsGroundingMetadataWebSearchQueries1() (GroundingMetadataWebSearchQueries1, error) {
	var body GroundingMetadataWebSearchQueries1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingMetadataWebSearchQueries1 overwrites any union data inside the GroundingMetadata_WebSearchQueries as the provided GroundingMetadataWebSearchQueries1
func (t *GroundingMetadata_WebSearchQueries) FromGroundingMetadataWebSearchQueries1(v GroundingMetadataWebSearchQueries1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingMetadataWebSearchQueries1 performs a merge with any union data inside the GroundingMetadata_WebSearchQueries, using the provided GroundingMetadataWebSearchQueries1
func (t *GroundingMetadata_WebSearchQueries) MergeGroundingMetadataWebSearchQueries1(v GroundingMetadataWebSearchQueries1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingMetadata_WebSearchQueries) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingMetadata_WebSearchQueries) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingSupportConfidenceScores0 returns the union data inside the GroundingSupport_ConfidenceScores as a GroundingSupportConfidenceScores0
func (t GroundingSupport_ConfidenceScores) AsGroundingSupportConfidenceScores0() (GroundingSupportConfidenceScores0, error) {
	var body GroundingSupportConfidenceScores0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingSupportConfidenceScores0 overwrites any union data inside the GroundingSupport_ConfidenceScores as the provided GroundingSupportConfidenceScores0
func (t *GroundingSupport_ConfidenceScores) FromGroundingSupportConfidenceScores0(v GroundingSupportConfidenceScores0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingSupportConfidenceScores0 performs a merge with any union data inside the GroundingSupport_ConfidenceScores, using the provided GroundingSupportConfidenceScores0
func (t *GroundingSupport_ConfidenceScores) MergeGroundingSupportConfidenceScores0(v GroundingSupportConfidenceScores0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingSupportConfidenceScores1 returns the union data inside the GroundingSupport_ConfidenceScores as a GroundingSupportConfidenceScores1
func (t GroundingSupport_ConfidenceScores) AsGroundingSupportConfidenceScores1() (GroundingSupportConfidenceScores1, error) {
	var body GroundingSupportConfidenceScores1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingSupportConfidenceScores1 overwrites any union data inside the GroundingSupport_ConfidenceScores as the provided GroundingSupportConfidenceScores1
func (t *GroundingSupport_ConfidenceScores) FromGroundingSupportConfidenceScores1(v GroundingSupportConfidenceScores1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingSupportConfidenceScores1 performs a merge with any union data inside the GroundingSupport_ConfidenceScores, using the provided GroundingSupportConfidenceScores1
func (t *GroundingSupport_ConfidenceScores) MergeGroundingSupportConfidenceScores1(v GroundingSupportConfidenceScores1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingSupport_ConfidenceScores) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingSupport_ConfidenceScores) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsGroundingSupportGroundingChunkIndices0 returns the union data inside the GroundingSupport_GroundingChunkIndices as a GroundingSupportGroundingChunkIndices0
func (t GroundingSupport_GroundingChunkIndices) AsGroundingSupportGroundingChunkIndices0() (GroundingSupportGroundingChunkIndices0, error) {
	var body GroundingSupportGroundingChunkIndices0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingSupportGroundingChunkIndices0 overwrites any union data inside the GroundingSupport_GroundingChunkIndices as the provided GroundingSupportGroundingChunkIndices0
func (t *GroundingSupport_GroundingChunkIndices) FromGroundingSupportGroundingChunkIndices0(v GroundingSupportGroundingChunkIndices0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingSupportGroundingChunkIndices0 performs a merge with any union data inside the GroundingSupport_GroundingChunkIndices, using the provided GroundingSupportGroundingChunkIndices0
func (t *GroundingSupport_GroundingChunkIndices) MergeGroundingSupportGroundingChunkIndices0(v GroundingSupportGroundingChunkIndices0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingSupportGroundingChunkIndices1 returns the union data inside the GroundingSupport_GroundingChunkIndices as a GroundingSupportGroundingChunkIndices1
func (t GroundingSupport_GroundingChunkIndices) AsGroundingSupportGroundingChunkIndices1() (GroundingSupportGroundingChunkIndices1, error) {
	var body GroundingSupportGroundingChunkIndices1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingSupportGroundingChunkIndices1 overwrites any union data inside the GroundingSupport_GroundingChunkIndices as the provided GroundingSupportGroundingChunkIndices1
func (t *GroundingSupport_GroundingChunkIndices) FromGroundingSupportGroundingChunkIndices1(v GroundingSupportGroundingChunkIndices1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingSupportGroundingChunkIndices1 performs a merge with any union data inside the GroundingSupport_GroundingChunkIndices, using the provided GroundingSupportGroundingChunkIndices1
func (t *GroundingSupport_GroundingChunkIndices) MergeGroundingSupportGroundingChunkIndices1(v GroundingSupportGroundingChunkIndices1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingSupport_GroundingChunkIndices) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingSupport_GroundingChunkIndices) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSegment returns the union data inside the GroundingSupport_Segment as a Segment
func (t GroundingSupport_Segment) AsSegment() (Segment, error) {
	var body Segment
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegment overwrites any union data inside the GroundingSupport_Segment as the provided Segment
func (t *GroundingSupport_Segment) FromSegment(v Segment) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegment performs a merge with any union data inside the GroundingSupport_Segment, using the provided Segment
func (t *GroundingSupport_Segment) MergeSegment(v Segment) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsGroundingSupportSegment1 returns the union data inside the GroundingSupport_Segment as a GroundingSupportSegment1
func (t GroundingSupport_Segment) AsGroundingSupportSegment1() (GroundingSupportSegment1, error) {
	var body GroundingSupportSegment1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromGroundingSupportSegment1 overwrites any union data inside the GroundingSupport_Segment as the provided GroundingSupportSegment1
func (t *GroundingSupport_Segment) FromGroundingSupportSegment1(v GroundingSupportSegment1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeGroundingSupportSegment1 performs a merge with any union data inside the GroundingSupport_Segment, using the provided GroundingSupportSegment1
func (t *GroundingSupport_Segment) MergeGroundingSupportSegment1(v GroundingSupportSegment1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t GroundingSupport_Segment) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *GroundingSupport_Segment) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsHTTPBaseDescription0 returns the union data inside the HTTPBase_Description as a HTTPBaseDescription0
func (t HTTPBase_Description) AsHTTPBaseDescription0() (HTTPBaseDescription0, error) {
	var body HTTPBaseDescription0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBaseDescription0 overwrites any union data inside the HTTPBase_Description as the provided HTTPBaseDescription0
func (t *HTTPBase_Description) FromHTTPBaseDescription0(v HTTPBaseDescription0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBaseDescription0 performs a merge with any union data inside the HTTPBase_Description, using the provided HTTPBaseDescription0
func (t *HTTPBase_Description) MergeHTTPBaseDescription0(v HTTPBaseDescription0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHTTPBaseDescription1 returns the union data inside the HTTPBase_Description as a HTTPBaseDescription1
func (t HTTPBase_Description) AsHTTPBaseDescription1() (HTTPBaseDescription1, error) {
	var body HTTPBaseDescription1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBaseDescription1 overwrites any union data inside the HTTPBase_Description as the provided HTTPBaseDescription1
func (t *HTTPBase_Description) FromHTTPBaseDescription1(v HTTPBaseDescription1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBaseDescription1 performs a merge with any union data inside the HTTPBase_Description, using the provided HTTPBaseDescription1
func (t *HTTPBase_Description) MergeHTTPBaseDescription1(v HTTPBaseDescription1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t HTTPBase_Description) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *HTTPBase_Description) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsHTTPBearerBearerFormat0 returns the union data inside the HTTPBearer_BearerFormat as a HTTPBearerBearerFormat0
func (t HTTPBearer_BearerFormat) AsHTTPBearerBearerFormat0() (HTTPBearerBearerFormat0, error) {
	var body HTTPBearerBearerFormat0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBearerBearerFormat0 overwrites any union data inside the HTTPBearer_BearerFormat as the provided HTTPBearerBearerFormat0
func (t *HTTPBearer_BearerFormat) FromHTTPBearerBearerFormat0(v HTTPBearerBearerFormat0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBearerBearerFormat0 performs a merge with any union data inside the HTTPBearer_BearerFormat, using the provided HTTPBearerBearerFormat0
func (t *HTTPBearer_BearerFormat) MergeHTTPBearerBearerFormat0(v HTTPBearerBearerFormat0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHTTPBearerBearerFormat1 returns the union data inside the HTTPBearer_BearerFormat as a HTTPBearerBearerFormat1
func (t HTTPBearer_BearerFormat) AsHTTPBearerBearerFormat1() (HTTPBearerBearerFormat1, error) {
	var body HTTPBearerBearerFormat1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBearerBearerFormat1 overwrites any union data inside the HTTPBearer_BearerFormat as the provided HTTPBearerBearerFormat1
func (t *HTTPBearer_BearerFormat) FromHTTPBearerBearerFormat1(v HTTPBearerBearerFormat1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBearerBearerFormat1 performs a merge with any union data inside the HTTPBearer_BearerFormat, using the provided HTTPBearerBearerFormat1
func (t *HTTPBearer_BearerFormat) MergeHTTPBearerBearerFormat1(v HTTPBearerBearerFormat1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t HTTPBearer_BearerFormat) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *HTTPBearer_BearerFormat) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsHTTPBearerDescription0 returns the union data inside the HTTPBearer_Description as a HTTPBearerDescription0
func (t HTTPBearer_Description) AsHTTPBearerDescription0() (HTTPBearerDescription0, error) {
	var body HTTPBearerDescription0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBearerDescription0 overwrites any union data inside the HTTPBearer_Description as the provided HTTPBearerDescription0
func (t *HTTPBearer_Description) FromHTTPBearerDescription0(v HTTPBearerDescription0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBearerDescription0 performs a merge with any union data inside the HTTPBearer_Description, using the provided HTTPBearerDescription0
func (t *HTTPBearer_Description) MergeHTTPBearerDescription0(v HTTPBearerDescription0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHTTPBearerDescription1 returns the union data inside the HTTPBearer_Description as a HTTPBearerDescription1
func (t HTTPBearer_Description) AsHTTPBearerDescription1() (HTTPBearerDescription1, error) {
	var body HTTPBearerDescription1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHTTPBearerDescription1 overwrites any union data inside the HTTPBearer_Description as the provided HTTPBearerDescription1
func (t *HTTPBearer_Description) FromHTTPBearerDescription1(v HTTPBearerDescription1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHTTPBearerDescription1 performs a merge with any union data inside the HTTPBearer_Description, using the provided HTTPBearerDescription1
func (t *HTTPBearer_Description) MergeHTTPBearerDescription1(v HTTPBearerDescription1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t HTTPBearer_Description) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *HTTPBearer_Description) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsHttpCredentialsPassword0 returns the union data inside the HttpCredentials_Password as a HttpCredentialsPassword0
func (t HttpCredentials_Password) AsHttpCredentialsPassword0() (HttpCredentialsPassword0, error) {
	var body HttpCredentialsPassword0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHttpCredentialsPassword0 overwrites any union data inside the HttpCredentials_Password as the provided HttpCredentialsPassword0
func (t *HttpCredentials_Password) FromHttpCredentialsPassword0(v HttpCredentialsPassword0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHttpCredentialsPassword0 performs a merge with any union data inside the HttpCredentials_Password, using the provided HttpCredentialsPassword0
func (t *HttpCredentials_Password) MergeHttpCredentialsPassword0(v HttpCredentialsPassword0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHttpCredentialsPassword1 returns the union data inside the HttpCredentials_Password as a HttpCredentialsPassword1
func (t HttpCredentials_Password) AsHttpCredentialsPassword1() (HttpCredentialsPassword1, error) {
	var body HttpCredentialsPassword1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHttpCredentialsPassword1 overwrites any union data inside the HttpCredentials_Password as the provided HttpCredentialsPassword1
func (t *HttpCredentials_Password) FromHttpCredentialsPassword1(v HttpCredentialsPassword1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHttpCredentialsPassword1 performs a merge with any union data inside the HttpCredentials_Password, using the provided HttpCredentialsPassword1
func (t *HttpCredentials_Password) MergeHttpCredentialsPassword1(v HttpCredentialsPassword1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t HttpCredentials_Password) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *HttpCredentials_Password) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsHttpCredentialsToken0 returns the union data inside the HttpCredentials_Token as a HttpCredentialsToken0
func (t HttpCredentials_Token) AsHttpCredentialsToken0() (HttpCredentialsToken0, error) {
	var body HttpCredentialsToken0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHttpCredentialsToken0 overwrites any union data inside the HttpCredentials_Token as the provided HttpCredentialsToken0
func (t *HttpCredentials_Token) FromHttpCredentialsToken0(v HttpCredentialsToken0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHttpCredentialsToken0 performs a merge with any union data inside the HttpCredentials_Token, using the provided HttpCredentialsToken0
func (t *HttpCredentials_Token) MergeHttpCredentialsToken0(v HttpCredentialsToken0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHttpCredentialsToken1 returns the union data inside the HttpCredentials_Token as a HttpCredentialsToken1
func (t HttpCredentials_Token) AsHttpCredentialsToken1() (HttpCredentialsToken1, error) {
	var body HttpCredentialsToken1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHttpCredentialsToken1 overwrites any union data inside the HttpCredentials_Token as the provided HttpCredentialsToken1
func (t *HttpCredentials_Token) FromHttpCredentialsToken1(v HttpCredentialsToken1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHttpCredentialsToken1 performs a merge with any union data inside the HttpCredentials_Token, using the provided HttpCredentialsToken1
func (t *HttpCredentials_Token) MergeHttpCredentialsToken1(v HttpCredentialsToken1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t HttpCredentials_Token) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *HttpCredentials_Token) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsHttpCredentialsUsername0 returns the union data inside the HttpCredentials_Username as a HttpCredentialsUsername0
func (t HttpCredentials_Username) AsHttpCredentialsUsername0() (HttpCredentialsUsername0, error) {
	var body HttpCredentialsUsername0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHttpCredentialsUsername0 overwrites any union data inside the HttpCredentials_Username as the provided HttpCredentialsUsername0
func (t *HttpCredentials_Username) FromHttpCredentialsUsername0(v HttpCredentialsUsername0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHttpCredentialsUsername0 performs a merge with any union data inside the HttpCredentials_Username, using the provided HttpCredentialsUsername0
func (t *HttpCredentials_Username) MergeHttpCredentialsUsername0(v HttpCredentialsUsername0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsHttpCredentialsUsername1 returns the union data inside the HttpCredentials_Username as a HttpCredentialsUsername1
func (t HttpCredentials_Username) AsHttpCredentialsUsername1() (HttpCredentialsUsername1, error) {
	var body HttpCredentialsUsername1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromHttpCredentialsUsername1 overwrites any union data inside the HttpCredentials_Username as the provided HttpCredentialsUsername1
func (t *HttpCredentials_Username) FromHttpCredentialsUsername1(v HttpCredentialsUsername1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeHttpCredentialsUsername1 performs a merge with any union data inside the HttpCredentials_Username, using the provided HttpCredentialsUsername1
func (t *HttpCredentials_Username) MergeHttpCredentialsUsername1(v HttpCredentialsUsername1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t HttpCredentials_Username) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *HttpCredentials_Username) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2Description0 returns the union data inside the OAuth2_Description as a OAuth2Description0
func (t OAuth2_Description) AsOAuth2Description0() (OAuth2Description0, error) {
	var body OAuth2Description0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2Description0 overwrites any union data inside the OAuth2_Description as the provided OAuth2Description0
func (t *OAuth2_Description) FromOAuth2Description0(v OAuth2Description0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2Description0 performs a merge with any union data inside the OAuth2_Description, using the provided OAuth2Description0
func (t *OAuth2_Description) MergeOAuth2Description0(v OAuth2Description0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2Description1 returns the union data inside the OAuth2_Description as a OAuth2Description1
func (t OAuth2_Description) AsOAuth2Description1() (OAuth2Description1, error) {
	var body OAuth2Description1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2Description1 overwrites any union data inside the OAuth2_Description as the provided OAuth2Description1
func (t *OAuth2_Description) FromOAuth2Description1(v OAuth2Description1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2Description1 performs a merge with any union data inside the OAuth2_Description, using the provided OAuth2Description1
func (t *OAuth2_Description) MergeOAuth2Description1(v OAuth2Description1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2_Description) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2_Description) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthAuthCode0 returns the union data inside the OAuth2Auth_AuthCode as a OAuth2AuthAuthCode0
func (t OAuth2Auth_AuthCode) AsOAuth2AuthAuthCode0() (OAuth2AuthAuthCode0, error) {
	var body OAuth2AuthAuthCode0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthAuthCode0 overwrites any union data inside the OAuth2Auth_AuthCode as the provided OAuth2AuthAuthCode0
func (t *OAuth2Auth_AuthCode) FromOAuth2AuthAuthCode0(v OAuth2AuthAuthCode0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthAuthCode0 performs a merge with any union data inside the OAuth2Auth_AuthCode, using the provided OAuth2AuthAuthCode0
func (t *OAuth2Auth_AuthCode) MergeOAuth2AuthAuthCode0(v OAuth2AuthAuthCode0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthAuthCode1 returns the union data inside the OAuth2Auth_AuthCode as a OAuth2AuthAuthCode1
func (t OAuth2Auth_AuthCode) AsOAuth2AuthAuthCode1() (OAuth2AuthAuthCode1, error) {
	var body OAuth2AuthAuthCode1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthAuthCode1 overwrites any union data inside the OAuth2Auth_AuthCode as the provided OAuth2AuthAuthCode1
func (t *OAuth2Auth_AuthCode) FromOAuth2AuthAuthCode1(v OAuth2AuthAuthCode1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthAuthCode1 performs a merge with any union data inside the OAuth2Auth_AuthCode, using the provided OAuth2AuthAuthCode1
func (t *OAuth2Auth_AuthCode) MergeOAuth2AuthAuthCode1(v OAuth2AuthAuthCode1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_AuthCode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_AuthCode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthAuthResponseUri0 returns the union data inside the OAuth2Auth_AuthResponseUri as a OAuth2AuthAuthResponseUri0
func (t OAuth2Auth_AuthResponseUri) AsOAuth2AuthAuthResponseUri0() (OAuth2AuthAuthResponseUri0, error) {
	var body OAuth2AuthAuthResponseUri0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthAuthResponseUri0 overwrites any union data inside the OAuth2Auth_AuthResponseUri as the provided OAuth2AuthAuthResponseUri0
func (t *OAuth2Auth_AuthResponseUri) FromOAuth2AuthAuthResponseUri0(v OAuth2AuthAuthResponseUri0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthAuthResponseUri0 performs a merge with any union data inside the OAuth2Auth_AuthResponseUri, using the provided OAuth2AuthAuthResponseUri0
func (t *OAuth2Auth_AuthResponseUri) MergeOAuth2AuthAuthResponseUri0(v OAuth2AuthAuthResponseUri0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthAuthResponseUri1 returns the union data inside the OAuth2Auth_AuthResponseUri as a OAuth2AuthAuthResponseUri1
func (t OAuth2Auth_AuthResponseUri) AsOAuth2AuthAuthResponseUri1() (OAuth2AuthAuthResponseUri1, error) {
	var body OAuth2AuthAuthResponseUri1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthAuthResponseUri1 overwrites any union data inside the OAuth2Auth_AuthResponseUri as the provided OAuth2AuthAuthResponseUri1
func (t *OAuth2Auth_AuthResponseUri) FromOAuth2AuthAuthResponseUri1(v OAuth2AuthAuthResponseUri1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthAuthResponseUri1 performs a merge with any union data inside the OAuth2Auth_AuthResponseUri, using the provided OAuth2AuthAuthResponseUri1
func (t *OAuth2Auth_AuthResponseUri) MergeOAuth2AuthAuthResponseUri1(v OAuth2AuthAuthResponseUri1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_AuthResponseUri) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_AuthResponseUri) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthAuthUri0 returns the union data inside the OAuth2Auth_AuthUri as a OAuth2AuthAuthUri0
func (t OAuth2Auth_AuthUri) AsOAuth2AuthAuthUri0() (OAuth2AuthAuthUri0, error) {
	var body OAuth2AuthAuthUri0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthAuthUri0 overwrites any union data inside the OAuth2Auth_AuthUri as the provided OAuth2AuthAuthUri0
func (t *OAuth2Auth_AuthUri) FromOAuth2AuthAuthUri0(v OAuth2AuthAuthUri0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthAuthUri0 performs a merge with any union data inside the OAuth2Auth_AuthUri, using the provided OAuth2AuthAuthUri0
func (t *OAuth2Auth_AuthUri) MergeOAuth2AuthAuthUri0(v OAuth2AuthAuthUri0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthAuthUri1 returns the union data inside the OAuth2Auth_AuthUri as a OAuth2AuthAuthUri1
func (t OAuth2Auth_AuthUri) AsOAuth2AuthAuthUri1() (OAuth2AuthAuthUri1, error) {
	var body OAuth2AuthAuthUri1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthAuthUri1 overwrites any union data inside the OAuth2Auth_AuthUri as the provided OAuth2AuthAuthUri1
func (t *OAuth2Auth_AuthUri) FromOAuth2AuthAuthUri1(v OAuth2AuthAuthUri1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthAuthUri1 performs a merge with any union data inside the OAuth2Auth_AuthUri, using the provided OAuth2AuthAuthUri1
func (t *OAuth2Auth_AuthUri) MergeOAuth2AuthAuthUri1(v OAuth2AuthAuthUri1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_AuthUri) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_AuthUri) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthClientId0 returns the union data inside the OAuth2Auth_ClientId as a OAuth2AuthClientId0
func (t OAuth2Auth_ClientId) AsOAuth2AuthClientId0() (OAuth2AuthClientId0, error) {
	var body OAuth2AuthClientId0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthClientId0 overwrites any union data inside the OAuth2Auth_ClientId as the provided OAuth2AuthClientId0
func (t *OAuth2Auth_ClientId) FromOAuth2AuthClientId0(v OAuth2AuthClientId0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthClientId0 performs a merge with any union data inside the OAuth2Auth_ClientId, using the provided OAuth2AuthClientId0
func (t *OAuth2Auth_ClientId) MergeOAuth2AuthClientId0(v OAuth2AuthClientId0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthClientId1 returns the union data inside the OAuth2Auth_ClientId as a OAuth2AuthClientId1
func (t OAuth2Auth_ClientId) AsOAuth2AuthClientId1() (OAuth2AuthClientId1, error) {
	var body OAuth2AuthClientId1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthClientId1 overwrites any union data inside the OAuth2Auth_ClientId as the provided OAuth2AuthClientId1
func (t *OAuth2Auth_ClientId) FromOAuth2AuthClientId1(v OAuth2AuthClientId1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthClientId1 performs a merge with any union data inside the OAuth2Auth_ClientId, using the provided OAuth2AuthClientId1
func (t *OAuth2Auth_ClientId) MergeOAuth2AuthClientId1(v OAuth2AuthClientId1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_ClientId) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_ClientId) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthClientSecret0 returns the union data inside the OAuth2Auth_ClientSecret as a OAuth2AuthClientSecret0
func (t OAuth2Auth_ClientSecret) AsOAuth2AuthClientSecret0() (OAuth2AuthClientSecret0, error) {
	var body OAuth2AuthClientSecret0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthClientSecret0 overwrites any union data inside the OAuth2Auth_ClientSecret as the provided OAuth2AuthClientSecret0
func (t *OAuth2Auth_ClientSecret) FromOAuth2AuthClientSecret0(v OAuth2AuthClientSecret0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthClientSecret0 performs a merge with any union data inside the OAuth2Auth_ClientSecret, using the provided OAuth2AuthClientSecret0
func (t *OAuth2Auth_ClientSecret) MergeOAuth2AuthClientSecret0(v OAuth2AuthClientSecret0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthClientSecret1 returns the union data inside the OAuth2Auth_ClientSecret as a OAuth2AuthClientSecret1
func (t OAuth2Auth_ClientSecret) AsOAuth2AuthClientSecret1() (OAuth2AuthClientSecret1, error) {
	var body OAuth2AuthClientSecret1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthClientSecret1 overwrites any union data inside the OAuth2Auth_ClientSecret as the provided OAuth2AuthClientSecret1
func (t *OAuth2Auth_ClientSecret) FromOAuth2AuthClientSecret1(v OAuth2AuthClientSecret1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthClientSecret1 performs a merge with any union data inside the OAuth2Auth_ClientSecret, using the provided OAuth2AuthClientSecret1
func (t *OAuth2Auth_ClientSecret) MergeOAuth2AuthClientSecret1(v OAuth2AuthClientSecret1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_ClientSecret) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_ClientSecret) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthRedirectUri0 returns the union data inside the OAuth2Auth_RedirectUri as a OAuth2AuthRedirectUri0
func (t OAuth2Auth_RedirectUri) AsOAuth2AuthRedirectUri0() (OAuth2AuthRedirectUri0, error) {
	var body OAuth2AuthRedirectUri0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthRedirectUri0 overwrites any union data inside the OAuth2Auth_RedirectUri as the provided OAuth2AuthRedirectUri0
func (t *OAuth2Auth_RedirectUri) FromOAuth2AuthRedirectUri0(v OAuth2AuthRedirectUri0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthRedirectUri0 performs a merge with any union data inside the OAuth2Auth_RedirectUri, using the provided OAuth2AuthRedirectUri0
func (t *OAuth2Auth_RedirectUri) MergeOAuth2AuthRedirectUri0(v OAuth2AuthRedirectUri0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthRedirectUri1 returns the union data inside the OAuth2Auth_RedirectUri as a OAuth2AuthRedirectUri1
func (t OAuth2Auth_RedirectUri) AsOAuth2AuthRedirectUri1() (OAuth2AuthRedirectUri1, error) {
	var body OAuth2AuthRedirectUri1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthRedirectUri1 overwrites any union data inside the OAuth2Auth_RedirectUri as the provided OAuth2AuthRedirectUri1
func (t *OAuth2Auth_RedirectUri) FromOAuth2AuthRedirectUri1(v OAuth2AuthRedirectUri1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthRedirectUri1 performs a merge with any union data inside the OAuth2Auth_RedirectUri, using the provided OAuth2AuthRedirectUri1
func (t *OAuth2Auth_RedirectUri) MergeOAuth2AuthRedirectUri1(v OAuth2AuthRedirectUri1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_RedirectUri) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_RedirectUri) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthState0 returns the union data inside the OAuth2Auth_State as a OAuth2AuthState0
func (t OAuth2Auth_State) AsOAuth2AuthState0() (OAuth2AuthState0, error) {
	var body OAuth2AuthState0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthState0 overwrites any union data inside the OAuth2Auth_State as the provided OAuth2AuthState0
func (t *OAuth2Auth_State) FromOAuth2AuthState0(v OAuth2AuthState0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthState0 performs a merge with any union data inside the OAuth2Auth_State, using the provided OAuth2AuthState0
func (t *OAuth2Auth_State) MergeOAuth2AuthState0(v OAuth2AuthState0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthState1 returns the union data inside the OAuth2Auth_State as a OAuth2AuthState1
func (t OAuth2Auth_State) AsOAuth2AuthState1() (OAuth2AuthState1, error) {
	var body OAuth2AuthState1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthState1 overwrites any union data inside the OAuth2Auth_State as the provided OAuth2AuthState1
func (t *OAuth2Auth_State) FromOAuth2AuthState1(v OAuth2AuthState1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthState1 performs a merge with any union data inside the OAuth2Auth_State, using the provided OAuth2AuthState1
func (t *OAuth2Auth_State) MergeOAuth2AuthState1(v OAuth2AuthState1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_State) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_State) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuth2AuthToken0 returns the union data inside the OAuth2Auth_Token as a OAuth2AuthToken0
func (t OAuth2Auth_Token) AsOAuth2AuthToken0() (OAuth2AuthToken0, error) {
	var body OAuth2AuthToken0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthToken0 overwrites any union data inside the OAuth2Auth_Token as the provided OAuth2AuthToken0
func (t *OAuth2Auth_Token) FromOAuth2AuthToken0(v OAuth2AuthToken0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthToken0 performs a merge with any union data inside the OAuth2Auth_Token, using the provided OAuth2AuthToken0
func (t *OAuth2Auth_Token) MergeOAuth2AuthToken0(v OAuth2AuthToken0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuth2AuthToken1 returns the union data inside the OAuth2Auth_Token as a OAuth2AuthToken1
func (t OAuth2Auth_Token) AsOAuth2AuthToken1() (OAuth2AuthToken1, error) {
	var body OAuth2AuthToken1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuth2AuthToken1 overwrites any union data inside the OAuth2Auth_Token as the provided OAuth2AuthToken1
func (t *OAuth2Auth_Token) FromOAuth2AuthToken1(v OAuth2AuthToken1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuth2AuthToken1 performs a merge with any union data inside the OAuth2Auth_Token, using the provided OAuth2AuthToken1
func (t *OAuth2Auth_Token) MergeOAuth2AuthToken1(v OAuth2AuthToken1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuth2Auth_Token) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuth2Auth_Token) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowAuthorizationCodeRefreshUrl0 returns the union data inside the OAuthFlowAuthorizationCode_RefreshUrl as a OAuthFlowAuthorizationCodeRefreshUrl0
func (t OAuthFlowAuthorizationCode_RefreshUrl) AsOAuthFlowAuthorizationCodeRefreshUrl0() (OAuthFlowAuthorizationCodeRefreshUrl0, error) {
	var body OAuthFlowAuthorizationCodeRefreshUrl0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowAuthorizationCodeRefreshUrl0 overwrites any union data inside the OAuthFlowAuthorizationCode_RefreshUrl as the provided OAuthFlowAuthorizationCodeRefreshUrl0
func (t *OAuthFlowAuthorizationCode_RefreshUrl) FromOAuthFlowAuthorizationCodeRefreshUrl0(v OAuthFlowAuthorizationCodeRefreshUrl0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowAuthorizationCodeRefreshUrl0 performs a merge with any union data inside the OAuthFlowAuthorizationCode_RefreshUrl, using the provided OAuthFlowAuthorizationCodeRefreshUrl0
func (t *OAuthFlowAuthorizationCode_RefreshUrl) MergeOAuthFlowAuthorizationCodeRefreshUrl0(v OAuthFlowAuthorizationCodeRefreshUrl0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowAuthorizationCodeRefreshUrl1 returns the union data inside the OAuthFlowAuthorizationCode_RefreshUrl as a OAuthFlowAuthorizationCodeRefreshUrl1
func (t OAuthFlowAuthorizationCode_RefreshUrl) AsOAuthFlowAuthorizationCodeRefreshUrl1() (OAuthFlowAuthorizationCodeRefreshUrl1, error) {
	var body OAuthFlowAuthorizationCodeRefreshUrl1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowAuthorizationCodeRefreshUrl1 overwrites any union data inside the OAuthFlowAuthorizationCode_RefreshUrl as the provided OAuthFlowAuthorizationCodeRefreshUrl1
func (t *OAuthFlowAuthorizationCode_RefreshUrl) FromOAuthFlowAuthorizationCodeRefreshUrl1(v OAuthFlowAuthorizationCodeRefreshUrl1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowAuthorizationCodeRefreshUrl1 performs a merge with any union data inside the OAuthFlowAuthorizationCode_RefreshUrl, using the provided OAuthFlowAuthorizationCodeRefreshUrl1
func (t *OAuthFlowAuthorizationCode_RefreshUrl) MergeOAuthFlowAuthorizationCodeRefreshUrl1(v OAuthFlowAuthorizationCodeRefreshUrl1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlowAuthorizationCode_RefreshUrl) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlowAuthorizationCode_RefreshUrl) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowClientCredentialsRefreshUrl0 returns the union data inside the OAuthFlowClientCredentials_RefreshUrl as a OAuthFlowClientCredentialsRefreshUrl0
func (t OAuthFlowClientCredentials_RefreshUrl) AsOAuthFlowClientCredentialsRefreshUrl0() (OAuthFlowClientCredentialsRefreshUrl0, error) {
	var body OAuthFlowClientCredentialsRefreshUrl0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowClientCredentialsRefreshUrl0 overwrites any union data inside the OAuthFlowClientCredentials_RefreshUrl as the provided OAuthFlowClientCredentialsRefreshUrl0
func (t *OAuthFlowClientCredentials_RefreshUrl) FromOAuthFlowClientCredentialsRefreshUrl0(v OAuthFlowClientCredentialsRefreshUrl0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowClientCredentialsRefreshUrl0 performs a merge with any union data inside the OAuthFlowClientCredentials_RefreshUrl, using the provided OAuthFlowClientCredentialsRefreshUrl0
func (t *OAuthFlowClientCredentials_RefreshUrl) MergeOAuthFlowClientCredentialsRefreshUrl0(v OAuthFlowClientCredentialsRefreshUrl0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowClientCredentialsRefreshUrl1 returns the union data inside the OAuthFlowClientCredentials_RefreshUrl as a OAuthFlowClientCredentialsRefreshUrl1
func (t OAuthFlowClientCredentials_RefreshUrl) AsOAuthFlowClientCredentialsRefreshUrl1() (OAuthFlowClientCredentialsRefreshUrl1, error) {
	var body OAuthFlowClientCredentialsRefreshUrl1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowClientCredentialsRefreshUrl1 overwrites any union data inside the OAuthFlowClientCredentials_RefreshUrl as the provided OAuthFlowClientCredentialsRefreshUrl1
func (t *OAuthFlowClientCredentials_RefreshUrl) FromOAuthFlowClientCredentialsRefreshUrl1(v OAuthFlowClientCredentialsRefreshUrl1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowClientCredentialsRefreshUrl1 performs a merge with any union data inside the OAuthFlowClientCredentials_RefreshUrl, using the provided OAuthFlowClientCredentialsRefreshUrl1
func (t *OAuthFlowClientCredentials_RefreshUrl) MergeOAuthFlowClientCredentialsRefreshUrl1(v OAuthFlowClientCredentialsRefreshUrl1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlowClientCredentials_RefreshUrl) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlowClientCredentials_RefreshUrl) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowImplicitRefreshUrl0 returns the union data inside the OAuthFlowImplicit_RefreshUrl as a OAuthFlowImplicitRefreshUrl0
func (t OAuthFlowImplicit_RefreshUrl) AsOAuthFlowImplicitRefreshUrl0() (OAuthFlowImplicitRefreshUrl0, error) {
	var body OAuthFlowImplicitRefreshUrl0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowImplicitRefreshUrl0 overwrites any union data inside the OAuthFlowImplicit_RefreshUrl as the provided OAuthFlowImplicitRefreshUrl0
func (t *OAuthFlowImplicit_RefreshUrl) FromOAuthFlowImplicitRefreshUrl0(v OAuthFlowImplicitRefreshUrl0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowImplicitRefreshUrl0 performs a merge with any union data inside the OAuthFlowImplicit_RefreshUrl, using the provided OAuthFlowImplicitRefreshUrl0
func (t *OAuthFlowImplicit_RefreshUrl) MergeOAuthFlowImplicitRefreshUrl0(v OAuthFlowImplicitRefreshUrl0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowImplicitRefreshUrl1 returns the union data inside the OAuthFlowImplicit_RefreshUrl as a OAuthFlowImplicitRefreshUrl1
func (t OAuthFlowImplicit_RefreshUrl) AsOAuthFlowImplicitRefreshUrl1() (OAuthFlowImplicitRefreshUrl1, error) {
	var body OAuthFlowImplicitRefreshUrl1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowImplicitRefreshUrl1 overwrites any union data inside the OAuthFlowImplicit_RefreshUrl as the provided OAuthFlowImplicitRefreshUrl1
func (t *OAuthFlowImplicit_RefreshUrl) FromOAuthFlowImplicitRefreshUrl1(v OAuthFlowImplicitRefreshUrl1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowImplicitRefreshUrl1 performs a merge with any union data inside the OAuthFlowImplicit_RefreshUrl, using the provided OAuthFlowImplicitRefreshUrl1
func (t *OAuthFlowImplicit_RefreshUrl) MergeOAuthFlowImplicitRefreshUrl1(v OAuthFlowImplicitRefreshUrl1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlowImplicit_RefreshUrl) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlowImplicit_RefreshUrl) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowPasswordRefreshUrl0 returns the union data inside the OAuthFlowPassword_RefreshUrl as a OAuthFlowPasswordRefreshUrl0
func (t OAuthFlowPassword_RefreshUrl) AsOAuthFlowPasswordRefreshUrl0() (OAuthFlowPasswordRefreshUrl0, error) {
	var body OAuthFlowPasswordRefreshUrl0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowPasswordRefreshUrl0 overwrites any union data inside the OAuthFlowPassword_RefreshUrl as the provided OAuthFlowPasswordRefreshUrl0
func (t *OAuthFlowPassword_RefreshUrl) FromOAuthFlowPasswordRefreshUrl0(v OAuthFlowPasswordRefreshUrl0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowPasswordRefreshUrl0 performs a merge with any union data inside the OAuthFlowPassword_RefreshUrl, using the provided OAuthFlowPasswordRefreshUrl0
func (t *OAuthFlowPassword_RefreshUrl) MergeOAuthFlowPasswordRefreshUrl0(v OAuthFlowPasswordRefreshUrl0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowPasswordRefreshUrl1 returns the union data inside the OAuthFlowPassword_RefreshUrl as a OAuthFlowPasswordRefreshUrl1
func (t OAuthFlowPassword_RefreshUrl) AsOAuthFlowPasswordRefreshUrl1() (OAuthFlowPasswordRefreshUrl1, error) {
	var body OAuthFlowPasswordRefreshUrl1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowPasswordRefreshUrl1 overwrites any union data inside the OAuthFlowPassword_RefreshUrl as the provided OAuthFlowPasswordRefreshUrl1
func (t *OAuthFlowPassword_RefreshUrl) FromOAuthFlowPasswordRefreshUrl1(v OAuthFlowPasswordRefreshUrl1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowPasswordRefreshUrl1 performs a merge with any union data inside the OAuthFlowPassword_RefreshUrl, using the provided OAuthFlowPasswordRefreshUrl1
func (t *OAuthFlowPassword_RefreshUrl) MergeOAuthFlowPasswordRefreshUrl1(v OAuthFlowPasswordRefreshUrl1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlowPassword_RefreshUrl) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlowPassword_RefreshUrl) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowAuthorizationCode returns the union data inside the OAuthFlows_AuthorizationCode as a OAuthFlowAuthorizationCode
func (t OAuthFlows_AuthorizationCode) AsOAuthFlowAuthorizationCode() (OAuthFlowAuthorizationCode, error) {
	var body OAuthFlowAuthorizationCode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowAuthorizationCode overwrites any union data inside the OAuthFlows_AuthorizationCode as the provided OAuthFlowAuthorizationCode
func (t *OAuthFlows_AuthorizationCode) FromOAuthFlowAuthorizationCode(v OAuthFlowAuthorizationCode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowAuthorizationCode performs a merge with any union data inside the OAuthFlows_AuthorizationCode, using the provided OAuthFlowAuthorizationCode
func (t *OAuthFlows_AuthorizationCode) MergeOAuthFlowAuthorizationCode(v OAuthFlowAuthorizationCode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowsAuthorizationCode1 returns the union data inside the OAuthFlows_AuthorizationCode as a OAuthFlowsAuthorizationCode1
func (t OAuthFlows_AuthorizationCode) AsOAuthFlowsAuthorizationCode1() (OAuthFlowsAuthorizationCode1, error) {
	var body OAuthFlowsAuthorizationCode1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowsAuthorizationCode1 overwrites any union data inside the OAuthFlows_AuthorizationCode as the provided OAuthFlowsAuthorizationCode1
func (t *OAuthFlows_AuthorizationCode) FromOAuthFlowsAuthorizationCode1(v OAuthFlowsAuthorizationCode1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowsAuthorizationCode1 performs a merge with any union data inside the OAuthFlows_AuthorizationCode, using the provided OAuthFlowsAuthorizationCode1
func (t *OAuthFlows_AuthorizationCode) MergeOAuthFlowsAuthorizationCode1(v OAuthFlowsAuthorizationCode1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlows_AuthorizationCode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlows_AuthorizationCode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowClientCredentials returns the union data inside the OAuthFlows_ClientCredentials as a OAuthFlowClientCredentials
func (t OAuthFlows_ClientCredentials) AsOAuthFlowClientCredentials() (OAuthFlowClientCredentials, error) {
	var body OAuthFlowClientCredentials
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowClientCredentials overwrites any union data inside the OAuthFlows_ClientCredentials as the provided OAuthFlowClientCredentials
func (t *OAuthFlows_ClientCredentials) FromOAuthFlowClientCredentials(v OAuthFlowClientCredentials) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowClientCredentials performs a merge with any union data inside the OAuthFlows_ClientCredentials, using the provided OAuthFlowClientCredentials
func (t *OAuthFlows_ClientCredentials) MergeOAuthFlowClientCredentials(v OAuthFlowClientCredentials) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowsClientCredentials1 returns the union data inside the OAuthFlows_ClientCredentials as a OAuthFlowsClientCredentials1
func (t OAuthFlows_ClientCredentials) AsOAuthFlowsClientCredentials1() (OAuthFlowsClientCredentials1, error) {
	var body OAuthFlowsClientCredentials1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowsClientCredentials1 overwrites any union data inside the OAuthFlows_ClientCredentials as the provided OAuthFlowsClientCredentials1
func (t *OAuthFlows_ClientCredentials) FromOAuthFlowsClientCredentials1(v OAuthFlowsClientCredentials1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowsClientCredentials1 performs a merge with any union data inside the OAuthFlows_ClientCredentials, using the provided OAuthFlowsClientCredentials1
func (t *OAuthFlows_ClientCredentials) MergeOAuthFlowsClientCredentials1(v OAuthFlowsClientCredentials1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlows_ClientCredentials) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlows_ClientCredentials) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowImplicit returns the union data inside the OAuthFlows_Implicit as a OAuthFlowImplicit
func (t OAuthFlows_Implicit) AsOAuthFlowImplicit() (OAuthFlowImplicit, error) {
	var body OAuthFlowImplicit
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowImplicit overwrites any union data inside the OAuthFlows_Implicit as the provided OAuthFlowImplicit
func (t *OAuthFlows_Implicit) FromOAuthFlowImplicit(v OAuthFlowImplicit) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowImplicit performs a merge with any union data inside the OAuthFlows_Implicit, using the provided OAuthFlowImplicit
func (t *OAuthFlows_Implicit) MergeOAuthFlowImplicit(v OAuthFlowImplicit) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowsImplicit1 returns the union data inside the OAuthFlows_Implicit as a OAuthFlowsImplicit1
func (t OAuthFlows_Implicit) AsOAuthFlowsImplicit1() (OAuthFlowsImplicit1, error) {
	var body OAuthFlowsImplicit1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowsImplicit1 overwrites any union data inside the OAuthFlows_Implicit as the provided OAuthFlowsImplicit1
func (t *OAuthFlows_Implicit) FromOAuthFlowsImplicit1(v OAuthFlowsImplicit1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowsImplicit1 performs a merge with any union data inside the OAuthFlows_Implicit, using the provided OAuthFlowsImplicit1
func (t *OAuthFlows_Implicit) MergeOAuthFlowsImplicit1(v OAuthFlowsImplicit1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlows_Implicit) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlows_Implicit) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOAuthFlowPassword returns the union data inside the OAuthFlows_Password as a OAuthFlowPassword
func (t OAuthFlows_Password) AsOAuthFlowPassword() (OAuthFlowPassword, error) {
	var body OAuthFlowPassword
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowPassword overwrites any union data inside the OAuthFlows_Password as the provided OAuthFlowPassword
func (t *OAuthFlows_Password) FromOAuthFlowPassword(v OAuthFlowPassword) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowPassword performs a merge with any union data inside the OAuthFlows_Password, using the provided OAuthFlowPassword
func (t *OAuthFlows_Password) MergeOAuthFlowPassword(v OAuthFlowPassword) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOAuthFlowsPassword1 returns the union data inside the OAuthFlows_Password as a OAuthFlowsPassword1
func (t OAuthFlows_Password) AsOAuthFlowsPassword1() (OAuthFlowsPassword1, error) {
	var body OAuthFlowsPassword1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOAuthFlowsPassword1 overwrites any union data inside the OAuthFlows_Password as the provided OAuthFlowsPassword1
func (t *OAuthFlows_Password) FromOAuthFlowsPassword1(v OAuthFlowsPassword1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOAuthFlowsPassword1 performs a merge with any union data inside the OAuthFlows_Password, using the provided OAuthFlowsPassword1
func (t *OAuthFlows_Password) MergeOAuthFlowsPassword1(v OAuthFlowsPassword1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OAuthFlows_Password) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OAuthFlows_Password) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpenIdConnectDescription0 returns the union data inside the OpenIdConnect_Description as a OpenIdConnectDescription0
func (t OpenIdConnect_Description) AsOpenIdConnectDescription0() (OpenIdConnectDescription0, error) {
	var body OpenIdConnectDescription0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectDescription0 overwrites any union data inside the OpenIdConnect_Description as the provided OpenIdConnectDescription0
func (t *OpenIdConnect_Description) FromOpenIdConnectDescription0(v OpenIdConnectDescription0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectDescription0 performs a merge with any union data inside the OpenIdConnect_Description, using the provided OpenIdConnectDescription0
func (t *OpenIdConnect_Description) MergeOpenIdConnectDescription0(v OpenIdConnectDescription0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectDescription1 returns the union data inside the OpenIdConnect_Description as a OpenIdConnectDescription1
func (t OpenIdConnect_Description) AsOpenIdConnectDescription1() (OpenIdConnectDescription1, error) {
	var body OpenIdConnectDescription1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectDescription1 overwrites any union data inside the OpenIdConnect_Description as the provided OpenIdConnectDescription1
func (t *OpenIdConnect_Description) FromOpenIdConnectDescription1(v OpenIdConnectDescription1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectDescription1 performs a merge with any union data inside the OpenIdConnect_Description, using the provided OpenIdConnectDescription1
func (t *OpenIdConnect_Description) MergeOpenIdConnectDescription1(v OpenIdConnectDescription1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpenIdConnect_Description) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpenIdConnect_Description) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpenIdConnectWithConfigDescription0 returns the union data inside the OpenIdConnectWithConfig_Description as a OpenIdConnectWithConfigDescription0
func (t OpenIdConnectWithConfig_Description) AsOpenIdConnectWithConfigDescription0() (OpenIdConnectWithConfigDescription0, error) {
	var body OpenIdConnectWithConfigDescription0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigDescription0 overwrites any union data inside the OpenIdConnectWithConfig_Description as the provided OpenIdConnectWithConfigDescription0
func (t *OpenIdConnectWithConfig_Description) FromOpenIdConnectWithConfigDescription0(v OpenIdConnectWithConfigDescription0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigDescription0 performs a merge with any union data inside the OpenIdConnectWithConfig_Description, using the provided OpenIdConnectWithConfigDescription0
func (t *OpenIdConnectWithConfig_Description) MergeOpenIdConnectWithConfigDescription0(v OpenIdConnectWithConfigDescription0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectWithConfigDescription1 returns the union data inside the OpenIdConnectWithConfig_Description as a OpenIdConnectWithConfigDescription1
func (t OpenIdConnectWithConfig_Description) AsOpenIdConnectWithConfigDescription1() (OpenIdConnectWithConfigDescription1, error) {
	var body OpenIdConnectWithConfigDescription1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigDescription1 overwrites any union data inside the OpenIdConnectWithConfig_Description as the provided OpenIdConnectWithConfigDescription1
func (t *OpenIdConnectWithConfig_Description) FromOpenIdConnectWithConfigDescription1(v OpenIdConnectWithConfigDescription1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigDescription1 performs a merge with any union data inside the OpenIdConnectWithConfig_Description, using the provided OpenIdConnectWithConfigDescription1
func (t *OpenIdConnectWithConfig_Description) MergeOpenIdConnectWithConfigDescription1(v OpenIdConnectWithConfigDescription1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpenIdConnectWithConfig_Description) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpenIdConnectWithConfig_Description) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpenIdConnectWithConfigGrantTypesSupported0 returns the union data inside the OpenIdConnectWithConfig_GrantTypesSupported as a OpenIdConnectWithConfigGrantTypesSupported0
func (t OpenIdConnectWithConfig_GrantTypesSupported) AsOpenIdConnectWithConfigGrantTypesSupported0() (OpenIdConnectWithConfigGrantTypesSupported0, error) {
	var body OpenIdConnectWithConfigGrantTypesSupported0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigGrantTypesSupported0 overwrites any union data inside the OpenIdConnectWithConfig_GrantTypesSupported as the provided OpenIdConnectWithConfigGrantTypesSupported0
func (t *OpenIdConnectWithConfig_GrantTypesSupported) FromOpenIdConnectWithConfigGrantTypesSupported0(v OpenIdConnectWithConfigGrantTypesSupported0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigGrantTypesSupported0 performs a merge with any union data inside the OpenIdConnectWithConfig_GrantTypesSupported, using the provided OpenIdConnectWithConfigGrantTypesSupported0
func (t *OpenIdConnectWithConfig_GrantTypesSupported) MergeOpenIdConnectWithConfigGrantTypesSupported0(v OpenIdConnectWithConfigGrantTypesSupported0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectWithConfigGrantTypesSupported1 returns the union data inside the OpenIdConnectWithConfig_GrantTypesSupported as a OpenIdConnectWithConfigGrantTypesSupported1
func (t OpenIdConnectWithConfig_GrantTypesSupported) AsOpenIdConnectWithConfigGrantTypesSupported1() (OpenIdConnectWithConfigGrantTypesSupported1, error) {
	var body OpenIdConnectWithConfigGrantTypesSupported1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigGrantTypesSupported1 overwrites any union data inside the OpenIdConnectWithConfig_GrantTypesSupported as the provided OpenIdConnectWithConfigGrantTypesSupported1
func (t *OpenIdConnectWithConfig_GrantTypesSupported) FromOpenIdConnectWithConfigGrantTypesSupported1(v OpenIdConnectWithConfigGrantTypesSupported1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigGrantTypesSupported1 performs a merge with any union data inside the OpenIdConnectWithConfig_GrantTypesSupported, using the provided OpenIdConnectWithConfigGrantTypesSupported1
func (t *OpenIdConnectWithConfig_GrantTypesSupported) MergeOpenIdConnectWithConfigGrantTypesSupported1(v OpenIdConnectWithConfigGrantTypesSupported1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpenIdConnectWithConfig_GrantTypesSupported) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpenIdConnectWithConfig_GrantTypesSupported) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpenIdConnectWithConfigRevocationEndpoint0 returns the union data inside the OpenIdConnectWithConfig_RevocationEndpoint as a OpenIdConnectWithConfigRevocationEndpoint0
func (t OpenIdConnectWithConfig_RevocationEndpoint) AsOpenIdConnectWithConfigRevocationEndpoint0() (OpenIdConnectWithConfigRevocationEndpoint0, error) {
	var body OpenIdConnectWithConfigRevocationEndpoint0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigRevocationEndpoint0 overwrites any union data inside the OpenIdConnectWithConfig_RevocationEndpoint as the provided OpenIdConnectWithConfigRevocationEndpoint0
func (t *OpenIdConnectWithConfig_RevocationEndpoint) FromOpenIdConnectWithConfigRevocationEndpoint0(v OpenIdConnectWithConfigRevocationEndpoint0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigRevocationEndpoint0 performs a merge with any union data inside the OpenIdConnectWithConfig_RevocationEndpoint, using the provided OpenIdConnectWithConfigRevocationEndpoint0
func (t *OpenIdConnectWithConfig_RevocationEndpoint) MergeOpenIdConnectWithConfigRevocationEndpoint0(v OpenIdConnectWithConfigRevocationEndpoint0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectWithConfigRevocationEndpoint1 returns the union data inside the OpenIdConnectWithConfig_RevocationEndpoint as a OpenIdConnectWithConfigRevocationEndpoint1
func (t OpenIdConnectWithConfig_RevocationEndpoint) AsOpenIdConnectWithConfigRevocationEndpoint1() (OpenIdConnectWithConfigRevocationEndpoint1, error) {
	var body OpenIdConnectWithConfigRevocationEndpoint1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigRevocationEndpoint1 overwrites any union data inside the OpenIdConnectWithConfig_RevocationEndpoint as the provided OpenIdConnectWithConfigRevocationEndpoint1
func (t *OpenIdConnectWithConfig_RevocationEndpoint) FromOpenIdConnectWithConfigRevocationEndpoint1(v OpenIdConnectWithConfigRevocationEndpoint1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigRevocationEndpoint1 performs a merge with any union data inside the OpenIdConnectWithConfig_RevocationEndpoint, using the provided OpenIdConnectWithConfigRevocationEndpoint1
func (t *OpenIdConnectWithConfig_RevocationEndpoint) MergeOpenIdConnectWithConfigRevocationEndpoint1(v OpenIdConnectWithConfigRevocationEndpoint1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpenIdConnectWithConfig_RevocationEndpoint) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpenIdConnectWithConfig_RevocationEndpoint) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpenIdConnectWithConfigScopes0 returns the union data inside the OpenIdConnectWithConfig_Scopes as a OpenIdConnectWithConfigScopes0
func (t OpenIdConnectWithConfig_Scopes) AsOpenIdConnectWithConfigScopes0() (OpenIdConnectWithConfigScopes0, error) {
	var body OpenIdConnectWithConfigScopes0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigScopes0 overwrites any union data inside the OpenIdConnectWithConfig_Scopes as the provided OpenIdConnectWithConfigScopes0
func (t *OpenIdConnectWithConfig_Scopes) FromOpenIdConnectWithConfigScopes0(v OpenIdConnectWithConfigScopes0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigScopes0 performs a merge with any union data inside the OpenIdConnectWithConfig_Scopes, using the provided OpenIdConnectWithConfigScopes0
func (t *OpenIdConnectWithConfig_Scopes) MergeOpenIdConnectWithConfigScopes0(v OpenIdConnectWithConfigScopes0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectWithConfigScopes1 returns the union data inside the OpenIdConnectWithConfig_Scopes as a OpenIdConnectWithConfigScopes1
func (t OpenIdConnectWithConfig_Scopes) AsOpenIdConnectWithConfigScopes1() (OpenIdConnectWithConfigScopes1, error) {
	var body OpenIdConnectWithConfigScopes1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigScopes1 overwrites any union data inside the OpenIdConnectWithConfig_Scopes as the provided OpenIdConnectWithConfigScopes1
func (t *OpenIdConnectWithConfig_Scopes) FromOpenIdConnectWithConfigScopes1(v OpenIdConnectWithConfigScopes1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigScopes1 performs a merge with any union data inside the OpenIdConnectWithConfig_Scopes, using the provided OpenIdConnectWithConfigScopes1
func (t *OpenIdConnectWithConfig_Scopes) MergeOpenIdConnectWithConfigScopes1(v OpenIdConnectWithConfigScopes1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpenIdConnectWithConfig_Scopes) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpenIdConnectWithConfig_Scopes) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0 returns the union data inside the OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported as a OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0
func (t OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) AsOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0() (OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0, error) {
	var body OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0 overwrites any union data inside the OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported as the provided OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0
func (t *OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) FromOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0(v OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0 performs a merge with any union data inside the OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported, using the provided OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0
func (t *OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) MergeOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0(v OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1 returns the union data inside the OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported as a OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1
func (t OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) AsOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1() (OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1, error) {
	var body OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1 overwrites any union data inside the OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported as the provided OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1
func (t *OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) FromOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1(v OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1 performs a merge with any union data inside the OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported, using the provided OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1
func (t *OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) MergeOpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1(v OpenIdConnectWithConfigTokenEndpointAuthMethodsSupported1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpenIdConnectWithConfig_TokenEndpointAuthMethodsSupported) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsOpenIdConnectWithConfigUserinfoEndpoint0 returns the union data inside the OpenIdConnectWithConfig_UserinfoEndpoint as a OpenIdConnectWithConfigUserinfoEndpoint0
func (t OpenIdConnectWithConfig_UserinfoEndpoint) AsOpenIdConnectWithConfigUserinfoEndpoint0() (OpenIdConnectWithConfigUserinfoEndpoint0, error) {
	var body OpenIdConnectWithConfigUserinfoEndpoint0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigUserinfoEndpoint0 overwrites any union data inside the OpenIdConnectWithConfig_UserinfoEndpoint as the provided OpenIdConnectWithConfigUserinfoEndpoint0
func (t *OpenIdConnectWithConfig_UserinfoEndpoint) FromOpenIdConnectWithConfigUserinfoEndpoint0(v OpenIdConnectWithConfigUserinfoEndpoint0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigUserinfoEndpoint0 performs a merge with any union data inside the OpenIdConnectWithConfig_UserinfoEndpoint, using the provided OpenIdConnectWithConfigUserinfoEndpoint0
func (t *OpenIdConnectWithConfig_UserinfoEndpoint) MergeOpenIdConnectWithConfigUserinfoEndpoint0(v OpenIdConnectWithConfigUserinfoEndpoint0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsOpenIdConnectWithConfigUserinfoEndpoint1 returns the union data inside the OpenIdConnectWithConfig_UserinfoEndpoint as a OpenIdConnectWithConfigUserinfoEndpoint1
func (t OpenIdConnectWithConfig_UserinfoEndpoint) AsOpenIdConnectWithConfigUserinfoEndpoint1() (OpenIdConnectWithConfigUserinfoEndpoint1, error) {
	var body OpenIdConnectWithConfigUserinfoEndpoint1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromOpenIdConnectWithConfigUserinfoEndpoint1 overwrites any union data inside the OpenIdConnectWithConfig_UserinfoEndpoint as the provided OpenIdConnectWithConfigUserinfoEndpoint1
func (t *OpenIdConnectWithConfig_UserinfoEndpoint) FromOpenIdConnectWithConfigUserinfoEndpoint1(v OpenIdConnectWithConfigUserinfoEndpoint1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeOpenIdConnectWithConfigUserinfoEndpoint1 performs a merge with any union data inside the OpenIdConnectWithConfig_UserinfoEndpoint, using the provided OpenIdConnectWithConfigUserinfoEndpoint1
func (t *OpenIdConnectWithConfig_UserinfoEndpoint) MergeOpenIdConnectWithConfigUserinfoEndpoint1(v OpenIdConnectWithConfigUserinfoEndpoint1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t OpenIdConnectWithConfig_UserinfoEndpoint) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *OpenIdConnectWithConfig_UserinfoEndpoint) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCodeExecutionResult returns the union data inside the PartInput_CodeExecutionResult as a CodeExecutionResult
func (t PartInput_CodeExecutionResult) AsCodeExecutionResult() (CodeExecutionResult, error) {
	var body CodeExecutionResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCodeExecutionResult overwrites any union data inside the PartInput_CodeExecutionResult as the provided CodeExecutionResult
func (t *PartInput_CodeExecutionResult) FromCodeExecutionResult(v CodeExecutionResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCodeExecutionResult performs a merge with any union data inside the PartInput_CodeExecutionResult, using the provided CodeExecutionResult
func (t *PartInput_CodeExecutionResult) MergeCodeExecutionResult(v CodeExecutionResult) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputCodeExecutionResult1 returns the union data inside the PartInput_CodeExecutionResult as a PartInputCodeExecutionResult1
func (t PartInput_CodeExecutionResult) AsPartInputCodeExecutionResult1() (PartInputCodeExecutionResult1, error) {
	var body PartInputCodeExecutionResult1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputCodeExecutionResult1 overwrites any union data inside the PartInput_CodeExecutionResult as the provided PartInputCodeExecutionResult1
func (t *PartInput_CodeExecutionResult) FromPartInputCodeExecutionResult1(v PartInputCodeExecutionResult1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputCodeExecutionResult1 performs a merge with any union data inside the PartInput_CodeExecutionResult, using the provided PartInputCodeExecutionResult1
func (t *PartInput_CodeExecutionResult) MergePartInputCodeExecutionResult1(v PartInputCodeExecutionResult1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_CodeExecutionResult) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_CodeExecutionResult) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsExecutableCode returns the union data inside the PartInput_ExecutableCode as a ExecutableCode
func (t PartInput_ExecutableCode) AsExecutableCode() (ExecutableCode, error) {
	var body ExecutableCode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromExecutableCode overwrites any union data inside the PartInput_ExecutableCode as the provided ExecutableCode
func (t *PartInput_ExecutableCode) FromExecutableCode(v ExecutableCode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeExecutableCode performs a merge with any union data inside the PartInput_ExecutableCode, using the provided ExecutableCode
func (t *PartInput_ExecutableCode) MergeExecutableCode(v ExecutableCode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputExecutableCode1 returns the union data inside the PartInput_ExecutableCode as a PartInputExecutableCode1
func (t PartInput_ExecutableCode) AsPartInputExecutableCode1() (PartInputExecutableCode1, error) {
	var body PartInputExecutableCode1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputExecutableCode1 overwrites any union data inside the PartInput_ExecutableCode as the provided PartInputExecutableCode1
func (t *PartInput_ExecutableCode) FromPartInputExecutableCode1(v PartInputExecutableCode1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputExecutableCode1 performs a merge with any union data inside the PartInput_ExecutableCode, using the provided PartInputExecutableCode1
func (t *PartInput_ExecutableCode) MergePartInputExecutableCode1(v PartInputExecutableCode1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_ExecutableCode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_ExecutableCode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFileData returns the union data inside the PartInput_FileData as a FileData
func (t PartInput_FileData) AsFileData() (FileData, error) {
	var body FileData
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileData overwrites any union data inside the PartInput_FileData as the provided FileData
func (t *PartInput_FileData) FromFileData(v FileData) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileData performs a merge with any union data inside the PartInput_FileData, using the provided FileData
func (t *PartInput_FileData) MergeFileData(v FileData) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputFileData1 returns the union data inside the PartInput_FileData as a PartInputFileData1
func (t PartInput_FileData) AsPartInputFileData1() (PartInputFileData1, error) {
	var body PartInputFileData1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputFileData1 overwrites any union data inside the PartInput_FileData as the provided PartInputFileData1
func (t *PartInput_FileData) FromPartInputFileData1(v PartInputFileData1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputFileData1 performs a merge with any union data inside the PartInput_FileData, using the provided PartInputFileData1
func (t *PartInput_FileData) MergePartInputFileData1(v PartInputFileData1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_FileData) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_FileData) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionCall returns the union data inside the PartInput_FunctionCall as a FunctionCall
func (t PartInput_FunctionCall) AsFunctionCall() (FunctionCall, error) {
	var body FunctionCall
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCall overwrites any union data inside the PartInput_FunctionCall as the provided FunctionCall
func (t *PartInput_FunctionCall) FromFunctionCall(v FunctionCall) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCall performs a merge with any union data inside the PartInput_FunctionCall, using the provided FunctionCall
func (t *PartInput_FunctionCall) MergeFunctionCall(v FunctionCall) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputFunctionCall1 returns the union data inside the PartInput_FunctionCall as a PartInputFunctionCall1
func (t PartInput_FunctionCall) AsPartInputFunctionCall1() (PartInputFunctionCall1, error) {
	var body PartInputFunctionCall1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputFunctionCall1 overwrites any union data inside the PartInput_FunctionCall as the provided PartInputFunctionCall1
func (t *PartInput_FunctionCall) FromPartInputFunctionCall1(v PartInputFunctionCall1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputFunctionCall1 performs a merge with any union data inside the PartInput_FunctionCall, using the provided PartInputFunctionCall1
func (t *PartInput_FunctionCall) MergePartInputFunctionCall1(v PartInputFunctionCall1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_FunctionCall) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_FunctionCall) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionResponse returns the union data inside the PartInput_FunctionResponse as a FunctionResponse
func (t PartInput_FunctionResponse) AsFunctionResponse() (FunctionResponse, error) {
	var body FunctionResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponse overwrites any union data inside the PartInput_FunctionResponse as the provided FunctionResponse
func (t *PartInput_FunctionResponse) FromFunctionResponse(v FunctionResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponse performs a merge with any union data inside the PartInput_FunctionResponse, using the provided FunctionResponse
func (t *PartInput_FunctionResponse) MergeFunctionResponse(v FunctionResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputFunctionResponse1 returns the union data inside the PartInput_FunctionResponse as a PartInputFunctionResponse1
func (t PartInput_FunctionResponse) AsPartInputFunctionResponse1() (PartInputFunctionResponse1, error) {
	var body PartInputFunctionResponse1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputFunctionResponse1 overwrites any union data inside the PartInput_FunctionResponse as the provided PartInputFunctionResponse1
func (t *PartInput_FunctionResponse) FromPartInputFunctionResponse1(v PartInputFunctionResponse1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputFunctionResponse1 performs a merge with any union data inside the PartInput_FunctionResponse, using the provided PartInputFunctionResponse1
func (t *PartInput_FunctionResponse) MergePartInputFunctionResponse1(v PartInputFunctionResponse1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_FunctionResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_FunctionResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBlob returns the union data inside the PartInput_InlineData as a Blob
func (t PartInput_InlineData) AsBlob() (Blob, error) {
	var body Blob
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBlob overwrites any union data inside the PartInput_InlineData as the provided Blob
func (t *PartInput_InlineData) FromBlob(v Blob) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBlob performs a merge with any union data inside the PartInput_InlineData, using the provided Blob
func (t *PartInput_InlineData) MergeBlob(v Blob) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputInlineData1 returns the union data inside the PartInput_InlineData as a PartInputInlineData1
func (t PartInput_InlineData) AsPartInputInlineData1() (PartInputInlineData1, error) {
	var body PartInputInlineData1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputInlineData1 overwrites any union data inside the PartInput_InlineData as the provided PartInputInlineData1
func (t *PartInput_InlineData) FromPartInputInlineData1(v PartInputInlineData1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputInlineData1 performs a merge with any union data inside the PartInput_InlineData, using the provided PartInputInlineData1
func (t *PartInput_InlineData) MergePartInputInlineData1(v PartInputInlineData1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_InlineData) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_InlineData) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPartInputText0 returns the union data inside the PartInput_Text as a PartInputText0
func (t PartInput_Text) AsPartInputText0() (PartInputText0, error) {
	var body PartInputText0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputText0 overwrites any union data inside the PartInput_Text as the provided PartInputText0
func (t *PartInput_Text) FromPartInputText0(v PartInputText0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputText0 performs a merge with any union data inside the PartInput_Text, using the provided PartInputText0
func (t *PartInput_Text) MergePartInputText0(v PartInputText0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputText1 returns the union data inside the PartInput_Text as a PartInputText1
func (t PartInput_Text) AsPartInputText1() (PartInputText1, error) {
	var body PartInputText1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputText1 overwrites any union data inside the PartInput_Text as the provided PartInputText1
func (t *PartInput_Text) FromPartInputText1(v PartInputText1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputText1 performs a merge with any union data inside the PartInput_Text, using the provided PartInputText1
func (t *PartInput_Text) MergePartInputText1(v PartInputText1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_Text) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_Text) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPartInputThought0 returns the union data inside the PartInput_Thought as a PartInputThought0
func (t PartInput_Thought) AsPartInputThought0() (PartInputThought0, error) {
	var body PartInputThought0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputThought0 overwrites any union data inside the PartInput_Thought as the provided PartInputThought0
func (t *PartInput_Thought) FromPartInputThought0(v PartInputThought0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputThought0 performs a merge with any union data inside the PartInput_Thought, using the provided PartInputThought0
func (t *PartInput_Thought) MergePartInputThought0(v PartInputThought0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputThought1 returns the union data inside the PartInput_Thought as a PartInputThought1
func (t PartInput_Thought) AsPartInputThought1() (PartInputThought1, error) {
	var body PartInputThought1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputThought1 overwrites any union data inside the PartInput_Thought as the provided PartInputThought1
func (t *PartInput_Thought) FromPartInputThought1(v PartInputThought1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputThought1 performs a merge with any union data inside the PartInput_Thought, using the provided PartInputThought1
func (t *PartInput_Thought) MergePartInputThought1(v PartInputThought1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_Thought) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_Thought) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsVideoMetadata returns the union data inside the PartInput_VideoMetadata as a VideoMetadata
func (t PartInput_VideoMetadata) AsVideoMetadata() (VideoMetadata, error) {
	var body VideoMetadata
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVideoMetadata overwrites any union data inside the PartInput_VideoMetadata as the provided VideoMetadata
func (t *PartInput_VideoMetadata) FromVideoMetadata(v VideoMetadata) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVideoMetadata performs a merge with any union data inside the PartInput_VideoMetadata, using the provided VideoMetadata
func (t *PartInput_VideoMetadata) MergeVideoMetadata(v VideoMetadata) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartInputVideoMetadata1 returns the union data inside the PartInput_VideoMetadata as a PartInputVideoMetadata1
func (t PartInput_VideoMetadata) AsPartInputVideoMetadata1() (PartInputVideoMetadata1, error) {
	var body PartInputVideoMetadata1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartInputVideoMetadata1 overwrites any union data inside the PartInput_VideoMetadata as the provided PartInputVideoMetadata1
func (t *PartInput_VideoMetadata) FromPartInputVideoMetadata1(v PartInputVideoMetadata1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartInputVideoMetadata1 performs a merge with any union data inside the PartInput_VideoMetadata, using the provided PartInputVideoMetadata1
func (t *PartInput_VideoMetadata) MergePartInputVideoMetadata1(v PartInputVideoMetadata1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartInput_VideoMetadata) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartInput_VideoMetadata) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsCodeExecutionResult returns the union data inside the PartOutput_CodeExecutionResult as a CodeExecutionResult
func (t PartOutput_CodeExecutionResult) AsCodeExecutionResult() (CodeExecutionResult, error) {
	var body CodeExecutionResult
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromCodeExecutionResult overwrites any union data inside the PartOutput_CodeExecutionResult as the provided CodeExecutionResult
func (t *PartOutput_CodeExecutionResult) FromCodeExecutionResult(v CodeExecutionResult) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeCodeExecutionResult performs a merge with any union data inside the PartOutput_CodeExecutionResult, using the provided CodeExecutionResult
func (t *PartOutput_CodeExecutionResult) MergeCodeExecutionResult(v CodeExecutionResult) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputCodeExecutionResult1 returns the union data inside the PartOutput_CodeExecutionResult as a PartOutputCodeExecutionResult1
func (t PartOutput_CodeExecutionResult) AsPartOutputCodeExecutionResult1() (PartOutputCodeExecutionResult1, error) {
	var body PartOutputCodeExecutionResult1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputCodeExecutionResult1 overwrites any union data inside the PartOutput_CodeExecutionResult as the provided PartOutputCodeExecutionResult1
func (t *PartOutput_CodeExecutionResult) FromPartOutputCodeExecutionResult1(v PartOutputCodeExecutionResult1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputCodeExecutionResult1 performs a merge with any union data inside the PartOutput_CodeExecutionResult, using the provided PartOutputCodeExecutionResult1
func (t *PartOutput_CodeExecutionResult) MergePartOutputCodeExecutionResult1(v PartOutputCodeExecutionResult1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_CodeExecutionResult) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_CodeExecutionResult) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsExecutableCode returns the union data inside the PartOutput_ExecutableCode as a ExecutableCode
func (t PartOutput_ExecutableCode) AsExecutableCode() (ExecutableCode, error) {
	var body ExecutableCode
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromExecutableCode overwrites any union data inside the PartOutput_ExecutableCode as the provided ExecutableCode
func (t *PartOutput_ExecutableCode) FromExecutableCode(v ExecutableCode) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeExecutableCode performs a merge with any union data inside the PartOutput_ExecutableCode, using the provided ExecutableCode
func (t *PartOutput_ExecutableCode) MergeExecutableCode(v ExecutableCode) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputExecutableCode1 returns the union data inside the PartOutput_ExecutableCode as a PartOutputExecutableCode1
func (t PartOutput_ExecutableCode) AsPartOutputExecutableCode1() (PartOutputExecutableCode1, error) {
	var body PartOutputExecutableCode1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputExecutableCode1 overwrites any union data inside the PartOutput_ExecutableCode as the provided PartOutputExecutableCode1
func (t *PartOutput_ExecutableCode) FromPartOutputExecutableCode1(v PartOutputExecutableCode1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputExecutableCode1 performs a merge with any union data inside the PartOutput_ExecutableCode, using the provided PartOutputExecutableCode1
func (t *PartOutput_ExecutableCode) MergePartOutputExecutableCode1(v PartOutputExecutableCode1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_ExecutableCode) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_ExecutableCode) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFileData returns the union data inside the PartOutput_FileData as a FileData
func (t PartOutput_FileData) AsFileData() (FileData, error) {
	var body FileData
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFileData overwrites any union data inside the PartOutput_FileData as the provided FileData
func (t *PartOutput_FileData) FromFileData(v FileData) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFileData performs a merge with any union data inside the PartOutput_FileData, using the provided FileData
func (t *PartOutput_FileData) MergeFileData(v FileData) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputFileData1 returns the union data inside the PartOutput_FileData as a PartOutputFileData1
func (t PartOutput_FileData) AsPartOutputFileData1() (PartOutputFileData1, error) {
	var body PartOutputFileData1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputFileData1 overwrites any union data inside the PartOutput_FileData as the provided PartOutputFileData1
func (t *PartOutput_FileData) FromPartOutputFileData1(v PartOutputFileData1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputFileData1 performs a merge with any union data inside the PartOutput_FileData, using the provided PartOutputFileData1
func (t *PartOutput_FileData) MergePartOutputFileData1(v PartOutputFileData1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_FileData) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_FileData) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionCall returns the union data inside the PartOutput_FunctionCall as a FunctionCall
func (t PartOutput_FunctionCall) AsFunctionCall() (FunctionCall, error) {
	var body FunctionCall
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionCall overwrites any union data inside the PartOutput_FunctionCall as the provided FunctionCall
func (t *PartOutput_FunctionCall) FromFunctionCall(v FunctionCall) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionCall performs a merge with any union data inside the PartOutput_FunctionCall, using the provided FunctionCall
func (t *PartOutput_FunctionCall) MergeFunctionCall(v FunctionCall) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputFunctionCall1 returns the union data inside the PartOutput_FunctionCall as a PartOutputFunctionCall1
func (t PartOutput_FunctionCall) AsPartOutputFunctionCall1() (PartOutputFunctionCall1, error) {
	var body PartOutputFunctionCall1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputFunctionCall1 overwrites any union data inside the PartOutput_FunctionCall as the provided PartOutputFunctionCall1
func (t *PartOutput_FunctionCall) FromPartOutputFunctionCall1(v PartOutputFunctionCall1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputFunctionCall1 performs a merge with any union data inside the PartOutput_FunctionCall, using the provided PartOutputFunctionCall1
func (t *PartOutput_FunctionCall) MergePartOutputFunctionCall1(v PartOutputFunctionCall1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_FunctionCall) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_FunctionCall) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsFunctionResponse returns the union data inside the PartOutput_FunctionResponse as a FunctionResponse
func (t PartOutput_FunctionResponse) AsFunctionResponse() (FunctionResponse, error) {
	var body FunctionResponse
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromFunctionResponse overwrites any union data inside the PartOutput_FunctionResponse as the provided FunctionResponse
func (t *PartOutput_FunctionResponse) FromFunctionResponse(v FunctionResponse) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeFunctionResponse performs a merge with any union data inside the PartOutput_FunctionResponse, using the provided FunctionResponse
func (t *PartOutput_FunctionResponse) MergeFunctionResponse(v FunctionResponse) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputFunctionResponse1 returns the union data inside the PartOutput_FunctionResponse as a PartOutputFunctionResponse1
func (t PartOutput_FunctionResponse) AsPartOutputFunctionResponse1() (PartOutputFunctionResponse1, error) {
	var body PartOutputFunctionResponse1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputFunctionResponse1 overwrites any union data inside the PartOutput_FunctionResponse as the provided PartOutputFunctionResponse1
func (t *PartOutput_FunctionResponse) FromPartOutputFunctionResponse1(v PartOutputFunctionResponse1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputFunctionResponse1 performs a merge with any union data inside the PartOutput_FunctionResponse, using the provided PartOutputFunctionResponse1
func (t *PartOutput_FunctionResponse) MergePartOutputFunctionResponse1(v PartOutputFunctionResponse1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_FunctionResponse) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_FunctionResponse) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsBlob returns the union data inside the PartOutput_InlineData as a Blob
func (t PartOutput_InlineData) AsBlob() (Blob, error) {
	var body Blob
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromBlob overwrites any union data inside the PartOutput_InlineData as the provided Blob
func (t *PartOutput_InlineData) FromBlob(v Blob) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeBlob performs a merge with any union data inside the PartOutput_InlineData, using the provided Blob
func (t *PartOutput_InlineData) MergeBlob(v Blob) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputInlineData1 returns the union data inside the PartOutput_InlineData as a PartOutputInlineData1
func (t PartOutput_InlineData) AsPartOutputInlineData1() (PartOutputInlineData1, error) {
	var body PartOutputInlineData1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputInlineData1 overwrites any union data inside the PartOutput_InlineData as the provided PartOutputInlineData1
func (t *PartOutput_InlineData) FromPartOutputInlineData1(v PartOutputInlineData1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputInlineData1 performs a merge with any union data inside the PartOutput_InlineData, using the provided PartOutputInlineData1
func (t *PartOutput_InlineData) MergePartOutputInlineData1(v PartOutputInlineData1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_InlineData) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_InlineData) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPartOutputText0 returns the union data inside the PartOutput_Text as a PartOutputText0
func (t PartOutput_Text) AsPartOutputText0() (PartOutputText0, error) {
	var body PartOutputText0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputText0 overwrites any union data inside the PartOutput_Text as the provided PartOutputText0
func (t *PartOutput_Text) FromPartOutputText0(v PartOutputText0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputText0 performs a merge with any union data inside the PartOutput_Text, using the provided PartOutputText0
func (t *PartOutput_Text) MergePartOutputText0(v PartOutputText0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputText1 returns the union data inside the PartOutput_Text as a PartOutputText1
func (t PartOutput_Text) AsPartOutputText1() (PartOutputText1, error) {
	var body PartOutputText1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputText1 overwrites any union data inside the PartOutput_Text as the provided PartOutputText1
func (t *PartOutput_Text) FromPartOutputText1(v PartOutputText1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputText1 performs a merge with any union data inside the PartOutput_Text, using the provided PartOutputText1
func (t *PartOutput_Text) MergePartOutputText1(v PartOutputText1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_Text) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_Text) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsPartOutputThought0 returns the union data inside the PartOutput_Thought as a PartOutputThought0
func (t PartOutput_Thought) AsPartOutputThought0() (PartOutputThought0, error) {
	var body PartOutputThought0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputThought0 overwrites any union data inside the PartOutput_Thought as the provided PartOutputThought0
func (t *PartOutput_Thought) FromPartOutputThought0(v PartOutputThought0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputThought0 performs a merge with any union data inside the PartOutput_Thought, using the provided PartOutputThought0
func (t *PartOutput_Thought) MergePartOutputThought0(v PartOutputThought0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputThought1 returns the union data inside the PartOutput_Thought as a PartOutputThought1
func (t PartOutput_Thought) AsPartOutputThought1() (PartOutputThought1, error) {
	var body PartOutputThought1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputThought1 overwrites any union data inside the PartOutput_Thought as the provided PartOutputThought1
func (t *PartOutput_Thought) FromPartOutputThought1(v PartOutputThought1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputThought1 performs a merge with any union data inside the PartOutput_Thought, using the provided PartOutputThought1
func (t *PartOutput_Thought) MergePartOutputThought1(v PartOutputThought1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_Thought) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_Thought) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsVideoMetadata returns the union data inside the PartOutput_VideoMetadata as a VideoMetadata
func (t PartOutput_VideoMetadata) AsVideoMetadata() (VideoMetadata, error) {
	var body VideoMetadata
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVideoMetadata overwrites any union data inside the PartOutput_VideoMetadata as the provided VideoMetadata
func (t *PartOutput_VideoMetadata) FromVideoMetadata(v VideoMetadata) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVideoMetadata performs a merge with any union data inside the PartOutput_VideoMetadata, using the provided VideoMetadata
func (t *PartOutput_VideoMetadata) MergeVideoMetadata(v VideoMetadata) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsPartOutputVideoMetadata1 returns the union data inside the PartOutput_VideoMetadata as a PartOutputVideoMetadata1
func (t PartOutput_VideoMetadata) AsPartOutputVideoMetadata1() (PartOutputVideoMetadata1, error) {
	var body PartOutputVideoMetadata1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromPartOutputVideoMetadata1 overwrites any union data inside the PartOutput_VideoMetadata as the provided PartOutputVideoMetadata1
func (t *PartOutput_VideoMetadata) FromPartOutputVideoMetadata1(v PartOutputVideoMetadata1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergePartOutputVideoMetadata1 performs a merge with any union data inside the PartOutput_VideoMetadata, using the provided PartOutputVideoMetadata1
func (t *PartOutput_VideoMetadata) MergePartOutputVideoMetadata1(v PartOutputVideoMetadata1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t PartOutput_VideoMetadata) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *PartOutput_VideoMetadata) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsRetrievalMetadataGoogleSearchDynamicRetrievalScore0 returns the union data inside the RetrievalMetadata_GoogleSearchDynamicRetrievalScore as a RetrievalMetadataGoogleSearchDynamicRetrievalScore0
func (t RetrievalMetadata_GoogleSearchDynamicRetrievalScore) AsRetrievalMetadataGoogleSearchDynamicRetrievalScore0() (RetrievalMetadataGoogleSearchDynamicRetrievalScore0, error) {
	var body RetrievalMetadataGoogleSearchDynamicRetrievalScore0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRetrievalMetadataGoogleSearchDynamicRetrievalScore0 overwrites any union data inside the RetrievalMetadata_GoogleSearchDynamicRetrievalScore as the provided RetrievalMetadataGoogleSearchDynamicRetrievalScore0
func (t *RetrievalMetadata_GoogleSearchDynamicRetrievalScore) FromRetrievalMetadataGoogleSearchDynamicRetrievalScore0(v RetrievalMetadataGoogleSearchDynamicRetrievalScore0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRetrievalMetadataGoogleSearchDynamicRetrievalScore0 performs a merge with any union data inside the RetrievalMetadata_GoogleSearchDynamicRetrievalScore, using the provided RetrievalMetadataGoogleSearchDynamicRetrievalScore0
func (t *RetrievalMetadata_GoogleSearchDynamicRetrievalScore) MergeRetrievalMetadataGoogleSearchDynamicRetrievalScore0(v RetrievalMetadataGoogleSearchDynamicRetrievalScore0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsRetrievalMetadataGoogleSearchDynamicRetrievalScore1 returns the union data inside the RetrievalMetadata_GoogleSearchDynamicRetrievalScore as a RetrievalMetadataGoogleSearchDynamicRetrievalScore1
func (t RetrievalMetadata_GoogleSearchDynamicRetrievalScore) AsRetrievalMetadataGoogleSearchDynamicRetrievalScore1() (RetrievalMetadataGoogleSearchDynamicRetrievalScore1, error) {
	var body RetrievalMetadataGoogleSearchDynamicRetrievalScore1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromRetrievalMetadataGoogleSearchDynamicRetrievalScore1 overwrites any union data inside the RetrievalMetadata_GoogleSearchDynamicRetrievalScore as the provided RetrievalMetadataGoogleSearchDynamicRetrievalScore1
func (t *RetrievalMetadata_GoogleSearchDynamicRetrievalScore) FromRetrievalMetadataGoogleSearchDynamicRetrievalScore1(v RetrievalMetadataGoogleSearchDynamicRetrievalScore1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeRetrievalMetadataGoogleSearchDynamicRetrievalScore1 performs a merge with any union data inside the RetrievalMetadata_GoogleSearchDynamicRetrievalScore, using the provided RetrievalMetadataGoogleSearchDynamicRetrievalScore1
func (t *RetrievalMetadata_GoogleSearchDynamicRetrievalScore) MergeRetrievalMetadataGoogleSearchDynamicRetrievalScore1(v RetrievalMetadataGoogleSearchDynamicRetrievalScore1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t RetrievalMetadata_GoogleSearchDynamicRetrievalScore) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *RetrievalMetadata_GoogleSearchDynamicRetrievalScore) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSearchEntryPointRenderedContent0 returns the union data inside the SearchEntryPoint_RenderedContent as a SearchEntryPointRenderedContent0
func (t SearchEntryPoint_RenderedContent) AsSearchEntryPointRenderedContent0() (SearchEntryPointRenderedContent0, error) {
	var body SearchEntryPointRenderedContent0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSearchEntryPointRenderedContent0 overwrites any union data inside the SearchEntryPoint_RenderedContent as the provided SearchEntryPointRenderedContent0
func (t *SearchEntryPoint_RenderedContent) FromSearchEntryPointRenderedContent0(v SearchEntryPointRenderedContent0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSearchEntryPointRenderedContent0 performs a merge with any union data inside the SearchEntryPoint_RenderedContent, using the provided SearchEntryPointRenderedContent0
func (t *SearchEntryPoint_RenderedContent) MergeSearchEntryPointRenderedContent0(v SearchEntryPointRenderedContent0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsSearchEntryPointRenderedContent1 returns the union data inside the SearchEntryPoint_RenderedContent as a SearchEntryPointRenderedContent1
func (t SearchEntryPoint_RenderedContent) AsSearchEntryPointRenderedContent1() (SearchEntryPointRenderedContent1, error) {
	var body SearchEntryPointRenderedContent1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSearchEntryPointRenderedContent1 overwrites any union data inside the SearchEntryPoint_RenderedContent as the provided SearchEntryPointRenderedContent1
func (t *SearchEntryPoint_RenderedContent) FromSearchEntryPointRenderedContent1(v SearchEntryPointRenderedContent1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSearchEntryPointRenderedContent1 performs a merge with any union data inside the SearchEntryPoint_RenderedContent, using the provided SearchEntryPointRenderedContent1
func (t *SearchEntryPoint_RenderedContent) MergeSearchEntryPointRenderedContent1(v SearchEntryPointRenderedContent1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t SearchEntryPoint_RenderedContent) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SearchEntryPoint_RenderedContent) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSearchEntryPointSdkBlob0 returns the union data inside the SearchEntryPoint_SdkBlob as a SearchEntryPointSdkBlob0
func (t SearchEntryPoint_SdkBlob) AsSearchEntryPointSdkBlob0() (SearchEntryPointSdkBlob0, error) {
	var body SearchEntryPointSdkBlob0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSearchEntryPointSdkBlob0 overwrites any union data inside the SearchEntryPoint_SdkBlob as the provided SearchEntryPointSdkBlob0
func (t *SearchEntryPoint_SdkBlob) FromSearchEntryPointSdkBlob0(v SearchEntryPointSdkBlob0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSearchEntryPointSdkBlob0 performs a merge with any union data inside the SearchEntryPoint_SdkBlob, using the provided SearchEntryPointSdkBlob0
func (t *SearchEntryPoint_SdkBlob) MergeSearchEntryPointSdkBlob0(v SearchEntryPointSdkBlob0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsSearchEntryPointSdkBlob1 returns the union data inside the SearchEntryPoint_SdkBlob as a SearchEntryPointSdkBlob1
func (t SearchEntryPoint_SdkBlob) AsSearchEntryPointSdkBlob1() (SearchEntryPointSdkBlob1, error) {
	var body SearchEntryPointSdkBlob1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSearchEntryPointSdkBlob1 overwrites any union data inside the SearchEntryPoint_SdkBlob as the provided SearchEntryPointSdkBlob1
func (t *SearchEntryPoint_SdkBlob) FromSearchEntryPointSdkBlob1(v SearchEntryPointSdkBlob1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSearchEntryPointSdkBlob1 performs a merge with any union data inside the SearchEntryPoint_SdkBlob, using the provided SearchEntryPointSdkBlob1
func (t *SearchEntryPoint_SdkBlob) MergeSearchEntryPointSdkBlob1(v SearchEntryPointSdkBlob1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t SearchEntryPoint_SdkBlob) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *SearchEntryPoint_SdkBlob) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSegmentEndIndex0 returns the union data inside the Segment_EndIndex as a SegmentEndIndex0
func (t Segment_EndIndex) AsSegmentEndIndex0() (SegmentEndIndex0, error) {
	var body SegmentEndIndex0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentEndIndex0 overwrites any union data inside the Segment_EndIndex as the provided SegmentEndIndex0
func (t *Segment_EndIndex) FromSegmentEndIndex0(v SegmentEndIndex0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentEndIndex0 performs a merge with any union data inside the Segment_EndIndex, using the provided SegmentEndIndex0
func (t *Segment_EndIndex) MergeSegmentEndIndex0(v SegmentEndIndex0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsSegmentEndIndex1 returns the union data inside the Segment_EndIndex as a SegmentEndIndex1
func (t Segment_EndIndex) AsSegmentEndIndex1() (SegmentEndIndex1, error) {
	var body SegmentEndIndex1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentEndIndex1 overwrites any union data inside the Segment_EndIndex as the provided SegmentEndIndex1
func (t *Segment_EndIndex) FromSegmentEndIndex1(v SegmentEndIndex1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentEndIndex1 performs a merge with any union data inside the Segment_EndIndex, using the provided SegmentEndIndex1
func (t *Segment_EndIndex) MergeSegmentEndIndex1(v SegmentEndIndex1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Segment_EndIndex) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Segment_EndIndex) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSegmentPartIndex0 returns the union data inside the Segment_PartIndex as a SegmentPartIndex0
func (t Segment_PartIndex) AsSegmentPartIndex0() (SegmentPartIndex0, error) {
	var body SegmentPartIndex0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentPartIndex0 overwrites any union data inside the Segment_PartIndex as the provided SegmentPartIndex0
func (t *Segment_PartIndex) FromSegmentPartIndex0(v SegmentPartIndex0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentPartIndex0 performs a merge with any union data inside the Segment_PartIndex, using the provided SegmentPartIndex0
func (t *Segment_PartIndex) MergeSegmentPartIndex0(v SegmentPartIndex0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsSegmentPartIndex1 returns the union data inside the Segment_PartIndex as a SegmentPartIndex1
func (t Segment_PartIndex) AsSegmentPartIndex1() (SegmentPartIndex1, error) {
	var body SegmentPartIndex1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentPartIndex1 overwrites any union data inside the Segment_PartIndex as the provided SegmentPartIndex1
func (t *Segment_PartIndex) FromSegmentPartIndex1(v SegmentPartIndex1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentPartIndex1 performs a merge with any union data inside the Segment_PartIndex, using the provided SegmentPartIndex1
func (t *Segment_PartIndex) MergeSegmentPartIndex1(v SegmentPartIndex1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Segment_PartIndex) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Segment_PartIndex) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSegmentStartIndex0 returns the union data inside the Segment_StartIndex as a SegmentStartIndex0
func (t Segment_StartIndex) AsSegmentStartIndex0() (SegmentStartIndex0, error) {
	var body SegmentStartIndex0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentStartIndex0 overwrites any union data inside the Segment_StartIndex as the provided SegmentStartIndex0
func (t *Segment_StartIndex) FromSegmentStartIndex0(v SegmentStartIndex0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentStartIndex0 performs a merge with any union data inside the Segment_StartIndex, using the provided SegmentStartIndex0
func (t *Segment_StartIndex) MergeSegmentStartIndex0(v SegmentStartIndex0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsSegmentStartIndex1 returns the union data inside the Segment_StartIndex as a SegmentStartIndex1
func (t Segment_StartIndex) AsSegmentStartIndex1() (SegmentStartIndex1, error) {
	var body SegmentStartIndex1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentStartIndex1 overwrites any union data inside the Segment_StartIndex as the provided SegmentStartIndex1
func (t *Segment_StartIndex) FromSegmentStartIndex1(v SegmentStartIndex1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentStartIndex1 performs a merge with any union data inside the Segment_StartIndex, using the provided SegmentStartIndex1
func (t *Segment_StartIndex) MergeSegmentStartIndex1(v SegmentStartIndex1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Segment_StartIndex) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Segment_StartIndex) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsSegmentText0 returns the union data inside the Segment_Text as a SegmentText0
func (t Segment_Text) AsSegmentText0() (SegmentText0, error) {
	var body SegmentText0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentText0 overwrites any union data inside the Segment_Text as the provided SegmentText0
func (t *Segment_Text) FromSegmentText0(v SegmentText0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentText0 performs a merge with any union data inside the Segment_Text, using the provided SegmentText0
func (t *Segment_Text) MergeSegmentText0(v SegmentText0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsSegmentText1 returns the union data inside the Segment_Text as a SegmentText1
func (t Segment_Text) AsSegmentText1() (SegmentText1, error) {
	var body SegmentText1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromSegmentText1 overwrites any union data inside the Segment_Text as the provided SegmentText1
func (t *Segment_Text) FromSegmentText1(v SegmentText1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeSegmentText1 performs a merge with any union data inside the Segment_Text, using the provided SegmentText1
func (t *Segment_Text) MergeSegmentText1(v SegmentText1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t Segment_Text) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *Segment_Text) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsServiceAccountCredential returns the union data inside the ServiceAccount_ServiceAccountCredential as a ServiceAccountCredential
func (t ServiceAccount_ServiceAccountCredential) AsServiceAccountCredential() (ServiceAccountCredential, error) {
	var body ServiceAccountCredential
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromServiceAccountCredential overwrites any union data inside the ServiceAccount_ServiceAccountCredential as the provided ServiceAccountCredential
func (t *ServiceAccount_ServiceAccountCredential) FromServiceAccountCredential(v ServiceAccountCredential) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeServiceAccountCredential performs a merge with any union data inside the ServiceAccount_ServiceAccountCredential, using the provided ServiceAccountCredential
func (t *ServiceAccount_ServiceAccountCredential) MergeServiceAccountCredential(v ServiceAccountCredential) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsServiceAccountServiceAccountCredential1 returns the union data inside the ServiceAccount_ServiceAccountCredential as a ServiceAccountServiceAccountCredential1
func (t ServiceAccount_ServiceAccountCredential) AsServiceAccountServiceAccountCredential1() (ServiceAccountServiceAccountCredential1, error) {
	var body ServiceAccountServiceAccountCredential1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromServiceAccountServiceAccountCredential1 overwrites any union data inside the ServiceAccount_ServiceAccountCredential as the provided ServiceAccountServiceAccountCredential1
func (t *ServiceAccount_ServiceAccountCredential) FromServiceAccountServiceAccountCredential1(v ServiceAccountServiceAccountCredential1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeServiceAccountServiceAccountCredential1 performs a merge with any union data inside the ServiceAccount_ServiceAccountCredential, using the provided ServiceAccountServiceAccountCredential1
func (t *ServiceAccount_ServiceAccountCredential) MergeServiceAccountServiceAccountCredential1(v ServiceAccountServiceAccountCredential1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ServiceAccount_ServiceAccountCredential) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ServiceAccount_ServiceAccountCredential) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsServiceAccountUseDefaultCredential0 returns the union data inside the ServiceAccount_UseDefaultCredential as a ServiceAccountUseDefaultCredential0
func (t ServiceAccount_UseDefaultCredential) AsServiceAccountUseDefaultCredential0() (ServiceAccountUseDefaultCredential0, error) {
	var body ServiceAccountUseDefaultCredential0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromServiceAccountUseDefaultCredential0 overwrites any union data inside the ServiceAccount_UseDefaultCredential as the provided ServiceAccountUseDefaultCredential0
func (t *ServiceAccount_UseDefaultCredential) FromServiceAccountUseDefaultCredential0(v ServiceAccountUseDefaultCredential0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeServiceAccountUseDefaultCredential0 performs a merge with any union data inside the ServiceAccount_UseDefaultCredential, using the provided ServiceAccountUseDefaultCredential0
func (t *ServiceAccount_UseDefaultCredential) MergeServiceAccountUseDefaultCredential0(v ServiceAccountUseDefaultCredential0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsServiceAccountUseDefaultCredential1 returns the union data inside the ServiceAccount_UseDefaultCredential as a ServiceAccountUseDefaultCredential1
func (t ServiceAccount_UseDefaultCredential) AsServiceAccountUseDefaultCredential1() (ServiceAccountUseDefaultCredential1, error) {
	var body ServiceAccountUseDefaultCredential1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromServiceAccountUseDefaultCredential1 overwrites any union data inside the ServiceAccount_UseDefaultCredential as the provided ServiceAccountUseDefaultCredential1
func (t *ServiceAccount_UseDefaultCredential) FromServiceAccountUseDefaultCredential1(v ServiceAccountUseDefaultCredential1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeServiceAccountUseDefaultCredential1 performs a merge with any union data inside the ServiceAccount_UseDefaultCredential, using the provided ServiceAccountUseDefaultCredential1
func (t *ServiceAccount_UseDefaultCredential) MergeServiceAccountUseDefaultCredential1(v ServiceAccountUseDefaultCredential1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ServiceAccount_UseDefaultCredential) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ServiceAccount_UseDefaultCredential) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsValidationErrorLoc0 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc0
func (t ValidationError_Loc_Item) AsValidationErrorLoc0() (ValidationErrorLoc0, error) {
	var body ValidationErrorLoc0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc0 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) FromValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc0 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc0
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc0(v ValidationErrorLoc0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsValidationErrorLoc1 returns the union data inside the ValidationError_Loc_Item as a ValidationErrorLoc1
func (t ValidationError_Loc_Item) AsValidationErrorLoc1() (ValidationErrorLoc1, error) {
	var body ValidationErrorLoc1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromValidationErrorLoc1 overwrites any union data inside the ValidationError_Loc_Item as the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) FromValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeValidationErrorLoc1 performs a merge with any union data inside the ValidationError_Loc_Item, using the provided ValidationErrorLoc1
func (t *ValidationError_Loc_Item) MergeValidationErrorLoc1(v ValidationErrorLoc1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t ValidationError_Loc_Item) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *ValidationError_Loc_Item) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsVideoMetadataEndOffset0 returns the union data inside the VideoMetadata_EndOffset as a VideoMetadataEndOffset0
func (t VideoMetadata_EndOffset) AsVideoMetadataEndOffset0() (VideoMetadataEndOffset0, error) {
	var body VideoMetadataEndOffset0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVideoMetadataEndOffset0 overwrites any union data inside the VideoMetadata_EndOffset as the provided VideoMetadataEndOffset0
func (t *VideoMetadata_EndOffset) FromVideoMetadataEndOffset0(v VideoMetadataEndOffset0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVideoMetadataEndOffset0 performs a merge with any union data inside the VideoMetadata_EndOffset, using the provided VideoMetadataEndOffset0
func (t *VideoMetadata_EndOffset) MergeVideoMetadataEndOffset0(v VideoMetadataEndOffset0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsVideoMetadataEndOffset1 returns the union data inside the VideoMetadata_EndOffset as a VideoMetadataEndOffset1
func (t VideoMetadata_EndOffset) AsVideoMetadataEndOffset1() (VideoMetadataEndOffset1, error) {
	var body VideoMetadataEndOffset1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVideoMetadataEndOffset1 overwrites any union data inside the VideoMetadata_EndOffset as the provided VideoMetadataEndOffset1
func (t *VideoMetadata_EndOffset) FromVideoMetadataEndOffset1(v VideoMetadataEndOffset1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVideoMetadataEndOffset1 performs a merge with any union data inside the VideoMetadata_EndOffset, using the provided VideoMetadataEndOffset1
func (t *VideoMetadata_EndOffset) MergeVideoMetadataEndOffset1(v VideoMetadataEndOffset1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t VideoMetadata_EndOffset) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *VideoMetadata_EndOffset) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// AsVideoMetadataStartOffset0 returns the union data inside the VideoMetadata_StartOffset as a VideoMetadataStartOffset0
func (t VideoMetadata_StartOffset) AsVideoMetadataStartOffset0() (VideoMetadataStartOffset0, error) {
	var body VideoMetadataStartOffset0
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVideoMetadataStartOffset0 overwrites any union data inside the VideoMetadata_StartOffset as the provided VideoMetadataStartOffset0
func (t *VideoMetadata_StartOffset) FromVideoMetadataStartOffset0(v VideoMetadataStartOffset0) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVideoMetadataStartOffset0 performs a merge with any union data inside the VideoMetadata_StartOffset, using the provided VideoMetadataStartOffset0
func (t *VideoMetadata_StartOffset) MergeVideoMetadataStartOffset0(v VideoMetadataStartOffset0) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

// AsVideoMetadataStartOffset1 returns the union data inside the VideoMetadata_StartOffset as a VideoMetadataStartOffset1
func (t VideoMetadata_StartOffset) AsVideoMetadataStartOffset1() (VideoMetadataStartOffset1, error) {
	var body VideoMetadataStartOffset1
	err := json.Unmarshal(t.union, &body)
	return body, err
}

// FromVideoMetadataStartOffset1 overwrites any union data inside the VideoMetadata_StartOffset as the provided VideoMetadataStartOffset1
func (t *VideoMetadata_StartOffset) FromVideoMetadataStartOffset1(v VideoMetadataStartOffset1) error {
	b, err := json.Marshal(v)
	t.union = b
	return err
}

// MergeVideoMetadataStartOffset1 performs a merge with any union data inside the VideoMetadata_StartOffset, using the provided VideoMetadataStartOffset1
func (t *VideoMetadata_StartOffset) MergeVideoMetadataStartOffset1(v VideoMetadataStartOffset1) error {
	b, err := json.Marshal(v)
	if err != nil {
		return err
	}

	merged, err := runtime.JsonMerge(t.union, b)
	t.union = merged
	return err
}

func (t VideoMetadata_StartOffset) MarshalJSON() ([]byte, error) {
	b, err := t.union.MarshalJSON()
	return b, err
}

func (t *VideoMetadata_StartOffset) UnmarshalJSON(b []byte) error {
	err := t.union.UnmarshalJSON(b)
	return err
}

// RequestEditorFn  is the function signature for the RequestEditor callback function
type RequestEditorFn func(ctx context.Context, req *http.Request) error

// Doer performs HTTP requests.
//
// The standard http.Client implements this interface.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// Client which conforms to the OpenAPI3 specification for this service.
type Client struct {
	// The endpoint of the server conforming to this interface, with scheme,
	// https://api.deepmap.com for example. This can contain a path relative
	// to the server, such as https://api.deepmap.com/dev-test, and all the
	// paths in the swagger spec will be appended to the server.
	Server string

	// Doer for performing requests, typically a *http.Client with any
	// customized settings, such as certificate chains.
	Client HttpRequestDoer

	// A list of callbacks for modifying requests which are generated before sending over
	// the network.
	RequestEditors []RequestEditorFn
}

// ClientOption allows setting custom parameters during construction
type ClientOption func(*Client) error

// Creates a new Client, with reasonable defaults
func NewClient(server string, opts ...ClientOption) (*Client, error) {
	// create a client with sane default values
	client := Client{
		Server: server,
	}
	// mutate client and add all optional params
	for _, o := range opts {
		if err := o(&client); err != nil {
			return nil, err
		}
	}
	// ensure the server URL always has a trailing slash
	if !strings.HasSuffix(client.Server, "/") {
		client.Server += "/"
	}
	// create httpClient, if not already present
	if client.Client == nil {
		client.Client = &http.Client{}
	}
	return &client, nil
}

// WithHTTPClient allows overriding the default Doer, which is
// automatically created using http.Client. This is useful for tests.
func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *Client) error {
		c.Client = doer
		return nil
	}
}

// WithRequestEditorFn allows setting up a callback function, which will be
// called right before sending the request. This can be used to mutate the request.
func WithRequestEditorFn(fn RequestEditorFn) ClientOption {
	return func(c *Client) error {
		c.RequestEditors = append(c.RequestEditors, fn)
		return nil
	}
}

// The interface specification for the client above.
type ClientInterface interface {
	// RedirectToDevUi request
	RedirectToDevUi(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListEvalSets request
	ListEvalSets(ctx context.Context, appName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateEvalSet request
	CreateEvalSet(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// AddSessionToEvalSetWithBody request with any body
	AddSessionToEvalSetWithBody(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	AddSessionToEvalSet(ctx context.Context, appName string, evalSetId string, body AddSessionToEvalSetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListEvalsInEvalSet request
	ListEvalsInEvalSet(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RunEvalWithBody request with any body
	RunEvalWithBody(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	RunEval(ctx context.Context, appName string, evalSetId string, body RunEvalJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListSessions request
	ListSessions(ctx context.Context, appName string, userId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateSessionWithBody request with any body
	CreateSessionWithBody(ctx context.Context, appName string, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateSession(ctx context.Context, appName string, userId string, body CreateSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteSession request
	DeleteSession(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetSession request
	GetSession(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// CreateSessionWithIdWithBody request with any body
	CreateSessionWithIdWithBody(ctx context.Context, appName string, userId string, sessionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	CreateSessionWithId(ctx context.Context, appName string, userId string, sessionId string, body CreateSessionWithIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListArtifactNames request
	ListArtifactNames(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DeleteArtifact request
	DeleteArtifact(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// LoadArtifact request
	LoadArtifact(ctx context.Context, appName string, userId string, sessionId string, artifactName string, params *LoadArtifactParams, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListArtifactVersions request
	ListArtifactVersions(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// LoadArtifactVersion request
	LoadArtifactVersion(ctx context.Context, appName string, userId string, sessionId string, artifactName string, versionId int, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetEventGraph request
	GetEventGraph(ctx context.Context, appName string, userId string, sessionId string, eventId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// GetTraceDict request
	GetTraceDict(ctx context.Context, eventId string, reqEditors ...RequestEditorFn) (*http.Response, error)

	// DevUi request
	DevUi(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// ListApps request
	ListApps(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RunWithBody request with any body
	RunWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	Run(ctx context.Context, body RunJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)

	// RunSseWithBody request with any body
	RunSseWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error)

	RunSse(ctx context.Context, body RunSseJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error)
}

func (c *Client) RedirectToDevUi(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRedirectToDevUiRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListEvalSets(ctx context.Context, appName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListEvalSetsRequest(c.Server, appName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateEvalSet(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateEvalSetRequest(c.Server, appName, evalSetId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddSessionToEvalSetWithBody(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddSessionToEvalSetRequestWithBody(c.Server, appName, evalSetId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) AddSessionToEvalSet(ctx context.Context, appName string, evalSetId string, body AddSessionToEvalSetJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewAddSessionToEvalSetRequest(c.Server, appName, evalSetId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListEvalsInEvalSet(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListEvalsInEvalSetRequest(c.Server, appName, evalSetId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RunEvalWithBody(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunEvalRequestWithBody(c.Server, appName, evalSetId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RunEval(ctx context.Context, appName string, evalSetId string, body RunEvalJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunEvalRequest(c.Server, appName, evalSetId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListSessions(ctx context.Context, appName string, userId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListSessionsRequest(c.Server, appName, userId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateSessionWithBody(ctx context.Context, appName string, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSessionRequestWithBody(c.Server, appName, userId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateSession(ctx context.Context, appName string, userId string, body CreateSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSessionRequest(c.Server, appName, userId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteSession(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteSessionRequest(c.Server, appName, userId, sessionId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetSession(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetSessionRequest(c.Server, appName, userId, sessionId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateSessionWithIdWithBody(ctx context.Context, appName string, userId string, sessionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSessionWithIdRequestWithBody(c.Server, appName, userId, sessionId, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) CreateSessionWithId(ctx context.Context, appName string, userId string, sessionId string, body CreateSessionWithIdJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewCreateSessionWithIdRequest(c.Server, appName, userId, sessionId, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListArtifactNames(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListArtifactNamesRequest(c.Server, appName, userId, sessionId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DeleteArtifact(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDeleteArtifactRequest(c.Server, appName, userId, sessionId, artifactName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LoadArtifact(ctx context.Context, appName string, userId string, sessionId string, artifactName string, params *LoadArtifactParams, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoadArtifactRequest(c.Server, appName, userId, sessionId, artifactName, params)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListArtifactVersions(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListArtifactVersionsRequest(c.Server, appName, userId, sessionId, artifactName)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) LoadArtifactVersion(ctx context.Context, appName string, userId string, sessionId string, artifactName string, versionId int, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewLoadArtifactVersionRequest(c.Server, appName, userId, sessionId, artifactName, versionId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetEventGraph(ctx context.Context, appName string, userId string, sessionId string, eventId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetEventGraphRequest(c.Server, appName, userId, sessionId, eventId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) GetTraceDict(ctx context.Context, eventId string, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewGetTraceDictRequest(c.Server, eventId)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) DevUi(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewDevUiRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) ListApps(ctx context.Context, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewListAppsRequest(c.Server)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RunWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) Run(ctx context.Context, body RunJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RunSseWithBody(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunSseRequestWithBody(c.Server, contentType, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

func (c *Client) RunSse(ctx context.Context, body RunSseJSONRequestBody, reqEditors ...RequestEditorFn) (*http.Response, error) {
	req, err := NewRunSseRequest(c.Server, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	if err := c.applyEditors(ctx, req, reqEditors); err != nil {
		return nil, err
	}
	return c.Client.Do(req)
}

// NewRedirectToDevUiRequest generates requests for RedirectToDevUi
func NewRedirectToDevUiRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListEvalSetsRequest generates requests for ListEvalSets
func NewListEvalSetsRequest(server string, appName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/eval_sets", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateEvalSetRequest generates requests for CreateEvalSet
func NewCreateEvalSetRequest(server string, appName string, evalSetId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "eval_set_id", runtime.ParamLocationPath, evalSetId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/eval_sets/%s", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewAddSessionToEvalSetRequest calls the generic AddSessionToEvalSet builder with application/json body
func NewAddSessionToEvalSetRequest(server string, appName string, evalSetId string, body AddSessionToEvalSetJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewAddSessionToEvalSetRequestWithBody(server, appName, evalSetId, "application/json", bodyReader)
}

// NewAddSessionToEvalSetRequestWithBody generates requests for AddSessionToEvalSet with any type of body
func NewAddSessionToEvalSetRequestWithBody(server string, appName string, evalSetId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "eval_set_id", runtime.ParamLocationPath, evalSetId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/eval_sets/%s/add_session", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListEvalsInEvalSetRequest generates requests for ListEvalsInEvalSet
func NewListEvalsInEvalSetRequest(server string, appName string, evalSetId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "eval_set_id", runtime.ParamLocationPath, evalSetId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/eval_sets/%s/evals", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRunEvalRequest calls the generic RunEval builder with application/json body
func NewRunEvalRequest(server string, appName string, evalSetId string, body RunEvalJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRunEvalRequestWithBody(server, appName, evalSetId, "application/json", bodyReader)
}

// NewRunEvalRequestWithBody generates requests for RunEval with any type of body
func NewRunEvalRequestWithBody(server string, appName string, evalSetId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "eval_set_id", runtime.ParamLocationPath, evalSetId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/eval_sets/%s/run_eval", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListSessionsRequest generates requests for ListSessions
func NewListSessionsRequest(server string, appName string, userId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateSessionRequest calls the generic CreateSession builder with application/json body
func NewCreateSessionRequest(server string, appName string, userId string, body CreateSessionJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateSessionRequestWithBody(server, appName, userId, "application/json", bodyReader)
}

// NewCreateSessionRequestWithBody generates requests for CreateSession with any type of body
func NewCreateSessionRequestWithBody(server string, appName string, userId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions", pathParam0, pathParam1)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewDeleteSessionRequest generates requests for DeleteSession
func NewDeleteSessionRequest(server string, appName string, userId string, sessionId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetSessionRequest generates requests for GetSession
func NewGetSessionRequest(server string, appName string, userId string, sessionId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewCreateSessionWithIdRequest calls the generic CreateSessionWithId builder with application/json body
func NewCreateSessionWithIdRequest(server string, appName string, userId string, sessionId string, body CreateSessionWithIdJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewCreateSessionWithIdRequestWithBody(server, appName, userId, sessionId, "application/json", bodyReader)
}

// NewCreateSessionWithIdRequestWithBody generates requests for CreateSessionWithId with any type of body
func NewCreateSessionWithIdRequestWithBody(server string, appName string, userId string, sessionId string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewListArtifactNamesRequest generates requests for ListArtifactNames
func NewListArtifactNamesRequest(server string, appName string, userId string, sessionId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s/artifacts", pathParam0, pathParam1, pathParam2)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDeleteArtifactRequest generates requests for DeleteArtifact
func NewDeleteArtifactRequest(server string, appName string, userId string, sessionId string, artifactName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	var pathParam3 string

	pathParam3, err = runtime.StyleParamWithLocation("simple", false, "artifact_name", runtime.ParamLocationPath, artifactName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s/artifacts/%s", pathParam0, pathParam1, pathParam2, pathParam3)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("DELETE", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewLoadArtifactRequest generates requests for LoadArtifact
func NewLoadArtifactRequest(server string, appName string, userId string, sessionId string, artifactName string, params *LoadArtifactParams) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	var pathParam3 string

	pathParam3, err = runtime.StyleParamWithLocation("simple", false, "artifact_name", runtime.ParamLocationPath, artifactName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s/artifacts/%s", pathParam0, pathParam1, pathParam2, pathParam3)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	if params != nil {
		queryValues := queryURL.Query()

		if params.Version != nil {

			if queryFrag, err := runtime.StyleParamWithLocation("form", true, "version", runtime.ParamLocationQuery, *params.Version); err != nil {
				return nil, err
			} else if parsed, err := url.ParseQuery(queryFrag); err != nil {
				return nil, err
			} else {
				for k, v := range parsed {
					for _, v2 := range v {
						queryValues.Add(k, v2)
					}
				}
			}

		}

		queryURL.RawQuery = queryValues.Encode()
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListArtifactVersionsRequest generates requests for ListArtifactVersions
func NewListArtifactVersionsRequest(server string, appName string, userId string, sessionId string, artifactName string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	var pathParam3 string

	pathParam3, err = runtime.StyleParamWithLocation("simple", false, "artifact_name", runtime.ParamLocationPath, artifactName)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s/artifacts/%s/versions", pathParam0, pathParam1, pathParam2, pathParam3)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewLoadArtifactVersionRequest generates requests for LoadArtifactVersion
func NewLoadArtifactVersionRequest(server string, appName string, userId string, sessionId string, artifactName string, versionId int) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	var pathParam3 string

	pathParam3, err = runtime.StyleParamWithLocation("simple", false, "artifact_name", runtime.ParamLocationPath, artifactName)
	if err != nil {
		return nil, err
	}

	var pathParam4 string

	pathParam4, err = runtime.StyleParamWithLocation("simple", false, "version_id", runtime.ParamLocationPath, versionId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s/artifacts/%s/versions/%s", pathParam0, pathParam1, pathParam2, pathParam3, pathParam4)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetEventGraphRequest generates requests for GetEventGraph
func NewGetEventGraphRequest(server string, appName string, userId string, sessionId string, eventId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "app_name", runtime.ParamLocationPath, appName)
	if err != nil {
		return nil, err
	}

	var pathParam1 string

	pathParam1, err = runtime.StyleParamWithLocation("simple", false, "user_id", runtime.ParamLocationPath, userId)
	if err != nil {
		return nil, err
	}

	var pathParam2 string

	pathParam2, err = runtime.StyleParamWithLocation("simple", false, "session_id", runtime.ParamLocationPath, sessionId)
	if err != nil {
		return nil, err
	}

	var pathParam3 string

	pathParam3, err = runtime.StyleParamWithLocation("simple", false, "event_id", runtime.ParamLocationPath, eventId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/apps/%s/users/%s/sessions/%s/events/%s/graph", pathParam0, pathParam1, pathParam2, pathParam3)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewGetTraceDictRequest generates requests for GetTraceDict
func NewGetTraceDictRequest(server string, eventId string) (*http.Request, error) {
	var err error

	var pathParam0 string

	pathParam0, err = runtime.StyleParamWithLocation("simple", false, "event_id", runtime.ParamLocationPath, eventId)
	if err != nil {
		return nil, err
	}

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/debug/trace/%s", pathParam0)
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewDevUiRequest generates requests for DevUi
func NewDevUiRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/dev-ui")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewListAppsRequest generates requests for ListApps
func NewListAppsRequest(server string) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/list-apps")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", queryURL.String(), nil)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// NewRunRequest calls the generic Run builder with application/json body
func NewRunRequest(server string, body RunJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRunRequestWithBody(server, "application/json", bodyReader)
}

// NewRunRequestWithBody generates requests for Run with any type of body
func NewRunRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/run")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

// NewRunSseRequest calls the generic RunSse builder with application/json body
func NewRunSseRequest(server string, body RunSseJSONRequestBody) (*http.Request, error) {
	var bodyReader io.Reader
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}
	bodyReader = bytes.NewReader(buf)
	return NewRunSseRequestWithBody(server, "application/json", bodyReader)
}

// NewRunSseRequestWithBody generates requests for RunSse with any type of body
func NewRunSseRequestWithBody(server string, contentType string, body io.Reader) (*http.Request, error) {
	var err error

	serverURL, err := url.Parse(server)
	if err != nil {
		return nil, err
	}

	operationPath := fmt.Sprintf("/run_sse")
	if operationPath[0] == '/' {
		operationPath = "." + operationPath
	}

	queryURL, err := serverURL.Parse(operationPath)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", queryURL.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", contentType)

	return req, nil
}

func (c *Client) applyEditors(ctx context.Context, req *http.Request, additionalEditors []RequestEditorFn) error {
	for _, r := range c.RequestEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	for _, r := range additionalEditors {
		if err := r(ctx, req); err != nil {
			return err
		}
	}
	return nil
}

// ClientWithResponses builds on ClientInterface to offer response payloads
type ClientWithResponses struct {
	ClientInterface
}

// NewClientWithResponses creates a new ClientWithResponses, which wraps
// Client with return type handling
func NewClientWithResponses(server string, opts ...ClientOption) (*ClientWithResponses, error) {
	client, err := NewClient(server, opts...)
	if err != nil {
		return nil, err
	}
	return &ClientWithResponses{client}, nil
}

// WithBaseURL overrides the baseURL.
func WithBaseURL(baseURL string) ClientOption {
	return func(c *Client) error {
		newBaseURL, err := url.Parse(baseURL)
		if err != nil {
			return err
		}
		c.Server = newBaseURL.String()
		return nil
	}
}

// ClientWithResponsesInterface is the interface specification for the client with responses above.
type ClientWithResponsesInterface interface {
	// RedirectToDevUiWithResponse request
	RedirectToDevUiWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RedirectToDevUiResponse, error)

	// ListEvalSetsWithResponse request
	ListEvalSetsWithResponse(ctx context.Context, appName string, reqEditors ...RequestEditorFn) (*ListEvalSetsResponse, error)

	// CreateEvalSetWithResponse request
	CreateEvalSetWithResponse(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*CreateEvalSetResponse, error)

	// AddSessionToEvalSetWithBodyWithResponse request with any body
	AddSessionToEvalSetWithBodyWithResponse(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddSessionToEvalSetResponse, error)

	AddSessionToEvalSetWithResponse(ctx context.Context, appName string, evalSetId string, body AddSessionToEvalSetJSONRequestBody, reqEditors ...RequestEditorFn) (*AddSessionToEvalSetResponse, error)

	// ListEvalsInEvalSetWithResponse request
	ListEvalsInEvalSetWithResponse(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*ListEvalsInEvalSetResponse, error)

	// RunEvalWithBodyWithResponse request with any body
	RunEvalWithBodyWithResponse(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunEvalResponse, error)

	RunEvalWithResponse(ctx context.Context, appName string, evalSetId string, body RunEvalJSONRequestBody, reqEditors ...RequestEditorFn) (*RunEvalResponse, error)

	// ListSessionsWithResponse request
	ListSessionsWithResponse(ctx context.Context, appName string, userId string, reqEditors ...RequestEditorFn) (*ListSessionsResponse, error)

	// CreateSessionWithBodyWithResponse request with any body
	CreateSessionWithBodyWithResponse(ctx context.Context, appName string, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSessionResponse, error)

	CreateSessionWithResponse(ctx context.Context, appName string, userId string, body CreateSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSessionResponse, error)

	// DeleteSessionWithResponse request
	DeleteSessionWithResponse(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*DeleteSessionResponse, error)

	// GetSessionWithResponse request
	GetSessionWithResponse(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*GetSessionResponse, error)

	// CreateSessionWithIdWithBodyWithResponse request with any body
	CreateSessionWithIdWithBodyWithResponse(ctx context.Context, appName string, userId string, sessionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSessionWithIdResponse, error)

	CreateSessionWithIdWithResponse(ctx context.Context, appName string, userId string, sessionId string, body CreateSessionWithIdJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSessionWithIdResponse, error)

	// ListArtifactNamesWithResponse request
	ListArtifactNamesWithResponse(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*ListArtifactNamesResponse, error)

	// DeleteArtifactWithResponse request
	DeleteArtifactWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*DeleteArtifactResponse, error)

	// LoadArtifactWithResponse request
	LoadArtifactWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, params *LoadArtifactParams, reqEditors ...RequestEditorFn) (*LoadArtifactResponse, error)

	// ListArtifactVersionsWithResponse request
	ListArtifactVersionsWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*ListArtifactVersionsResponse, error)

	// LoadArtifactVersionWithResponse request
	LoadArtifactVersionWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, versionId int, reqEditors ...RequestEditorFn) (*LoadArtifactVersionResponse, error)

	// GetEventGraphWithResponse request
	GetEventGraphWithResponse(ctx context.Context, appName string, userId string, sessionId string, eventId string, reqEditors ...RequestEditorFn) (*GetEventGraphResponse, error)

	// GetTraceDictWithResponse request
	GetTraceDictWithResponse(ctx context.Context, eventId string, reqEditors ...RequestEditorFn) (*GetTraceDictResponse, error)

	// DevUiWithResponse request
	DevUiWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DevUiResponse, error)

	// ListAppsWithResponse request
	ListAppsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListAppsResponse, error)

	// RunWithBodyWithResponse request with any body
	RunWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunResponse, error)

	RunWithResponse(ctx context.Context, body RunJSONRequestBody, reqEditors ...RequestEditorFn) (*RunResponse, error)

	// RunSseWithBodyWithResponse request with any body
	RunSseWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunSseResponse, error)

	RunSseWithResponse(ctx context.Context, body RunSseJSONRequestBody, reqEditors ...RequestEditorFn) (*RunSseResponse, error)
}

type RedirectToDevUiResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
}

// Status returns HTTPResponse.Status
func (r RedirectToDevUiResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RedirectToDevUiResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListEvalSetsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]string
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ListEvalSetsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListEvalSetsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateEvalSetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r CreateEvalSetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateEvalSetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type AddSessionToEvalSetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r AddSessionToEvalSetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r AddSessionToEvalSetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListEvalsInEvalSetResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]string
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ListEvalsInEvalSetResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListEvalsInEvalSetResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RunEvalResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]RunEvalResult
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r RunEvalResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RunEvalResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListSessionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Session
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ListSessionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListSessionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateSessionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Session
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r CreateSessionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateSessionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteSessionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r DeleteSessionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteSessionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetSessionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Session
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r GetSessionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetSessionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type CreateSessionWithIdResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *Session
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r CreateSessionWithIdResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r CreateSessionWithIdResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListArtifactNamesResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]string
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ListArtifactNamesResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListArtifactNamesResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DeleteArtifactResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r DeleteArtifactResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DeleteArtifactResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type LoadArtifactResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		union json.RawMessage
	}
	JSON422 *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r LoadArtifactResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r LoadArtifactResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListArtifactVersionsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]int
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r ListArtifactVersionsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListArtifactVersionsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type LoadArtifactVersionResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *struct {
		union json.RawMessage
	}
	JSON422 *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r LoadArtifactVersionResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r LoadArtifactVersionResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetEventGraphResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r GetEventGraphResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetEventGraphResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type GetTraceDictResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r GetTraceDictResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r GetTraceDictResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type DevUiResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
}

// Status returns HTTPResponse.Status
func (r DevUiResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r DevUiResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type ListAppsResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]string
}

// Status returns HTTPResponse.Status
func (r ListAppsResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r ListAppsResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RunResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *[]Event
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r RunResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RunResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

type RunSseResponse struct {
	Body         []byte
	HTTPResponse *http.Response
	JSON200      *interface{}
	JSON422      *HTTPValidationError
}

// Status returns HTTPResponse.Status
func (r RunSseResponse) Status() string {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.Status
	}
	return http.StatusText(0)
}

// StatusCode returns HTTPResponse.StatusCode
func (r RunSseResponse) StatusCode() int {
	if r.HTTPResponse != nil {
		return r.HTTPResponse.StatusCode
	}
	return 0
}

// RedirectToDevUiWithResponse request returning *RedirectToDevUiResponse
func (c *ClientWithResponses) RedirectToDevUiWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*RedirectToDevUiResponse, error) {
	rsp, err := c.RedirectToDevUi(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRedirectToDevUiResponse(rsp)
}

// ListEvalSetsWithResponse request returning *ListEvalSetsResponse
func (c *ClientWithResponses) ListEvalSetsWithResponse(ctx context.Context, appName string, reqEditors ...RequestEditorFn) (*ListEvalSetsResponse, error) {
	rsp, err := c.ListEvalSets(ctx, appName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListEvalSetsResponse(rsp)
}

// CreateEvalSetWithResponse request returning *CreateEvalSetResponse
func (c *ClientWithResponses) CreateEvalSetWithResponse(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*CreateEvalSetResponse, error) {
	rsp, err := c.CreateEvalSet(ctx, appName, evalSetId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateEvalSetResponse(rsp)
}

// AddSessionToEvalSetWithBodyWithResponse request with arbitrary body returning *AddSessionToEvalSetResponse
func (c *ClientWithResponses) AddSessionToEvalSetWithBodyWithResponse(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*AddSessionToEvalSetResponse, error) {
	rsp, err := c.AddSessionToEvalSetWithBody(ctx, appName, evalSetId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddSessionToEvalSetResponse(rsp)
}

func (c *ClientWithResponses) AddSessionToEvalSetWithResponse(ctx context.Context, appName string, evalSetId string, body AddSessionToEvalSetJSONRequestBody, reqEditors ...RequestEditorFn) (*AddSessionToEvalSetResponse, error) {
	rsp, err := c.AddSessionToEvalSet(ctx, appName, evalSetId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseAddSessionToEvalSetResponse(rsp)
}

// ListEvalsInEvalSetWithResponse request returning *ListEvalsInEvalSetResponse
func (c *ClientWithResponses) ListEvalsInEvalSetWithResponse(ctx context.Context, appName string, evalSetId string, reqEditors ...RequestEditorFn) (*ListEvalsInEvalSetResponse, error) {
	rsp, err := c.ListEvalsInEvalSet(ctx, appName, evalSetId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListEvalsInEvalSetResponse(rsp)
}

// RunEvalWithBodyWithResponse request with arbitrary body returning *RunEvalResponse
func (c *ClientWithResponses) RunEvalWithBodyWithResponse(ctx context.Context, appName string, evalSetId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunEvalResponse, error) {
	rsp, err := c.RunEvalWithBody(ctx, appName, evalSetId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunEvalResponse(rsp)
}

func (c *ClientWithResponses) RunEvalWithResponse(ctx context.Context, appName string, evalSetId string, body RunEvalJSONRequestBody, reqEditors ...RequestEditorFn) (*RunEvalResponse, error) {
	rsp, err := c.RunEval(ctx, appName, evalSetId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunEvalResponse(rsp)
}

// ListSessionsWithResponse request returning *ListSessionsResponse
func (c *ClientWithResponses) ListSessionsWithResponse(ctx context.Context, appName string, userId string, reqEditors ...RequestEditorFn) (*ListSessionsResponse, error) {
	rsp, err := c.ListSessions(ctx, appName, userId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListSessionsResponse(rsp)
}

// CreateSessionWithBodyWithResponse request with arbitrary body returning *CreateSessionResponse
func (c *ClientWithResponses) CreateSessionWithBodyWithResponse(ctx context.Context, appName string, userId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSessionResponse, error) {
	rsp, err := c.CreateSessionWithBody(ctx, appName, userId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSessionResponse(rsp)
}

func (c *ClientWithResponses) CreateSessionWithResponse(ctx context.Context, appName string, userId string, body CreateSessionJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSessionResponse, error) {
	rsp, err := c.CreateSession(ctx, appName, userId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSessionResponse(rsp)
}

// DeleteSessionWithResponse request returning *DeleteSessionResponse
func (c *ClientWithResponses) DeleteSessionWithResponse(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*DeleteSessionResponse, error) {
	rsp, err := c.DeleteSession(ctx, appName, userId, sessionId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteSessionResponse(rsp)
}

// GetSessionWithResponse request returning *GetSessionResponse
func (c *ClientWithResponses) GetSessionWithResponse(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*GetSessionResponse, error) {
	rsp, err := c.GetSession(ctx, appName, userId, sessionId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetSessionResponse(rsp)
}

// CreateSessionWithIdWithBodyWithResponse request with arbitrary body returning *CreateSessionWithIdResponse
func (c *ClientWithResponses) CreateSessionWithIdWithBodyWithResponse(ctx context.Context, appName string, userId string, sessionId string, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*CreateSessionWithIdResponse, error) {
	rsp, err := c.CreateSessionWithIdWithBody(ctx, appName, userId, sessionId, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSessionWithIdResponse(rsp)
}

func (c *ClientWithResponses) CreateSessionWithIdWithResponse(ctx context.Context, appName string, userId string, sessionId string, body CreateSessionWithIdJSONRequestBody, reqEditors ...RequestEditorFn) (*CreateSessionWithIdResponse, error) {
	rsp, err := c.CreateSessionWithId(ctx, appName, userId, sessionId, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseCreateSessionWithIdResponse(rsp)
}

// ListArtifactNamesWithResponse request returning *ListArtifactNamesResponse
func (c *ClientWithResponses) ListArtifactNamesWithResponse(ctx context.Context, appName string, userId string, sessionId string, reqEditors ...RequestEditorFn) (*ListArtifactNamesResponse, error) {
	rsp, err := c.ListArtifactNames(ctx, appName, userId, sessionId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListArtifactNamesResponse(rsp)
}

// DeleteArtifactWithResponse request returning *DeleteArtifactResponse
func (c *ClientWithResponses) DeleteArtifactWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*DeleteArtifactResponse, error) {
	rsp, err := c.DeleteArtifact(ctx, appName, userId, sessionId, artifactName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDeleteArtifactResponse(rsp)
}

// LoadArtifactWithResponse request returning *LoadArtifactResponse
func (c *ClientWithResponses) LoadArtifactWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, params *LoadArtifactParams, reqEditors ...RequestEditorFn) (*LoadArtifactResponse, error) {
	rsp, err := c.LoadArtifact(ctx, appName, userId, sessionId, artifactName, params, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoadArtifactResponse(rsp)
}

// ListArtifactVersionsWithResponse request returning *ListArtifactVersionsResponse
func (c *ClientWithResponses) ListArtifactVersionsWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, reqEditors ...RequestEditorFn) (*ListArtifactVersionsResponse, error) {
	rsp, err := c.ListArtifactVersions(ctx, appName, userId, sessionId, artifactName, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListArtifactVersionsResponse(rsp)
}

// LoadArtifactVersionWithResponse request returning *LoadArtifactVersionResponse
func (c *ClientWithResponses) LoadArtifactVersionWithResponse(ctx context.Context, appName string, userId string, sessionId string, artifactName string, versionId int, reqEditors ...RequestEditorFn) (*LoadArtifactVersionResponse, error) {
	rsp, err := c.LoadArtifactVersion(ctx, appName, userId, sessionId, artifactName, versionId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseLoadArtifactVersionResponse(rsp)
}

// GetEventGraphWithResponse request returning *GetEventGraphResponse
func (c *ClientWithResponses) GetEventGraphWithResponse(ctx context.Context, appName string, userId string, sessionId string, eventId string, reqEditors ...RequestEditorFn) (*GetEventGraphResponse, error) {
	rsp, err := c.GetEventGraph(ctx, appName, userId, sessionId, eventId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetEventGraphResponse(rsp)
}

// GetTraceDictWithResponse request returning *GetTraceDictResponse
func (c *ClientWithResponses) GetTraceDictWithResponse(ctx context.Context, eventId string, reqEditors ...RequestEditorFn) (*GetTraceDictResponse, error) {
	rsp, err := c.GetTraceDict(ctx, eventId, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseGetTraceDictResponse(rsp)
}

// DevUiWithResponse request returning *DevUiResponse
func (c *ClientWithResponses) DevUiWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*DevUiResponse, error) {
	rsp, err := c.DevUi(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseDevUiResponse(rsp)
}

// ListAppsWithResponse request returning *ListAppsResponse
func (c *ClientWithResponses) ListAppsWithResponse(ctx context.Context, reqEditors ...RequestEditorFn) (*ListAppsResponse, error) {
	rsp, err := c.ListApps(ctx, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseListAppsResponse(rsp)
}

// RunWithBodyWithResponse request with arbitrary body returning *RunResponse
func (c *ClientWithResponses) RunWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunResponse, error) {
	rsp, err := c.RunWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunResponse(rsp)
}

func (c *ClientWithResponses) RunWithResponse(ctx context.Context, body RunJSONRequestBody, reqEditors ...RequestEditorFn) (*RunResponse, error) {
	rsp, err := c.Run(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunResponse(rsp)
}

// RunSseWithBodyWithResponse request with arbitrary body returning *RunSseResponse
func (c *ClientWithResponses) RunSseWithBodyWithResponse(ctx context.Context, contentType string, body io.Reader, reqEditors ...RequestEditorFn) (*RunSseResponse, error) {
	rsp, err := c.RunSseWithBody(ctx, contentType, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunSseResponse(rsp)
}

func (c *ClientWithResponses) RunSseWithResponse(ctx context.Context, body RunSseJSONRequestBody, reqEditors ...RequestEditorFn) (*RunSseResponse, error) {
	rsp, err := c.RunSse(ctx, body, reqEditors...)
	if err != nil {
		return nil, err
	}
	return ParseRunSseResponse(rsp)
}

// ParseRedirectToDevUiResponse parses an HTTP response from a RedirectToDevUiWithResponse call
func ParseRedirectToDevUiResponse(rsp *http.Response) (*RedirectToDevUiResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RedirectToDevUiResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseListEvalSetsResponse parses an HTTP response from a ListEvalSetsWithResponse call
func ParseListEvalSetsResponse(rsp *http.Response) (*ListEvalSetsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListEvalSetsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseCreateEvalSetResponse parses an HTTP response from a CreateEvalSetWithResponse call
func ParseCreateEvalSetResponse(rsp *http.Response) (*CreateEvalSetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateEvalSetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseAddSessionToEvalSetResponse parses an HTTP response from a AddSessionToEvalSetWithResponse call
func ParseAddSessionToEvalSetResponse(rsp *http.Response) (*AddSessionToEvalSetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &AddSessionToEvalSetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseListEvalsInEvalSetResponse parses an HTTP response from a ListEvalsInEvalSetWithResponse call
func ParseListEvalsInEvalSetResponse(rsp *http.Response) (*ListEvalsInEvalSetResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListEvalsInEvalSetResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseRunEvalResponse parses an HTTP response from a RunEvalWithResponse call
func ParseRunEvalResponse(rsp *http.Response) (*RunEvalResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RunEvalResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []RunEvalResult
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseListSessionsResponse parses an HTTP response from a ListSessionsWithResponse call
func ParseListSessionsResponse(rsp *http.Response) (*ListSessionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListSessionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Session
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseCreateSessionResponse parses an HTTP response from a CreateSessionWithResponse call
func ParseCreateSessionResponse(rsp *http.Response) (*CreateSessionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateSessionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Session
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseDeleteSessionResponse parses an HTTP response from a DeleteSessionWithResponse call
func ParseDeleteSessionResponse(rsp *http.Response) (*DeleteSessionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteSessionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseGetSessionResponse parses an HTTP response from a GetSessionWithResponse call
func ParseGetSessionResponse(rsp *http.Response) (*GetSessionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetSessionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Session
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseCreateSessionWithIdResponse parses an HTTP response from a CreateSessionWithIdWithResponse call
func ParseCreateSessionWithIdResponse(rsp *http.Response) (*CreateSessionWithIdResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &CreateSessionWithIdResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest Session
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseListArtifactNamesResponse parses an HTTP response from a ListArtifactNamesWithResponse call
func ParseListArtifactNamesResponse(rsp *http.Response) (*ListArtifactNamesResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListArtifactNamesResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseDeleteArtifactResponse parses an HTTP response from a DeleteArtifactWithResponse call
func ParseDeleteArtifactResponse(rsp *http.Response) (*DeleteArtifactResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DeleteArtifactResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseLoadArtifactResponse parses an HTTP response from a LoadArtifactWithResponse call
func ParseLoadArtifactResponse(rsp *http.Response) (*LoadArtifactResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &LoadArtifactResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			union json.RawMessage
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseListArtifactVersionsResponse parses an HTTP response from a ListArtifactVersionsWithResponse call
func ParseListArtifactVersionsResponse(rsp *http.Response) (*ListArtifactVersionsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListArtifactVersionsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []int
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseLoadArtifactVersionResponse parses an HTTP response from a LoadArtifactVersionWithResponse call
func ParseLoadArtifactVersionResponse(rsp *http.Response) (*LoadArtifactVersionResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &LoadArtifactVersionResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest struct {
			union json.RawMessage
		}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseGetEventGraphResponse parses an HTTP response from a GetEventGraphWithResponse call
func ParseGetEventGraphResponse(rsp *http.Response) (*GetEventGraphResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetEventGraphResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseGetTraceDictResponse parses an HTTP response from a GetTraceDictWithResponse call
func ParseGetTraceDictResponse(rsp *http.Response) (*GetTraceDictResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &GetTraceDictResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseDevUiResponse parses an HTTP response from a DevUiWithResponse call
func ParseDevUiResponse(rsp *http.Response) (*DevUiResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &DevUiResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseListAppsResponse parses an HTTP response from a ListAppsWithResponse call
func ParseListAppsResponse(rsp *http.Response) (*ListAppsResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &ListAppsResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []string
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	}

	return response, nil
}

// ParseRunResponse parses an HTTP response from a RunWithResponse call
func ParseRunResponse(rsp *http.Response) (*RunResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RunResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest []Event
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}

// ParseRunSseResponse parses an HTTP response from a RunSseWithResponse call
func ParseRunSseResponse(rsp *http.Response) (*RunSseResponse, error) {
	bodyBytes, err := io.ReadAll(rsp.Body)
	defer func() { _ = rsp.Body.Close() }()
	if err != nil {
		return nil, err
	}

	response := &RunSseResponse{
		Body:         bodyBytes,
		HTTPResponse: rsp,
	}

	switch {
	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 200:
		var dest interface{}
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON200 = &dest

	case strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode == 422:
		var dest HTTPValidationError
		if err := json.Unmarshal(bodyBytes, &dest); err != nil {
			return nil, err
		}
		response.JSON422 = &dest

	}

	return response, nil
}
