package buyandsellstock;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class BuyAndSellStockTest {
    private BuyAndSellStock buyAndSellStock;

    @BeforeEach
    void setUp() {
        buyAndSellStock  = new BuyAndSellStock();
    }

    @Test
    @DisplayName("Example one")
    void testExampleOne() {
        int[] prices = {7,1,5,3,6,4};

        int expected = 5;
        int actual = buyAndSellStock.maxProfit(prices);

        assertEquals(expected, actual, "Should return " + expected);
    }

    @Test
    @DisplayName("Example two")
    void testExampleTwo() {
        int[] prices = {7,6,4,3,1};

        int expected = 0;
        int actual = buyAndSellStock.maxProfit(prices);

        assertEquals(expected, actual, "Should return " + expected);
    }
}