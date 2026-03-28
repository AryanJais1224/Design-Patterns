type DB interface {
    Connect()
}

type MySQL struct{}
func (m MySQL) Connect() {}

type Mongo struct{}
func (m Mongo) Connect() {}

func NewDB(dbType string) DB {
    switch dbType {
    case "mysql":
        return MySQL{}
    case "mongo":
        return Mongo{}
    }
    return nil
}
