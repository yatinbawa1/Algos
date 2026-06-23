package mergesortedlinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func isAvailable(item *ListNode) bool {
	return item != nil
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	mergedList := &ListNode{}
	tail := mergedList

	for isAvailable(list1) && isAvailable(list2) {
		if list1.Val <= list2.Val {
			tail.Next = list1
			list1 = list1.Next
		} else {
			tail.Next = list2
			list2 = list2.Next
		}

		tail = tail.Next
	}

	// If one list finishes first
	// then these will add
	// rest of other lists items
	// to the merge list

	if isAvailable(list1) {
		tail.Next = list1
	} else {
		tail.Next = list2
	}

	return mergedList.Next
}
