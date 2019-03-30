package models

type HelloworldDao struct {
}

func (dao *HelloworldDao) Hello() *map[string]string {
	res := make(map[string]string)
	res["hello"] = "world"
	return &res
}
