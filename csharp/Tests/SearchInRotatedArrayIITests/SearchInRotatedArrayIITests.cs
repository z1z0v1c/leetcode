using Solutions.SearchInRotatedArrayII;

namespace Tests.SearchInRotatedArrayIITests;

public class SearchInRotatedArrayIITests
{
    private SearchInRotatedArrayII searchInRotatedArrayII;

    [SetUp]
    public void Setup() => searchInRotatedArrayII = new();

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [2,5,6,0,0,1,2];
        int target = 0;
        
        Assert.That(searchInRotatedArrayII.Search(nums, target), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [2,5,6,0,0,1,2];
        int target = 3;
        
        Assert.That(searchInRotatedArrayII.Search(nums, target), Is.False);
    }
}

