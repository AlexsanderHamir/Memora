package memcontext

func NewContextManager() *DefaultContextManager {
	return &DefaultContextManager{}
}

func (dcm *DefaultContextManager) CreateContext(name string) Context {
	ctx := &DefaultContext{name: name}
	dcm.contexts.Store(name, ctx)
	return ctx
}

func (dcm *DefaultContextManager) GetContext(name string) (Context, error) {
	ctx, ok := dcm.contexts.Load(name)
	if !ok {
		return nil, ErrContextNotFound
	}
	return ctx.(Context), nil
}

func (dcm *DefaultContextManager) GetOrCreateContext(name string) (Context, error) {
	ctx, err := dcm.GetContext(name)
	if err != nil {
		return dcm.CreateContext(name), nil
	}
	return ctx, nil
}

func (dcm *DefaultContextManager) DeleteContext(name string) {
	dcm.contexts.Delete(name)
}
