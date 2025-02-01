package opt

import (
	"fmt"

	"github.com/tniedbala/tempe-go/tempe/api"
)

type Options struct {
	whitespace whitespaceOptions
}

func DefaultOptions() *Options {
	return &Options{
		whitespace: defaultWhitespaceOptions(),
	}
}

func (o *Options) Set(opt api.Option) error {
	if _, err := opt(o); err != nil {
		return err
	}
	return nil
}

func (o *Options) Get(opt api.Option) (any, bool) {
	value, err := opt(o)
	return value, err != nil
}

func (o Options) String() string {
	return fmt.Sprintf("Options: {\n    Whitespace: %s\n}", o.whitespace)
}

type TemplateOption func(opts *Options) (any, error)

func Option(option TemplateOption) api.Option {
	return func(opts api.Options) (any, error) {
		o, ok := opts.(*Options)
		if !ok {
			return nil, fmt.Errorf("invalid options: %v", opts)
		}
		return option(o)
	}
}
