import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
import math



def plotting_seperately(df):
    fig, ((ax1,ax2),(ax3,ax4)) = plt.subplots(2,2,figsize=(12, 10))
    
    fig.suptitle('comapre sorting algorithm')
    ax1.plot(df.index, df["bubble"] / (1000))

    ax1.set(xlabel='item in list', ylabel='times taken micro seconds',
        title='Bubble sort plotting')
    ax1.grid()
    ax2.plot(df.index, df["merge"] / (1000))

    ax2.set(xlabel='item in list', ylabel='times taken micro seconds',
        title='Merge sort plotting')
    ax2.grid()
    ax3.plot(df.index, df["quick"] / (1000))

    ax3.set(xlabel='item in list', ylabel='times taken micro seconds',
        title='Quick sort plotting')
    ax3.grid()
    ax4.plot(df.index, df["sort"] / (1000))

    ax4.set(xlabel='item in list', ylabel='times taken micro seconds',
        title='sort.Ints plotting')
    ax4.grid()
    
    
    fig.savefig("comparison-between-sort.png")
    plt.show()

def plotting_combined(df):
    plt.figure(figsize=(10,10))    
    plt.xlabel("size of random array")
    plt.ylabel("time taken in milisecond")
    plt.title("Sorting algorithm comparison")
    plt.plot(df.index, df["bubble"] / (1000),label="bubble")
    plt.plot(df.index, df["merge"] / (1000),label="merge")
    plt.plot(df.index, df["quick"] / (1000),label="quick")
    plt.plot(df.index, df["sort"] / (1000),label="sort")
    plt.legend()
    plt.savefig("combined-comparison-between-sort.png")
    plt.show()

def fix_outlier(df_series):
    Q1 = df_series.quantile(0.25)
    Q3 = df_series.quantile(0.75)
    replacement_value = df_series.quantile(0.5)
    IQR = Q3 - Q1
    lower_bound = Q1 - 1.5 * IQR
    upper_bound = Q3 + 1.5 * IQR
    outliers_mask = (df_series < lower_bound) | (df_series > upper_bound)
    return np.where(outliers_mask, (df_series.shift(1) + df_series.shift(-1)) / 2,
                              df_series)
    # return np.where((df_series < lower_bound) | (df_series > upper_bound), replacement_value, df_series)

def plotting_theory(df):
    plt.figure(figsize=(10,10))    
    plt.xlabel("size of random array")
    plt.ylabel("operation did")
    plt.title("Sorting algorithm theory comparison")
    plt.plot(df.index, df["log_n"] ,label="log_n")
    plt.plot(df.index, df["n_squared"],label="n_squared")
    plt.legend()
    plt.savefig("theory-comparison-sort.png") 
    plt.show()


# Define the updated cal function
def cal_log_n(n):
    return n * math.log(n)

def cal_squared_n(n):
    return n ** 2




if __name__ == "__main__":
    df = pd.read_json("../sorting/sort-test.json")

    df["merge"] = fix_outlier(df["merge"])
    df["bubble"] = fix_outlier(df["bubble"])
    df["sort"] = fix_outlier(df["sort"])
    df["quick"] = fix_outlier(df["quick"])

    plotting_seperately(df)
    plotting_combined(df)


    # Use a list comprehension to apply the updated cal function 300 times
    resultsLogN = [cal_log_n(i) for i in range(1, 301)]  # Avoid log(0)
    resultsSquaredN = [cal_squared_n(i) for i in range(1, 301)]  # Avoid log(0)

    # Create a DataFrame from the results list
    df = pd.DataFrame({'log_n': resultsLogN, "n_squared": resultsSquaredN})

    # Print the DataFrame
    plotting_theory(df)

