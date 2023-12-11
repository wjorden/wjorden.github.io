---
slug: scikit-learn
title: Diabetes XGBoost testing
date: 2023-12-11
lastmod: ["lastmod", ":git", "date", "publishDate"]
tags: ["python", "statistics", "scikit-learn", "xgboost"]
---

## Steps

1.  Load the “diabetes.csv” dataset using Pandas ✓  
2.  Split the data into 80% for training and 20% for testing ✓
3.  Train an XG-Boost classifier model using SK-Learn Library ✓  
4.  Assess trained model performance  
5.  Plot the confusion matrix  
6.  Print the classification report

Before beginning, I am using RStudio so loading R’s reticulate library is first.

``` r
library(reticulate)
```

Now, import all of the required Python libraries.

``` python
import pandas as pd
import numpy as np
import seaborn as sns
import matplotlib.pyplot as plt
from sklearn.model_selection import train_test_split
from xgboost import XGBClassifier
```

Import the dataset.

``` python
df = pd.read_csv('dataset/diabetes.csv')
df.info()
```

    ## <class 'pandas.core.frame.DataFrame'>
    ## RangeIndex: 768 entries, 0 to 767
    ## Data columns (total 9 columns):
    ##  #   Column                    Non-Null Count  Dtype  
    ## ---  ------                    --------------  -----  
    ##  0   Pregnancies               768 non-null    int64  
    ##  1   Glucose                   768 non-null    int64  
    ##  2   BloodPressure             768 non-null    int64  
    ##  3   SkinThickness             768 non-null    int64  
    ##  4   Insulin                   768 non-null    int64  
    ##  5   BMI                       768 non-null    float64
    ##  6   DiabetesPedigreeFunction  768 non-null    float64
    ##  7   Age                       768 non-null    int64  
    ##  8   Outcome                   768 non-null    int64  
    ## dtypes: float64(2), int64(7)
    ## memory usage: 54.1 KB

Histogram plots to check distribution.

``` python
df.hist(bins=30, figsize=(20,20), color='b')
```

<img src="{{< blogdown/postref >}}index_files/figure-html/unnamed-chunk-4-1.png" width="1920" />
Check Correlation.

``` python
correlation = df.corr()
f, ax = plt.subplots(figsize=(20,20))
sns.heatmap(correlation, annot=True)
```

<img src="{{< blogdown/postref >}}index_files/figure-html/unnamed-chunk-5-3.png" width="1920" />
Split the data into testing and training.

``` python
y=df['Outcome']
X=df.drop(['Outcome'], axis=1)

X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=42)
print(f'X Train: {X_train.shape}, X_test: {X_test.shape}, y_train: {y_train.shape}, y_test: {y_test.shape}')
```

    ## X Train: (614, 8), X_test: (154, 8), y_train: (614,), y_test: (154,)

Train XGBoost classifier.

``` python
xgb_classifier = XGBClassifier(objective='binary:logistic', eval_metric='error', learning_rate=0.1, max_depth=5, n_estimators=10, use_label_encoder=False)
xgb_classifier.fit(X_train, y_train)
```

