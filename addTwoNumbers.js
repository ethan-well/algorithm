/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var addTwoNumbers = function(l1, l2) {
  var arr1 = [];
  var arr2 = [];

  do {
    arr1.push(l1.val);
    l1 = l1.next;
  } while ( l1 != null );

  do {
    arr2.push(l2.val);
    l2 = l2.next;
  } while ( l2 != null );

  return arr1;
};
