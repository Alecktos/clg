package config

type DeviceType int64

const (
	DeviceTypeMobile DeviceType = iota
	DeviceTypeDesktop
)

const (
	WindowWidth  = 1290 / 4
	WindowHeight = 2532 / 4
)

const CurrentDevice DeviceType = DeviceTypeDesktop
