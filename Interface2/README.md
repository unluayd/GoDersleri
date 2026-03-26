# Interface2 — `IShippable` ve dilim üzerinden hesaplama

Bu örnek, gerçek hayata yakın bir **“şunun kargosu ne kadar?”** sözleşmesini arayüzle modellemeyi ve **fonksiyonun somut türleri bilmeden** `[]IShippable` ile çalışmasını gösterir.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `2.iface.go` | `IShippable`, `Books`, `Electronics`, `main`, `CalculateShippingCost`. |

## `IShippable`

Tek metot: `ShippingCost() int`. Kargo ücreti hesaplanabilen her ürün bu sözleşmeyi karşılayabilir.

## Ürün türleri

- **`Books`** — `desi` alanı; maliyet `desi * 10`.  
- **`Electronics`** — yine `desi`; maliyet `desi * 20`.  

Her iki metot da **pointer alıcı** ile tanımlandığı için (`*Books`, `*Electronics`), arayüz değişkenine ve dilime **`&Books{...}`**, **`&Electronics{...}`** konmalıdır. Sadece değer (`Books{...}`) yazmak, bu örnekte arayüz uyumsuzluğuna yol açar.

## `main` içeriği

1. `var product IShippable` — önce `&Books`, sonra `&Electronics` atanır; `ShippingCost` çıktıları farklıdır.  
2. `cart := []IShippable{&Books{...}, &Electronics{...}}` — tek dilimde karışık türler.  
3. `CalculateShippingCost(cart)` — döngüde yalnızca arayüz metodu çağrılır; fonksiyon `Books` / `Electronics` import etmek zorunda değildir.

## Çalıştırma

```bash
cd Interface2
go run .
```
