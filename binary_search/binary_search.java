package binary_search;

public class binary_search {
    public static void main(String[] args) {
        int[] arr = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};

        System.out.println(binarySearch(arr, 7));
    }

    public static int binarySearch(int[] arr, int number) {
        int low = 0;
        int high = arr.length - 1;
        int count = 0;

        while(low <= high) {
            count++;
            double middle = Math.floor((low + high) / 2);
            int guess = arr[(int) middle];

            if(guess == number) {
                System.out.println("Performed " + count + " actions");
                return (int) middle;
            }
            if(guess < number) {
                low = (int) middle + 1;
            }
            if(guess > number) {
                high = (int) middle - 1;
            }

        }

        System.out.println("Performed " + count + " actions");
        return 0;
    }
}
