package main

import "fmt"

func main() {
	bakiye := 1000
	fmt.Printf("Hesabınızda %d TL bulunmaktadır.\n", bakiye)
	fmt.Println("Para mı çekmek istersiniz yoksa para mı yatırmak istersiniz?")
	fmt.Println("1-Para çekmek istiyorum.\n2-Para yatırmak istiyorum.")

	var secenek int
	var tutar int

	fmt.Print("Seçiminiz: ")
	fmt.Scan(&secenek)

	if secenek == 1 {
		fmt.Println("Ne kadar para çekmek istersiniz?")
		fmt.Print("Çekmek istediğiniz tutarı giriniz: ")
		fmt.Scan(&tutar)
		bakiye := 1000
		if tutar > bakiye || tutar < 0 {
			fmt.Println("Bu işlem yapılamaz.")
			return
		}
		bakiye = bakiye - tutar
		fmt.Printf("Hesabınızda artık %d TL bulunmaktadır.", bakiye)

	} else if secenek == 2 {
		fmt.Println("Ne kadar para yatırmak istersiniz?")
		fmt.Print("Yatırmak istediğiniz tutarı giriniz:")
		fmt.Scan(&tutar)

		bakiye := 1000
		bakiye = bakiye + tutar
		fmt.Printf("Hesabınızda artık %d TL bulunmaktadır.", bakiye)
	} else {
		fmt.Println("Geçersiz giriş yaptınız!")
	}
}
