package toml

import (
	"io"
	"io/ioutil"
	"os"

	"github.com/BurntSushi/toml"
)

// Configurer contains from a static slice of bytes used to configure helpers
type Configurer struct {
	Value []byte
}

// New creates a new TOML configuration helper using the reader as the data source
func New(r io.Reader) (cfg *Configurer, err error) {
	cfg = new(Configurer)
	if cfg.Value, err = ioutil.ReadAll(r); err != nil {
		return
	}
	return
}

// NewFromFile creates a new TOML configuration helper using the file as the data source
func NewFromFile(name string) (*Configurer, error) {
	var f *os.File
	var err error
	if f, err = os.Open(name); err != nil {
		return nil, err
	}
	defer f.Close()
	return New(f)
}

// Configure one ore more items using the standard library's JSON decoder
func (z Configurer) Configure(items ...interface{}) (err error) {
	for _, item := range items {
		if _, err = toml.Decode(string(z.Value), item); err != nil {
			return
		}
	}
	return
}
