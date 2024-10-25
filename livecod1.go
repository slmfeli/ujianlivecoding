package main
import "fmt"

//2311102216

func jenisKursi(harga  int) {
	var m, v, b string 
	
	if m == "member" {
		if b == "biasa"  {
			harga = 40000
		} else if v == "VIP" {
			harga = 60000
		}
		
	} else if v == "nonmember" {
		if b == "biasa" {
			harga = 50000
		} else if v == "VIP" {
			harga = 70000
		}
	}

} 

func diskon(n, d int) {
	if n > 2 {
		d = 15/100*n
		fmt.Println("VIP")
	} else {
		fmt.Println("biasa")
	}
}

func main() {
	var n int
	var v, t string
	fmt.Println("Masukan jumlah tiket: ")
	fmt.Scan(&n)

	fmt.Println("Masukan jenis kursi (biasa/VIP): ")
	fmt.Scan(&v)

	fmt.Println("Apakah anda member? (true/false)" )
	fmt.Scan(&t)

	fmt.Println("Total harga: ", jenisKursi)

}