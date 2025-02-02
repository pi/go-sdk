// Code generated by sdkgen. DO NOT EDIT.

//nolint
package devices

import (
	"context"

	"google.golang.org/grpc"

	devices "github.com/yandex-cloud/go-genproto/yandex/cloud/iot/devices/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// RegistryServiceClient is a devices.RegistryServiceClient with
// lazy GRPC connection initialization.
type RegistryServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ devices.RegistryServiceClient = &RegistryServiceClient{}

// AddCertificate implements devices.RegistryServiceClient
func (c *RegistryServiceClient) AddCertificate(ctx context.Context, in *devices.AddRegistryCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).AddCertificate(ctx, in, opts...)
}

// Create implements devices.RegistryServiceClient
func (c *RegistryServiceClient) Create(ctx context.Context, in *devices.CreateRegistryRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements devices.RegistryServiceClient
func (c *RegistryServiceClient) Delete(ctx context.Context, in *devices.DeleteRegistryRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteCertificate implements devices.RegistryServiceClient
func (c *RegistryServiceClient) DeleteCertificate(ctx context.Context, in *devices.DeleteRegistryCertificateRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).DeleteCertificate(ctx, in, opts...)
}

// Get implements devices.RegistryServiceClient
func (c *RegistryServiceClient) Get(ctx context.Context, in *devices.GetRegistryRequest, opts ...grpc.CallOption) (*devices.Registry, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).Get(ctx, in, opts...)
}

// List implements devices.RegistryServiceClient
func (c *RegistryServiceClient) List(ctx context.Context, in *devices.ListRegistriesRequest, opts ...grpc.CallOption) (*devices.ListRegistriesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).List(ctx, in, opts...)
}

// ListCertificates implements devices.RegistryServiceClient
func (c *RegistryServiceClient) ListCertificates(ctx context.Context, in *devices.ListRegistryCertificatesRequest, opts ...grpc.CallOption) (*devices.ListRegistryCertificatesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).ListCertificates(ctx, in, opts...)
}

// ListDeviceTopicAliases implements devices.RegistryServiceClient
func (c *RegistryServiceClient) ListDeviceTopicAliases(ctx context.Context, in *devices.ListDeviceTopicAliasesRequest, opts ...grpc.CallOption) (*devices.ListDeviceTopicAliasesResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).ListDeviceTopicAliases(ctx, in, opts...)
}

// ListOperations implements devices.RegistryServiceClient
func (c *RegistryServiceClient) ListOperations(ctx context.Context, in *devices.ListRegistryOperationsRequest, opts ...grpc.CallOption) (*devices.ListRegistryOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).ListOperations(ctx, in, opts...)
}

// Update implements devices.RegistryServiceClient
func (c *RegistryServiceClient) Update(ctx context.Context, in *devices.UpdateRegistryRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return devices.NewRegistryServiceClient(conn).Update(ctx, in, opts...)
}
