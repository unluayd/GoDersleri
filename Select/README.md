# Select — Birden fazla kanalı aynı anda dinlemek

`select`, bir goroutine’in **birden fazla channel işleminden** hangisinin hazır olduğuna göre dallanmasını sağlar. Bu örnekte iki string kanalı farklı gecikmelerle dolar; döngü içinde `select` ile ikisinden gelen veriler toplanır.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `main.go` | `channel1`, `channel2`, iki gönderici goroutine, `for` + `select` + `default`. |

## Akış

1. Bir goroutine **10 saniye** sonra `channel1 <- "Hello"` yollar.  
2. Diğeri **5 saniye** sonra `channel2 <- "World"` yollar.  
3. `for len(data1) == 0 || len(data2) == 0` — iki veri de dolana kadar döngü sürer.  
4. `select` içinde hazır olan `case` çalışır; hiçbiri hazır değilse `default` çalışır (`"Veri Gelmedi."`).  
5. Döngünün CPU’yu yakmaması için her turda `time.Sleep(1 * time.Second)` vardır.

## `default` notu

`default` olduğunda `select` bloklanmaz; bu yüzden örnek **aktif bekleme** (polling) karakterindedir. Üretimde genelde bloklayıcı `case <-ch` ile beklemek veya `ctx.Done()` gibi iptal kanalları kullanmak daha yaygındır.

## Çalıştırma

```bash
cd Select
go run .
```
