// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: rpc/order/v1/order.proto

package orderv1connect

import (
	connect "connectrpc.com/connect"
	context "context"
	errors "errors"
	v1 "github.com/imkrish7/ecoship-api/gen/rpc/order/v1"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect.IsAtLeastVersion1_13_0

const (
	// OrderServiceName is the fully-qualified name of the OrderService service.
	OrderServiceName = "order.v1.OrderService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// OrderServiceCreateOrderProcedure is the fully-qualified name of the OrderService's CreateOrder
	// RPC.
	OrderServiceCreateOrderProcedure = "/order.v1.OrderService/CreateOrder"
	// OrderServiceAddToCartProcedure is the fully-qualified name of the OrderService's AddToCart RPC.
	OrderServiceAddToCartProcedure = "/order.v1.OrderService/AddToCart"
	// OrderServiceGetOrderByIdProcedure is the fully-qualified name of the OrderService's GetOrderById
	// RPC.
	OrderServiceGetOrderByIdProcedure = "/order.v1.OrderService/GetOrderById"
	// OrderServiceGetOrdersProcedure is the fully-qualified name of the OrderService's GetOrders RPC.
	OrderServiceGetOrdersProcedure = "/order.v1.OrderService/GetOrders"
	// OrderServiceCancelOrderProcedure is the fully-qualified name of the OrderService's CancelOrder
	// RPC.
	OrderServiceCancelOrderProcedure = "/order.v1.OrderService/CancelOrder"
	// OrderServiceCancelOrderItemProcedure is the fully-qualified name of the OrderService's
	// CancelOrderItem RPC.
	OrderServiceCancelOrderItemProcedure = "/order.v1.OrderService/CancelOrderItem"
	// OrderServiceRemoveItemFromCartProcedure is the fully-qualified name of the OrderService's
	// RemoveItemFromCart RPC.
	OrderServiceRemoveItemFromCartProcedure = "/order.v1.OrderService/RemoveItemFromCart"
	// OrderServiceReturnOrderItemsProcedure is the fully-qualified name of the OrderService's
	// ReturnOrderItems RPC.
	OrderServiceReturnOrderItemsProcedure = "/order.v1.OrderService/ReturnOrderItems"
	// OrderServiceIncreaseItemQuantityProcedure is the fully-qualified name of the OrderService's
	// IncreaseItemQuantity RPC.
	OrderServiceIncreaseItemQuantityProcedure = "/order.v1.OrderService/IncreaseItemQuantity"
	// OrderServiceDecreaseItemQuantityProcedure is the fully-qualified name of the OrderService's
	// DecreaseItemQuantity RPC.
	OrderServiceDecreaseItemQuantityProcedure = "/order.v1.OrderService/DecreaseItemQuantity"
)

// These variables are the protoreflect.Descriptor objects for the RPCs defined in this package.
var (
	orderServiceServiceDescriptor                    = v1.File_rpc_order_v1_order_proto.Services().ByName("OrderService")
	orderServiceCreateOrderMethodDescriptor          = orderServiceServiceDescriptor.Methods().ByName("CreateOrder")
	orderServiceAddToCartMethodDescriptor            = orderServiceServiceDescriptor.Methods().ByName("AddToCart")
	orderServiceGetOrderByIdMethodDescriptor         = orderServiceServiceDescriptor.Methods().ByName("GetOrderById")
	orderServiceGetOrdersMethodDescriptor            = orderServiceServiceDescriptor.Methods().ByName("GetOrders")
	orderServiceCancelOrderMethodDescriptor          = orderServiceServiceDescriptor.Methods().ByName("CancelOrder")
	orderServiceCancelOrderItemMethodDescriptor      = orderServiceServiceDescriptor.Methods().ByName("CancelOrderItem")
	orderServiceRemoveItemFromCartMethodDescriptor   = orderServiceServiceDescriptor.Methods().ByName("RemoveItemFromCart")
	orderServiceReturnOrderItemsMethodDescriptor     = orderServiceServiceDescriptor.Methods().ByName("ReturnOrderItems")
	orderServiceIncreaseItemQuantityMethodDescriptor = orderServiceServiceDescriptor.Methods().ByName("IncreaseItemQuantity")
	orderServiceDecreaseItemQuantityMethodDescriptor = orderServiceServiceDescriptor.Methods().ByName("DecreaseItemQuantity")
)

// OrderServiceClient is a client for the order.v1.OrderService service.
type OrderServiceClient interface {
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error)
	AddToCart(context.Context, *connect.Request[v1.AddToCartRequest]) (*connect.Response[v1.AddToCartResponse], error)
	GetOrderById(context.Context, *connect.Request[v1.GetOrderByIdRequest]) (*connect.Response[v1.GetOrderByIdResponse], error)
	GetOrders(context.Context, *connect.Request[v1.GetOrdersRequest]) (*connect.Response[v1.GetOrdersResponse], error)
	CancelOrder(context.Context, *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error)
	CancelOrderItem(context.Context, *connect.Request[v1.CancelOrderItemRequest]) (*connect.Response[v1.CancelOrderItemResponse], error)
	RemoveItemFromCart(context.Context, *connect.Request[v1.RemoveItemFromCartRequest]) (*connect.Response[v1.RemoveItemFromCartResponse], error)
	ReturnOrderItems(context.Context, *connect.Request[v1.ReturnOrderItemsRequest]) (*connect.Response[v1.ReturnOrderItemsResponse], error)
	IncreaseItemQuantity(context.Context, *connect.Request[v1.IncreaseItemQuantityRequest]) (*connect.Response[v1.IncreaseItemQuantityResponse], error)
	DecreaseItemQuantity(context.Context, *connect.Request[v1.DecreaseItemQuantityRequest]) (*connect.Response[v1.DecreaseItemQuantityResponse], error)
}

