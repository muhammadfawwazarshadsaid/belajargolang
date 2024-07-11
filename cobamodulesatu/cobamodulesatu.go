package cobamodulesatu

import "fmt"

type Operator struct{
	namaoperator string
	tugas string
}

func SetOperator(namaoperator string, tugas string) Operator{
	return Operator{namaoperator: namaoperator, tugas: tugas}
}

func (o Operator) GetOperator(){
	fmt.Println(o.namaoperator,o.tugas)
}

func TestHello(hasil string) {
	fmt.Println(hasil)
}


type Lingkaran struct{
	r int
}
func SetLingkaran(r int) Lingkaran{
	return Lingkaran{r:r}
}
func (l Lingkaran) Luas() int{
	return l.r
}

type BangunRuang interface{
	Luas() int
	Keliling() int
}

type Prisma struct{
	r int
}

func SetPrisma(r int) Prisma{
	return Prisma{r:r}
}

func (p Prisma) Luas() int{
	return p.r
}