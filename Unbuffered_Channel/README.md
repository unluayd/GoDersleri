# Unbuffered Channel Ornekleri

Bu klasor, Go dilinde `unbuffered channel` kullaniminin farkli senaryolarini gosteren ornekler icerir.

## Dosyalar

- `main.go`: Tek bir gonderici ve alici arasindaki temel senkron veri aktarimini gosterir.
- `main2.go`: Kanala birden fazla veri gonderme, alma ve kanal kapatma akisini gosterir.
- `main3.go`: Ayni kanalda birden fazla alici varken tek bir veri gonderilmesinin sonucunu gosterir.

## Kodun Ozeti

- Unbuffered channel yapisinda gonderme ve alma islemleri birbirini bekler.
- `main.go`, `done` kanali ile goroutine tamamlanmasini izler.
- `main2.go`, kanalin kapatilmasi ve `isOpen` kontrolu ile veri alma mantigini aciklar.
- `main3.go`, birden fazla alici arasinda tek bir verinin yalnizca bir goroutine tarafindan alinacagini gostermeye yardimci olur.

## Calistirma

```bash
go run main.go
go run main2.go
go run main3.go
```
