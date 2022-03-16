package config

type AppConfig struct {
	KafkaConf `ini:"kafka"`
	TailConf  `ini:"taillog"`
}

type KafkaConf struct {
	Address string `ini:"addrs"`
	Topic   string `ini:"topic"`
}

type TailConf struct {
	Path string `ini:"path"`
}
