package main

import (
	"fmt"
	"math/rand"
)

// Sosyal ağdaki kullanıcı bağlantılarını tutan bir harita. bunu global tanımlamamız tüm fonksiyonların buna erişmesini sağlayacaktır
var baglantilar = make(map[int][]int)

// KullaniciOlustur fonksiyonu, bir kullanıcı oluşturur ve onu haritaya ekler.
func KullaniciOlustur(id int) {
	// Yeni bir kullanıcı oluşturur ve bağlantılarını boş bir dilim olarak ayarlar
	baglantilar[id] = []int{}
}

// BaglantiOlustur iki kullanıcı arasında karşılıklı bağlantı oluşturur
func BaglantiOlustur(a, b int) {
	// Bağlantı oluştur fonksiyonu
	baglantilar[a] = append(baglantilar[a], b)
	baglantilar[b] = append(baglantilar[b], a)
}

// BaglantiKopar iki kullanıcı arasındaki bağlantıyı koparır
func BaglantiKopar(a, b int) {
	// Bağlantı Kopar Fonksiyonu
	baglantilar[a] = removeFromSlice(baglantilar[a], b)
	baglantilar[b] = removeFromSlice(baglantilar[b], a)
}

// remove bir slice'tan belirli bir öğeyi kaldırır
func removeFromSlice(slice []int, elem int) []int {
	// tekil bağlantı silme işlemi
	for i, v := range slice {
		if v == elem {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// YuzKullaniciOlustur 100 farklı kullanıcı oluşturur
func YuzKullaniciOlustur() {
	// rastgele 100 kullanıcı oluşturunuz
	for i := 0; i < 100; i++ {
		id := 10000000 + rand.Intn(90000000)
		baglantilar[id] = []int{}
	}

}

// OrtakBaglantiTespiti iki kullanıcının ortak arkadaşlarını bulur
func OrtakBaglantiTespiti(a, b int) []int {
	// Ortak bağlantıları tutacak slice
	var ortaklar []int
	for _, v := range baglantilar[a] {
		if contains(baglantilar[b], v) {
			ortaklar = append(ortaklar, v)
		}
	}
	return ortaklar
}

func contains(slice []int, elem int) bool {
	for _, v := range slice {
		if v == elem {
			return true
		}
	}
	return false
}

func main() {
	KullaniciOlustur(55005500)
	KullaniciOlustur(45001500)
	KullaniciOlustur(35021300)
	fmt.Println("Dost Meclisine Hoşgeldiniz!")
	for {
		var secim1, secim2 int
		fmt.Println("Id'leri verilen bu kullanıcılardan hangilerinin bağlantı kurmasını istersiniz?\n1) 55005500\n2) 45001500\n3) 35021300")
		fmt.Print("1. Kullanıcı: ")
		fmt.Scan(&secim1)
		fmt.Print("2. Kullanıcı: ")
		fmt.Scan(&secim2)

		if secim1 == 1 && secim2 == 2 {
			BaglantiOlustur(55005500, 45001500)
		} else if secim1 == 2 && secim2 == 1 {
			BaglantiOlustur(55005500, 45001500)
		} else if secim1 == 1 && secim2 == 3 {
			BaglantiOlustur(55005500, 35021300)
		} else if secim1 == 3 && secim2 == 1 {
			BaglantiOlustur(55005500, 35021300)
		} else if secim1 == 2 && secim2 == 3 {
			BaglantiOlustur(45001500, 35021300)
		} else if secim1 == 3 && secim2 == 2 {
			BaglantiOlustur(45001500, 35021300)
		}
		// Bağlantıları görüntüle
		fmt.Println("55005500 bağlantıları:", baglantilar[55005500])
		fmt.Println("45001500 bağlantıları:", baglantilar[45001500])
		fmt.Println("35021300 bağlantıları:", baglantilar[35021300])
		var sec int
		fmt.Println("Kullanıcı oluşturma işleminiz bitti mi?\n1-evet\n2-hayır")
		fmt.Scan(&sec)
		if sec == 1 {
			break
		} else {
			continue
		}
	}

	// 100 farklı kullanıcı oluştur
	YuzKullaniciOlustur()
	fmt.Println("100 kullanıcı oluşturuldu:", len(baglantilar))

	// Ortak bağlantı tespiti
	ortaklar := OrtakBaglantiTespiti(55005500, 45001500)
	fmt.Println("55005500 ve 45001500 arasındaki ortak arkadaşlar:", ortaklar)
}
