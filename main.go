package main

import (
	"fmt"
	"strconv"
)

// Вам даны "головы" двух отсортированных по возрастанию связанных списков l1 и l2.
// Объедините два списка в один отсортированный список.
// Список должен быть составлен путем соединения вместе узлов первых двух списков.

// Результат – "голова" объединенного связанного списка.

// Пример:
// l1: 2 -> 3 -> 5
// l2: 1 -> 3 -> 5

// Ожидаемый результат:
// 1 -> 2 -> 3 -> 3 -> 5 -> 5

// ListNode Узел списка выглядит следующим образом
type ListNode struct {
	Val  int
	Next *ListNode
}

func (list ListNode) String() string {
	str := strconv.Itoa(list.Val)
	l := list.Next
	for l != nil {
		str = str + fmt.Sprintf(" -> %d", l.Val)
		l = l.Next
	}
	return str
}

// main run me
func main() {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}
	l3 := mergeTwoLists(l1, l2)
	fmt.Println("res: ", l3)
}

// mergeTwoLists merge 2 linked lists
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil {
		return nil
	}

	var tmp = new(ListNode)
	var cursor = tmp
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			cursor.Next = l1
			l1 = l1.Next
		} else {
			cursor.Next = l2
			l2 = l2.Next
		}
		// next step
		cursor = cursor.Next
	}

	// last
	if l1 != nil {
		cursor.Next = l1
	} else {
		cursor.Next = l2
	}

	return tmp.Next
}
