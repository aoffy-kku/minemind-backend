import numpy as np
import pandas as pd
import joblib


def model(id, phq9):
    """
    Classification Model using st5 data
    id: GROUP_USER_ID
    """

    # ===== load data ===== #
    X = np.array([[phq9]], dtype=float)

    # ===== load model and predict ===== #
    clf = joblib.load('./minemind_analysis/model/model_phq9V1_lr.joblib')
    # y_pred = int(clf.predict(X)[0])
    # score = max(np.round(clf.predict_proba(X) * 100, 2)[0])
    score = np.round(clf.predict_proba(X) * 100, 2)[0][1]

    if score < 25.0:
        y_pred = 1
    elif (score >= 25.0) and (score < 50.0):
        y_pred = 2
    elif (score >= 50.0) and (score < 75.0):
        y_pred = 3
    else:
        y_pred = 4

    prediction = dict({'id': id, 
			'class': y_pred, 
			'score': score})

    return prediction

