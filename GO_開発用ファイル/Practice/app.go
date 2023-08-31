// パッケージ
package main

import (
	"fmt"
	"math"
)

func main() {
   fmt.Println(math.Pi) //=> 3.141592653589793
}

// 関数
func main() {
    fmt.Println(hello("Hello World")) //=> Hello World
    fmt.Println(hello2("Hello", "World")) //=> Hello World
    fmt.Println(multipleArgs("Hello", "World")) //=> World Hello
}
//ex1:引数の戻り値の型を書く必要がある
func hello(arg string) string{
	return arg
}
//ex2:関数が２つ以上の同種類の引数を伴う際、型の省略する事が可能
func hello2(arg1, arg2 string) string {
    return arg1 + " " + arg2
}
//ex3:関数は複数の値を返すことができる
func multipleArgs(arg1, arg2 string)(string, string){
    return arg2, arg1
}

// 変数
//ex1:`var`によって変数を宣言し、型の明示が必要である。
var lang string 
//※変数に初期値を与えないとゼロ値（Zero values)が設定されます。(数値型には0、bool型にはfalse、string型には""(空の文字列)が与えられる。

//ex2:初期値を渡した状態で変数を宣言すると型の明示を省略が可能。
var lang2 = "Golang"//初期値を渡す。

func main() {
    //ex3:関数内では｀:=`を利用することでより短いコードで変数の宣言を行うことが可能
    lang3 := "JS" 
    fmt.Println(lang, lang2, lang3) //=> Golang JS
}

// 定数
const Num = 2

func main(){
    const Greetings = "Hello World"
    fmt.Println(Greetings) // => Hello World
    fmt.Println(Num) // => 2
}

// IF文
package main 

import "fmt"

func condition(arg string)string{
	if v := "GO"; arg == v {
    	return "This is Golang"
	}else{
    	return "This is not Golang"
	}
}
func main(){
    fmt.Println(condition("Swift")) //=> This is not Golang
}

// For文
func main(){
    sum := 0
    for i := 0; i < 5; i++ {
        sum += i 
    }
    fmt.Println(sum) //=> 10
}

// Switch文
func main(){
    fmt.Println(condition("Go")) //This is Go
}

func condition(arg string) string{
    switch arg {
    case "Ruby":
        return "This is Ruby"
    case "Go": //これ以降のcaseやdefaultは実行されない。
        return "This is Go"
    case "JS":
        return "This is JS"
    default:
        return "I don't know what this is"
    }
}

// Defer(遅延実行)
package main 

import "fmt"

func main(){
    defer fmt.Println("Golang") //defer1
    defer fmt.Println("Ruby") //defer2
    fmt.Println("JS") 
    //=> JS
    //=> Ruby
    //=> Golang
}

// ポインタ
func main(){
    var lang string 
    lang = "Go" 
    //ex1:`&`を用いることで変数のアドレスの取得が可能
    fmt.Println(&lang) //=> 0x1040c128

    //ex2:`*`を使用する事でポインタを値にとったポインタ変数の宣言が可能
    var langP *string 
    langP = &lang
    fmt.Println(langP)//=> 0x1040c128

    //ex3:ポインタ型変数名の前に`*`をつけることで変数の中身へのアクセスが可能
    fmt.Println(*langP) //=> Go
}

// 構造体
// ex1: typeとstructを使用して定義する。
type Person struct {
	name string 
	age int
 }
 
 func main(){
 //ex2:複数の初期化方法が存在する
	 //初期化方法①:変数定義後にフィールドを設定する
	 var mike Person
	 mike.name = "Mike"
	 mike.age = 23
 
	 //初期化方法②: {} で順番にフィールドの値を渡す
	 var bob = Person{"Bob", 35}
 
	 //初期化方法③:フィールド名を ： で指定する方法
	 var sam = Person{age:89, name: "Sam"}
 
	 fmt.Println(mike, bob, sam) //=> {Mike 23} {Bob 35} {Sam 89}
 }

// 配列
//ex1:配列とは、同じ型を持つ値（要素）を並べたもの
//ex2:複数の宣言方法がある

//宣言方法(1)
var arr1 [2]string 
//宣言方法(2)
var arr2 [2]string = [2]string{"Golang", "Ruby"}
//宣言方法(3)
var arr3 = [...]string{"Golang", "Ruby"}

func main(){
    arr1[0] = "Golang"
    arr1[1] = "Ruby"
    fmt.Println(arr1, arr2, arr3) //=> [Golang Ruby] [Golang Ruby] [Golang Ruby]
}

