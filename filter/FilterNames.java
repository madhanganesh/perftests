import java.util.List;
import java.util.ArrayList;
import java.util.Arrays;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.stream.Stream;

class FilterNames {
    public static void main(String[] args) {

        try {

            String fileName = "../data/names.txt";
            Stream<String> stream  = Files.lines(Paths.get(fileName));

            List<String> names = new ArrayList<>();
            stream.forEach(name -> names.add(name));


            long startTime = System.currentTimeMillis();

            names.removeIf(n -> n.contains("A"));

            long stopTime = System.currentTimeMillis();
            long elapsedTime = stopTime - startTime;
            System.out.println(elapsedTime + " ms");

            System.out.println(names.size());

        } catch (Exception e) {
            e.printStackTrace();
        }
    }
}
