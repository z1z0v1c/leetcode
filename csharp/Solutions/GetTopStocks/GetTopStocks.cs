/// <summary>
/// Live coding asignment
/// 
/// In this problem, you'll be given a list of stock prices, and you'll be asked to
/// return the three stocks with the highest average price.
/// 
/// Implement the method GetTopStocks(stocks, prices) which takes as input:
///     - an array of strings (stocks), representing the considered stocks.
///     - an array of two dimensions (prices), representing the stock prices (inner lists) for
///       each day (outer list).
/// 
/// An example input would look like this:
///     ["AMZN", "CACC", "EQIXS", "GOOG", "ORLY", "ULTA"]
///     [11.93, 11.09, 12.11, 10.93, 9.83, 8.14], [10.34, 10.56, 10.14, 12.17, 13.12, 8,90]
/// 
/// Your GetTopStocks method should return an array containing the names of the three
/// stocks with the highest average value. The array should be sorted by decreasing
/// average value. You're guaranteed that each stock will have a unique average value.
/// 
/// For the above example, the correct output would be:
///     ["GOOG", "ORLY", "AMZN"]
/// </summary>

namespace Solutions.GetTopStocks;

public class GetTopStocksSolution
{
    public string[] GetTopStocks(string[] stocks, double[][] prices)
    {
        var topStocks = new string[3];
        var average = new double[stocks.Length];

        for (int i = 0; i < stocks.Length; i++)
        {
            var sum = 0d;
            for (int j = 0; j < prices.Length; j++)
            {
                sum += prices[j][i];
            }

            average[i] = sum / prices.Length;
        }

        var orderedStocks = new Dictionary<double, string>();

        for (int i = 0; i < stocks.Length; i++)
        {
            orderedStocks.Add(average[i], stocks[i]);
        }

        List<double> ordered = [.. orderedStocks.Keys];
        ordered.Sort();

        for (int i = 0; i < 3; i++)
        {
            topStocks[i] = orderedStocks.GetValueOrDefault(ordered.ElementAt(stocks.Length - i - 1))!;
        }

        return topStocks;
    }
} 