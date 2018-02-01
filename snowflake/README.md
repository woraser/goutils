# snowflake
这是对SnowFlake算法的一个改进版，在原算法基础上提供了对各域的位数的自定义设置。
如在一些场景下，dateCenterId、workerId并不需要占太多的位数，反而sequence部分需要占较多的位数。此时就可以根据具体业务场景自己设置。
并支持在多线程访问时的安全问题。自己测试时连续生成了大约一亿多的ID，并没有发现重复的。详见相关的测试代码。
```
//默认生成器，默认的各位数按原生的算法
gentor1, err := NewIDGenerator().SetWorkerId(100).Init()
if err != nil {
    fmt.Println(err)
}

//自定义位数的生成器
gentor2, err := NewIDGenerator().
    SetTimeBitSize(48).
    SetSequenceBitSize(10).
    SetWorkerIdBitSize(5).
    SetWorkerId(30).
    Init()
if err != nil {
    fmt.Println(err)
}

for i := 0; i < 100; i++ {
    id1, err := gentor1.NextId()
    if err != nil {
        fmt.Println(err)
        return
    }
    id2, err := gentor2.NextId()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%d%s%d\n", id1, b2, id2)
}
```