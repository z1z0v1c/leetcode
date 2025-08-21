using Solutions.CommonClasses;
using Solutions.RemoveDuplicatesFromList;

namespace Tests.RemoveDuplicatesFromList;

public class RemoveDuplicatesFromListTests
{
    private RemoveDuplicatesFromListSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new RemoveDuplicatesFromListSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        ListNode head = new(1, new ListNode(1, new ListNode(2)));
        ListNode? expected = new(1, new ListNode(2));

        // Act
        var actual = solution.DeleteDuplicates(head);

        // Assert
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
        // Arrange
        ListNode head = new(1, new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(3)))));
        ListNode? expected = new(1, new ListNode(2, new ListNode(3)));

        // Act
        var actual = solution.DeleteDuplicates(head);

        // Assert
        while (expected != null)
        {
            Assert.That(expected.Val, Is.EqualTo(actual?.Val));

            expected = expected?.Next;
            actual = actual?.Next;
        }
    }
}