func stage1(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for x := range in {
            out <- x * 2
        }
        close(out)
    }()
    return out
}