<style>#sk-container-id-1 {color: black;}#sk-container-id-1 pre{padding: 0;}#sk-container-id-1 div.sk-toggleable {background-color: white;}#sk-container-id-1 label.sk-toggleable__label {cursor: pointer;display: block;width: 100%;margin-bottom: 0;padding: 0.3em;box-sizing: border-box;text-align: center;}#sk-container-id-1 label.sk-toggleable__label-arrow:before {content: "▸";float: left;margin-right: 0.25em;color: #696969;}#sk-container-id-1 label.sk-toggleable__label-arrow:hover:before {color: black;}#sk-container-id-1 div.sk-estimator:hover label.sk-toggleable__label-arrow:before {color: black;}#sk-container-id-1 div.sk-toggleable__content {max-height: 0;max-width: 0;overflow: hidden;text-align: left;background-color: #f0f8ff;}#sk-container-id-1 div.sk-toggleable__content pre {margin: 0.2em;color: black;border-radius: 0.25em;background-color: #f0f8ff;}#sk-container-id-1 input.sk-toggleable__control:checked~div.sk-toggleable__content {max-height: 200px;max-width: 100%;overflow: auto;}#sk-container-id-1 input.sk-toggleable__control:checked~label.sk-toggleable__label-arrow:before {content: "▾";}#sk-container-id-1 div.sk-estimator input.sk-toggleable__control:checked~label.sk-toggleable__label {background-color: #d4ebff;}#sk-container-id-1 div.sk-label input.sk-toggleable__control:checked~label.sk-toggleable__label {background-color: #d4ebff;}#sk-container-id-1 input.sk-hidden--visually {border: 0;clip: rect(1px 1px 1px 1px);clip: rect(1px, 1px, 1px, 1px);height: 1px;margin: -1px;overflow: hidden;padding: 0;position: absolute;width: 1px;}#sk-container-id-1 div.sk-estimator {font-family: monospace;background-color: #f0f8ff;border: 1px dotted black;border-radius: 0.25em;box-sizing: border-box;margin-bottom: 0.5em;}#sk-container-id-1 div.sk-estimator:hover {background-color: #d4ebff;}#sk-container-id-1 div.sk-parallel-item::after {content: "";width: 100%;border-bottom: 1px solid gray;flex-grow: 1;}#sk-container-id-1 div.sk-label:hover label.sk-toggleable__label {background-color: #d4ebff;}#sk-container-id-1 div.sk-serial::before {content: "";position: absolute;border-left: 1px solid gray;box-sizing: border-box;top: 0;bottom: 0;left: 50%;z-index: 0;}#sk-container-id-1 div.sk-serial {display: flex;flex-direction: column;align-items: center;background-color: white;padding-right: 0.2em;padding-left: 0.2em;position: relative;}#sk-container-id-1 div.sk-item {position: relative;z-index: 1;}#sk-container-id-1 div.sk-parallel {display: flex;align-items: stretch;justify-content: center;background-color: white;position: relative;}#sk-container-id-1 div.sk-item::before, #sk-container-id-1 div.sk-parallel-item::before {content: "";position: absolute;border-left: 1px solid gray;box-sizing: border-box;top: 0;bottom: 0;left: 50%;z-index: -1;}#sk-container-id-1 div.sk-parallel-item {display: flex;flex-direction: column;z-index: 1;position: relative;background-color: white;}#sk-container-id-1 div.sk-parallel-item:first-child::after {align-self: flex-end;width: 50%;}#sk-container-id-1 div.sk-parallel-item:last-child::after {align-self: flex-start;width: 50%;}#sk-container-id-1 div.sk-parallel-item:only-child::after {width: 0;}#sk-container-id-1 div.sk-dashed-wrapped {border: 1px dashed gray;margin: 0 0.4em 0.5em 0.4em;box-sizing: border-box;padding-bottom: 0.4em;background-color: white;}#sk-container-id-1 div.sk-label label {font-family: monospace;font-weight: bold;display: inline-block;line-height: 1.2em;}#sk-container-id-1 div.sk-label-container {text-align: center;}#sk-container-id-1 div.sk-container {/* jupyter's `normalize.less` sets `[hidden] { display: none; }` but bootstrap.min.css set `[hidden] { display: none !important; }` so we also need the `!important` here to be able to override the default hidden behavior on the sphinx rendered scikit-learn.org. See: https://github.com/scikit-learn/scikit-learn/issues/21755 */display: inline-block !important;position: relative;}#sk-container-id-1 div.sk-text-repr-fallback {display: none;}</style><div id="sk-container-id-1" class="sk-top-container"><div class="sk-text-repr-fallback"><pre>XGBClassifier(base_score=None, booster=None, callbacks=None,
              colsample_bylevel=None, colsample_bynode=None,
              colsample_bytree=None, device=None, early_stopping_rounds=None,
              enable_categorical=False, eval_metric=&#x27;error&#x27;, feature_types=None,
              gamma=None, grow_policy=None, importance_type=None,
              interaction_constraints=None, learning_rate=0.1, max_bin=None,
              max_cat_threshold=None, max_cat_to_onehot=None,
              max_delta_step=None, max_depth=5, max_leaves=None,
              min_child_weight=None, missing=nan, monotone_constraints=None,
              multi_strategy=None, n_estimators=10, n_jobs=None,
              num_parallel_tree=None, random_state=None, ...)</pre><b>In a Jupyter environment, please rerun this cell to show the HTML representation or trust the notebook. <br />On GitHub, the HTML representation is unable to render, please try loading this page with nbviewer.org.</b></div><div class="sk-container" hidden><div class="sk-item"><div class="sk-estimator sk-toggleable"><input class="sk-toggleable__control sk-hidden--visually" id="sk-estimator-id-1" type="checkbox" checked><label for="sk-estimator-id-1" class="sk-toggleable__label sk-toggleable__label-arrow">XGBClassifier</label><div class="sk-toggleable__content"><pre>XGBClassifier(base_score=None, booster=None, callbacks=None,
              colsample_bylevel=None, colsample_bynode=None,
              colsample_bytree=None, device=None, early_stopping_rounds=None,
              enable_categorical=False, eval_metric=&#x27;error&#x27;, feature_types=None,
              gamma=None, grow_policy=None, importance_type=None,
              interaction_constraints=None, learning_rate=0.1, max_bin=None,
              max_cat_threshold=None, max_cat_to_onehot=None,
              max_delta_step=None, max_depth=5, max_leaves=None,
              min_child_weight=None, missing=nan, monotone_constraints=None,
              multi_strategy=None, n_estimators=10, n_jobs=None,
              num_parallel_tree=None, random_state=None, ...)</pre></div></div></div></div></div>

Check results.

``` python
result = xgb_classifier.score(X_test, y_test)
print(f'Accuracy: {result}')
```

    ## Accuracy: 0.7207792207792207

``` python

y_predict = xgb_classifier.predict(X_test)
```

Predictions on test data.

``` python
from sklearn.metrics import classification_report
print(classification_report(y_test, y_predict))
```

    ##               precision    recall  f1-score   support
    ## 
    ##            0       0.76      0.83      0.79        99
    ##            1       0.63      0.53      0.57        55
    ## 
    ##     accuracy                           0.72       154
    ##    macro avg       0.69      0.68      0.68       154
    ## weighted avg       0.71      0.72      0.71       154