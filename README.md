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
