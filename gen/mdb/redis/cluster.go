// Code generated by sdkgen. DO NOT EDIT.

//nolint
package redis

import (
	"context"

	"google.golang.org/grpc"

	redis "github.com/yandex-cloud/go-genproto/yandex/cloud/mdb/redis/v1"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/operation"
)

//revive:disable

// ClusterServiceClient is a redis.ClusterServiceClient with
// lazy GRPC connection initialization.
type ClusterServiceClient struct {
	getConn func(ctx context.Context) (*grpc.ClientConn, error)
}

var _ redis.ClusterServiceClient = &ClusterServiceClient{}

// AddHosts implements redis.ClusterServiceClient
func (c *ClusterServiceClient) AddHosts(ctx context.Context, in *redis.AddClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).AddHosts(ctx, in, opts...)
}

// AddShard implements redis.ClusterServiceClient
func (c *ClusterServiceClient) AddShard(ctx context.Context, in *redis.AddClusterShardRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).AddShard(ctx, in, opts...)
}

// Backup implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Backup(ctx context.Context, in *redis.BackupClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Backup(ctx, in, opts...)
}

// Create implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Create(ctx context.Context, in *redis.CreateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Create(ctx, in, opts...)
}

// Delete implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Delete(ctx context.Context, in *redis.DeleteClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Delete(ctx, in, opts...)
}

// DeleteHosts implements redis.ClusterServiceClient
func (c *ClusterServiceClient) DeleteHosts(ctx context.Context, in *redis.DeleteClusterHostsRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).DeleteHosts(ctx, in, opts...)
}

// DeleteShard implements redis.ClusterServiceClient
func (c *ClusterServiceClient) DeleteShard(ctx context.Context, in *redis.DeleteClusterShardRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).DeleteShard(ctx, in, opts...)
}

// Get implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Get(ctx context.Context, in *redis.GetClusterRequest, opts ...grpc.CallOption) (*redis.Cluster, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Get(ctx, in, opts...)
}

// GetShard implements redis.ClusterServiceClient
func (c *ClusterServiceClient) GetShard(ctx context.Context, in *redis.GetClusterShardRequest, opts ...grpc.CallOption) (*redis.Shard, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).GetShard(ctx, in, opts...)
}

// List implements redis.ClusterServiceClient
func (c *ClusterServiceClient) List(ctx context.Context, in *redis.ListClustersRequest, opts ...grpc.CallOption) (*redis.ListClustersResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).List(ctx, in, opts...)
}

// ListBackups implements redis.ClusterServiceClient
func (c *ClusterServiceClient) ListBackups(ctx context.Context, in *redis.ListClusterBackupsRequest, opts ...grpc.CallOption) (*redis.ListClusterBackupsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).ListBackups(ctx, in, opts...)
}

// ListHosts implements redis.ClusterServiceClient
func (c *ClusterServiceClient) ListHosts(ctx context.Context, in *redis.ListClusterHostsRequest, opts ...grpc.CallOption) (*redis.ListClusterHostsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).ListHosts(ctx, in, opts...)
}

// ListLogs implements redis.ClusterServiceClient
func (c *ClusterServiceClient) ListLogs(ctx context.Context, in *redis.ListClusterLogsRequest, opts ...grpc.CallOption) (*redis.ListClusterLogsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).ListLogs(ctx, in, opts...)
}

// ListOperations implements redis.ClusterServiceClient
func (c *ClusterServiceClient) ListOperations(ctx context.Context, in *redis.ListClusterOperationsRequest, opts ...grpc.CallOption) (*redis.ListClusterOperationsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).ListOperations(ctx, in, opts...)
}

// ListShards implements redis.ClusterServiceClient
func (c *ClusterServiceClient) ListShards(ctx context.Context, in *redis.ListClusterShardsRequest, opts ...grpc.CallOption) (*redis.ListClusterShardsResponse, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).ListShards(ctx, in, opts...)
}

// Move implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Move(ctx context.Context, in *redis.MoveClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Move(ctx, in, opts...)
}

// Rebalance implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Rebalance(ctx context.Context, in *redis.RebalanceClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Rebalance(ctx, in, opts...)
}

// Restore implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Restore(ctx context.Context, in *redis.RestoreClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Restore(ctx, in, opts...)
}

// Start implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Start(ctx context.Context, in *redis.StartClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Start(ctx, in, opts...)
}

// StartFailover implements redis.ClusterServiceClient
func (c *ClusterServiceClient) StartFailover(ctx context.Context, in *redis.StartClusterFailoverRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).StartFailover(ctx, in, opts...)
}

// Stop implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Stop(ctx context.Context, in *redis.StopClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Stop(ctx, in, opts...)
}

// Update implements redis.ClusterServiceClient
func (c *ClusterServiceClient) Update(ctx context.Context, in *redis.UpdateClusterRequest, opts ...grpc.CallOption) (*operation.Operation, error) {
	conn, err := c.getConn(ctx)
	if err != nil {
		return nil, err
	}
	return redis.NewClusterServiceClient(conn).Update(ctx, in, opts...)
}
