package model

import pb "albelt.top/mxshop_protos/albelt/good_srv/go"

type Brand struct {
	Id   int32
	Name string
	Logo string
}

func NewBrandFromPb(b *pb.Brand) *Brand {
	if b == nil {
		return nil
	}

	return &Brand{
		Id:   b.Id,
		Name: b.Name,
		Logo: b.Logo,
	}
}
