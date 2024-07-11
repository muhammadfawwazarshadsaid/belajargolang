package main

import (
	"belajargolang/cobamodulesatu"
	"fmt"
	"strconv"
	"strings"
)

type Student struct{
	nama string
	kelas int
	nik int
}

type StudentID struct{
	Student
	nik int
	id int
}

type Guru struct{
	nama string
	nik int
	Sekolah struct{
		namasekolah string
		idsekolah int
	}

}

func (g *Guru) setGuru(nama string, nik int){
	g.nama = nama
	g.nik = nik
}

func main(){
	
	cobamodulesatu.TestHello("arshad")

	var operator1 = cobamodulesatu.SetOperator("arshad","belajar")
	operator1.GetOperator()

	var lingkaran = cobamodulesatu.SetLingkaran(8)
	var prisma = cobamodulesatu.SetPrisma(8)
	lingkaran.Luas()
	prisma.Luas()


	var sekolah1 = struct {
	namasekolah string
	idsekolah   int
	}{
	namasekolah: "MAN 2 Kota Bekasi",
	idsekolah:   232,
	}

	var guruakun1 = Guru{"aads",2,struct{namasekolah string; idsekolah int}{namasekolah:sekolah1.namasekolah,}}
	guruakun1.setGuru("arshad", 2206082511)
	fmt.Println(guruakun1)


	var guru1 = Guru{
		nama: "arshad",
		nik: 2206082511,
		Sekolah: struct{
			namasekolah string; 
			idsekolah int}{
				namasekolah: "MAN 1 Kota Bekasi",
				idsekolah: 12,
			}}
	var guru2 = Guru{
		nama: "arshad",
		nik: 2206082511,
		Sekolah: struct{
			namasekolah string; 
			idsekolah int}{
				namasekolah: sekolah1.namasekolah,
				idsekolah: sekolah1.idsekolah,
			}}
	fmt.Println(guru1)
	fmt.Println(guru2)

	var student1 Student
	student1.nama = "arshad"
	fmt.Println(student1.nama)

	var student2 *Student = &student1
	student2.kelas = 4
	student2.nik = 2206082512
	fmt.Println(student2)

	var student3 StudentID
	student3.nik = 2206082511
	fmt.Println(student3.Student.nik)

	var student4 = StudentID{nik: 2206082511, id:1}
	var student5 *StudentID = &student4
	student5.nama = "arshad"
	student5.kelas = 12
	fmt.Println(student5)




	// var nama string = "arshad"
	// var nama2 = "arshad"
	// nama3 := 12
	// fmt.Printf("coba belajar golang")
	// fmt.Printf("%s %s %s %b",nama,nama2,nama2, nama3)

	/* 
	Contoh multiline
	*/

	// const(
	// 	nama string = "arshad"
	// 	nama2
	// )

	var valuesofindex = [...]int{1,2,3,4,5}

	for _,value := range valuesofindex[0:2]{
		fmt.Println(value)
	}

	var mappingvalue = map[string]int{"a":1,"b":2}

	for _,value := range mappingvalue{
		fmt.Println(value)
	}

	var multidimensi = [2][3]int{{1,2,3},{1,2,3}}
	fmt.Println(multidimensi)

	var alokasi = make([]int,2)
	alokasi[1] = 2
	fmt.Println(alokasi[1])

	var usingcaplen = []int{1,2,3,4,5}
	indexcapln := usingcaplen[5:]
	fmt.Println(cap(indexcapln))
	fmt.Println(len(usingcaplen))

	cobacopy := []int{4,6,6,6,6,6,7}
	cobacopy1 := []int{6,5,4,3,2}
	cobacopy2 := copy(cobacopy,cobacopy1)
	fmt.Println(cobacopy2)

	var checkExist = map[string]int{"a":1,"b":2}
	var value,isExist = checkExist["a"]
	if isExist{
		fmt.Println(value)
	}
	fmt.Println(isExist)

	var extendedMap = []map[string]int{{"a":1,"b":2},{"a":4,"b":2}}
	for _,value := range extendedMap{
		fmt.Println(value["a"])
	}

	printSesuatu("Nilai Basdat",[]int{96,99,100})
	printSesuatuReturn("Nilai Basdat",[]int{96,99,100})
	printSesuatuReturn2("Nilai Basdat",[]int{96,99,100})

	fmt.Println(printFuncVariadic(1,2,3,4,5))
	fmt.Println(printFuncVariadicS("Muhammad","Fawwaz","Arshad","Said"))
	nama := []string{"Muhammad","Fawwaz","Arshad","Said"}
	fmt.Println(printFuncVariadicS(nama...))

	fmt.Println(cobaTipeData(1))


	var numberA int = 4
	var numberB *int = &numberA

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	fmt.Println("")

	numberA = 5

	fmt.Println("numberA (value)   :", numberA)
	fmt.Println("numberA (address) :", &numberA)
	fmt.Println("numberB (value)   :", *numberB)
	fmt.Println("numberB (address) :", numberB)

	var numberC *int = &numberA
	fmt.Println(*numberC)

	var listx []int = []int{1,2,3,4,5,6}
	var x,y int = titiktitik(listx...)
	fmt.Println(x,y)
	
	var namax = "arshad"
	pointer(&namax)

	var namazz *string = &namax
	fmt.Println(*namazz)

	
}

func printSesuatu(name string, arr []int){
	for _,value := range arr{
		fmt.Printf("%s: %d\n",name,value)
	}

	var arrstring = make([]string,len(arr))
	for index,x := range arr{
		arrstring[index] = strconv.Itoa(x)
	}

	var gabungNilai = strings.Join(arrstring, ",")
	fmt.Printf("%s: %s\n", name, gabungNilai)
}

func printSesuatuReturn(name string, arr []int) string{
	for _,value := range arr{
		fmt.Printf("%s: %d\n",name,value)
	}

	var arrstring = make([]string,len(arr))
	for index,x := range arr{
		arrstring[index] = strconv.Itoa(x)
	}

	var gabungNilai = strings.Join(arrstring, ",")
	return fmt.Sprintf("%s: %s", name, gabungNilai)
}

func printSesuatuReturn2(name string, arr []int) (string,string){
	for _,value := range arr{
		fmt.Printf("%s: %d\n",name,value)
	}

	var arrstring = make([]string,len(arr))
	for index,x := range arr{
		arrstring[index] = strconv.Itoa(x)
	}

	var gabungNilai = strings.Join(arrstring, ",")
	return name, gabungNilai
}

func printFuncVariadic(number ...int) int{
	sum := 0
	for _,value := range number{
		sum+=value
	}
	return sum
}

func printFuncVariadicS(nama ...string) string{
	namasaya := strings.Join(nama, " ")
	return namasaya
}


type IntegerBaru int

func cobaTipeData(number IntegerBaru) IntegerBaru{
	return number
}

func titiktitik(listx ...int) (int,int){
	var x int = 0

	for _,value := range (listx){
		x += value
	}
	var y int = x
	return x,y
}

func pointer(valuePointer *string) {
	for _, value := range *valuePointer {
		fmt.Println(string(value))
	}
}