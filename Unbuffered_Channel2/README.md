# Unbuffered_Channel2 — Döngü ile gönderim, `close` ve `ok` kontrolü

Bu klasör, tamponsuz kanala **ardışık sayı gönderme**, gönderim bitince **`close`** etme ve alıcıda **`data, ok := <-ch`** ile kanalın kapalı olup olmadığını anlama örneğidir.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `main2.go` | `chan int`, gönderici goroutine (`0…10`), ana goroutine’de sonsuz döngü + `isOpen` kontrolü. |

## Gönderici goroutine

- `for i := 0; i <= 10; i++` ile her değer `myChannel <- i` ile iletilir; her gönderimde tamponsuz kanal nedeniyle bir alıcının hazır olması gerekir.  
- Ardından `close(myChannel)` çağrılır; artık kanala yeni değer gönderilmez, alıcılar kalanları okuyup `ok == false` görür.

## Alıcı (ana goroutine)

- `for { ... }` içinde `data, isOpen := <-myChannel`.  
- `isOpen == false` ise kanal kapanmış ve tamponda veri kalmamış demektir; döngü `break` ile biter.  
- Aksi halde `data` yazdırılır.

## Çalıştırma

```bash
cd Unbuffered_Channel2
go run .
```

veya

```bash
go run main2.go
```
