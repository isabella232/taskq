package memqueue

import "gopkg.in/queue.v1/processor"

type Storager interface {
	Exists(key string) bool
}

type Options struct {
	Name    string
	Storage Storager

	AlwaysSync bool // if true messages are processed synchronously

	Processor processor.Options
}

func (opt *Options) init() {}