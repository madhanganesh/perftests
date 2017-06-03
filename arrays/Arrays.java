class Arrays {
    public static void main(String[] args) {
        long sum = 0;
        long startTime = System.currentTimeMillis();

        for (int e=0; e<30; e++) {
            long[] x = new long[1000000];
            for (int i=0; i<1000000; i++) {
                x[i] = i;
            }

            long[] y = new long[1000000-1];
            for (int i=0; i<1000000-1; i++) {
                y[i] = x[i] + x[i+1];
            }

            sum = 0;
            for (int i=0; i<1000000; i+=100) {
                sum += y[i];
            }
        }
        long stopTime = System.currentTimeMillis();
        long elapsedTime = stopTime - startTime;
        System.out.println(elapsedTime + " ms");
        System.out.println(sum);
    }
}