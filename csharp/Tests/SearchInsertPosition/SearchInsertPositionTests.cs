using Solutions.SearchInsertPosition;

namespace Tests.SearchInsertPosition;

public class SearchInsertPositionTests
{
    private SearchInsertPositionSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new SearchInsertPositionSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [1, 3, 5, 6];
        const int target = 5;

        const int expected = 2;
        var actual = solution.SearchInsert(nums, target);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [1, 3, 5, 6];
        const int target = 2;

        const int expected = 1;
        var actual = solution.SearchInsert(nums, target);

        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        int[] nums = [1, 3, 5, 6];
        const int target = 7;

        const int expected = 4;
        var actual = solution.SearchInsert(nums, target);

        Assert.That(actual, Is.EqualTo(expected));
    }
}