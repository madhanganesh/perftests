import time

with open('../data/names.txt') as f:
    lines = f.readlines()

start = time.time()
names = [name for name in lines if 'A' not in name]
end = time.time()
print(str(end - start) + ' sec')
print(len(names))
