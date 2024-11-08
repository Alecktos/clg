package misc

type DeviceType int64

const (
	DeviceTypeMobile DeviceType = iota
	DeviceTypeDesktop
)

const CurrentDevice DeviceType = DeviceTypeDesktop
