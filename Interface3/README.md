# Interface3 — `ICoder`: çok metotlu arayüz ve codec türleri

Bir arayüzde **birden fazla metot** olduğunda, o arayüzü karşılayan her somut türün **hepsini** uygulaması gerekir. Bu yüzden yalnızca `Encode` veya yalnızca `Decode` bilen ayrı struct’lar tek başına `ICoder` olamaz; örnekte her aile **`XCodec` / `YCodec`** gibi birleşik türlerle modellenir.

## Dosya

| Dosya | Açıklama |
|--------|-----------|
| `main.go` | `ICoder`, `XCodec`, `YCodec`, metotlar ve `main`. |

## `ICoder`

- `Encode(value string) string`  
- `Decode(value string) string`  

## `XCodec` ve `YCodec`

Boş struct’lar; davranış tamamen metotlarda. Her ikisi de `Encode` ve `Decode` için `fmt.Println` ile hangi codec’in çalıştığını gösterir ve sabit bir string döner (derslik basitlik).

## `main` akışı

`var coder ICoder` — önce `coder = &XCodec{}`, `Encode` / `Decode` çağrılır; sonra `coder = &YCodec{}` ile aynı değişkene başka somut tür verilir ve çağrılar Y implementasyonuna gider.

## Çalıştırma

```bash
cd Interface3
go run .
```
