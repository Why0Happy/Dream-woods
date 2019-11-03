package Section_1

//1.4 变量和常量

//1.4.1 变量
//var varNAME dateType = value (显示完整声明)示例：
var a1 int = 1
var b1 int = 2*3
var c1 int = b1
var d1 int

//varNAME := value (短声明)示例：
func dsm() {       //短声明只能出现在函数（方法）内
	a := 2
	b,c := 1,"abc"
	d := a+b
	println(a,b,c,d)
}

//1.4.2 常量
const a2  = 1
const b2  = 2
const c2  = a2 + b2

const(                        //运用iota声明一组常量
	a3 = iota                 // a3==0 (iota == 0)
	b3						  // b3==1 (iota == 1)
	c3 int = 2                // c3==2 (iota == 3,unused)
	d3 float64 = iota * 2     // d3==8.0 (iota == 4)
)

//1.5 基本数据类型

//布尔类型 bool
func be()  {
	var a bool
	a=true
	b := false
	c := a1<d1
	var d bool        //d==false 布尔型默认值为false
	println(a,b,c,d)
}

//整型 byte,int,int8,int16,int32,int64,uint,uint8,uint16,uint32,uint64,uintptr
func zx()  {
	var a byte = 1    //byte是uint8的别名
	var b int = 32
	println(a,b)
}

//浮点型 float32,float64
func fdx() {
	var a float32 = 1.555
	var b float64 = 1
	c := 1.66666            //浮点型字面量被自动推断为float64类型
	println(a,b,c)
}

//复数 complex64,complex128
func fs()  {
	var a complex64 = 3.1 + 6i     //complex64由两个float32构成
	var b complex128 = 3.2 + 7i    //complex128由两个float64构成

//处理复数的三个内置函数
	var c= complex(4.3,8)     //构造一个复数
	d := real(c)                   //返回复数实部，返回值为对应的浮点型
	e := imag(c)				   //返回复数虚部，返回值为对应的浮点型
	println(a,b,d,e)
}

