package main

import (
	"bufio"
	"io"
	"os"
	"sync"
	"time"
)

type Aof struct {
	file *os.File
	rd   *bufio.Reader
	mu   sync.Mutex
}

func NewAof(path string) (*Aof, error) {
	f, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR, 0666)

	if err != nil {
		return nil, err
	}

	aof := &Aof{
		file: f,
		rd:   bufio.NewReader(f),
	}

	go func() {
		for {
			aof.mu.Lock()
			aof.file.Sync()
			aof.mu.Unlock()

			time.Sleep(time.Second)
		}
	}()

	return aof, nil
}

func (a *Aof) Write(value Value) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	_, err := a.file.Write(value.Marshal())

	if err != nil {
		return err
	}

	return nil
}

func (a *Aof) Read(cb func(value Value)) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	resp := NewResp(a.file)

	for {
		value, err := resp.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		cb(value)
	}

	return nil
}

func (a *Aof) Close() error {
	a.mu.Lock()
	defer a.mu.Unlock()

	return a.file.Close()
}
