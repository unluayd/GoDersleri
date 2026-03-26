# Unbuffered_Channel3 — Tek değer, iki alıcı goroutine

Bu örnek, **aynı kanalı dinleyen birden fazla goroutine** olduğunda Go’nun kanalı nasıl paylaştığını gözlemlemek içindir: **bir gönderilen değer yalnızca bir alıcı tarafından alınır.**

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `main3.go` | `chan int`, iki alıcı goroutine, `main` içinde tek `myChannel <- 10`. |

## Davranış

1. İki goroutine de `<-myChannel` ile bloklanır ve değer bekler.  
2. `main` `10` gönderdiğinde **yalnızca biri** bu değeri alır; diğeri hâlâ bloklu kalır.  
3. İkinci bir değer gönderilmediği için program, ikinci goroutine yüzünden **sonlanmayabilir** (deadlock veya asılı goroutine). Bu, eğitim amaçlı “tek mesajın tek tüketiciye gittiği” kuralını göstermek içindir; üretim kodunda tüm goroutine’lerin tamamlanacağı tasarım gerekir.

## Çalıştırma

```bash
cd Unbuffered_Channel3
go run .
```

veya

```bash
go run main3.go
```
