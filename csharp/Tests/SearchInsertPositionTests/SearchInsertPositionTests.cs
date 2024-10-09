using Solutions.SearchInsertPosition;

namespace Tests.SearchInsertPositionTests;

public class SearchInsertPositionTests
{
    private SearchInsertPosition searchInsertPosition;

    [SetUp]
    public void Setup() => searchInsertPosition = new();

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [1,3,5,6];
        int target = 5;

        int expected = 2;
        int actual = searchInsertPosition.SearchInsert(nums, target);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [1,3,5,6];
        int target = 2;

        int expected = 1;
        int actual = searchInsertPosition.SearchInsert(nums, target);
        
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        int[] nums = [1,3,5,6];
        int target = 7;

        int expected = 4;
        int actual = searchInsertPosition.SearchInsert(nums, target);
        
        Assert.That(actual, Is.EqualTo(expected));
    }
}
