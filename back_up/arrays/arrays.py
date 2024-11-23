import time

sum = 0
start = time.time()
for e in range(30):
    sum = 0

    x = [0] * 1000000
    for i in range(1000000):
        #x.append(i)
        x[i] = i

    y = [0] * 1000000
    for i in range(1000000 - 1):
        #y.append(x[i] + x[i+1])
        y[i] = x[i] + x[i+1]

    for i in range(0, 1000000, 100):
        sum += y[i]

end = time.time()
print('Elapsed ' + str(end-start))
print(sum)
