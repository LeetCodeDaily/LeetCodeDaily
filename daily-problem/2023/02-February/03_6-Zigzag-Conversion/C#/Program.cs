var s = "PAYPALISHIRING";
var numRows = 3;
var expected = "PAHNAPLSIIGYIR";
Console.WriteLine(new Solution().Convert(s, numRows) == expected ? "PASSED" : "FAILED");


s = "PAYPALISHIRING";
numRows = 4;
expected = "PINALSIGYAHRPI";
Console.WriteLine(new Solution().Convert(s, numRows) == expected ? "PASSED" : "FAILED");


s = "A";
numRows = 1;
expected = "A";
Console.WriteLine(new Solution().Convert(s, numRows) == expected ? "PASSED" : "FAILED");