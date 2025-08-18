using Solutions.LuckyMoney;

namespace Tests.LuckyMoneyTests;

public class LuckyMoneyTests
{
    private LuckyMoneySolution solution;
    
    [SetUp]
    public void SetUp() => solution = new LuckyMoneySolution();

    [Test]
    public void TestClassicalBudget()
    {
        // Arrange
        var money = 11;
        var giftees = 2;
        var expected = 1;
        
        // Act
        var actual = solution.LuckyMoney(money, giftees);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestPerfectBudget()
    {
        // Arrange
        var money = 32;
        var giftees = 4;
        var expected = 4;
        
        // Act
        var actual = solution.LuckyMoney(money, giftees);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestBudgetWithAFour()
    {
        // Arrange
        var money = 12;
        var giftees = 2;
        var expected = 0;
        
        // Act
        var actual = solution.LuckyMoney(money, giftees);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestExampleFour()
    {
        // Arrange
        var money = 13;
        var giftees = 3;
        var expected = 1;
        
        // Act
        var actual = solution.LuckyMoney(money, giftees);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}