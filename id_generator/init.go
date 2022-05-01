package id_generator

import (
	"github.com/bwmarrin/snowflake"
)

var Node *snowflake.Node

func Init() {
	var err error
	Node, err = snowflake.NewNode(1)
	if err != nil {
		panic("node初始化失败")
	}
}

func GenerateId() string {
	return Node.Generate().String()
}
