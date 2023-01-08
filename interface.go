package id

type IGeneratorId interface {
	NextId() Uid
}
