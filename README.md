# Task
Нужно написать класс на Golang со следующим интерфейсом (псевдокод):
## Class
### Incrementor {
#### int getNumber();
* Возвращает текущее число. В самом начале это ноль.
#### void incrementNumber();
* Увеличивает текущее число на один. После каждого вызова этого метода getNumber() будет возвращать число на один больше.
#### setMaximumValue(int maximumValue);
* Устанавливает максимальное значение текущего числа.
* Хранимое число не может превышать установленное максимальное значение.
* Когда при вызове incrementNumber() текущее число достигает этого значения, оно обнуляется, т.е. getNumber() начинает
снова возвращать ноль, и снова один после следующего вызова incrementNumber() и так далее.
* По умолчанию максимум -- максимальное значение int.
* Нельзя позволять установить тут число меньше нуля.
### }

# Tests
## cower testing
go test -coverprofile fmtcoverage.html fmt
## start testing
go test -v ./...
