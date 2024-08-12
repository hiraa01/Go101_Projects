package main

import "fmt"

var class1 = make(map[string]bool)
var class2 = make(map[string]bool)
var class3 = make(map[string]bool)
var class4 = make(map[string]bool)
var graduated = make(map[string]bool)

func moveStudents(fromClass, toClass *map[string]bool) {
	for student := range *fromClass {
		(*toClass)[student] = true
		delete(*fromClass, student)
	}
}
func newSemester() {
	// 4. sınıftan mezunlara taşı
	moveStudents(&class4, &graduated)
	// 3. sınıftan 4. sınıfa taşı
	moveStudents(&class3, &class4)
	// 2. sınıftan 3. sınıfa taşı
	moveStudents(&class2, &class3)
	// 1. sınıftan 2. sınıfa taşı
	moveStudents(&class1, &class2)

}

// Belirli sınıfların öğrencilerini listeleyen fonksiyon
func listClass(class *map[string]bool, className string) {
	fmt.Println(className)
	for student := range *class {
		fmt.Println(student)
	}
	fmt.Println()
}

// Tüm sınıfları listeleyen fonksiyon
func listAllClasses() {
	listClass(&class1, "1. Sınıf")
	listClass(&class2, "2. Sınıf")
	listClass(&class3, "3. Sınıf")
	listClass(&class4, "4. Sınıf")
	listClass(&graduated, "Mezun")
}

func main() {

	class1["Kemal Saygın"] = true
	class1["Derin Kuzu"] = true

	class2["Furkan Sönmez"] = true
	class2["Gül Çağla"] = true

	class3["Semih Aydın"] = true
	class3["Ayşe Yılmaz"] = true

	class4["Zeynep Bucak"] = true
	class4["Salih Kara"] = true

	graduated[""] = true

	listClass(&class1, "1. Sınıf")
	listClass(&class2, "2. Sınıf")
	listClass(&class3, "3. Sınıf")
	listClass(&class4, "4. Sınıf")
	listClass(&graduated, "Mezun")

	newSemester()

	// Güncellenmiş sınıfları listele
	fmt.Println("Dönem Sonu Güncellenen Sınıflar:")
	listAllClasses()

}
