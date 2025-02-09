package options

import (
	"encoding/json"
	"fmt"

	"github.com/tniedbala/tempe-go/tempe/api"
)

// private struct so json marshalling/unmarshalling can be done without exposing public properties
type optionsMarshaller struct {
	Whitespace *WhitespaceOptions `json:"whitespace"`
}

type Options struct {
	whitespace *WhitespaceOptions 
}

func DefaultOptions() *Options {
	return &Options{
		whitespace: DefaultWhitespaceOptions(),
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
	return value, err == nil
}

func (o *Options) String() string {
	b, err := json.MarshalIndent(o, "", "  ")
	if err != nil {
		return fmt.Sprintf(`Options{error: "%s"}`, err)
	}
	return fmt.Sprintf("Options%s", string(b))
}

func (o *Options) MarshalJSON() ([]byte, error) {
	return json.Marshal(optionsMarshaller{Whitespace: o.whitespace})
}

func (o *Options) UnmarshalJSON(data []byte) error {
	var opts optionsMarshaller
	if err := json.Unmarshal(data, &opts); err != nil {
		return err
	}
	*o = Options{whitespace: opts.Whitespace}
	return nil
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
