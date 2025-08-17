/// <sumary>
///
/// You work for a shop that wishes to give a discount of discount% to the most expensive item purchased
/// by a given customer during the sales period. Only one product can benefit from the discount.
///
/// You are tasked by the shop owner to implement the method CalculateTotalPrice(prices, discount) which
/// takes the list of prices of the products purchased by the customer end the percentage 'discount' as
/// parameters and returns the total purchase price as an integer (rounded down if the total is a float
/// number.
///
/// Constraints:
///     - 0 <= discount <= 100
///     - 0 < price of a product < 100000
///     - 0 < number of products < 100
/// </sumary>
namespace Solutions.CalculateTotalPrice;

public class CalculateTotalPriceSolution
{
    public int CalculateTotalPrice(int[] prices, int discount)
    {
        var mostExpensive = 0;
        var totalAmount = 0;

        foreach (var price in prices)
        {
            totalAmount += price;

            if (price > mostExpensive)
            {
                mostExpensive = price;
            }
        }

        var discountAmount = (mostExpensive * discount) / 100;

        return totalAmount - discountAmount;
    }
}