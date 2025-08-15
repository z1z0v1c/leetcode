namespace Tests.GetTopStocksTests;

public class GetTopStocksTests
{
    private GetTopStocks getTopStocks;

    [SetUp]
    public void Setup() => getTopStocks = new();

    [Test]
    public void TestExampleOne()
    {
        var stocks = new string { "AMZN", "CACC", "EQIXS", "GOOG", "ORLY", "ULTA" };
        var prices = new double
        {
            {12.81, 11.09, 12.11, 10.93, 9.83, 8.14},
            {10.34, 10.56, 10.14, 12.17, 9.66, 8,90}
        };

        var expected = { "GOOG", "ORLY", "AMZN" };
        var actual = getTopStocks.GetTopStocks(stocks, prices);

        Assert.That(actual, Is.EqualTo(expected));
    }

}