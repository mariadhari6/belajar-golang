package calculation

func Add(num1 int, num2 int) int { // public function diawali dengan huruf besar, dapat diakses oleh semua package
	return add(num1, num2)
}
func add(num1 int, num2 int) int { // private function diawali huruf kecil, hanya dapat diakses pada pcakage yang sama
	return num1 + num2
}
