package test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestMongoDBIntegrationExample(t *testing.T) {
	ctx := context.Background()

	container, err := setupMongoContainer(ctx)
	assert.NoError(t, err)

	t.Cleanup(func() {
		err := container.Terminate(ctx)
		assert.NoErrorf(t, err, "error terminating container")
	})

	endpoint, err := container.Endpoint(ctx, "mongodb")
	assert.NoErrorf(t, err, "error getting endpoint")

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(endpoint))
	assert.NoError(t, err, "error creating mongo client")

	err = mongoClient.Ping(ctx, nil)
	assert.NoError(t, err, "error pinging mongo")
}

type mongoContainer struct {
	testcontainers.Container
}

func setupMongoContainer(ctx context.Context) (*mongoContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "mongo:6",
		ExposedPorts: []string{"27017/tcp"},
		WaitingFor: wait.ForAll(
			wait.ForLog("Waiting for connections"),
			wait.ForListeningPort("27017/tcp"),
		),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	return &mongoContainer{Container: container}, nil
}
