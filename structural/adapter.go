type OldAPI struct{}

func (o OldAPI) OldMethod() string {
    return "old"
}

type Adapter struct {
    old OldAPI
}

func (a Adapter) NewMethod() string {
    return a.old.OldMethod()
}
