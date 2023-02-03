using System.Text;

public class Solution {
    public string Convert(string s, int numRows) {
        if (numRows == 1)
        {
            return s;
        }

        var lines = new StringBuilder[numRows];
        for (int i = 0; i < numRows; i++)
        {
            lines[i] = new StringBuilder();
        }

        int processedCharacters = 0;
        int ix = 0;
        int iy = 0;
        while (processedCharacters < s.Length)
        {
            if (ix%(numRows-1) == 0)
            {
                lines[iy].Append(s[processedCharacters]);
                processedCharacters++;
            }
            else
            {
                if (iy == (numRows-1) - ix%(numRows-1))
                {
                    lines[iy].Append(s[processedCharacters]);
                    processedCharacters++;
                }
            }

            if (iy < numRows-1)
            {
                iy++;
            }
            else
            {
                iy = 0;
                ix++;
            }
        }

        var sb = new StringBuilder();
        foreach (var line in lines)
        {
            sb.Append(line);
        }
        return sb.ToString();
    }
}