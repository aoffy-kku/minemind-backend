import numpy as np
import pandas as pd
from src.data_loader import load_sleep, load_stress, load_step
from src import dataprep
from src.feature_extract import stat_feature_extraction, sleep_quality_score, sleep_detail
from src.feature_extract import sax_bop_prop
import joblib


def model(id, clm, cortisol, timestamp, birthDate, start, end):
    """
    Classification Model using smart watch data
    id : GROUP_USER_ID
    start : first date and time which data will be used in model
    end : last date and time which data will be used in model
    """

    # ===== load data zensorium ===== #
    BIRTH_DATE = pd.to_datetime(birthDate)
    GROUP_USER_ID = str(clm)
    START_TIME = str(start)
    END_TIME = str(end)

    auth_token = "93aafa0e6cd844fd8332ae4027650bb8f08b5b2b4ab7958c5c57a2b5f84b7d84"
    head = {"Authorization": "Bearer "+auth_token}
    body = {
        "groupUserIds": [GROUP_USER_ID],
        "fromDateTime": START_TIME,
        "toDateTime": END_TIME
    }
    based_url = "https://info.minemind.net/zensoriumAPI/"

    # load data
    sleep = load_sleep(based_url, head, body)
    stress = load_stress(based_url, head, body)
    step = load_step(based_url, head, body)
    # tracking data
    tracking = dataprep.tracking_prep(stress, step, sleep)

    # user information
    # based_path = 'D:/MineMindScoreDemo/data/'
    # file1 = pd.ExcelFile(based_path + 'User.xlsx')
    # user = file1.parse('User_Info')
    # user = user[['GROUP_USER_ID', 'Birth_Date']]
    # user = user[user['GROUP_USER_ID'] == GROUP_USER_ID]
    user = pd.DataFrame({
        'GROUP_USER_ID': [GROUP_USER_ID], 
        'Birth_Date' :[BIRTH_DATE]
    })

    # ===== feature extraction ===== #
    var = ['HR(bpm)', 'PSV', 'PSF(mmHg)']
    stat = stat_feature_extraction(tracking, sleep, var=var, IDCol='Group_User_ID')
    sleep_score = sleep_quality_score(tracking, user, IDCol='Group_User_ID')
    sleepdetail = sleep_detail(sleep, IDCol='Group_User_ID')
    bop_prop = sax_bop_prop(tracking, var='HR(bpm)', IDCol='Group_User_ID')

    # ===== feature selection ===== #
    # create df model and store features
    df_zensorium = pd.concat([stat, sleep_score, sleepdetail, bop_prop], axis=1)

    features_zensorium = ['ad_kurt_HR(bpm)', 'ad_sd_HR(bpm)', 'ad_skew_HR(bpm)', 'mean_REM', 'ns_skew_HR(bpm)',
                         'sd_sleep_session', 'skew_sleep_session', 'start_day(%)', 'REM_20_25(%)', 'REM_Gt25(%)',
                         'REM_Lt20(%)', 'SS58(%)', 'SSLT3(%)', 'bca', 'cba', 'abc', 'cca', 'acc']
    # df_zensorium = df_zensorium[features_zensorium]



    # ===== load data cortisol ===== #
    df_cortisol = pd.DataFrame({
                'GROUP_USER_ID': [clm],
                'TIMESTAMP': [timestamp],
				'cortisol': [cortisol]
    })

    # add type column
    df_cortisol['type'] = pd.to_datetime(df_cortisol['TIMESTAMP'], errors='coerce', format='%d/%m/%Y %H:%M').dt.hour > 18

    # preprocessing
    df_cortisol['type'] = df_cortisol['type'].astype('category')
    df_cortisol['cortisol'] = df_cortisol['cortisol'].astype('float')

    features_cortisol = ['type', 'cortisol']


    # ===== merge data ===== #
    df_all = pd.merge(df_cortisol, df_zensorium, how='left', left_on='GROUP_USER_ID', right_index=True)


    # ===== select features ===== #
    X = df_all[features_zensorium + features_cortisol]


    # ===== load model and predict ===== #
    clf = joblib.load('./minemind_analysis/model/model_watchcortisolV1_rf.joblib')
    # y_pred = int(clf.predict(X)[0])
    # score = max(np.round(clf.predict_proba(X) * 100, 2)[0])

    score = np.round(clf.predict_proba(df) * 100, 2)[0][1]

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

