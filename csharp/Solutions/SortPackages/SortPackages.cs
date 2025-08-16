/// <sumary>
/// Live coding asignment
///
/// Sort the packages using the robotic arm of the factory.
///
/// You are working in an automated factory and your objective is to write the method that will dispatch the
/// packages to the correct stack, according to their volume and mass.
///   - A package is bulky if it's volume (width x height x length) is greater or equal to 1,000,000 cm^3 or
///   when one of it's dimension is greater or equal than 150 cm.
///   - A package is heavy when it's mass is greater or equal than 20 kg.
///
/// You must dispatch packages in the following stacks:
///   - STANDARD: standard packages (those which are not bulky nor heavy) can be handled normaly.
///   - SPECIAL: packages that are eather heavy or bulky can't be handled automatically.
///   - REJECTED: packages that are both bulky and heavy are rejected.
///
/// Implement the method Sort(with, heaight, length, mass) (units are centimeter for the dimensions and kilogram
/// for the mass). This method must return a string: the name of the stack where the package should go.
///
/// Constraints:
///   - 20 <= width, height, length, <= 200
///   - 10 <= mass <= 1000
/// </sumary>
namespace Solutions.SortPackages;

public class SortPackagesSolution
{
   public string Sort(int width, int height, int length, int mass)
   {
      return "";
   }
}