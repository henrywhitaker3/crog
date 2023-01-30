package client

import (
	"context"
	"time"

	"github.com/henrywhitaker3/crog/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	conn   *grpc.ClientConn
	client pb.CrogClient
}

func New(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewCrogClient(conn)

	return &Client{
		conn:   conn,
		client: client,
	}, nil
}

func (c *Client) Close() {
	c.conn.Close()
}

func (c *Client) GetActions() ([]string, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	resp, err := c.client.List(ctx, &pb.ListActionsRequest{})
	if err != nil {
		return nil, err
	}

	out := []string{}

	for _, act := range resp.Actions {
		out = append(out, act.Name)
	}

	return out, nil
}

func (c *Client) RunAction(name string) (*pb.RunActionResponse, error) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*120)
	defer cancel()

	return c.client.Run(ctx, &pb.RunActionRequest{Action: name})
}
