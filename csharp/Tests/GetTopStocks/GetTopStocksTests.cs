using Solutions.GetTopStocks;

namespace Tests.GetTopStocks;

public class GetTopStocksTests
{
    private GetTopStocksSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new GetTopStocksSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var stocks = new[] { "AMZN", "CACC", "EQIXS", "GOOG", "ORLY", "ULTA" };
        var prices = new double[][]
        {
            [11.93, 11.09, 12.11, 10.93, 9.83, 8.14],
            [10.34, 10.56, 10.14, 12.17, 13.12, 8, 90]
        };
        var expected = new[] { "GOOG", "ORLY", "AMZN" };

        // Act
        var actual = solution.GetTopStocks(stocks, prices);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}