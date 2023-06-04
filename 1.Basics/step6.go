package main

func main() { //STRING, NUMBER, BOOLEAN

//STRING

//o'zg - nomi - tipi  -  VALUE - qiymat
	var message string = "Hello Student"
	println(message)

	var message1 string 	// DECLARE --> tanitdim
	message1 = ==+"Hello World" // ASSIGMAND --> TENGLASHTIRDIM
	priprintln(message1)
	
	var box1, box2 string = "first",  "second"
	println(box1,box2)

// NOTE !

var box = "This is variable" // TIPI berilmagan ammo bu ham islaydi !
printlydin(box) 			 // tipi berilmasa o'zi avtomat topib sekin ishlaydi
						// agar tipi berilasa kod tezlik jihatdan yahshi ishlaydi
//_______________________________________________________________________________

// NUMBER
// INIT	
	var a,b,c, int	// BUTUN SON Ex: 7, -5, -11, 23, 94, 768, 811
	println(a,b,c)  // --> 0 0 0

	var n1, n2, n3, int = 11, 22, 33,
	println(n1, n2, n3)

	var num1, num2, num3 int = 11, 22, 33
	println(num1, num2, num3)
//_______________________________________________________________________________

// FLOAT

	var x, y, z float32 // QOLODIQ SON Ex: 1.5, -11.6, 123.5, 4567.11, -12.5
	println(x,y,z)     // --> +0.000000e+000 +0.000000e+000 +0.000000e+000

	var f1, f2, f3 float64 = 11.1, 22.2, 33.3
	println(f1,f2,f3)

	var flo1, flo2, flo3 float64 = 11.1, 22.2, 33.3
	//____________________________________________________________________________
// BOOLEAN

	var erkakMi bool = true // false
	println(erkakMi)

//_________________________________________________________________________________

// HAMMA O'ZGARUVCHILAR
 // YANI bitta satrga tog'ridan tog'ri harhil tipda qiymat birlashtirish munkun !

	var q, w, e, r = 11.1, 22.3, "Hello", true
	println(q, w, e, r)

//	_________________________________________________________________________________

//NEW NOTE !
// := --> BU OPEEATOR ISHLATILSA VAR va TIPIyozilishi shaart emas !

	num, flo, str, truFal := 11, 22.3, "Hello", true
	println(num, flo, str, truFal) 	// bir satirlik o'zgaruvchi typelar ko'di

	var bag1, bag2, bag3, bag3, bag4 = 33, 44.5, "Hello", true
	println(bag1, bag2, bag3, bag4)
	
}