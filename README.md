# spotify-time-search-api

## これは何
WINCのハッカソンの助っ人で作ったAPI
SpotifyのAPIを叩いてデータを集めてきて、時間で検索できるようにしたもの

## APIの設計
時間で検索
/search?time={time}
trackid指定でdbに追加
/track?id={trackid}
playlistid指定でdbに追加
/playlist?id={playlistid}

idはspotifyをwebで開いた時のpathの末尾

## 使い方(多分変える)

```
docker compose up -d
cd api
go mod tidy
go run main.go
```