using Solutions.CommonClasses;
using Solutions.CountCompleteTreeNodes;

namespace Tests.CountCompleteTreeNodes;

public class CountCompleteTreeNodesTests
{
    private CountCompleteTreeNodesSolution solution;

    [SetUp]
    public void Setup()
    {
        solution = new CountCompleteTreeNodesSolution();
    }

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var root = new TreeNode(
            1,
            new TreeNode(
                2,
                new TreeNode(4),
                new TreeNode(5)
            ),
            new TreeNode(
                3,
                new TreeNode(6),
                null
            )
        );
        const int expected = 6;

        // Act
        var actual = solution.CountNodes(root);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleTwo()
    {
        // Arrange
        TreeNode? root = null;
        const int expected = 0;

        // Act
        var actual = solution.CountNodes(root!);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }

    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var root = new TreeNode(1);
        const int expected = 1;

        // Act
        var actual = solution.CountNodes(root);

        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}