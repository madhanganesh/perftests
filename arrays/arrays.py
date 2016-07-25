import time

sum = 0
start = time.time()
for e in xrange(30):
    sum = 0

    x = []
    for i in xrange(1000000):
        x.append(i)

    y = []
    for i in xrange(1000000 - 1):
        y.append(x[i] + x[i+1])

    for i in xrange(0, 1000000,  100):
        sum +=  y[i]

end = time.time()
print('Elapsed ' + str(end-start))
print sum
