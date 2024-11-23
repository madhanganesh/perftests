using System;
using System.Diagnostics;
using System.Collections.Generic;

namespace arrays
{
    class Test
    {
        public static void Main(String[] args)
        {
            Stopwatch sw = new Stopwatch();
            sw.Start();
            long sum = 0;

            for (int e=0; e<30; e++)
            {
                var x = new long[1000000];
                for (var i=0; i<1000000; i++)
                {
                    x[i] = i;
                }

                var y = new long[1000000-1];
                for (var i=0; i<1000000-1; i++)
                {
                    y[i] = x[i] + x[i+1];
                }

                sum = 0;
                for (var i=0; i<1000000; i+=100)
                {
                    sum += y[i];
                }
            }
            sw.Stop();

            Console.WriteLine("Elapsed={0} ms",sw.Elapsed.Milliseconds);
            Console.WriteLine(sum);
        }
    }
}