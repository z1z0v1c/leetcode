using Solutions.CalculateTotalPrice;

namespace Tests.CalculateTotalPrice;

public class CalculateTotalPriceTests
{
    private CalculateTotalPriceSolution solution;

    [SetUp]
    public void SetUp() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var prices = new[] { 10, 20, 30 };
        var discount = 50;
        var expected = 45;
        
        // Act
        var actual = solution.CalculateTotalPrice(prices, discount);
        
        // Assert
        Assert.That(expected, Is.EqualTo(actual));
    }

}