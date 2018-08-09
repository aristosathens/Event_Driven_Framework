package JSONSaver

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"sync"
)

var (
	lock sync.Mutex
)

// Credit: https://medium.com/@matryer/golang-advent-calendar-day-eleven-persisting-go-objects-to-disk-7caf1ee3d11d

// ------------------------------------------- Public ------------------------------------------- //

func Save(path string, object interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	r, err := marshal(object)
	if err != nil {
		return err
	}
	_, err = io.Copy(f, r)
	return err
}

func Load(path string, v interface{}) error {
	lock.Lock()
	defer lock.Unlock()
	if !fileExists(path) {
		return nil
	}
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	// err = unmarshal(f, v)
	// return v, err
	return unmarshal(f, v)
}

// ------------------------------------------- Utilities ------------------------------------------- //

var marshal = func(v interface{}) (io.Reader, error) {
	b, err := json.MarshalIndent(v, "", "\t")
	if err != nil {
		return nil, err
	}
	return bytes.NewReader(b), nil
}

var unmarshal = func(r io.Reader, v interface{}) error {
	return json.NewDecoder(r).Decode(v)
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
		// return true, nil
	}
	if os.IsNotExist(err) {
		return false
		// return false, nil
	}
	return false
	// return true, err
}
