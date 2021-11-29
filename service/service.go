package service

import "aurora/dao"

type Service struct {
	dao *dao.Dao
}

func New() *Service {
	s := &Service{
		dao: dao.New(),
	}
	return s
}
