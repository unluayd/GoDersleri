# Buffered Channel Ornegi

Bu klasor, Go dilinde `buffered channel` kullanimini gosteren basit bir ornek icerir.

## Dosyalar

- `main.go`: Buffered channel olusturma, veri gonderme ve veri alma mantigini gosteren ornek kod.

## Kodun Ozeti

- `make(chan int, 4)` ile kapasitesi 4 olan bir kanal olusturulur.
- Ilk goroutine, `0` ile `9` arasindaki sayilari kanala gonderir.
- Gonderim bittiginde kanal `close(myChannel)` ile kapatilir.
- Ikinci goroutine, kanaldaki verileri `range` ile okuyup ekrana yazar.
- `time.Sleep` cagrilari, ornekte goroutine calismasini gozlemlemeyi kolaylastirmak icin kullanilmistir.

## Calistirma

```bash
go run main.go
```
