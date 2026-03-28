var instance *Config
var once sync.Once

func GetInstance() *Config {
    once.Do(func() {
        instance = &Config{}
    })
    return instance
}
