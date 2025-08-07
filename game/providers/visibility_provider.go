package providers

type VisibilityProvider interface {
	SetVisibility(visible bool)
	IsVisible() bool
	Reset()
}

type visibilityProvider struct {
	visible bool
}

func NewVisibilityProvider() VisibilityProvider {
	return &visibilityProvider{
		visible: true,
	}
}

func (v *visibilityProvider) SetVisibility(visible bool) {
	v.visible = visible
}

func (v *visibilityProvider) IsVisible() bool {
	return v.visible
}

func (v *visibilityProvider) Reset() {
	v.visible = true
}
