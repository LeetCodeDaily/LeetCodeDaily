var words = new String[] { "hello", "leetcode" };
var order = "hlabcdefgijkmnopqrstuvwxyz";
var expected = true;

Console.WriteLine(new Solution().IsAlienSorted(words, order) == expected ? "PASSED" : "FAILED");


words = new String[] { "word","world","row" };
order = "worldabcefghijkmnpqstuvxyz";
expected = false;

Console.WriteLine(new Solution().IsAlienSorted(words, order) == expected ? "PASSED" : "FAILED");


words = new String[] { "apple","app" };
order = "abcdefghijklmnopqrstuvwxyz";
expected = false;

Console.WriteLine(new Solution().IsAlienSorted(words, order) == expected ? "PASSED" : "FAILED");