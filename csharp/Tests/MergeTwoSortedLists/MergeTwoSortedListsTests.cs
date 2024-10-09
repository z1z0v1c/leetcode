using Solutions.CommonClases;
using Solutions.MergeTwoSortedLists;

namespace Tests.MergeTwoSortedListsTests;

public class MergeTwoSortedListsTests
{
    private MergeTwoSortedLists mergeTwoSortedLists;

    [SetUp]
    public void Setup() => mergeTwoSortedLists = new();

    [Test]
    public void TestExampleOne()
    {
        ListNode l1 = new(1, new(2, new(4)));
        ListNode l2 = new(1, new(3, new(4)));

        ListNode expected = new(1, new(1, new(2, new(3, new(4, new(4))))));
        ListNode actual = mergeTwoSortedLists.MergeTwoLists(l1, l2);

        while (expected != null)
        {
            Assert.That(expected.val, Is.EqualTo(actual.val));

            expected = expected.next;
            actual = actual.next;
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        Assert.That(mergeTwoSortedLists.MergeTwoLists(null, null), Is.EqualTo(null));
    }

    [Test]
    public void TestExampleThre()
    {
        ListNode l2 = new(0);

        ListNode expected = new(0);
        ListNode actual = mergeTwoSortedLists.MergeTwoLists(null, l2);

        Assert.That(expected.val, Is.EqualTo(actual.val));
    }
}
