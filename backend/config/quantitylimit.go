package config

type QuantityLimit struct {
	Page  int64
	Limit int64
}

func (q *QuantityLimit) Drop() int64 {
	return (q.Page - 1) * q.Limit
}
