using System;
using System.IO;
using System.Linq;
using System.Diagnostics;

namespace filter
{
    class Program
    {
        static void Main(string[] args)
        {
            try 
            {
                var fileName = "../data/names.txt";
                var line = File.ReadAllLines(fileName);
                var originalSize = line.Length;

                Stopwatch sw = new Stopwatch();
                sw.Start();

                line = line.Where(l => !l.Contains("A")).ToArray();

                sw.Stop();
                Console.WriteLine("Elapsed={0} ms",sw.Elapsed.Milliseconds);

                Console.WriteLine("Original Size : " + originalSize);
                Console.WriteLine("Filtered Size : " + line.Length);

            }
            catch(Exception e)
            {
                Console.WriteLine(e.ToString());
            }
        }
    }
}
