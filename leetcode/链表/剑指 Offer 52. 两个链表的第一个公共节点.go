package 链表

/**
输入两个链表，找出它们的第一个公共节点。

思路：双指针 T = O(m+n) S= O(1)
*/
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		/*假设 较长的链表m, 较短的是n, 公共长度c, 较长的前面部分是a, 较短的前面是b, 那么有啊a+c = m; b+c = n
		那么 2个链表，分别走 a+c+b  == b+c+a 就会在公共节点相遇【非常的牛逼】
			这一步非常的精髓，因为较短的一条链表，肯定会先走完，然后从较长的一条链表开始走
		*/
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
