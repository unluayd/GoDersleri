# OneVideoGo

Bu proje, Go dilinde goroutine, channel, `select`, `context` ve `WaitGroup` gibi eszamanlilik konularini kucuk orneklerle anlatir.

## Klasorler

- `main.go`: `sync.WaitGroup` ile birden fazla goroutine'in tamamlanmasini bekleme ornegi.
- `Buffered_Channel/`: Buffered channel kullaniminin temel ornegi.
- `Unbuffered_Channel/`: Unbuffered channel ile ilgili farkli senaryolar.
- `Select/`: `select` yapisi ile birden fazla kanali dinleme ornegi.
- `Context/`: `context` ile veri tasima ve iptal mantiginin giris seviyesi ornegi.

## Amac

- Her ornekte kodun icine Turkce aciklama yorumlari eklendi.
- Her alt klasor icin o klasordeki ornegi aciklayan bir `README.md` dosyasi hazirlandi.
- Ornekler, Go'da eszamanli programlama konularini daha kolay takip edebilmek icin sade tutuldu.

## Calistirma

Her ornegi ilgili klasore girip su komutla calistirabilirsiniz:

```bash
go run main.go
```
