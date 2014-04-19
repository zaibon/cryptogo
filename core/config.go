package core

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

type Config struct {
	Wallets map[string]*Wallet
}

const defaultConfigFolder = "conf/wallets/"

//ParseConfig parses the Configuration file and return pointer of Config struct
func ParseConfig() (*Config, error) {
	dirName := defaultConfigFolder

	//FIXME not portable
	if !strings.HasSuffix(dirName, "/") {
		dirName += "/"
	}

	pattern := dirName + "*.json"
	filenames, err := filepath.Glob(dirName + "*.json")
	if err != nil {
		return nil, err
	}
	if len(filenames) == 0 {
		return nil, fmt.Errorf("%s matches no files: %s", pattern)
	}
	return parseFiles(filenames)
}

// parseFiles parse a Config file and return a pointer to
//a filled struc decoded from the file
func parseFiles(filenames []string) (*Config, error) {
	if len(filenames) == 0 {
		// Not really a problem, but be consistent.
		return nil, fmt.Errorf("no files")
	}

	wallets := make(map[string]*Wallet, len(filenames))

	for _, filename := range filenames {
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		w := &Wallet{}
		if err := w.Unmarshal(b); err != nil {
			return nil, err
		}
		wallets[w.Name] = w
	}

	conf := &Config{
		Wallets: wallets,
	}
	return conf, nil
}
