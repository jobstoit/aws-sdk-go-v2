// Code generated by smithy-go-codegen DO NOT EDIT.
package restxml

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_HttpPayloadTraitsWithMediaType_awsRestxmlSerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *HttpPayloadTraitsWithMediaTypeInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes a blob in the HTTP payload with a content-type
		"HttpPayloadTraitsWithMediaTypeWithBlob": {
			Params: &HttpPayloadTraitsWithMediaTypeInput{
				Foo:  ptr.String("Foo"),
				Blob: []byte("blobby blob blob"),
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/HttpPayloadTraitsWithMediaType",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"text/plain"},
				"X-Foo":        []string{"Foo"},
			},
			RequireHeader: []string{
				"Content-Length",
			},
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareReaderBytes(actual, []byte(`blobby blob blob`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.HttpPayloadTraitsWithMediaType(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if actualReq.Body != nil {
				defer actualReq.Body.Close()
			}
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}

func TestClient_HttpPayloadTraitsWithMediaType_awsRestxmlDeserialize(t *testing.T) {
	cases := map[string]struct {
		StatusCode    int
		Header        http.Header
		BodyMediaType string
		Body          []byte
		ExpectResult  *HttpPayloadTraitsWithMediaTypeOutput
	}{
		// Serializes a blob in the HTTP payload with a content-type
		"HttpPayloadTraitsWithMediaTypeWithBlob": {
			StatusCode: 200,
			Header: http.Header{
				"Content-Type": []string{"text/plain"},
				"X-Foo":        []string{"Foo"},
			},
			Body: []byte(`blobby blob blob`),
			ExpectResult: &HttpPayloadTraitsWithMediaTypeOutput{
				Foo:  ptr.String("Foo"),
				Blob: []byte("blobby blob blob"),
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				for k, vs := range c.Header {
					for _, v := range vs {
						w.Header().Add(k, v)
					}
				}
				if len(c.BodyMediaType) != 0 && len(w.Header().Values("Content-Type")) == 0 {
					w.Header().Set("Content-Type", c.BodyMediaType)
				}
				if len(c.Body) != 0 {
					w.Header().Set("Content-Length", strconv.Itoa(len(c.Body)))
				}
				w.WriteHeader(c.StatusCode)
				if len(c.Body) != 0 {
					if _, err := io.Copy(w, bytes.NewReader(c.Body)); err != nil {
						t.Errorf("failed to write response body, %v", err)
					}
				}
			}))
			defer server.Close()
			client := New(Options{
				APIOptions: []APIOptionFunc{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: aws.EndpointResolverFunc(func(service, region string) (e aws.Endpoint, err error) {
					e.URL = server.URL
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               aws.NewBuildableHTTPClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			var params HttpPayloadTraitsWithMediaTypeInput
			result, err := client.HttpPayloadTraitsWithMediaType(context.Background(), &params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if diff := cmp.Diff(c.ExpectResult, result, cmpopts.IgnoreUnexported(middleware.Metadata{})); len(diff) != 0 {
				t.Errorf("expect c.ExpectResult value match:\n%s", diff)
			}
		})
	}
}
