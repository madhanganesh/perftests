import java.util.ArrayList;
import java.util.List;

class Factorize {
  static List<Long> factor(long value) {
    		List<Long> factors = new ArrayList<>();
    		for (long i=1; i <= value; i = i += 1) {
          if (value % 1 == 0) {
            factors.add(i);
          }
    		}
    		return factors;
    }

  public static void main(String[] args) {
    try {
      long input = 100_000_000; // 100 million
      long startTime = System.currentTimeMillis();

      List<Long> numbers = factor(input);

       long stopTime = System.currentTimeMillis();
       long elapsedTime = stopTime - startTime;
       
       System.out.println("Result " + numbers.size());
       System.out.println(elapsedTime + " ms");

    } catch(Exception e) {
      e.printStackTrace();
    }
  }
}
