import numpy as np
import pandas as pd
import matplotlib.pyplot as plt
from scipy.optimize import curve_fit
import csv
import math
import glob

def main():
    k = 0
    shell_sorts = {"Shell Sort(prattGaps)", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)"}
    nlogn_sorts = {"Quick Sort", "Merge Sort", "Heap Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", "Shell Sort(prattGaps)"}
    types = {"Best", "Worst", "Almost", "Average"}
    # sort_names = {"Selection Sort", "Insertion Sort", "Quick Sort", "Bubble Sort", 
    #               "Merge Sort", "Shell Sort(shellGaps)", "Shell Sort(hibbardGaps)", 
    #               "Shell Sort(prattGaps)", "Heap Sort"}
    sort_names = {"Heap Sort"}
    for type in types:
        for sort_name in sort_names:
            x_values = []
            y_values = []
            try:
                # Constructing file name based on sort name and type
                filename = f'csvData/{sort_name}{type}Data.csv'
                with open(filename, 'r') as file:
                    reader = csv.DictReader(file)

                    # Read each row in the CSV file
                    for row in reader:
                        # print(row)  # Output each row as a dictionary
                        y_values.append(float(row['execTimes']))  # Add execTimes to list
                        x_values.append(float(row['sizes']))       # Add sizes to list

                # Check if we have data to process
                if not x_values or not y_values:
                    print(f"No data found in {filename}. Skipping.")
                    continue

                # Convert lists to NumPy arrays
                y = np.array(y_values)
                x = np.array(x_values)
                
                # Fit model based on sorting algorithm type
                if sort_name == "Quick Sort" and type == "Worst":
                    param, _ = curve_fit(n2, x, y)
                    ans = n2(x, param[0], param[1])
                elif sort_name in nlogn_sorts:
                    if sort_name in shell_sorts:
                        if sort_name == "Shell Sort(shellGaps)":
                            if type == "Best":
                                param, _ = curve_fit(nLogN, x, y)
                                ans = nLogN(x, param[0], param[1])
                            else:
                                param, _ = curve_fit(n2, x, y)
                                ans = n2(x, param[0], param[1])
                        elif sort_name == "Shell Sort(hibbardGaps)":
                            if type == "Best":
                                param, _ = curve_fit(nLogN, x, y)
                                ans = nLogN(x, param[0], param[1])
                            else:
                                param, _ = curve_fit(n2, x, y)
                                ans = n2(x, param[0], param[1])
                        #pratt gaps
                        else:
                            param, _ = curve_fit(nLog2N, x, y)
                            ans = nLog2N(x, param[0], param[1])

                        param, _ = curve_fit(n2, x, y)
                        ans = n2(x, param[0], param[1])
                    else:
                        param, _ = curve_fit(nLogN, x, y)
                        ans = nLogN(x, param[0], param[1])
                elif sort_name == "Insertion Sort" and type == "Best":
                        param, _ = curve_fit(const, x, y)
                        ans = const(x, param[0], param[1])
                else:
                    param, _ = curve_fit(n2, x, y)
                    ans = n2(x, param[0], param[1])

                # Plotting results
                plt.scatter(x, y, label='Данные', color='blue')
                plt.plot(x, ans, '--', label='Подогнанная модель', color='red')
                plt.xlabel('Размеры (sizes)')
                plt.ylabel('Время выполнения (execTimes)')
                plt.title(f"Регрессия {sort_name} {type}")
                plt.legend()
                plt.grid()
                plt.show()
                x_values = []
                y_values = []
                x = []
                y = []
                k+=1
                file.close()
            except FileNotFoundError:
                print(f"File {filename} not found. Skipping.")
            except Exception as e:
                print(f"An error occurred while processing {filename}: {e}")

    print(k)

def zero(x, a, b):
    return a*0*x+b

def n32(x, a, b):
    return a *(x**(3/2)) + b

def n2(x, a, b):
    return a * (x ** 2) + b

def f(x, a, b):
    x = np.asarray(x)
    value = x * b
    if np.any(value <= 0):
        raise ValueError("Input to log must be positive.")
    return a * (x * np.log(value) / np.log(2))

def nLogN(x, a, b):
    return a * (x * np.log(x)) + b
def nLog2N(x, a, b):
    return a * (x * (np.log(x))**2) + b
def const(x, a, b):
    return a * x + b 

def plotN2():
    n = np.arange(2, 500000) 
    # Предположим, что у вас есть соответствующие значения времени выполнения (например, O(n^2))
    # Для примера мы можем использовать квадратичную зависимость
    # exec_times = n * (np.log(n)/np.log(2))**2  # Это просто для примера; замените на ваши реальные данные
    exec_times = n
    # Создаем график
    # plt.figure(figsize=(10, 6))
    plt.plot(n, exec_times, '-', label='Время выполнения (O(n))', color='blue')

    # Настраиваем график
    plt.title('График времени выполнения в зависимости от размера входных данных')
    plt.xlabel('Размер входных данных (sizes)')
    plt.ylabel('Время выполнения')
    plt.grid()
    plt.legend()

    # Показываем график
    plt.show()



if __name__ == "__main__":
    main()
    # plotN2()
    # plot()
