public class Solution {
    public bool IsAlienSorted(string[] words, string order) {
        var map = new Dictionary<char, int>();

        for (int i = 0; i < order.Length; i++)
        {
            map.Add(order[i], i);
        }

        for (int i = 0; i < words.Length - 1; i++)
        {
            if (!IsLexicographicallyOrder(words[i], words[i+1], map))
            {
                return false;
            }
        }
        return true;
    }

    private bool IsLexicographicallyOrder(string firstString, string secondString, Dictionary<char, int> charOrders)
    {
        for (int i = 0; i < Math.Min(firstString.Length, secondString.Length); i++)
        {
           if (firstString[i] != secondString[i])
           {
                return charOrders[firstString[i]] < charOrders[secondString[i]];
           }
        }
        return firstString.Length <= secondString.Length;
    }
}