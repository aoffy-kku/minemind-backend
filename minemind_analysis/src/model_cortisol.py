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
    df['type'] = pd.to_datetime(df['TIMESTAMP']).dt.hour > 18

    # preprocessing
    df['type'] = df['type'].astype('category')
    df['cortisol'] = df['cortisol'].astype('float')


    # ===== select feature ===== #
    X = df[['type', 'cortisol']]


    # ===== load model and predict ===== #
    clf = joblib.load('./minemind_analysis/model/model_cortisolV1_dt.joblib')
    y_pred = int(clf.predict(X)[0])
    score = max(np.round(clf.predict_proba(X) * 100, 2)[0])

    prediction = dict({'id': id, 
			'class': y_pred, 
			'score': score})

    return prediction


