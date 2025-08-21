using Solutions.AddTwoNumbers;
using Solutions.CommonClasses;

namespace Tests.AddTwoNumbers;

public class AddTwoNumbersTests
{
    private AddTwoNumbersSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new AddTwoNumbersSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        ListNode l1 = new(2, new ListNode(4, new ListNode(3)));
        ListNode l2 = new(5, new ListNode(6, new ListNode(4)));
        ListNode? expected = new(7, new ListNode(0, new ListNode(8)));

        // Act
        var actual = solution.AddTwoNumbers(l1, l2);

        // Assert
        while (expected != null)
        {
            Assert.That(expected.Val, Is.EqualTo(actual!.Val));

            expected = expected.Next;
            actual = actual.Next;
        }
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        ListNode l1 = new();
        ListNode l2 = new();
        ListNode? expected = new();

        // Act
        var actual = solution.AddTwoNumbers(l1, l2);

        // Assert
        while (expected != null)
        {
            Assert.That(expected.Val, Is.EqualTo(actual!.Val));

            expected = expected.Next;
            actual = actual.Next;
        }
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        ListNode l1 = new(9,
            new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9, new ListNode(9)))))));
        ListNode l2 = new(9, new ListNode(9, new ListNode(9, new ListNode(9))));
        ListNode? expected = new(8,
            new ListNode(9,
                new ListNode(9, new ListNode(9, new ListNode(0, new ListNode(0, new ListNode(0, new ListNode(1))))))));

        // Act
        var actual = solution.AddTwoNumbers(l1, l2);

        // Assert
        while (expected != null)
        {
            Assert.That(expected.Val, Is.EqualTo(actual!.Val));

            expected = expected.Next;
            actual = actual.Next;
        }
    }
}