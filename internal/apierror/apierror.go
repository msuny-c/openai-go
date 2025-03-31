// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"fmt"
	"net/http"

	// Removing httputil import as it's not supported by TinyGo

	"github.com/openai/openai-go/internal/apijson"
	"github.com/openai/openai-go/packages/resp"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	Code    string `json:"code,required"`
	Message string `json:"message,required"`
	Param   string `json:"param,required"`
	Type    string `json:"type,required"`
	// Metadata for the response, check the presence of optional fields with the
	// [resp.Field.IsPresent] method.
	JSON struct {
		Code        resp.Field
		Message     resp.Field
		Param       resp.Field
		Type        resp.Field
		ExtraFields map[string]resp.Field
		raw         string
	} `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s %q: %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.raw)
}

// TinyGo-compatible version that doesn't use httputil
func (r *Error) DumpRequest(body bool) []byte {
	if r.Request == nil {
		return []byte("Request is nil")
	}

	var result string
	result = fmt.Sprintf("%s %s HTTP/%d.%d\r\n", r.Request.Method, r.Request.URL.Path, r.Request.ProtoMajor, r.Request.ProtoMinor)

	// Add headers
	for key, values := range r.Request.Header {
		for _, value := range values {
			result += fmt.Sprintf("%s: %s\r\n", key, value)
		}
	}

	// We're not including body content as it would require more complex handling
	return []byte(result)
}

// TinyGo-compatible version that doesn't use httputil
func (r *Error) DumpResponse(body bool) []byte {
	if r.Response == nil {
		return []byte("Response is nil")
	}

	var result string
	result = fmt.Sprintf("HTTP/%d.%d %d %s\r\n", r.Response.ProtoMajor, r.Response.ProtoMinor, r.Response.StatusCode, r.Response.Status)

	// Add headers
	for key, values := range r.Response.Header {
		for _, value := range values {
			result += fmt.Sprintf("%s: %s\r\n", key, value)
		}
	}

	// We're not including body content as it would require more complex handling
	return []byte(result)
}
