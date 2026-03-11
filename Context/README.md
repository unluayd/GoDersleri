# Context Ornegi

Bu klasor, Go dilinde `context` kullaniminin temel mantigini gosteren basit bir ornek icerir.

## Dosyalar

- `main.go`: Context icine veri ekleme, fonksiyonlar arasinda tasima ve timeout senaryosunu yorum satirlariyla gosteren ornek kod.

## Kodun Ozeti

- `context.Background()` ile temel bir context olusturulur.
- `context.WithValue` kullanilarak context icine `correlation-id` eklenir.
- `F1`, `F2` ve `F3` fonksiyonlari ayni context'i alip bu degeri okur.
- Dosyada yorum satiri olarak timeout ve `select` ile iptal yonetimi ornegi de bulunur.
- `getProductDetailsFromExternalService`, gecikmeli bir dis servis cagrisini simule eder.

## Calistirma

```bash
go run main.go
```
