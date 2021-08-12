package impl

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewIncrementerGood godoc
// @Summary Создает новый instance Incrementer с валидными данными
// @Description проверка на отсутствие ошибки.
// @Tags Test
func TestNewIncrementerGood(t *testing.T) {
	_, err := NewIncrementer(0, 10)
	assert.NoError(t, err)
}

// TestNewIncrementerBad godoc
// @Summary Создает новый instance Incrementer с не валидными данными
// @Description проверка на наличие ошибки.
// @Tags Test
func TestNewIncrementerBad(t *testing.T) {
	_, err := NewIncrementer(9, 1)
	assert.Error(t, err)
}

// TestNewIncrementerNullMax godoc
// @Summary Создает новый instance Incrementer с максимальным значением 0
// @Description проверка на наличие ошибки.
// @Tags Test
func TestNewIncrementerNullMax(t *testing.T) {
	_, err := NewIncrementer(9, 0)
	assert.Error(t, err)
}

// TestGetNumberDefaultCases godoc
// @Summary Серия тестов, проверяющая работу GetNumber
// @Description Создает новый instance Incrementer, не изменяя его вызывает GetNumber
// проверяет, что возвращается стартовое значение
// @Tags Test
func TestGetNumberDefaultCases(t *testing.T) {
	cases := []struct {
		inStart, inMax, want uint32
	}{
		{0, 100, 0},
		{MaxUint - 1, MaxUint, MaxUint - 1},
		{1000, 9999, 1000},
	}
	for _, c := range cases {
		incr, _ := NewIncrementer(c.inStart, c.inMax)
		got := incr.GetNumber()
		if got != c.want {
			t.Errorf("Reverse(%v) == %v, want %v", c.inStart, got, c.want)
		}
	}
}

// TestIncrementNumberCases godoc
// @Summary Серия тестов, проверяющая работу IncrementNumber
// @Description Создает новый instance Incrementer, увеличивает значение переменной с помощью IncrementNumber
// проверяет, что возвращается стартовое значение + 1
// @Tags Test
func TestIncrementNumberCases(t *testing.T) {
	cases := []struct {
		inStart, inMax, want uint32
	}{
		{0, 100, 1},
		{MaxUint - 2, MaxUint, MaxUint - 1},
		{1000, 9999, 1001},
	}
	for _, c := range cases {
		incr, _ := NewIncrementer(c.inStart, c.inMax)
		incr.IncrementNumber()
		got := incr.GetNumber()
		if got != c.want {
			t.Errorf("Reverse(%v) == %v, want %v", c.inStart, got, c.want)
		}
	}
}

// TestIncrementNumberOverflowCases godoc
// @Summary Серия тестов, проверяющая работу IncrementNumber
// @Description Создает новый instance Incrementer, увеличивает значение переменной с помощью IncrementNumber
// проверяет, переполнение максимального значения
// @Tags Test
func TestIncrementNumberOverflowCases(t *testing.T) {
	cases := []struct {
		inStart, inMax, want uint32
	}{
		{0, 1, 0},
		{MaxUint - 1, MaxUint, 0},
		{1000, 9999, 1001},
	}
	for _, c := range cases {
		incr, _ := NewIncrementer(c.inStart, c.inMax)
		incr.IncrementNumber()
		got := incr.GetNumber()
		if got != c.want {
			t.Errorf("Reverse(%v) == %v, want %v", c.inStart, got, c.want)
		}
	}
}

// TestSetMaximumValueCasesErrors godoc
// @Summary Серия тестов, проверяющая работу SetMaximumValue на ошибку
// @Description Создает новый instance Incrementer, увеличивает значение переменной с помощью IncrementNumber
// Устанавливает максимальное значение меньше или равно текущему. Проверяет ошибку
// @Tags Test
func TestSetMaximumValueCasesErrors(t *testing.T) {
	cases := []struct {
		inStart, inMax, setMax uint32
	}{
		{1, 5, 0},
		{MaxUint - 1, MaxUint, MaxUint - 2},
		{1000, 9999, 1000},
	}
	for _, c := range cases {
		incr, _ := NewIncrementer(c.inStart, c.inMax)
		err := incr.SetMaximumValue(c.setMax)
		assert.Error(t, err)
	}
}

// TestSetMaximumValueCasesGood godoc
// @Summary Серия тестов с переполнением и без
// @Description Создает новый instance Incrementer, увеличивает значение переменной с помощью IncrementNumber
// Устанавливает максимальное значение больше текущего. Увеличивает значение ещё раз
// @Tags Test
func TestSetMaximumValueCasesGood(t *testing.T) {
	cases := []struct {
		inStart, inMax, setMax, want uint32
	}{
		{1, 5, 7, 3},
		{MaxUint - 3, MaxUint, MaxUint - 1, 0},
		{1000, 9999, 1002, 0},
	}
	for _, c := range cases {
		incr, _ := NewIncrementer(c.inStart, c.inMax)
		incr.IncrementNumber()
		err := incr.SetMaximumValue(c.setMax)
		if err != nil {
			t.Errorf("Error: (%v) +1 <= %v", c.inStart, c.setMax)
		}
		incr.IncrementNumber()
		got := incr.GetNumber()
		if got != c.want {
			t.Errorf("Reverse(%v) == %v, want %v", c.inStart, got, c.want)
		}
	}
}
