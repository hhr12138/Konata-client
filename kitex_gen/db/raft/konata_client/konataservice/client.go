// Code generated by Kitex v0.9.1. DO NOT EDIT.

package konataservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	konata_client "github.com/hhr12138/Konata-client/kitex_gen/db/raft/konata_client"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Get(ctx context.Context, args_ *konata_client.GetArgs_, callOptions ...callopt.Option) (r *konata_client.Reply, err error)
	PutAppend(ctx context.Context, args_ *konata_client.PutAppendArgs_, callOptions ...callopt.Option) (r *konata_client.Reply, err error)
	RemoveReqId(ctx context.Context, args_ *konata_client.GetArgs_, callOptions ...callopt.Option) (r *konata_client.Reply, err error)
	RequestVote(ctx context.Context, args_ *konata_client.RequestVoteArgs_, callOptions ...callopt.Option) (r *konata_client.RequestVoteReply, err error)
	AppendEntries(ctx context.Context, args_ *konata_client.RequestAppendArgs_, callOptions ...callopt.Option) (r *konata_client.RequestAppendReply, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kKonataServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kKonataServiceClient struct {
	*kClient
}

func (p *kKonataServiceClient) Get(ctx context.Context, args_ *konata_client.GetArgs_, callOptions ...callopt.Option) (r *konata_client.Reply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Get(ctx, args_)
}

func (p *kKonataServiceClient) PutAppend(ctx context.Context, args_ *konata_client.PutAppendArgs_, callOptions ...callopt.Option) (r *konata_client.Reply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PutAppend(ctx, args_)
}

func (p *kKonataServiceClient) RemoveReqId(ctx context.Context, args_ *konata_client.GetArgs_, callOptions ...callopt.Option) (r *konata_client.Reply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RemoveReqId(ctx, args_)
}

func (p *kKonataServiceClient) RequestVote(ctx context.Context, args_ *konata_client.RequestVoteArgs_, callOptions ...callopt.Option) (r *konata_client.RequestVoteReply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.RequestVote(ctx, args_)
}

func (p *kKonataServiceClient) AppendEntries(ctx context.Context, args_ *konata_client.RequestAppendArgs_, callOptions ...callopt.Option) (r *konata_client.RequestAppendReply, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.AppendEntries(ctx, args_)
}
