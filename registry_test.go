package registry_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/nilslice/schema-registry/server"

	"github.com/stretchr/testify/assert"
	pb "github.com/nilslice/schema-registry/go/registrypb"
)

func TestRegistryProtoCode(t *testing.T) {
	assert.Empty(t, pb.InitRequest{})
	s := server.Service{}
	ctx := context.TODO()
	schema := &pb.Schema{
		Lockfile: []byte("{}"),
	}
	lockfile, err := ioutil.ReadFile("proto.lock")
	assert.NoError(t, err)

	resp, err := s.Status(ctx, &pb.StatusRequest{
		Schema:   schema,
		Lockfile: lockfile,
	})
	for _, warn := range resp.Warnings {
		fmt.Println(warn.Message)
	}
	assert.NoError(t, err)
	assert.Equal(t, 42, len(resp.Warnings))
}
