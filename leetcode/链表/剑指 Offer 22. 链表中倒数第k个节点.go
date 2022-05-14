package 链表

/**
输入一个链表，输出该链表中倒数第k个节点。为了符合大多数人的习惯，本题从1开始计数，即链表的尾节点是倒数第1个节点。

* 快慢双指针 T = O(n) S = O(1)
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getKthFromEnd(head *ListNode, k int) *ListNode {
	fast, slow := head, head
	for fast != nil && k > 0 {
		fast = fast.Next
		k--
	}
	//最后fast指向尾指针后的nil,slow指向倒数第二个节点
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
