import numpy as np
from scipy.optimize import curve_fit
 
from matplotlib import pyplot as plt
import pandas as pd

import csv

with open('Merge SortData.csv', 'r') as file:
    reader = csv.DictReader(file)
    # y = row['execTimes']
    # x = row['sizes']
    for row in reader:
        print(row)  # Выводит каждую строку как словарь
        # y.append(row['execTimes'])   # Доступ к значению по имени столбца
        # x.append(row['sizes'])
        y = row['execTimes']
        x = row['sizes']
# x = np.linspace(1000, 300000, num = 40)

# y = x**2 + np.random.normal(size = 40)

# df = pd.read_csv('Bubble SortData.csv')
# x = df['sizes'] 
# y = df['execTimes'] 

def test(x, a, b):
    return a * (x*b)**2

param, param_cov = curve_fit(test, x, y)

print("Sine function coefficients:")
print(param)
print("Covariance of coefficients:")
print(param_cov)

ans = (param[0]*((param[1]*x)**2))

plt.plot(x, y, 'o', color ='red', label ="data")
plt.plot(x, ans, '--', color ='blue', label ="optimized data")
plt.legend()
plt.show()