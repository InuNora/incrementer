package impl

import "errors"

const MaxUint = ^uint32(0)

type Incrementer struct {
	Value    uint32
	MaxValue uint32
}

// NewIncrementer godoc
// @Summary Создает новый instance Incrementer
// @Description Максимальное значение не должно быть меньше стартового или нулем
// @Tags Incrementer
func NewIncrementer(startValue uint32, maxValue uint32) (*Incrementer, error) {
	if startValue > maxValue {
		return nil, errors.New("Maximum value cannot be less than the start value")
	}
	if maxValue == 0 {
		return nil, errors.New("Maximum value cannot be null")
	}
	return &Incrementer{
		Value:    startValue,
		MaxValue: maxValue,
	}, nil
}

// GetNumber godoc
// @Summary Возвращает текущее число.
// @Description В самом начале это ноль.
// @Tags Incrementer
func (inc *Incrementer) GetNumber() uint32 {
	return inc.Value
}

// IncrementNumber godoc
// @Summary Увеличивает текущее число на один.
// @Description После каждого вызова этого метода getNumber() будет возвращать число на один больше.
// @Tags Incrementer
func (inc *Incrementer) IncrementNumber() {
	inc.Value += 1
	if inc.Value >= inc.MaxValue {
		inc.Value = 0
	}
	return
}

// SetMaximumValue godoc
// @Summary Устанавливает максимальное значение текущего числа.
// @Description Хранимое число не может превышать установленное максимальное значение.
// Когда при вызове incrementNumber() текущее число достигает
// этого значения, оно обнуляется, т.е. getNumber() начинает
// снова возвращать ноль, и снова один после следующего
// вызова incrementNumber() и так далее.
// По умолчанию максимум -- максимальное значение int.
// Нельзя позволять установить тут число меньше нуля.
// @Tags Incrementer
func (inc *Incrementer) SetMaximumValue(maxValue uint32) error {
	if maxValue == 0 {
		return errors.New("Maximum value cannot be less than the start value")
	}
	if maxValue <= inc.Value {
		return errors.New("Maximum value cannot be null")
	}
	inc.MaxValue = maxValue
	return nil
}
