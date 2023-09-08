package main

type Kendaraan struct {
	totalRoda int

	kecepatanPerJam int
}

type Mobil struct {
	Kendaraan
}

func (m *Mobil) Berjalan() {

	m.TambahKecepatan(10)

}

func (m *Mobil) TambahKecepatan(kecepatanBaru int) {

	m.kecepatanPerJam += kecepatanBaru

}

func main() {

	mobilCepat := Mobil{}

	mobilCepat.Berjalan()

	mobilCepat.Berjalan()

	mobilCepat.Berjalan()

	mobilLambat := Mobil{}

	mobilLambat.Berjalan()

}
