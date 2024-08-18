package stringtointeger;

public class StringToInteger {
    public static void main(String[] args) {
        var sti = new StringToInteger();

        System.out.println(sti.myAtoi(""));
        System.out.println(sti.myAtoi("   "));
        System.out.println(sti.myAtoi("-"));
    }

    public int myAtoi(String s) {
        char[] str = s.toCharArray();

        for (char c : str) {
            if (c == ' ') {
                continue;
            }
            if (c == '-') {
                return -1;
            }
            if (c == '+') {
                return 0;
            }
        }

        return 0;
    }
}

