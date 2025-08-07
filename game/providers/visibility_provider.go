package providers

type VisibilityProvider interface {
	SetVisibility(visible bool)
	IsVisible() bool
}

type visibilityProvider struct {
	visible bool
}

func NewVisibilityProvider() VisibilityProvider {
	return &visibilityProvider{
		visible: true, // Default visibility is true
	}
}

func (v *visibilityProvider) SetVisibility(visible bool) {
	v.visible = visible
}

func (v *visibilityProvider) IsVisible() bool {
	return v.visible
}
