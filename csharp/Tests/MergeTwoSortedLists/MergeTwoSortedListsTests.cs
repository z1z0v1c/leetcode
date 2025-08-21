using Solutions.CommonClasses;
using Solutions.MergeTwoSortedLists;

namespace Tests.MergeTwoSortedLists;

public class MergeTwoSortedListsTests
{
    private MergeTwoSortedListsSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new MergeTwoSortedListsSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        ListNode l1 = new(1, new ListNode(2, new ListNode(4)));
        ListNode l2 = new(1, new ListNode(3, new ListNode(4)));
        ListNode? expected = new(1,
            new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(4))))));

        // Act
        var actual = solution.MergeTwoLists(l1, l2);

        // Assert
        while (expected != null)
        {
            Assert.That(expected.Val, Is.EqualTo(actual?.Val));

            expected = expected.Next;
            actual = actual?.Next;
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange-Act-Assert :)
        Assert.That(solution.MergeTwoLists(null, null), Is.EqualTo(null));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        ListNode l2 = new();
        ListNode? expected = new();

        // Act
        var actual = solution.MergeTwoLists(null, l2);

        // Assert
        Assert.That(expected.Val, Is.EqualTo(actual?.Val));
    }
}