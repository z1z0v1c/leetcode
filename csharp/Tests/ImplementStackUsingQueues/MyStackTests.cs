using Solutions.ImplementStackUsingQueues;

namespace Tests.ImplementStackUsingQueues;

public class MyStackTests
{
    private MyStack solution;

    [SetUp]
    public void Setup()
    {
        solution = new MyStack();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        const int first = 1, second = 2;
        const int expectedTop = 2, expectedPop = 2;
        
        // Act
        solution.Push(1);
        solution.Push(2);
        var actualTop = solution.Top();
        var actualPop = solution.Pop();
        var actualEmpty = solution.Empty();
        
        // Assert
        Assert.Multiple(() =>
        {
            Assert.That(actualTop, Is.EqualTo(expectedTop));
            Assert.That(actualPop, Is.EqualTo(expectedPop));
            Assert.That(actualEmpty, Is.True);
        });
    }
}