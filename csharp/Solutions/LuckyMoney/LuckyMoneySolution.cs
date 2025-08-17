/// <sumary>
/// Live coding assignment
///
/// In Chinese culture, it is common during celebrations to give 'red envelopes' containing some money.
/// Most often, the adult generations give to the younger generations.
///
/// You want to build an app to help grandparents divide their giving budget among their grandchildren.
/// Write a program that calculates the number of 'lucky' gifts (equal to 8) based on the budget, money,
/// and the number of grandchildren, giftees.
///
/// Many rules, mixing tradition and superstition, govern this gift:
///     - Gifts must not contain the amount 4, as this sounds like 'death'.
///     - It is suspicious to give an amount of 8, as this sounds like 'fortune'.
///     - It would be wrong to give nothing to any of granchildren.
///
/// So you algorithm should return the number of gifts wqual to 8, following these rules:
///     - Spend the entire budget (unless there is enough budget to to give 8 to everyone).
///     - Do not give any 4 (by tradition, the budget will never be 4).
///     - Do not give any 0 (unless there is not enough money for all giftees).
///     - Give a maximum amount of 8 once the above rules are met.
///
/// Classical budget:
///     - money = 11
///     - giftees = 2
///
///     One lucky grandchild could get 8, the other could get 3. So, the method must return 1.
///
/// Perfect budget:
///     - money = 32
///     - giftees = 4
///
///     It is possible to give 8 to all the grandchildren. So, we must return 4.
///
/// Budget with a four:
///     - money = 12
///     - giftees = 2
///
///     If the first grandchild gets 8, then the secont would get 4, which is forbbiden. Whatever the
///     choice (7 to the first and 5 to the second, or 6 to both), nobody will get an amount of 8. So
///     we must return 0.
///
/// Constraints:
///     - 0 <= money <= 100
///     - 1 <= giftees <= 10
///     - Available RAM: 512 MB
///     - Timeout: 2 seconds
namespace Solutions.LuckyMoney;

public class LuckyMoneySolution
{
    public int LuckyMoney(int money, int giftees)
    {
        return 0;
    }
}