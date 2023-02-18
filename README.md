# spotify-time-search-api

## これは何
WINCのハッカソンの助っ人で作ったAPI

SpotifyのAPIを叩いてデータを集めてきて、時間で検索できるようにしたもの

## APIの設計

```
時間で検索
/search
Query一覧
time : ちょうど同じ時間を持つ曲を返す
min_time : 指定以上の時間を持つ曲を返す
max_time : 指定以下の時間を持つ曲を返す
around : 指定に最も近い時間を持つ曲を返す

trackid指定でdbに追加
/track?id={trackid}

playlistid指定でdbに追加
/playlist?id={playlistid}
idは複数指定が可能
ex) id=xxxxxxxx&id=yyyyyyy&id=zzzzzzz
```

idはspotifyをwebで開いた時のpathの末尾

## 使い方

1. `/api/.env` を作る
```
/api/.env.sample を参考に。
https://developer.spotify.com/dashboard/applications
で取得できるClient IDと Secret IDを記入
```
2.`docker compose up -d`
M1Macの場合は、`docker-compose.yml`を次のように書き換えてください。
```
version: '3'  
services:
  go:
    build: .
    tty: true
    volumes:
      - ./api:/api
    ports:
      - 3000:3000
    depends_on:
      - "db"
  db:
    platform: linux/x86_64
    image: mysql:latest
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: db
      MYSQL_USER: user
      MYSQL_PASSWORD: password
      TZ: 'Asia/Tokyo'
    command: mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
    volumes:
      - mysql-data:/var/lib/mysql
      - ./db/my.cnf:/etc/mysql/conf.d/my.cnf
    ports:
      - 3306:3306
volumes:
  mysql-data:
    driver: local
```
3. `localhost:3000`に開きます

