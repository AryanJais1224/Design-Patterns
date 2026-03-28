type User struct {
    Name  string
    Email string
}

type UserBuilder struct {
    user User
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
    b.user.Name = name
    return b
}

func (b *UserBuilder) Build() User {
    return b.user
}
