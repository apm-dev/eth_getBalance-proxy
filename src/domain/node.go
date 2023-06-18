package domain

type Node struct {
	ID         string
	Url        string
	Blockchain string
	Weight     int
}

type NodeRepository interface {
	GetNodesByBlockchain(blockchain string) ([]Node, error)
}
