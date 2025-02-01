package tempe

type GoEngine struct {
	TemplateEngine
}

func DefaultEngine() GoEngine {
	return NewGoEngine()
}

func NewGoEngine() GoEngine {
	env, value := &GoEnv{}, GoValue{}
	return GoEngine{
		NewTemplateEngine(env, value),
	}
}

func (e GoEngine) String() string {
	return "GoEngine{}"
}
