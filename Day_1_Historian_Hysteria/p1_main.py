

with open('input.txt', 'r') as f:
    data = f.readlines()


left = []
right = []
for line in data:
    l,r = line.split()
    left.append(int(l))
    right.append(int(r))

left.sort()
right.sort()

answer = 0
for i in range(len(right)):
    offset = left[i] -  right[i]
    answer += abs(offset)

print(answer)