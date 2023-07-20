# GenshinCalender
~~※ heroku が有料になったので、現在使用不可。~~  
Google Cloud Run に移行できたので使用可能。

原神のキャラ育成をサポートする web アプリ

## 機能
育成したいキャラクターを選択して、天賦素材と必要樹脂数が表示されます。

## 使い方
1. 育てたいキャラクターを選択します。
2. 選択し終えたら、次へのボタンを押します。
3. レベルを上げたい部分を選択し、レベル上げ値を設定します。
4. 秘境 1回分のドロップ数を設定します。
5. 設定し終えたら、次へのボタンを押します。
6. 最後に、選択したキャラに必要な素材と、秘境の曜日、総必要樹脂数、濃縮樹脂数、樹脂回復の時間が表示されます。


## Deploy memo
```
gcloud run deploy genshin-calendar --region=asia-northeast1 --source=.
```