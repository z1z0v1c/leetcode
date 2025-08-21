using Solutions.CommonClasses;
using Solutions.RemoveDuplicatesFromList;

namespace Tests.RemoveDuplicatesFromList;

public class RemoveDuplicatesFromListTests
{
    private RemoveDuplicatesFromListSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        ListNode head = new(1, new(1, new(2)));

        ListNode? expected = new(1, new(2));
        ListNode? actual = solution.DeleteDuplicates(head);

        while (expected != null)
        {
            Assert.That(expected.Val, Is.EqualTo(actual?.Val));

            expected = expected?.Next;
            actual = actual?.Next;
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        ListNode head = new(1, new(1, new(2, new(3, new(3)))));

        ListNode? expected = new(1, new(2, new(3)));
        ListNode? actual = solution.DeleteDuplicates(head);

        while (expected != null)
        {
            Assert.That(expected.Val, Is.EqualTo(actual?.Val));

            expected = expected?.Next;
            actual = actual?.Next;
        }
    }
}