// 連想配列
func main(){
    //ex1複数の宣言方法が存在する

    //①組み込み関数make()を利用して宣言
    //make(map[キーの型]値の型, キャパシティの初期値)
    //make(map[キーの型]値の型)

    map1 := make(map[string]string)
    map1["Name"] = "Mike"
    map1["Gender"] = "Male"

    //②初期値を指定して宣言
    //var 変数名 map[key]value = map[key]value{key1: value1, key2: value2, ..., keyN: valueN}
    map2 := map[string]int{"Age":25,"UserId":2}

    //ex2初期値を指定しない場合、変数はnil(nil マップ)に初期化される
    var map3 map[string]string

    fmt.Println(map1, map2, map3) //=>map[Gender:Male Name:Mike] map[Age:25 UserId:2] map[]
}

// スライス
func main(){
    //参照用配列
    var arr[2]string = [2]string{"Ruby","Golang"}

    //ex1:配列とは異なり長さ指定の必要なし
    var slice1 []string //スライス(1)
    var slice2 []string = []string{"Ruby", "Golang"} //スライス(2)

    //ex2:配列から要素を取り出し参照する形での宣言が可能
    var slice3 = arr[0:2] //スライス(3)

    //ex2:make()を利用した宣言が可能
    var slice4 = make([]string,2,2) //スライス(4)

    //ex3:配列とは異なり要素の追加が可能
    //append は新しいスライスを返すことに注意
    slice5 := [] string{"JavaScript"}
    newSlice := append(slice5, "Ruby") //sliceに"Ruby"を追加

    fmt.Println(slice1,slice2,slice3,slice4, slice5, newSlice) //=>[] [Ruby Golang] [Ruby Golang] [ ] [JavaScript] [JavaScript Ruby]
}

// Range
//スライスとマップの作成 
var slice1 = []string{"Golang", "Ruby"}
var map1 = map[string]string{"Lang1":"Golang", "Lang2":"Ruby"}

func main(){
    //ex1:スライスやマップに使用すると反復毎に2つの変数を返す。
    //ex2:スライスの場合、1つ目の変数は `インデックス(index)`で、`2つ目は要素(value)`である。
    for index, value := range slice1{
        fmt.Println(index,value)
        //=> 0 Golang
        //=> 1 Ruby
    }

    //ex3:マップの場合、1つ目の変数は`キー(key)`で、２つ目の変数は`バリュー(value)`である。
    for key, value := range map1{
        fmt.Println(key, value)
        //=> Lang1 Golang
        //=> Lang2 Ruby
    }

    //ex4:インデックスや値は、 _ へ代入することで省略することが可能。

    for _,value := range map1{
        fmt.Println(value)
        //=> Golang
        //=> Ruby
    }
}

// Interface
type Person struct {} //Person構造体
type Person2 struct {} //Person2構造体

//Person構造体のメソッドintro()
func (p Person) intro() string{ 
    return "Hello World"
}

//Person2構造体のメソッドintro()
func (p Person2) intro() string{
     return "Hello World"
}

//Person構造体のメソッドintro()を実行するメソッド
func IntroForPerson(arg *Person){
    fmt.Println(arg.intro())
}

//Person2構造体のメソッドintro()を実行するメソッド
func IntroForPerson2(arg *Person2){
    fmt.Println(arg.intro())
}

func main(){
  bob := new(Person)
  mike := new(Person2) 

  IntroForPerson(bob) //=> Hello World
  IntroForPerson2(mike) //=> Hello World
}

// Go Routine
package main 

import( 
 "fmt"
 "log"
 "runtime"
)
func main(){
    fmt.Println("Hello World")
    //ex1:関数(またはメソッド)の呼び出しの前にgoを付けると、異なるgoroutineで関数を実行することが可能。
    go hello()
    go goodMorning()
    //ex2:runtime.NumGoroutine()を使用する事で現在起動しているgoroutine(ゴルーチン)の数を知ることが可能。
    log.Println(runtime.NumGoroutine()) 
    //=> Hello World
    //=> 2009/11/10 23:00:00 3
    //3がgoroutineの数。
    //ここではmain, hello, goodMorningの３つを指す。
}

func hello(){
    fmt.Println("Hello")
}

func goodMorning(){
    fmt.Println("Good Morning")
}

// Channel
func main(){
    //ex1:組み込み関数`make()`を使用する事で生成可能。
    ch := make(chan string) 

    //ex2:<-を用いる事で値の送受信が可能。
    //ex2+:作成したチャネルchに値を送信。
    go func(){ ch <- "str"}()

    //ex2+:チャネルchから値を受信
    msg := <- ch
    fmt.Println(msg) //=>str
}
