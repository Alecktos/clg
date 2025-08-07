package providers

type DisableProvider interface {
	SetDisabled(disabled bool)
	IsDisabled() bool
	Reset()
}

type disableProvider struct {
	disabled bool
}

func NewDisableProvider() DisableProvider {
	return &disableProvider{
		disabled: false,
	}
}

func (v *disableProvider) SetDisabled(disabled bool) {
	v.disabled = disabled
}

func (v *disableProvider) IsDisabled() bool {
	return v.disabled
}

func (v *disableProvider) Reset() {
	v.disabled = false
}
