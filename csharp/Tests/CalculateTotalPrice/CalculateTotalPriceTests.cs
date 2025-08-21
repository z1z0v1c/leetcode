using Solutions.CalculateTotalPrice;

namespace Tests.CalculateTotalPrice;

public class CalculateTotalPriceTests
{
    private CalculateTotalPriceSolution solution;

    [SetUp]
    public void SetUp()
    {
        solution = new CalculateTotalPriceSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var prices = new[] { 10, 20, 30 };
        const int discount = 50;
        const int expected = 45;

        // Act
        var actual = solution.CalculateTotalPrice(prices, discount);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}