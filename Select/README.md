# Select Ornegi

Bu klasor, Go dilinde `select` yapisinin birden fazla kanali ayni anda dinlemek icin nasil kullanildigini gosteren basit bir ornek icerir.

## Dosyalar

- `main.go`: Iki farkli kanaldan gelen verileri `select` ile takip eden ornek kod.

## Kodun Ozeti

- Iki ayri kanal olusturulur.
- Her kanal farkli gecikmelerle veri gonderir.
- `select` ifadesi, hangi kanalda veri hazirsa onu isler.
- `default` blogu sayesinde veri gelmedigi durumlar da gorulebilir.

## Calistirma

```bash
go run main.go
```
