package main
import (
	"fmt" 
    "math"
	"time"
)
var a=9   //global package level
// func calc(f,g int) (int,int,int,int,int){
// 	return (f+g),(f-g),(f*g),(f/g),(f%g)
// }
// func f2(){
// 	var a=10
// 	fmt.Println(a)
// }
func main(){
	
	// var num ,n int  // assigned  0 automatically
	// num , n =20,30
	// var num1 int= 10  //int is optional if you not init it
	// var num2 uint= 10  //uint for +ve integer
	// var result = num1+num
	// fmt.Println(num2)
	// fmt.Println(n)
	// fmt.Println("The sum of",num1,"and",num,"is = ",result)
	
	// //other wat for var
	// res :=100     //same as var result=100 shorthand
	// fmt.Println(res)
	// const a =11
	// //a=12   error
	// fmt.Println(a)

	// var sum,dif,mul,div,mod = calc(9,3)
	// fmt.Println(sum,dif,mul,div,mod)

	//FOR LOOP
	// for{
	// 	fmt.Println("HELLO")     //INFINITE
	// }
	// for i:=1;i<=5;i++{
	// 	fmt.Println("hello",i)
	// }
	// for i:=1;i<=100;i++{
	// 	fmt.Print(i," ")
	// }



	// f2()
	// var a=11     //function level
	// fmt.Println(a)
	 var num float64 = 90
	 var result=math.Sqrt(num)
	 var roundresult=math.Round(result)    //Floor,Ceil can be used
	 fmt.Printf("The output is %.2f Thank you",roundresult)
	 fmt.Println()


	 //IF
    a:=10

	if a==1{
		fmt.Println("one")
	}else if a==2{
		fmt.Println("two")
	}else if a==3{
		fmt.Println("three")
	}else if a==4{
		fmt.Println("four")
	}else if a==5{
		fmt.Println("five")
	}else if a==6{
		fmt.Println("six")
	}else{
		fmt.Println("greater than 6")
	}

	if a%2==0{
		fmt.Println("even")
	} else{
		fmt.Println("odd")
	}
	var p,q,r,s float64=10,0,99,10
	var m=math.Max( p,q)
	var n=math.Max( r,s)
	var o=math.Max( m,n)
	fmt.Println(o)
	var j=55
	switch j {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	case 5:
		fmt.Println("five")
	case 6:
		fmt.Println("six")	
	default:
		fmt.Println("None")
	}

	fmt.Println("When's Saturday ?")
	today:=time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tommorow")
	case today + 2:
		fmt.Println("In two days")
	case today + 3:
		fmt.Println("In three days")
	case today + 4:
		fmt.Println("In four days")
	case today + 5:
		fmt.Println("in five days")	
	default:
		fmt.Println("Too far away")
	}
	c()
	for aa:=0;aa<=10;aa++{
		fmt.Println(aa)
	}
	fmt.Println("bye....")
	for ab:=0;ab<=10;ab++{
		defer fmt.Println(ab)
	}
	fmt.Println("bye....")
    var s1 = student{101,"abhi",99.99}
	fmt.Println(s1)
	fmt.Println(s1.roll)
	fmt.Println(s1.name)
	fmt.Println(s1.marks)
	var s2 = student{roll:102,marks:100}
	fmt.Println(s2)

}

type student struct{
	roll int
	name string
	marks float64
}

func c(){
	fmt.Println("a begins")
	defer b()        //execute after a only
	fmt.Println("a ends")
}
func b(){
	fmt.Println("in b")
}
