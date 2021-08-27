package util

import (
	"context"
	"io"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PushImage(c *client.Client, image string, opt types.ImagePushOptions) (io.ReadCloser, error) {
	return c.ImagePush(context.Background(), image, opt)
}
