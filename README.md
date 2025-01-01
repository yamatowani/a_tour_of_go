# A Tour of Go学習メモ
## Packages
- Goのプログラムはパッケージ(package)で構成される
- プログラムはmainパッケージから開始される
- パッケージ名はインポートパスの最後の要素と同じ名前になる
  - ex: math/randパッケージはpackage randステートメントで始まるファイル群で構成される
## Imports
- factored import statement: 括弧でパッケージのインポートステートメントをグループ化すること
  - factoredは要素化, グループ化という意味
