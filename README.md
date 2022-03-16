### kafka实现一个简单的日志收集系统
+ 出于对kafka在日志收集方面的兴趣，故实现了一个简单的demo见识一下kafka在这方面强大的性能
+ 安装kafka，同时进入server.properties配置文件修改一些配置，主要是端口号
+ 启动zookeeper`./bin/zookeeper-server-start.sh -daemon ./config/zookeeper.properties`
+ 启动kafka服务`./bin/kafka-server-start.sh ./config/server.properties`
+ 开启消费者，监听消息`./bin/kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic web_log --from-beginning`
+ 修改项目的配置文件config.ini，主要是配合本机情况进行配置
+ 启动项目`go run main.go`
+ 在my.log文件手动写入数据，kafka消费者把一行当作一条消息。当我们输入一行数据之后，按下回车，数据就会被消费者进行消费了，可以在消费者对应的终端开到消费的消息
+ 如果想应用到项目里面，就只要将日志的输出文件指定为配置文件里面的path就行了。大概的流程就是项目产生的日志->监听的日志文件->kafka