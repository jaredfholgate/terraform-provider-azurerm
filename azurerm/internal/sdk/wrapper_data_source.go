package sdk

import (
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

type DataSourceWrapper struct {
	dataSource DataSource
	logger     Logger
}

func NewDataSourceWrapper(dataSource DataSource) DataSourceWrapper {
	return DataSourceWrapper{
		dataSource: dataSource,
		logger:     ConsoleLogger{},
	}
}

func (rw *DataSourceWrapper) DataSource() (*schema.Resource, error) {
	resourceSchema, err := combineSchema(rw.dataSource.Arguments(), rw.dataSource.Attributes())
	if err != nil {
		return nil, fmt.Errorf("building Schema: %+v", err)
	}

	var d = func(duration time.Duration) *time.Duration {
		return &duration
	}

	resource := schema.Resource{
		Schema: *resourceSchema,
		Read: func(d *schema.ResourceData, meta interface{}) error {
			ctx, metaData := runArgs(d, meta, rw.logger)
			return rw.dataSource.Read().Func(ctx, metaData)
		},
		Timeouts: &schema.ResourceTimeout{
			Read: d(rw.dataSource.Read().Timeout),
		},
	}

	return &resource, nil
}
