package main

import (
	"fmt"
)

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func Mati(b *Bebek) {
	b.hidup = false
}

func Terbang(b *Bebek) {
	if b.energi > 0 && b.hidup && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			Mati(b)
		}
	}
}

func Makan(b *Bebek) {
	if b.energi > 0 && b.hidup {
		b.energi += 1
	}
}

func main() {
	b := Bebek{
		energi:       4,
		hidup:        true,
		bisaTerbang:  true,
		suaraTerbang: "Kwek kwek!",
	}

	Terbang(&b)
	Terbang(&b)
	Terbang(&b)

	Makan(&b)

	fmt.Printf("Energi Bebek: %d\n", b.energi)
	fmt.Printf("Bebek Hidup: %t\n", b.hidup)
}
