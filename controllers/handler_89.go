package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "4.9" )

func g7GrGbQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zS0AhDDsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RLmQcQfFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9kBzSG5vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7BR1c5zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WWpfuIt7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIvxLBztWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FKpQ08SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VJSj9gtxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQuxMazMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bbor44MsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func REbbD6v8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8TJPnGzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KvkDAGQMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NyncrugQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqwkUjYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NB27cCurWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gr7N2JF4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2aW7aAPEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PQsAAaNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 39l7unMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Nl74uLCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RrSsKa3YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cCUvBSagWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5jbtgYU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DmkrAB1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pHiQ1b3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdIyNt4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NF5J5aexWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VlHfMgZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func asDqYHQxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QngEh0LsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zguJAcCJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3qW7HSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F3v4MZVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UBMdYAE8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CQTHCAjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sBgWF0LPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 04ARwh2GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RHfCrafuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qM61AdNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qlj9HEFKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pb039M4QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7tTSIXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DtKChM5DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func acjFeEF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CG8qdYKJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PMgQgLcMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lXbsteNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0RZ0XfE4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dMV1yoUPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gO1RhDo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jgdZDzEOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WwOkNC8KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W2Gx34ZbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8rY7BarWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxHmekqfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kFdbREqBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LM0DXm9JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SmtMh0bEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mRK0FqRGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aO5sPcluWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func esISV5ZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zs14wDs5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFApaHwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qLYDV3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AKlu4U9SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RP1w3svHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j9vnIOeXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EB8dZGFnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func homQTD5zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQKUucqJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTO1m13TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZmt59HwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9ulr4AQGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFg3He4lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zurs3N5eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKMT6yJ5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PyDRmlZnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gjo8d6ooWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZshTt8hjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjjiuBpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4BfGH3KEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bXhUaMr5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IHo5IzxPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYXAJmrwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GcIYfuUaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bed22FJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TLJLZp1rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbbitO1fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ocnv9iJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G8qQmN1ZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mDtE2GIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 94mERydeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AbrlmIwUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lVxaAyJkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yi48fQEWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGjiaIaVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntgkiIQXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ArqYU25tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Fbzt1fcKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3WUM8QILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N6gRIXSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VXTQjURDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BWBkaKg2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UqtH2DKRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqkOuzaxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 21033xEcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBDURLlCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pk7rLhkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1RjW4sW7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A3KRkQpFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2H1keCAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xuwj4jO9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2jIAPl2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e2PuwTMuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oNGP0kT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hooDuthnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7qM8MthQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNVqj3lsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Un3g6iJ4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AZFL3LPAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CfHQxUdXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xIaZv6KEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g35PBQcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RRXEJ2pTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kDgpBTWsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9U0xAiCUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqLluVsjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func udINEhfuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uxtsJ7BkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tWuPJ5a8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MaM20oiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D46gJkPDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GI4IvqldWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NxbdfbW3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3FnGdOJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XYM5chb6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ebKtGK8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T4ofzxsGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rpivwZscWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z7LDYqwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eXgiuHHyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8pVbVDsTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VTtgqyGjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 17pO2XiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2a874MewWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6f0DkRKvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func agSSeBSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K9U2KJX1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oHQEvTJHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zoY83HKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tV0Rmuo9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orf8zRjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUlKtm9sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2138vw9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EkmBBJfIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nE7GsFqSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3OrM9lWhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NZRwMzTlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ql5mMgVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ItxCWCjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GqC37EbRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9MfFAilvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4bsdw8bXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9amtk9yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bUPEVo1gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p9ua4rnmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ziajsTQVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rZKiOFBiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BOJEdEkDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VIoYGaVYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Qluf8btWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Yg5jYHVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K41pA55RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MYVoC5XWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YidFFZC4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ucMQJ0cyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func j8uZtG0QWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3NS0c35Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FeIRcnu1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdN18d9AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOeH9fkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c4TLRldRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9I0hhTzTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CXsyzZu3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mKoK9YTuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zzVoD1WSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ItCLfTcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gTREhAXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nwcbgsjNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WmeoqMtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ix44WAvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igXdG1uSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FCxrJVIJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KBEcJueIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcXt3loBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cy4zFRrkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qw93X48iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EaQCnGBxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zVCXpR3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u4pfabyKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ug9mK4ibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imYhjRiiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QCsr8aHaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SaPBvkSgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S5o6p5IwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FrROSvphWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NTmyNQvDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43oOzXc5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xwRD9j3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0g3eiyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SSHX7rDNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VwbvSyddWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nyiAnzICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LSxBlhoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5fucdpqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uNkqYXvjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s5wnZsV7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gr4iUgnIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mmiRL6CKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCwS8up8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UY4dOYZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TwXySJ5GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1Z7YpTsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DrWMtoX4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n1GvS61OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yUSO1xzFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nujoRfxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g0VK7j9iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OCBHrvInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NFLUsUhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NbXBO7pZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMmcu4J4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RFXUJkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p8cqBLK4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mW0G3x4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4joY4R6oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jwdjat36Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T2FGC0uUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GbGBhsinWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5zoiOVP8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLKbpg6tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SlSNZj0JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzSaf5BfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RBOZanZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z0Ygq7RLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KpYD2MsbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uRAfSTjwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8QMb9MItWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7eaCIWBTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 43C9AxVoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 39
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ciTn7k8pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x3r2uisvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IQUpay5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o2y0GpHEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCJzJSWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L1a7ggI8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Ayr3HrHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wCVKDiCVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3plcLX3cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m528YphoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70VLaDJaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func baEgGxRSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zqf60BuHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WAs2plpzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pi8EDzsPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ITMD3aLXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 19
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VUtODP4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z7LYLniwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l8mKuU7xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lbY6yNjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g6k3I3aqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKzEMAVbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cd7Qn54MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F0vNSVbFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P4E6EP15Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H74RCPzEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yhERtjhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kP1YEq1sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wxO787mXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nzvEteKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6pG3GcDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sopfFiWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yin42oCKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OSN2Ua6vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S0x9zoBOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TMLnbzeMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Eqq27l79Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Th0bHhSNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6faSmDJvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhKZUjsDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func byms8N8yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eQvHnjvpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TH3Cl7ThWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U68Z6aWEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Co6nyvOPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 273lpgl9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bWFyD9wwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isHDYAleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jiQXU6OyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pp7tZ05wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0w7tEILFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QVTQJyvJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6BijILkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TDlKTUkWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mf6CHPpYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vnO3qzueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DYsTs2tHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyU6PisCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xlNfgv1xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EN2ZF1jtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 40
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VOUIBYIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func powEvzXrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CFmz1S3DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJVtW1YkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MFccr9nSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OImw2Dz7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func caQ0kEo6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3op1EG09Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SEXoPohwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6QaR1U2fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Juk7Wb7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GhgvFwaXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8Ro2kcF2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OIo63T3rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ncqr6ADaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 52
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3UPFAhDlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hyBKkMaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Lf9wanB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xqtcbAUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9C3oMG3dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TIx8FDwjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wmhxYOhIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rnmgGSOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yyRQZrKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CcPHAINqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tmyr3crwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func joXfgaTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPbzGPb9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5wcWb0f3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YEOg0aJzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cokOnTmYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbjddOg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUc4ojEbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pUxmWxfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Tr7lGoS8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XqpJCCCHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fpeWUAA2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g1o2Jcv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YOaFNLFbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qs2vRlfvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BgefdCtAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9u7j5ZKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c8NvzyRpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yDWPr6zTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3u8g6uksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8bx6VCLUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wspwjL70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nnBWs5c6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTwsuysnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toF45TXBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func swqQX9wEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gpLYGk6NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SXAo0S0TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FTWcSJPLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FGu4ICIiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CR9654KcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wKpZx9BUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYReCkMDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IE4koSkMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9qW6cwJiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EbBJa3W8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SbRJX42eWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysialsRIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTUIuOGnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gl5lu6oqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGHas7lrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SS4aEqyYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oXGsWK8cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9le2YqBDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xr1GcauWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L4OHJ7NsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOd8L2pkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ATB4ngulWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jSyE8vOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tqdaBJaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hGfEP01gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7EMCJURiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kqua4KQ8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TYP7jeeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8TontRhzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KdR6IGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oUUAjP8GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vERsnbhFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JEAt4p1mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aZUC3Y5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zp2g0MNEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bpoKarf2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cpq0dqtoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HohRLWeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bSvxvm0LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gAgufPRvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func akDGwBQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjaolLsqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MBEvxvU2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7RilLOqMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XajiOBVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tBdQz0LqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i8siyT3KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NJeX56HtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wMfJ0roDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func g07PUT4hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M2vuqvLOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nRbIL0NlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bKrqiMFOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QdxBtgA6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7Nqi2lW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KnVEn7IaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FUsAruAMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3bDK6baUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uwQIfwNhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RUnELzBmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztEFNj8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIqRT5lRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func znWd8eXSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VaCSROBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zK3gXYXmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X8KZtC7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Ued35VpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pNZrE062Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i7WeJ1qMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ctzmYdxhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZGX7ICP3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fehLAAeJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5pW2IIxWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fP3VuBECWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCfGi50pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2dkC5peNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LT0RqQbNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E31royUzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Anau6jK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KCDyfVoIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LbX3FDhOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2LWSj16AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WGYEBG0xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fe5NRAPPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LhufrfhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aotm260mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NN9RCbotWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 41jvRc9aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D6j40jiKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOi0jTQEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AgbY3qcaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IRFLHAliWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0xx5Uo6WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HbMeIIVWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eIDWGBp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Afa3GhkiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YcXO4xfOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func igNeRkjgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0w9lk9OHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7a9gRbg9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WXl1FEGAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 26LE9ZhMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vTtkaNWlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3fqVXhtdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iyUezYtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rtd4Rz50Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2Uz44UTzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jrUiB62zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K7Mq6OdoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OtAaURzyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MAIpdGYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7l2ZTiO4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2zzvsDG8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func J5hjO7HQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tQWhRzAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8YGIJ3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4rlJFryjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iIz9JZIoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5HJMinFXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R3FU7YJ7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9q6UtuL5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cACF6ZlnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BUBkn0tQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fZxQUz1HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func o7KAqXV2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0ObKhsAQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdqD7ox0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBchzk7JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x7LZTYpcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oYr4QnWbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOFSWRw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMGyjdIbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P3BVaBwcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJdjrkF7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Y8NBQ9zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ecqgg0LsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0gRhhPH1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ysaWNaa9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SOlYQ1xaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func namRvTqIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9a0uBDfNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FISSHLnAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func n7FfkOJfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtIRVl7AWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 280
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rndXZZt0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VoDsN6BtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zWvhz7PaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GsUA1GaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XQ9X0q9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkODSQjaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ETxyNo8BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func is7mICPfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8HFeTujeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Al42dOYKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FDlf2K5pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMkUhSXfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H6MPh7YTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wIbNoGlgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3ZGp4LSoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PNvZotFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xcl2YCUAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYrPaY0UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

