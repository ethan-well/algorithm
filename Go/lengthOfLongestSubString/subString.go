func lengthOfLongestSubString(s string) int {
  m := map[byte]int{}
  n := len(s)

  r = -1
  subLen = 0
 
  for i := 0; i < n; i ++ {
    if i != 0 {
      delete(m, s[i-1])
    }

    for r + 1 < n && m[r+1] == 0 {
      m[s[r+1]]++
      r++
    }
    
    subLen = max(subLen, r - i + 1)
  }
}

func max(x, y int) int {
  if x < y {
    return y
  }
  
  return x
}
