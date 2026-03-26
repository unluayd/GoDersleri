# Interface1 — `iletisim` arayüzü ve çok biçimlilik

Bu klasördeki örnek, Go’da **interface** tanımının ne işe yaradığını ve **aynı arayüz değişkenine** farklı somut türlerin atanabilmesini gösterir.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `1.iface.go` | `iletisim` arayüzü, `sayilar` struct’ı, `sayi` tip takma adı ve `main`. |

## `iletisim` arayüzü

`topla() int` metodunu bildirir. Bu imzayı sağlayan her tür, derleyici açısından `iletisim` ile uyumludur.

## Somut türler

### `sayilar`

Üç alan: `bir`, `iki`, `uc`. `topla` metodu **pointer alıcı** ile tanımlıdır: `func (sayi *sayilar) topla() int`. Bu yüzden arayüze atanırken `&sayilar{...}` veya mevcut değişkenin adresi (`&v`) kullanılır; yalnız `sayilar` değeri atanamaz.

### `sayi`

`type sayi int` ile `int` için yeni bir isim verilir. `topla` **değer alıcı** ile tanımlıdır ve bu örnekte `int(f)` döndürür (yani saklanan sayının kendisi). `sayi(10)` doğrudan `iletisim` değişkenine atanabilir.

## `main` akışı

1. `var a iletisim` — kutunun dış tipi her zaman arayüz.  
2. `a = &v` — içeride `*sayilar`; `a.topla()` üç alanın toplamını verir.  
3. `a = sayi(10)` — içeride `sayi`; `a.topla()` `10` döner.  

Aynı `a` değişkeni, farklı anda farklı somut tür taşır; çağrılan metot **dinamik tipe** göre seçilir.

## Çalıştırma

```bash
cd Interface1
go run .
```

veya

```bash
go run 1.iface.go
```
