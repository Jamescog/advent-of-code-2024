

with open('input.txt', 'r') as f:
    data = f.readlines()


left = []
right = []
for line in data:
    l,r = line.split()
    left.append(int(l))
    right.append(int(r))

simlariy_score = 0
for i in left:

    c = right.count(i)
    if c != 0:
        simlariy_score += i * c

print(simlariy_score)
