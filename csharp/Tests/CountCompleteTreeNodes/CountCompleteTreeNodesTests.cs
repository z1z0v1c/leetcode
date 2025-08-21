using Solutions.CommonClases;
using Solutions.CountCompleteTreeNodes;

namespace Tests.CountCompleteTreeNodes;

public class CountCompleteTreeNodesTests
{
    private CountCompleteTreeNodesSolution solution;

    [SetUp]
    public void Setup() => solution = new();

    [Test]
    public void TestExampleOne()
    {
        // Arrange
        var root = new TreeNode(
            val: 1,
            left: new TreeNode(
                val: 2,
                left: new TreeNode(val: 4),
                right: new TreeNode(val: 5)
            ),
            right: new TreeNode(
                val: 3,
                left: new TreeNode(val: 6),
                right: null
            )
        );
        var expected = 6;
        
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
        var expected = 0;
        
        // Act
        var actual = solution.CountNodes(root!);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
    
    [Test]
    public void TestExampleThree()
    {
        // Arrange
        var root = new TreeNode(val: 1);
        var expected = 1;
        
        // Act
        var actual = solution.CountNodes(root);
        
        // Assert
        Assert.That(actual, Is.EqualTo(expected));
    }
}