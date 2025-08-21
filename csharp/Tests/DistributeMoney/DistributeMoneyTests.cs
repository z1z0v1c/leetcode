using Solutions.DistributeMoney;

namespace Tests.DistributeMoney;

public class DistributeMoneyTests
{
    private DistributeMoneySolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new DistributeMoneySolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const int money = 20;
        const int children = 3;
        const int expected = 1;

        // Act
        var actual = solution.DistMoney(money, children);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        const int money = 16;
        const int children = 2;
        const int expected = 2;

        // Act
        var actual = solution.DistMoney(money, children);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        const int money = 1;
        const int children = 2;
        const int expected = -1;

        // Act
        var actual = solution.DistMoney(money, children);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        const int money = 17;
        const int children = 2;
        const int expected = 1;

        // Act
        var actual = solution.DistMoney(money, children);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}