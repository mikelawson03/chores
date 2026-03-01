import time

week = int(time.strftime("%U"))
print(type(week))

if week % 2 == 0:
    print("Even")

else:
    print("Odd")