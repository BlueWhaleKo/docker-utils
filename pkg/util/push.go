package util

import (
	"context"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func PushImage(c *client.Client, image string, opt types.ImagePushOptions) error {
	reader, err := c.ImagePush(context.Background(), image, opt)
	if err != nil {
		return err
	}

	Print(reader)

	return nil
}
