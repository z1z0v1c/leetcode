using Solutions.DistributeMoney;

namespace Tests.DistributeMoney;

public class DistributeMoneyTests
{
    private DistributeMoneySolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var money = 20;
        var children = 3;
        var expected = 1;
        
        // Act
        var actual = solution.DistMoney(money,  children);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        var money = 16;
        var children = 2;
        var expected = 2;
        
        // Act
        var actual = solution.DistMoney(money,  children);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var money = 1;
        var children = 2;
        var expected = -1;
        
        // Act
        var actual = solution.DistMoney(money,  children);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}