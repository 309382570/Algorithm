package palindromeLinkedList

/*
Given a singly linked list, determine if it is a palindrome.

Example 1:

Input: 1->2
Output: false
Example 2:

Input: 1->2->2->1
Output: true
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

// 判断列表是否回文， 注意！！不是回环
func isPalindrome(head *ListNode) bool {

}
