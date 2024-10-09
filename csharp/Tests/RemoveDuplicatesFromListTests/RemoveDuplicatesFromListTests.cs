using Solutions.CommonClases;
using Solutions.RemoveDuplicatesFromList;

namespace Tests.RemoveDuplicatesFromListTests;

public class RemoveDuplicatesFromListTests
{
    private RemoveDuplicatesFromList removeDuplicates;

    [SetUp]
    public void Setup() => removeDuplicates = new();

    [Test]
    public void TestExampleOne()
    {
        ListNode head = new(1, new(1, new(2)));

        ListNode? expected = new(1, new(2));
        ListNode? actual = removeDuplicates.DeleteDuplicates(head);

        while (expected != null)
        {
            Assert.That(expected.val, Is.EqualTo(actual?.val));

            expected = expected?.next;
            actual = actual?.next;
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        ListNode head = new(1, new(1, new(2, new(3, new(3)))));

        ListNode? expected = new(1, new(2, new(3)));
        ListNode? actual = removeDuplicates.DeleteDuplicates(head);

        while (expected != null)
        {
            Assert.That(expected.val, Is.EqualTo(actual?.val));

            expected = expected?.next;
            actual = actual?.next;
        }
    }
}
