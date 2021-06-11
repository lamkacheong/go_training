作业：基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

抽象了一个brother server的struct，可以处理这些同生共死http server。
main.go是使用这个struct的例子。
