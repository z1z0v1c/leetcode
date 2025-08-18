using Solutions.SearchInRotatedArrayII;

namespace Tests.SearchInRotatedArrayII;

public class SearchInRotatedArrayIiTests
{
    private SearchInRotatedArrayIISolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        int[] nums = [2,5,6,0,0,1,2];
        int target = 0;
        
        Assert.That(solution.Search(nums, target), Is.True);
    }

    [Test]
    public void TestExampleTwo()
    {
        int[] nums = [2,5,6,0,0,1,2];
        int target = 3;
        
        Assert.That(solution.Search(nums, target), Is.False);
    }
}

