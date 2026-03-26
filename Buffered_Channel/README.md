# Buffered_Channel — Kapasiteli kanal

Tamponlu (**buffered**) channel, belirli sayıda değeri **alıcı olmadan geçici olarak** tutabilir. Bu örnek, kapasite `4` olan bir `chan int` üzerinden gönderimin nasıl ilerlediğini ve `range` ile okumanın nasıl yapıldığını gösterir.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `main.go` | `make(chan int, 4)`, gönderici goroutine, alıcı goroutine, `close`. |

## Önemli noktalar

- `make(chan int, 4)` — kanal en fazla 4 `int` tutabilir; dolmadıkça gönderici bloklanmayabilir (alıcı yavaş olsa bile, tampon dolana kadar).  
- Gönderici: `0…9` değerlerini yollar, bittiğinde `close(myChannel)` çağırır. Kapalı kanala yazmak panic üretir; bu yüzden kapatma genelde sadece gönderen tarafta yapılır.  
- Alıcı goroutine: `for data := range myChannel` ile kapanana kadar okur.  
- `time.Sleep` çağrıları çıktıyı ayırmak ve gönderim/alım hız farkını gözlemlemek için eklenmiştir.

## Tamponsuz kanal ile fark

Tamponsuz kanalda her `send`, eşzamanlı bir `receive` olmadan ilerlemez. Burada ise tampon dolu değilse gönderici kısa süre ilerleyebilir.

## Çalıştırma

```bash
cd Buffered_Channel
go run .
```
