package modules

import (
	"fmt"
	"github.com/davidscholberg/go-i3barjson"
	"os"
)

type Interface struct {
	BlockIndex     int    `mapstructure:"block_index"`
	UpdateInterval int    `mapstructure:"update_interval"`
	UpdateSignal   int    `mapstructure:"update_signal"`
	IfaceName      string `mapstructure:"interface_name"`
}

func (c Interface) GetBlockIndex() int {
	return c.BlockIndex
}

func (c Interface) GetUpdateFunc() func(b *i3barjson.Block, c BlockConfig) {
	return updateIfaceBlock
}

func (c Interface) GetUpdateInterval() int {
	return c.UpdateInterval
}

func (c Interface) GetUpdateSignal() int {
	return c.UpdateSignal
}

func updateIfaceBlock(b *i3barjson.Block, c BlockConfig) {
	cfg := c.(Interface)
	var statusStr string
	fullTextFmt := "E: %s"
	// TODO: make interface name configurable
	sysFilePath := fmt.Sprintf("/sys/class/net/%s/operstate", cfg.IfaceName)
	r, err := os.Open(sysFilePath)
	if err != nil {
		b.Urgent = true
		b.FullText = fmt.Sprintf(fullTextFmt, err.Error())
		return
	}
	_, err = fmt.Fscanf(r, "%s", &statusStr)
	if err != nil {
		b.Urgent = true
		b.FullText = fmt.Sprintf(fullTextFmt, err.Error())
		return
	}
	r.Close()
	if statusStr == "up" {
		b.Urgent = false
	} else {
		b.Urgent = true
	}
	b.FullText = fmt.Sprintf(fullTextFmt, statusStr)
}
