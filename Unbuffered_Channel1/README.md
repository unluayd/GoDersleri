# Unbuffered_Channel1 — Tamponsuz kanal ve `done` sinyali

Bu klasörde **yalnızca** aşağıdaki dosya bulunur. Tamponsuz (unbuffered) kanalda veri aktarımının **gönderen ve alıcıyla aynı anda eşleşmesi** gerektiği vurgulanır.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `main.go` | `chan string`, bir alıcı ve bir gönderici goroutine; tamamlanmayı `done chan bool` ile bekletme. |

## Akış

1. `myChannel := make(chan string)` — tampon yok; `make(chan string, 0)` ile aynı anlam.  
2. Alıcı goroutine: `<-myChannel` ile mesajı okur, `fmt.Println` yapar, ardından `done <- true` ile ana akışa “bitti” sinyali verir.  
3. Gönderici goroutine: `myChannel <- "Hello from goroutine 1"` ile iletir.  
4. `main`: `<-done` ile alıcının işini bitirmesini bekler; aksi halde program alıcıdan önce çıkabilir.  
5. Son `time.Sleep(1 * time.Second)` çıktının görünmesi için kısa süre tanır.

## Diğer tamponsuz örnekler

Çoklu gönderim / `close` ve çoklu alıcı senaryoları sırasıyla **`Unbuffered_Channel2`** ve **`Unbuffered_Channel3`** klasörlerindedir.

## Çalıştırma

```bash
cd Unbuffered_Channel1
go run .
```
