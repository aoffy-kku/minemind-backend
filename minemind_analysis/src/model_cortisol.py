import numpy as np
import pandas as pd
import joblib


def model(id, cortisol, timestamp):
    """
    Classification Model using cortisol data
    id: GROUP_USER_ID
    """

    # ===== load data ===== #
    df = pd.DataFrame({'TIMESTAMP' : [timestamp], 
			'cortisol' : [cortisol]})
    
    # add type column
    df['type'] = pd.to_datetime(df['TIMESTAMP'], errors='coerce', format='%d/%m/%Y %H:%M').dt.hour > 18

    # preprocessing
    df['type'] = df['type'].astype('category')
    df['cortisol'] = df['cortisol'].astype('float')


    # ===== select feature ===== #
    X = df[['type', 'cortisol']]


    # ===== load model and predict ===== #
    clf = joblib.load('./minemind_analysis/model/model_cortisolV1_dt.joblib')
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


