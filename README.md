
# Quick Name Gen

これはChatGPT API(Chat Completions API)を利用したアプリ名付けツールです。  
3つの質問に答えるだけでAIが最適な名前の候補を考えてくれます。  

このツールの名前もChatGPT APIに考えてもらいました🎉  

## 環境変数の設定

このツールはChat Completions APIを使用しているため[Open AIのAPIキー](https://platform.openai.com/account/api-keys)が必要です。

APIキーを取得したらプロジェクト内の.devcontainer/example.envを.devcontainer/.envに変更してください。  
その後ファイル内の「OPEN_AI_SECRET」の値を取得したAPIキーに置き換えてください。

## プロジェクトを開く

このプロジェクトを開くにはVS Codeにdevcontainers拡張がインストールされている必要があります。  

VS Codeでプロジェクトを開いたらコマンドパレットから「Dev Containers: Reopen in Container」を実行してください。  
しばらくするとGoの環境が起動します。

## 実行方法

アプリの名付けツールを実行するには以下のコマンドを実行します。

```bash
make run
```

アプリについて3つ質問されます。  
それぞれ誰のためのアプリか、何をするアプリか、どのように実現するかという質問です。

解答が終わりしばらく待つと解答データに沿った10個のアプリ名候補が表示されるはずです🎉
