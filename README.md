# OneVideoGo

Bu depo, Go’da **eşzamanlılık** (goroutine, channel, `select`, `context`, `WaitGroup`) ve **arayüz (interface)** konularını küçük, çalıştırılabilir örneklerle toplar.

## Klasörler

### Kök dizin

- `main.go`: `sync.WaitGroup` ile birden fazla goroutine’in bitmesini bekleme örneği.

### Arayüz örnekleri

- `Interface1/`: `iletisim` arayüzü; `*sayilar` ve `sayi` türleriyle aynı arayüz değişkeninde çok biçimlilik.
- `Interface2/`: `IShippable` ve farklı ürün türleri (`Books`, `Electronics`); dilim üzerinden toplam kargo hesabı.
- `Interface3/`: `ICoder` arayüzü; `XCodec` ve `YCodec` ile tek kutuda hem `Encode` hem `Decode`.

### Kanallar

- `Buffered_Channel/`: Tamponlu channel kullanımı.
- `Unbuffered_Channel1/`, `Unbuffered_Channel2/`, `Unbuffered_Channel3/`: Tamponsuz channel senaryoları (her biri ayrı giriş dosyası: `main.go`, `main2.go`, `main3.go`).

### Diğer

- `Select/`: Birden fazla kanalı `select` ile dinleme.
- `Context/`: `context` ile iptal ve değer taşıma girişi.

Bazı alt klasörlerde konuya özel `README.md` dosyaları da vardır.

## Amaç

- Örnek kodlarda Türkçe açıklama yorumları kullanılmaya çalışıldı.
- Eşzamanlılık ve arayüzleri ayrı klasörlerde izole ederek takip etmeyi kolaylaştırmak.

## Çalıştırma

`go.mod` dosyası proje kökündedir. Bir örneği çalıştırmak için o örneğin klasörüne girip `go run .` kullanın; paketteki tüm `*.go` dosyaları birlikte derlenir.

Örnek:

```bash
cd Interface3
go run .
```

Kök dizindeki `WaitGroup` örneği için proje kökünde:

```bash
go run .
```

`Interface1` gibi tek dosya adı farklı olan paketlerde de aynı klasörde `go run .` yeterlidir. İsterseniz doğrudan dosya yolu da verebilirsiniz:

```bash
go run Interface1/1.iface.go
```
