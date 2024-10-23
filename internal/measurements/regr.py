import numpy as np
import matplotlib.pyplot as plt
from sklearn.preprocessing import PolynomialFeatures
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, r2_score
import operator
# import csv
# import operator
import pandas as pd

np.random.seed(0)

# Create a function to generate a polynomial dataset
def create_polynomial_data(n):
    X = np.linspace(-3, 3, n)
    Y = X ** 3 + 2 * X**2 - 3 * X + 2
    Y += np.random.normal(0, 1, n) # add some noise
    return X, Y

def get_polynomial_data():
    df = pd.read_csv('Bubble SortData.csv')

    X = df['sizes'].values.reshape(-1, 1)  # Преобразуем в двумерный массив
    Y = df['execTimes'].values.reshape(-1, 1)
    return X, Y


# X, Y = create_polynomial_data(100)
X, Y = get_polynomial_data()




# Предсказание значений для графика
X_grid = np.arange(0, 10, 0.1).reshape(-1, 1)  # Создаем более плотную сетку для плавной линии

# Преобразование входных данных в полиномиальные признаки
degree = 2  # Степень полинома
poly_features = PolynomialFeatures(degree=degree)
X_poly = poly_features.fit_transform(X)

# Create polynomial features
polynomial_features = PolynomialFeatures(degree=3)
X_poly = polynomial_features.fit_transform(X)
X_grid_poly = poly_features.transform(X_grid)
# Обучение модели линейной регрессии на полиномиальных признаках
model = LinearRegression()
model.fit(X_poly, Y)

y_poly_pred = model.predict(X_grid_poly)

# Sort the values of x before line plot
sort_axis = operator.itemgetter(0)
sorted_zip = sorted(zip(X, y_poly_pred), key=sort_axis)
X, y_poly_pred = zip(*sorted_zip)

plt.scatter(X, Y, s=10)
plt.plot(X, y_poly_pred, color='r')
plt.show()

# # Визуализация результатов
# plt.scatter(X, Y, color='blue', label='Данные')
# plt.plot(X_grid, y_pred, color='red', label='Полиномиальная регрессия')
# plt.title('Полиномиальная регрессия')
# plt.xlabel('X')
# plt.ylabel('Y')
# plt.legend()
# plt.grid()
# plt.show()