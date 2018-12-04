package single

import (
	"errors"
)

var (
	ErrListInvalid = errors.New("链表无效")
	ErrNodeInvalid = errors.New("节点无效")
)

// 单链表
type SingleLinkedList struct {
	Data_ interface{}
	Next_ *SingleLinkedList
}

// 新建

// 初始化

// 插入
func (sll *SingleLinkedList) Insert(source, node *SingleLinkedList) error {
	if source == nil || node == nil {
		return ErrNodeInvalid
	}

	// 该链表中不包含带插入位置的节点
	if !sll.IsHas(source) || sll == nil {
		return ErrListInvalid
	}

	// 执行插入
	node.Next_ = source.Next_
	source.Next_ = node
	return nil
}

// 删除

// 查找
func (sll *SingleLinkedList) IsHas(node *SingleLinkedList) bool {

	// 复制一份，防止修改原有sll的头节点地址
	temp := sll

	// 节点是空节点
	if node == nil {
		return false
	}

	// 遍历链表
	for {

		// 节点为空
		if temp == nil {
			return false
		}

		// 数据相等
		if temp.Data_ == node.Data_ {
			return true
		}
		temp = temp.Next_
	}
}
