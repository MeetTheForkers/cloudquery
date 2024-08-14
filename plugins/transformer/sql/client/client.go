package client

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/apache/arrow/go/v17/arrow"
	"github.com/cloudquery/cloudquery/plugins/transformer/sql/client/spec"
	"github.com/cloudquery/cloudquery/plugins/transformer/sql/client/transformers"
	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/rs/zerolog"
)

type Client struct {
	plugin.UnimplementedSource
	plugin.UnimplementedDestination

	logger zerolog.Logger
	spec   spec.Spec
	tfs    []*transformers.Transformer
}

func New(ctx context.Context, logger zerolog.Logger, s []byte, opts plugin.NewClientOptions) (plugin.Client, error) {
	c := &Client{
		logger: logger.With().Str("module", opts.PluginMeta.Name).Logger(),
	}
	if opts.NoConnection {
		return c, nil
	}

	if err := json.Unmarshal(s, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file spec: %w", err)
	}
	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}

	for _, transformationSpec := range c.spec.TransformationSpecs {
		tf, err := transformers.NewFromSpec(ctx, transformationSpec)
		if err != nil {
			return nil, err
		}
		c.tfs = append(c.tfs, tf)
	}

	return c, nil
}

func (c *Client) Transform(ctx context.Context, recvRecords <-chan arrow.Record, sendRecords chan<- arrow.Record) error {
	for {
		select {
		case record, ok := <-recvRecords:
			if !ok {
				return nil
			}

			fmt.Println("received record")
			var records []arrow.Record
			// Run all transformers sequentially on the record
			for _, tf := range c.tfs {
				var err error
				records, err = tf.Transform(record)
				if err != nil {
					fmt.Println(err)
					return err
				}
			}
			fmt.Println("transformed record")

			for _, record := range records {
				sendRecords <- record
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (c *Client) TransformSchema(ctx context.Context, schema *arrow.Schema) (*arrow.Schema, error) {
	for _, tf := range c.tfs {
		var err error
		schema, err = tf.TransformSchema(schema)
		if err != nil {
			return nil, err
		}
	}
	return schema, nil
}

func (*Client) Close(ctx context.Context) error {
	return nil
}
