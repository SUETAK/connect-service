// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: eliza/v1/eliza.proto

package elizav1connect

import (
	context "context"
	errors "errors"
	v1 "example/gen/eliza/v1"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// ElizaServiceName is the fully-qualified name of the ElizaService service.
	ElizaServiceName = "eliza.v1.ElizaService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// ElizaServiceSayProcedure is the fully-qualified name of the ElizaService's Say RPC.
	ElizaServiceSayProcedure = "/eliza.v1.ElizaService/Say"
	// ElizaServiceHelloProcedure is the fully-qualified name of the ElizaService's Hello RPC.
	ElizaServiceHelloProcedure = "/eliza.v1.ElizaService/Hello"
)

// ElizaServiceClient is a client for the eliza.v1.ElizaService service.
type ElizaServiceClient interface {
	Say(context.Context, *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error)
	Hello(context.Context, *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error)
}

// NewElizaServiceClient constructs a client for the eliza.v1.ElizaService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewElizaServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) ElizaServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &elizaServiceClient{
		say: connect_go.NewClient[v1.SayRequest, v1.SayResponse](
			httpClient,
			baseURL+ElizaServiceSayProcedure,
			opts...,
		),
		hello: connect_go.NewClient[v1.HelloRequest, v1.HelloResponse](
			httpClient,
			baseURL+ElizaServiceHelloProcedure,
			opts...,
		),
	}
}

// elizaServiceClient implements ElizaServiceClient.
type elizaServiceClient struct {
	say   *connect_go.Client[v1.SayRequest, v1.SayResponse]
	hello *connect_go.Client[v1.HelloRequest, v1.HelloResponse]
}

// Say calls eliza.v1.ElizaService.Say.
func (c *elizaServiceClient) Say(ctx context.Context, req *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error) {
	return c.say.CallUnary(ctx, req)
}

// Hello calls eliza.v1.ElizaService.Hello.
func (c *elizaServiceClient) Hello(ctx context.Context, req *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error) {
	return c.hello.CallUnary(ctx, req)
}

// ElizaServiceHandler is an implementation of the eliza.v1.ElizaService service.
type ElizaServiceHandler interface {
	Say(context.Context, *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error)
	Hello(context.Context, *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error)
}

// NewElizaServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewElizaServiceHandler(svc ElizaServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	elizaServiceSayHandler := connect_go.NewUnaryHandler(
		ElizaServiceSayProcedure,
		svc.Say,
		opts...,
	)
	elizaServiceHelloHandler := connect_go.NewUnaryHandler(
		ElizaServiceHelloProcedure,
		svc.Hello,
		opts...,
	)
	return "/eliza.v1.ElizaService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case ElizaServiceSayProcedure:
			elizaServiceSayHandler.ServeHTTP(w, r)
		case ElizaServiceHelloProcedure:
			elizaServiceHelloHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedElizaServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedElizaServiceHandler struct{}

func (UnimplementedElizaServiceHandler) Say(context.Context, *connect_go.Request[v1.SayRequest]) (*connect_go.Response[v1.SayResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("eliza.v1.ElizaService.Say is not implemented"))
}

func (UnimplementedElizaServiceHandler) Hello(context.Context, *connect_go.Request[v1.HelloRequest]) (*connect_go.Response[v1.HelloResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("eliza.v1.ElizaService.Hello is not implemented"))
}
