using Solutions.ClimbingStairs;

namespace Tests.ClimbingStairs;

public class ClimbingStairsTests
{
    private Solutions.ClimbingStairs.ClimbingStairsSolution climbingStairsSolution;

    [SetUp]
    public void Setup() => climbingStairsSolution = new();

    [Test]
    public void TestExampleOne()
    {
        int n = 2;

        int expected = 2;
        int actual = climbingStairsSolution.ClimbStairs(n);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int n = 3;

        int expected = 3;
        int actual = climbingStairsSolution.ClimbStairs(n);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

