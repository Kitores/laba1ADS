import numpy as np
import matplotlib.pyplot as plt
from sklearn.preprocessing import PolynomialFeatures
from sklearn.linear_model import LinearRegression
from sklearn.metrics import mean_squared_error, r2_score
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

# Преобразование входных данных в полиномиальные признаки
degree = 2  # Степень полинома
poly_features = PolynomialFeatures(degree=degree)
X_poly = poly_features.fit_transform(X)

# Обучение модели линейной регрессии на полиномиальных признаках
model = LinearRegression()
model.fit(X_poly, Y)

# Предсказание значений для графика
X_grid = np.arange(0, 10, 0.1).reshape(-1, 1)  # Создаем более плотную сетку для плавной линии
X_grid_poly = poly_features.transform(X_grid)
y_pred = model.predict(X_grid_poly)

# Визуализация результатов
plt.scatter(X, Y, color='blue', label='Данные')
plt.plot(X_grid, y_pred, color='red', label='Полиномиальная регрессия')
plt.title('Полиномиальная регрессия')
plt.xlabel('X')
plt.ylabel('Y')
plt.legend()
plt.grid()
plt.show()
