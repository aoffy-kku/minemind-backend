import numpy as np
import pandas as pd
import joblib


def model(id, st5):
    """
    Classification Model using st5 data
    id: GROUP_USER_ID
    """
    # ===== load data ===== #
    X = np.array([[st5]], dtype=float)

    # ===== load model and predict ===== #
    clf = joblib.load('./minemind_analysis/model/model_st5V1_lr.joblib')
    y_pred = int(clf.predict(X)[0])
    score = max(np.round(clf.predict_proba(X) * 100, 2)[0])

    prediction = dict({'id': id, 
			'class': y_pred, 
			'score': score})

    return prediction

