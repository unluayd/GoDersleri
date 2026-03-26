# GoDersleri — Go Eğitim Serisi

Bu depo, **Go programlama dili** için adım adım ilerleyebileceğiniz küçük örneklerden oluşan bir **eğitim serisidir**. Amaç; **eşzamanlılık** (goroutine, channel, `select`, `context`, `WaitGroup`) ve **arayüz (interface)** gibi konuları yalın kod ve Türkçe yorumlarla pekiştirmektir.

Modül adı: `onevideogo` (`go.mod`). Her alt klasördeki örneğin ne yaptığını o klasördeki `README.md` dosyası ayrıntılı anlatır.

---

## Seride neler var?

| Konu | Klasör | Dosya |
|------|--------|--------|
| Goroutine’leri topluca beklemek | *(kök)* | `main.go` |
| Tamponlu kanal | `Buffered_Channel/` | `main.go` |
| Tamponsuz kanal — temel eşleşme | `Unbuffered_Channel1/` | `main.go` |
| Tamponsuz kanal — döngü, kapatma | `Unbuffered_Channel2/` | `main2.go` |
| Tamponsuz kanal — çoklu alıcı | `Unbuffered_Channel3/` | `main3.go` |
| `select` ile çoklu kanal | `Select/` | `main.go` |
| `context` — değer taşıma | `Context/` | `main.go` |
| Interface — tek metot, iki tür | `Interface1/` | `1.iface.go` |
| Interface — dilim ve işlev | `Interface2/` | `2.iface.go` |
| Interface — çok metotlu sözleşme | `Interface3/` | `main.go` |

---

## Kök dizin: `main.go` (`sync.WaitGroup`)

Üç anonim goroutine paralel başlatılır; her biri `defer wg.Done()` ile sayacı azaltır. `wg.Add(3)` ile beklenen iş sayısı ayarlanır; `wg.Wait()` ana goroutine’i üçü de bitene kadar bloklar. Böylece **“arka plandaki işler bitsin, sonra devam edeyim”** deseni gösterilir. Çıktıda toplam süre `time.Since` ile yazdırılır (üç `Sleep(3s)` paralel olduğu için yaklaşık 3 saniye sürer, sırayla değil).

---

## Alt klasörler (özet)

Ayrıntı için ilgili klasördeki **README.md** dosyasına bakın.

- **`Buffered_Channel/`** — Kapasiteli kanal (`make(chan T, n)`); gönderici `close` eder, alıcı `range` veya döngü ile okur. Tampon sayesinde gönderen, alıcı hazır olmasa da (tampon dolmadıkça) ilerleyebilir.
- **`Unbuffered_Channel1/`** — Tamponsuz kanalda gönderim–alım **eşzamanlı eşleşmesi**; `done` kanalı ile alıcı bittiğinde `main`’in ilerlemesi.
- **`Unbuffered_Channel2/`** — `0…10` gönderimi, `close`, alıcı tarafta `data, ok := <-ch` ile kanal kapanınca döngüden çıkma.
- **`Unbuffered_Channel3/`** — İki alıcı goroutine aynı kanalı dinler; **tek bir değer yalnızca bir alıcıya** gider (diğeri bloklanır — bu örnekte ikinci değer gönderilmediği için dikkat: pratikte deadlock riski vardır, eğitim amaçlı davranış gözlemi).
- **`Select/`** — İki kanaldan hangisi önce hazırsa `select` ile o case çalışır; `default` ile beklemeden dönüş (busy-wait + `Sleep` ile örnek yumuşatılmış).
- **`Context/`** — `Background`, `WithValue`; `correlation-id` gibi istek bağlamını `F1→F2→F3` zincirinde taşıma. Yorum satırlarında timeout + `select` ile iptal deseni de gösterilir.
- **`Interface1/`** — Küçük `iletisim` arayüzü; `*sayilar` (pointer alıcı) ve `sayi` (int takma adı, değer alıcı) aynı değişkende kullanılır.
- **`Interface2/`** — `IShippable`; `Books` / `Electronics` ve `[]IShippable` üzerinden `CalculateShippingCost` ile arayüze dayalı fonksiyon.
- **`Interface3/`** — `ICoder` ile `Encode` + `Decode`; `XCodec` / `YCodec` çok biçimliliği.

---

## Önerilen çalışma sırası

1. Kök `main.go` (WaitGroup)  
2. `Unbuffered_Channel1` → `2` → `3`  
3. `Buffered_Channel`  
4. `Select`  
5. `Context`  
6. `Interface1` → `Interface2` → `Interface3`  

Sıra ihtiyaca göre değiştirilebilir; kanalları `select` ve `context` öncesinde görmek genelde daha kolaydır.

---

## Çalıştırma

Proje kökünde `go.mod` bulunur. Örnek klasörüne girip paketi çalıştırın:

```bash
cd Buffered_Channel
go run .
```

Kökteki WaitGroup örneği:

```bash
cd /path/to/GoDersleri
go run .
```

Belirli bir dosya:

```bash
go run Interface1/1.iface.go
go run Unbuffered_Channel2/main2.go
```

---

## Katkı ve notlar

- Kod içi açıklamalar mümkün olduğunca **Türkçe** tutulmuştur.  
- Bu seri, resmi dokümantasyonun yerine geçmez; `https://go.dev/doc/` ile birlikte kullanılmalıdır.
