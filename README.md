# 概要

競技プログラミングサイト等で、使用する Golang 用のテンプレート

# 使用方法

* 事前準備
    1. 本リポジトリのルートフォルダでコマンドプロンプトを開く
    2. `go mod init` を実行
* 使用時
    1. 各ファイルに以下の値を入力する
        * 入力値 → input.txt
        * 出力値 → expected.txt
    2. 問題を解く
        * main.go の MakeAnswer() に解答を入力する
            * 「入力用の Scanner 定義処理」は解答提出時に置き換える必要があるため、注意
        * go test で解答が正しいかを確認する
            * 例として `.\example\101` の動作を確認するためのコマンドは以下
                * `go test .\example\101`
    3. 解答を提出する
        * MakeAnswer() 内の処理を提出する

# 背景

以下のプログラミング練習のためのサイトで問題を解く際、入力値・出力値をコピペするのが煩わしかったため、作成。

* paiza ラーニング
    * https://paiza.jp/challenges
* yukicoder
    * https://yukicoder.me/problems
