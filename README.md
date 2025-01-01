# A Tour of Go学習メモ
## Packages
- Goのプログラムはパッケージ(package)で構成される
- プログラムはmainパッケージから開始される
- パッケージ名はインポートパスの最後の要素と同じ名前になる
  - ex: math/randパッケージはpackage randステートメントで始まるファイル群で構成される
## Imports
- factored import statement: 括弧でパッケージのインポートステートメントをグループ化すること
  - factoredは要素化, グループ化という意味
## Exported names
- パッケージをインポートすると、そのパッケージでエクスポートされている名前を参照できる
- エクスポートされた名前は必ず大文字から始まる
## Functions
- 関数は0個上の引数を取る
- 変数名の**後ろ**に型名を書く
## Functions continued
- 関数の2つ以上の引数が同じ型である場合は、最後の型を残して省略して記述できる
```go
func add(x, y int) int {
	return x + y
}
```
## Multiple results
- 関数は複数の戻り値を返すことができる
## Named reutn values
- 戻り値となる変数に名前をつけれる
- naked return: 名前をつけた戻り値の変数を使うとreturnステートメントに何も書かずに戻せる
  - 短い関数でのみ使うことを推奨, 長い関数で使うと読みにくい
```go
func split(sum int) (x, y int) {
	x = sum * 4/ 9
	y = sum - x
	return
}
```
## Variables
- varステートメント: 変数を宣言
  - 宣言した変数の後ろに型を書く
  - パッケージ, 関数で利用できる
```go
var c, python , java bool
```
## Variables with initializers
- var宣言では、initializer(初期化子)を与えることができる
  - 変数に初期化子が与えられている場合は、型を省略でき初期化子が持つ型になる
## Short variable declarations
- **関数の中では**、var宣言の代わりに:=の代入文を使うことができる
- 関数の外ではvar, funcなどのキーワードで始まる宣言しか使えない
```go
func main() {
  k := 3
  c, python, java := true, false, "No!"
}
```
## Basic Type
### Goの組み込み型
- bool
- stirng
- 符号付き整数
  - int, int8, int16, int32(rune), int64
- 符号なし整数
  - uint, uint8(byte), int16, int32, int64, uintptr
- float32, float64
- complex64, complex128
- Error
## Zero Values
- 変数に初期値を与えないとゼロ値が与えらえる
  - 0, false, ""(空文字列)など
## Type conversions
- 変数v, 型Tの場合, T(v)で型変換を行うことができる
```go
i := 42
f := float64(i)
u := uint(f)
```
型変換を行わない場合、float64型の変数fはuintの値として宣言できないよ!と怒られてしまう
```bash
# command-line-arguments
./type-conversions.go:11:15: cannot use f (variable of type float64) as uint value in variable declaration
```
## Type inference
- 型を指定せずに変数を宣言すると、右側の変数から型推論が行われる
- 右側の変数が型を持っている場合, 左側の新しい変数は同じ型になる
```go
i := 35 // int
f := 3.14159 // float64
g := 0.867 + 0.5i //complex128
```
## Constants
- 定数はconstキーワードを使って宣言
- 文字, 文字列, boolean, 数値のみで使える
- :=による宣言代入はできない
## Numeric Constants
- 定数は高精度な値
- 型のない定数は状況に応じて必要な型を取る
- intは最大64bitの整数を保持できるが、環境によっては精度が低い
## For
- 初期化文; 条件式; 後処理; {}の形で書く
- この3つの部分を括る括弧がないのが特徴, 中括弧は常に必要
## For continued
- 初期化と後処理ステートメントの記述は任意である
```go
func main() {
	sum := 1
	for ; sum < 1000; {
		sum += sum
	}
	fmt.Println(sum)
}
```
## For is Go's "while"
セミコロンを省略して、他の言語のwhile文のように利用できる
```go
func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
```
## Forever
- 条件式を省略すれば無限ループになる
## If
- if文に括弧はいらない, 中括弧は必要
## If with a short statement
- if文の中で変数を宣言し、ifのスコープ内, ブロック内でのみ使える
## If and else
- if文で宣言された変数はelseブロック内でも使える
