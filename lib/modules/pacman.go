package modules

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/davidscholberg/go-i3barjson"
)

// Pacman represents the configuration for the pacman display block.
type Pacman struct {
	BlockConfigBase `yaml:",inline"`
}

// UpdateBlock updates the pacman block.
func (c Pacman) UpdateBlock(b *i3barjson.Block) {
	b.Color = c.Color
	fullTextFmt := fmt.Sprintf("%s%%s", c.Label)

	out, _ := exec.Command("checkupdates").Output()
	count := bytes.Count(out, []byte{10})

	b.FullText = fmt.Sprintf(fullTextFmt, fmt.Sprintf("%d", count))
}
