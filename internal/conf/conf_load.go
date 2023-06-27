package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"path/filepath"
)

var Active string

func NewConfig(path string) Bootstrap {
	fileName := "config.yaml"
	if Active != "" {
		fileName = "config-" + Active + ".yaml"
	}
	c := config.New(
		config.WithSource(
			file.NewSource(filepath.Join(path, fileName)),
		),
	)
	defer c.Close()
	if err := c.Load(); err != nil {
		panic(err)
	}
	var bc Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}
	bc.Folder = path
	return bc
}
