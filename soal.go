/*
//soal nomor 1
package main
import "fmt"
func main() {
	var a,b,total int
	fmt.Scan(&a,&b)
	if a > b {
		total = a - b
		fmt.Println("Selisih angka Heiji dan Shinichi adalah" , total)
	}
	if b > a {
		total = b - a
		fmt.Println("Selisih angka Heiji dan Shinichi adalah" , total)
	}
}
*/

/*
//soal nomor 2
package main

import "fmt"

func main() {
	var ab, bc, cd, ad int
	var hasil string

	fmt.Scan(&ab, &bc, &cd, &ad)
	if ab > 0 && bc > 0 && cd > 0 && ad > 0 {
		if ab == cd && bc == ad && ((ab > bc) && (cd > ad) || (bc > ab) && (ad > cd)) {
			hasil = "Angka dapat membentuk persegi panjang"
			fmt.Println(hasil)
		} else if ab == bc && bc == cd && cd == ad {
			hasil = "Angka dapat membentuk persegi"
			fmt.Println(hasil)
		} else {
			hasil = "Angka tidak membentuk bidang apapun"
			fmt.Println(hasil)
		}
	} else {
		hasil = "Angka harus lebih dari 0"
		fmt.Println(hasil)
	}
}
*/

/*
//soal nomor 3
package main

import "fmt"
func main() {
	var a,b,total int
	fmt.Scan(&a,&b)
	total = a - b
	if (total != 0 ) {
		if (total <= -1) && (total >= -16) && (total % 2 == -1) {
			fmt.Println("Heiji Menang")
		} else if (total > 0) && (total <= 45) && (total % 2 == 0 ) {
			fmt.Println("Shinichi menang")
		}else {
			fmt.Println("Shinci dan Heiji kalah")
		}
	}else if  total == 0 {
		fmt.Println("shinichi dan Heiji seri")
	}
}
*/

// soal nomor 4
package main

import "fmt"

func main(){
	var sHp, sAtt, sDef, iv float64

	fmt.Scan(&sHp, &sAtt, &sDef)
	if (sHp >= 0 && sHp <=15) && (sAtt >= 0 && sAtt <=15) && (sDef >= 0 && sDef <=15) {
		iv = ((sHp + sAtt + sDef)/45) *100
		if (iv >= 0 && iv <= 100) {
			if (iv >= 0 && iv <= 48.90){
				fmt.Println("Overall, your pokemon has room for improvement as far as batting goes")
				fmt.Println("0 stars")
			} else if (iv >= 51.10 && iv <= 64.45) {
				fmt.Println("Overall, your pokemon is pretty decent!")
				fmt.Println("1 stars")
			} else if (iv >= 66.60 && iv <= 80) {
				fmt.Println("Overall, your pokemon is really strong!")
				fmt.Println("2 stars")
			} else if (iv >= 82.2 && iv <= 97.8) {
				fmt.Println("Overall, your pokemon looks like it can really battle with the best of them!")
				fmt.Println("3 stars")
			} else if iv == 100 {
				fmt.Println("Overall, your pokemon looks like it can really battle with the best of them!")
				fmt.Println("4 stars")
			}
		} else {
			fmt.Println("Not a Pokemon")
		}
	} else {
		fmt.Println("Not a Pokemon")
	}
}
