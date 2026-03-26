# Context — İstek bağlamı ve değer taşıma

`context` paketi, özellikle ağ servisleri ve goroutine ağaçlarında **iptal süresi**, **deadline** ve **istek kapsamındaki anahtar–değer** bilgisini taşımak için kullanılır. Bu klasördeki örnek, öncelikle **`WithValue` ile zincir boyunca veri taşımayı** gösterir.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `main.go` | `Product`, global `productChannel`, `main`, `F1`–`F3`, yorum satırlarında timeout senaryosu, `getProductDetailsFromExternalService`. |

## Aktif olarak çalışan kısım

1. `ctx := context.Background()` — kök context.  
2. `ctx = context.WithValue(ctx, "correlation-id", "abc123")` — (dikkat: üretimde genelde özel tip anahtar kullanılır; string anahtar çakışmaya açıktır, burada öğretim basitliği var.)  
3. `F1(ctx)` → `F2(ctx)` → `F3(ctx)` — aynı `ctx` iletilir; her biri `ctx.Value("correlation-id")` okur ve yazdırır.

## Yorum satırlarındaki bölüm

`WithTimeout`, `defer cancel()`, goroutine ile uzun süren iş ve `select` içinde `case <-ctx.Done()` deseni, zaman aşımında ana akışın nasıl kurtulacağını anlatmak için bırakılmıştır; çalıştırmak için yorumları kaldırıp uyumlu hale getirmeniz gerekir.

## `getProductDetailsFromExternalService`

10 saniye uyuyup sonra `productChannel` üzerinden `Product` gönderir; uzun süren dış servis çağrısının kaba bir simülasyonudur.

## Çalıştırma

```bash
cd Context
go run .
```

## Okuma önerisi

Resmi rehber: [https://pkg.go.dev/context](https://pkg.go.dev/context)
