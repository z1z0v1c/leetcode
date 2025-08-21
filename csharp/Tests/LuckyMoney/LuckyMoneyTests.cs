using Solutions.LuckyMoney;

namespace Tests.LuckyMoney;

public class LuckyMoneyTests
{
    private LuckyMoneySolution solution;

    [SetUp]
    public void SetUp()
    {
        solution = new LuckyMoneySolution();
    }

    [Test]
    public void TestClassicalBudget()
    {
        // Arrange
        const int money = 11;
        const int giftees = 2;
        const int expected = 1;

        // Act
        var actual = solution.LuckyMoney(money, giftees);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestPerfectBudget()
    {
        // Arrange
        const int money = 32;
        const int giftees = 4;
        const int expected = 4;

        // Act
        var actual = solution.LuckyMoney(money, giftees);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestBudgetWithAFour()
    {
        // Arrange
        const int money = 12;
        const int giftees = 2;
        const int expected = 0;

        // Act
        var actual = solution.LuckyMoney(money, giftees);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleFour()
    {
        // Arrange
        const int money = 13;
        const int giftees = 3;
        const int expected = 1;

        // Act
        var actual = solution.LuckyMoney(money, giftees);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}