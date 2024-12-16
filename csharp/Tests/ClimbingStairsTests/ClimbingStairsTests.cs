using Solutions.ClimbingStairs;

namespace Tests.ClimbingStairsTests;

public class ClimbingStairsTests
{
    private ClimbingStairs climbingStairs;

    [SetUp]
    public void Setup() => climbingStairs = new();

    [Test]
    public void TestExampleOne()
    {
        int n = 2;

        int expected = 2;
        int actual = climbingStairs.ClimbStairs(n);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int n = 3;

        int expected = 3;
        int actual = climbingStairs.ClimbStairs(n);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

