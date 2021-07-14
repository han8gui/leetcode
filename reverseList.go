package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode
	newHead = reverse(head.Next)
	head.Next.Next = head // 链表指向之前一个结点
	head.Next = nil // 将关系断开
	return newHead
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	for head != nil {
		// 指针指向之前的
		nextNode := head.Next // 下一个节点的位置
		head.Next = newHead   // 当前节点向前一个节点
		newHead = head        // 后移指针指向前一个
		head = nextNode       // 指针后羿，处理下一个节点
	}
	return newHead
}

func ShowNode(p *ListNode) { //遍历
	for p != nil {
		fmt.Println(*p)
		p = p.Next //移动指针
	}
}

func main() {
	var head = new(ListNode)
	head.Val = 0
	var tail *ListNode
	tail = head //tail用于记录最末尾的结点的地址，刚开始tail的的指针指向头结点
	for i := 1; i < 10; i++ {
		var node = ListNode{Val: i}
		(*tail).Next = &node
		tail = &node
	}
	ShowNode(head)
	res := reverse(head)
	ShowNode(res)
}
