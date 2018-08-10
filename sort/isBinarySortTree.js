function isBinarySortTree(node) {
  result = true;
  if(node.left !== null) {
    if(node.left.value > node.value) {
      return false;
    } else {
      result = isBinarySortTree(node.left);
    }
  }

  if(result && node.right !== null) {
    if(node.value > node.right.value) {
      return false;
    } else {
      result = isBinarySortTree(node.right);
    }
  }

  return result;
}