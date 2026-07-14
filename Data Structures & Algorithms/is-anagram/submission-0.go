func isAnagram(s string, t string) bool {
    // If lengths are different, they cannot be anagrams
    if len(s) != len(t) {
        return false
    }

    // Fixed array for 26 lowercase English letters
    var counts [26]int

    // Increment for string s, decrement for string t
    for i := 0; i < len(s); i++ {
        counts[s[i]-'a']++
        counts[t[i]-'a']--
    }

    // If all counts are 0, it's a valid anagram
    for _, count := range counts {
        if count != 0 {
            return false
        }
    }

    return true
}