//code A:
var (
	x   int
	i   uint64
	n   uint64
	tmp []byte
)

func test() {
	for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}

	func() {
		size := len(tmp)
		for i := 0; i == uint64(size); i++ {
		}
	}()

	for i := 0; func(j int) bool {
		return j > 100
	}(int(i)); i++ {
		k := 3
		k--
	}
}

//code B:
var (
	x   int
	i   uint64
	n   uint64
	tmp []byte
)

func test() {
	for {
		i++
		tmp = append(tmp, 'a')
		if i == 1000 {
			break
		}
	}

	func() {
		for i := 0; i == len(tmp); i++ {
		}
	}()

	for i := 0; func(j int) bool {
		return j > 100
	}(int(i)); i++ {
		k := 3
		k--
	}
}

// Code A sedikit lebih baik daripada Code B karena penyimpanan eksplisit dari ‘len(tmp)’ dalam variabel,
//yang membuat kode lebih mudah dibaca dan dipelihara. Meskipun kedua kode memiliki masalah logis yang sama
//(loop dengan kondisi ‘i == size’ dan ‘func(j int) bool { return j > 100 }’ tidak akan pernah dieksekusi),
//Code A lebih mudah dibaca dan dipahami karena memisahkan perhitungan dan perbandingan.