// NewOrderServiceClient constructs a client for the order.v1.OrderService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewOrderServiceClient(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) OrderServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &orderServiceClient{
		createOrder: connect.NewClient[v1.CreateOrderRequest, v1.CreateOrderResponse](
			httpClient,
			baseURL+OrderServiceCreateOrderProcedure,
			connect.WithSchema(orderServiceCreateOrderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		addToCart: connect.NewClient[v1.AddToCartRequest, v1.AddToCartResponse](
			httpClient,
			baseURL+OrderServiceAddToCartProcedure,
			connect.WithSchema(orderServiceAddToCartMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOrderById: connect.NewClient[v1.GetOrderByIdRequest, v1.GetOrderByIdResponse](
			httpClient,
			baseURL+OrderServiceGetOrderByIdProcedure,
			connect.WithSchema(orderServiceGetOrderByIdMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		getOrders: connect.NewClient[v1.GetOrdersRequest, v1.GetOrdersResponse](
			httpClient,
			baseURL+OrderServiceGetOrdersProcedure,
			connect.WithSchema(orderServiceGetOrdersMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		cancelOrder: connect.NewClient[v1.CancelOrderRequest, v1.CancelOrderResponse](
			httpClient,
			baseURL+OrderServiceCancelOrderProcedure,
			connect.WithSchema(orderServiceCancelOrderMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		cancelOrderItem: connect.NewClient[v1.CancelOrderItemRequest, v1.CancelOrderItemResponse](
			httpClient,
			baseURL+OrderServiceCancelOrderItemProcedure,
			connect.WithSchema(orderServiceCancelOrderItemMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		removeItemFromCart: connect.NewClient[v1.RemoveItemFromCartRequest, v1.RemoveItemFromCartResponse](
			httpClient,
			baseURL+OrderServiceRemoveItemFromCartProcedure,
			connect.WithSchema(orderServiceRemoveItemFromCartMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		returnOrderItems: connect.NewClient[v1.ReturnOrderItemsRequest, v1.ReturnOrderItemsResponse](
			httpClient,
			baseURL+OrderServiceReturnOrderItemsProcedure,
			connect.WithSchema(orderServiceReturnOrderItemsMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		increaseItemQuantity: connect.NewClient[v1.IncreaseItemQuantityRequest, v1.IncreaseItemQuantityResponse](
			httpClient,
			baseURL+OrderServiceIncreaseItemQuantityProcedure,
			connect.WithSchema(orderServiceIncreaseItemQuantityMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
		decreaseItemQuantity: connect.NewClient[v1.DecreaseItemQuantityRequest, v1.DecreaseItemQuantityResponse](
			httpClient,
			baseURL+OrderServiceDecreaseItemQuantityProcedure,
			connect.WithSchema(orderServiceDecreaseItemQuantityMethodDescriptor),
			connect.WithClientOptions(opts...),
		),
	}
}

// orderServiceClient implements OrderServiceClient.
type orderServiceClient struct {
	createOrder          *connect.Client[v1.CreateOrderRequest, v1.CreateOrderResponse]
	addToCart            *connect.Client[v1.AddToCartRequest, v1.AddToCartResponse]
	getOrderById         *connect.Client[v1.GetOrderByIdRequest, v1.GetOrderByIdResponse]
	getOrders            *connect.Client[v1.GetOrdersRequest, v1.GetOrdersResponse]
	cancelOrder          *connect.Client[v1.CancelOrderRequest, v1.CancelOrderResponse]
	cancelOrderItem      *connect.Client[v1.CancelOrderItemRequest, v1.CancelOrderItemResponse]
	removeItemFromCart   *connect.Client[v1.RemoveItemFromCartRequest, v1.RemoveItemFromCartResponse]
	returnOrderItems     *connect.Client[v1.ReturnOrderItemsRequest, v1.ReturnOrderItemsResponse]
	increaseItemQuantity *connect.Client[v1.IncreaseItemQuantityRequest, v1.IncreaseItemQuantityResponse]
	decreaseItemQuantity *connect.Client[v1.DecreaseItemQuantityRequest, v1.DecreaseItemQuantityResponse]
}

// CreateOrder calls order.v1.OrderService.CreateOrder.
func (c *orderServiceClient) CreateOrder(ctx context.Context, req *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error) {
	return c.createOrder.CallUnary(ctx, req)
}

// AddToCart calls order.v1.OrderService.AddToCart.
func (c *orderServiceClient) AddToCart(ctx context.Context, req *connect.Request[v1.AddToCartRequest]) (*connect.Response[v1.AddToCartResponse], error) {
	return c.addToCart.CallUnary(ctx, req)
}

// GetOrderById calls order.v1.OrderService.GetOrderById.
func (c *orderServiceClient) GetOrderById(ctx context.Context, req *connect.Request[v1.GetOrderByIdRequest]) (*connect.Response[v1.GetOrderByIdResponse], error) {
	return c.getOrderById.CallUnary(ctx, req)
}

// GetOrders calls order.v1.OrderService.GetOrders.
func (c *orderServiceClient) GetOrders(ctx context.Context, req *connect.Request[v1.GetOrdersRequest]) (*connect.Response[v1.GetOrdersResponse], error) {
	return c.getOrders.CallUnary(ctx, req)
}

// CancelOrder calls order.v1.OrderService.CancelOrder.
func (c *orderServiceClient) CancelOrder(ctx context.Context, req *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error) {
	return c.cancelOrder.CallUnary(ctx, req)
}

// CancelOrderItem calls order.v1.OrderService.CancelOrderItem.
func (c *orderServiceClient) CancelOrderItem(ctx context.Context, req *connect.Request[v1.CancelOrderItemRequest]) (*connect.Response[v1.CancelOrderItemResponse], error) {
	return c.cancelOrderItem.CallUnary(ctx, req)
}

// RemoveItemFromCart calls order.v1.OrderService.RemoveItemFromCart.
func (c *orderServiceClient) RemoveItemFromCart(ctx context.Context, req *connect.Request[v1.RemoveItemFromCartRequest]) (*connect.Response[v1.RemoveItemFromCartResponse], error) {
	return c.removeItemFromCart.CallUnary(ctx, req)
}

// ReturnOrderItems calls order.v1.OrderService.ReturnOrderItems.
func (c *orderServiceClient) ReturnOrderItems(ctx context.Context, req *connect.Request[v1.ReturnOrderItemsRequest]) (*connect.Response[v1.ReturnOrderItemsResponse], error) {
	return c.returnOrderItems.CallUnary(ctx, req)
}

// IncreaseItemQuantity calls order.v1.OrderService.IncreaseItemQuantity.
func (c *orderServiceClient) IncreaseItemQuantity(ctx context.Context, req *connect.Request[v1.IncreaseItemQuantityRequest]) (*connect.Response[v1.IncreaseItemQuantityResponse], error) {
	return c.increaseItemQuantity.CallUnary(ctx, req)
}

// DecreaseItemQuantity calls order.v1.OrderService.DecreaseItemQuantity.
func (c *orderServiceClient) DecreaseItemQuantity(ctx context.Context, req *connect.Request[v1.DecreaseItemQuantityRequest]) (*connect.Response[v1.DecreaseItemQuantityResponse], error) {
	return c.decreaseItemQuantity.CallUnary(ctx, req)
}

// OrderServiceHandler is an implementation of the order.v1.OrderService service.
type OrderServiceHandler interface {
	CreateOrder(context.Context, *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error)
	AddToCart(context.Context, *connect.Request[v1.AddToCartRequest]) (*connect.Response[v1.AddToCartResponse], error)
	GetOrderById(context.Context, *connect.Request[v1.GetOrderByIdRequest]) (*connect.Response[v1.GetOrderByIdResponse], error)
	GetOrders(context.Context, *connect.Request[v1.GetOrdersRequest]) (*connect.Response[v1.GetOrdersResponse], error)
	CancelOrder(context.Context, *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error)
	CancelOrderItem(context.Context, *connect.Request[v1.CancelOrderItemRequest]) (*connect.Response[v1.CancelOrderItemResponse], error)
	RemoveItemFromCart(context.Context, *connect.Request[v1.RemoveItemFromCartRequest]) (*connect.Response[v1.RemoveItemFromCartResponse], error)
	ReturnOrderItems(context.Context, *connect.Request[v1.ReturnOrderItemsRequest]) (*connect.Response[v1.ReturnOrderItemsResponse], error)
	IncreaseItemQuantity(context.Context, *connect.Request[v1.IncreaseItemQuantityRequest]) (*connect.Response[v1.IncreaseItemQuantityResponse], error)
	DecreaseItemQuantity(context.Context, *connect.Request[v1.DecreaseItemQuantityRequest]) (*connect.Response[v1.DecreaseItemQuantityResponse], error)
}

// NewOrderServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewOrderServiceHandler(svc OrderServiceHandler, opts ...connect.HandlerOption) (string, http.Handler) {
	orderServiceCreateOrderHandler := connect.NewUnaryHandler(
		OrderServiceCreateOrderProcedure,
		svc.CreateOrder,
		connect.WithSchema(orderServiceCreateOrderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceAddToCartHandler := connect.NewUnaryHandler(
		OrderServiceAddToCartProcedure,
		svc.AddToCart,
		connect.WithSchema(orderServiceAddToCartMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceGetOrderByIdHandler := connect.NewUnaryHandler(
		OrderServiceGetOrderByIdProcedure,
		svc.GetOrderById,
		connect.WithSchema(orderServiceGetOrderByIdMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceGetOrdersHandler := connect.NewUnaryHandler(
		OrderServiceGetOrdersProcedure,
		svc.GetOrders,
		connect.WithSchema(orderServiceGetOrdersMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceCancelOrderHandler := connect.NewUnaryHandler(
		OrderServiceCancelOrderProcedure,
		svc.CancelOrder,
		connect.WithSchema(orderServiceCancelOrderMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceCancelOrderItemHandler := connect.NewUnaryHandler(
		OrderServiceCancelOrderItemProcedure,
		svc.CancelOrderItem,
		connect.WithSchema(orderServiceCancelOrderItemMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceRemoveItemFromCartHandler := connect.NewUnaryHandler(
		OrderServiceRemoveItemFromCartProcedure,
		svc.RemoveItemFromCart,
		connect.WithSchema(orderServiceRemoveItemFromCartMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceReturnOrderItemsHandler := connect.NewUnaryHandler(
		OrderServiceReturnOrderItemsProcedure,
		svc.ReturnOrderItems,
		connect.WithSchema(orderServiceReturnOrderItemsMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceIncreaseItemQuantityHandler := connect.NewUnaryHandler(
		OrderServiceIncreaseItemQuantityProcedure,
		svc.IncreaseItemQuantity,
		connect.WithSchema(orderServiceIncreaseItemQuantityMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	orderServiceDecreaseItemQuantityHandler := connect.NewUnaryHandler(
		OrderServiceDecreaseItemQuantityProcedure,
		svc.DecreaseItemQuantity,
		connect.WithSchema(orderServiceDecreaseItemQuantityMethodDescriptor),
		connect.WithHandlerOptions(opts...),
	)
	return "/order.v1.OrderService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case OrderServiceCreateOrderProcedure:
			orderServiceCreateOrderHandler.ServeHTTP(w, r)
		case OrderServiceAddToCartProcedure:
			orderServiceAddToCartHandler.ServeHTTP(w, r)
		case OrderServiceGetOrderByIdProcedure:
			orderServiceGetOrderByIdHandler.ServeHTTP(w, r)
		case OrderServiceGetOrdersProcedure:
			orderServiceGetOrdersHandler.ServeHTTP(w, r)
		case OrderServiceCancelOrderProcedure:
			orderServiceCancelOrderHandler.ServeHTTP(w, r)
		case OrderServiceCancelOrderItemProcedure:
			orderServiceCancelOrderItemHandler.ServeHTTP(w, r)
		case OrderServiceRemoveItemFromCartProcedure:
			orderServiceRemoveItemFromCartHandler.ServeHTTP(w, r)
		case OrderServiceReturnOrderItemsProcedure:
			orderServiceReturnOrderItemsHandler.ServeHTTP(w, r)
		case OrderServiceIncreaseItemQuantityProcedure:
			orderServiceIncreaseItemQuantityHandler.ServeHTTP(w, r)
		case OrderServiceDecreaseItemQuantityProcedure:
			orderServiceDecreaseItemQuantityHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedOrderServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedOrderServiceHandler struct{}

func (UnimplementedOrderServiceHandler) CreateOrder(context.Context, *connect.Request[v1.CreateOrderRequest]) (*connect.Response[v1.CreateOrderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.CreateOrder is not implemented"))
}

func (UnimplementedOrderServiceHandler) AddToCart(context.Context, *connect.Request[v1.AddToCartRequest]) (*connect.Response[v1.AddToCartResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.AddToCart is not implemented"))
}

func (UnimplementedOrderServiceHandler) GetOrderById(context.Context, *connect.Request[v1.GetOrderByIdRequest]) (*connect.Response[v1.GetOrderByIdResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.GetOrderById is not implemented"))
}

func (UnimplementedOrderServiceHandler) GetOrders(context.Context, *connect.Request[v1.GetOrdersRequest]) (*connect.Response[v1.GetOrdersResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.GetOrders is not implemented"))
}

func (UnimplementedOrderServiceHandler) CancelOrder(context.Context, *connect.Request[v1.CancelOrderRequest]) (*connect.Response[v1.CancelOrderResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.CancelOrder is not implemented"))
}

func (UnimplementedOrderServiceHandler) CancelOrderItem(context.Context, *connect.Request[v1.CancelOrderItemRequest]) (*connect.Response[v1.CancelOrderItemResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.CancelOrderItem is not implemented"))
}

func (UnimplementedOrderServiceHandler) RemoveItemFromCart(context.Context, *connect.Request[v1.RemoveItemFromCartRequest]) (*connect.Response[v1.RemoveItemFromCartResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.RemoveItemFromCart is not implemented"))
}

func (UnimplementedOrderServiceHandler) ReturnOrderItems(context.Context, *connect.Request[v1.ReturnOrderItemsRequest]) (*connect.Response[v1.ReturnOrderItemsResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.ReturnOrderItems is not implemented"))
}

func (UnimplementedOrderServiceHandler) IncreaseItemQuantity(context.Context, *connect.Request[v1.IncreaseItemQuantityRequest]) (*connect.Response[v1.IncreaseItemQuantityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.IncreaseItemQuantity is not implemented"))
}

func (UnimplementedOrderServiceHandler) DecreaseItemQuantity(context.Context, *connect.Request[v1.DecreaseItemQuantityRequest]) (*connect.Response[v1.DecreaseItemQuantityResponse], error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("order.v1.OrderService.DecreaseItemQuantity is not implemented"))
}
