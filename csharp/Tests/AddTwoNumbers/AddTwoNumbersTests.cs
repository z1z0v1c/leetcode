using Solutions.AddTwoNumbers;
using Solutions.CommonClases;

namespace Tests.AddTwoNumbers;

public class AddTwoNumbersTests
{
    private AddTwoNumbersSolution addTwoNumbers;

    [SetUp]
    public void Setup() => addTwoNumbers = new();

    [Test]
    public void TestExampleOne()
    {
        ListNode l1 = new(2, new(4, new(3)));
        ListNode l2 = new(5, new(6, new(4)));

        ListNode expected = new(7, new(0, new(8)));
        ListNode actual = addTwoNumbers.AddTwoNumbers(l1, l2);

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
        ListNode l1 = new(0);
        ListNode l2 = new(0);

        ListNode expected = new(0);
        ListNode actual = addTwoNumbers.AddTwoNumbers(l1, l2);

        while (expected != null)
        {
            Assert.That(expected.val, Is.EqualTo(actual.val));

            expected = expected.next;
            actual = actual.next;
        }
    }

    [Test]
    public void TestExampleThree()
    {
        ListNode l1 = new(9, new(9, new(9, new(9, new(9, new(9, new(9)))))));
        ListNode l2 = new(9, new(9, new(9, new(9))));

        ListNode expected = new(8, new(9, new(9, new(9, new(0, new(0, new(0, new(1))))))));
        ListNode actual = addTwoNumbers.AddTwoNumbers(l1, l2);

        while (expected != null)
        {
            Assert.That(expected.val, Is.EqualTo(actual.val));

            expected = expected.next;
            actual = actual.next;
        }
    }
}

