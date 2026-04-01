type UserRepo interface {
    GetUser(id int) string
}

type MySQLRepo struct{}

func (m MySQLRepo) GetUser(id int) string {
    return "user from mysql"
}
