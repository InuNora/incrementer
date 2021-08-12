package service

type IncrementerService interface {
	GetNumber() uint32
	IncrementNumber()
	SetMaximumValue(maxValue uint32) error
}
