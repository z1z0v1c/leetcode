package org.example;

import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.DisplayName;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class StringToIntegerTest {
    private StringToInteger stringToInteger;

   @BeforeEach
    void setUp() {
        stringToInteger = new StringToInteger();
    }

    @AfterEach
    void tearDown() {
    }

    @Test
    @DisplayName("Empty string")
    void testMyAtoiEmptyString() {
        assertEquals(0, stringToInteger.myAtoi(""),
                "Empty string should return 0");
    }
}