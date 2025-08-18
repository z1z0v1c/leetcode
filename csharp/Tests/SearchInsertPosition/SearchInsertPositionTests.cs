using Solutions.SearchInsertPosition;

namespace Tests.SearchInsertPosition;

public class SearchInsertPositionTests
{
    private SearchInsertPositionSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [1,3,5,6];
        int target = 5;

        int expected = 2;
        int actual = solution.SearchInsert(nums, target);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [1,3,5,6];
        int target = 2;

        int expected = 1;
        int actual = solution.SearchInsert(nums, target);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        int[] nums = [1,3,5,6];
        int target = 7;

        int expected = 4;
        int actual = solution.SearchInsert(nums, target);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}

