import requests
import pandas as pd
from datetime import datetime


def load_sleep(based_url, head, body):
    # request from zensoriumAPI
    path = "measurement/getSleepData"
    response = requests.post(based_url+path, headers=head, json=body)
    data = response.json()
    # data preprocessing
    sleep = pd.DataFrame(data[0]['data'])
    sleep = sleep.drop(columns='key_time')
    sleep = sleep.rename(columns={'GROUP_USER_ID':'Group_User_ID',
                                  'REM':'REM(%)',
                                  'TTB':'SleepSession(min)'})

    sleep['TIMESTAMP'] = sleep['TIMESTAMP'].apply(lambda x: datetime.strptime(x, '%Y-%m-%dT%H:%M:%SZ'))
    sleep['TIMESTART'] = sleep['TIMESTART'].apply(lambda x: datetime.strptime(x, '%Y-%m-%dT%H:%M:%SZ'))

    sleep['Start_Date'] = sleep['TIMESTART'].dt.to_period('D')
    sleep['Start_Time'] = sleep['TIMESTART'].apply(lambda x: x.strftime('%H:%M:%S'))

    sleep['End_Date'] = sleep['TIMESTAMP'].dt.to_period('D')
    sleep['End_Time'] = sleep['TIMESTAMP'].apply(lambda x: x.strftime('%H:%M:%S'))

    sleep = sleep.drop(columns=['TIMESTAMP', 'TIMESTART'])

    # preprocessing
    sleep = sleep.drop(sleep[(sleep['SleepSession(min)'] <= 60) & (sleep['REM(%)'] == 0)].index)
    sleep['Date'] = sleep['End_Date'].astype('str')

    return sleep


def load_stress(based_url, head, body):
    # request from zensoriumAPI
    path = "measurement/getStressData"
    response = requests.post(based_url+path, json=body, headers=head)
    data = response.json()
    # data preprocessing
    stress = pd.DataFrame(data[0]['data'])
    stress = stress.drop(columns='key_time')
    stress = stress.rename(columns={'GROUP_USER_ID': 'Group_User_ID',
                                    'HR': 'HR(bpm)',
                                    'PSF': 'PSF(mmHg)',
                                    'QUADRANT': 'Stress'})

    stress['TIMESTAMP'] = stress['TIMESTAMP'].apply(lambda x: datetime.strptime(x, '%Y-%m-%dT%H:%M:%SZ'))

    stress['Date'] = stress['TIMESTAMP'].dt.to_period('D')
    stress['Time'] = stress['TIMESTAMP'].apply(lambda x: x.strftime('%H:%M:%S'))

    stress = stress.drop(columns='TIMESTAMP')
    return stress


def load_step(based_url, head, body):
    # request from zensoriumAPI
    path = "measurement/getStepData"
    response = requests.post(based_url + path, json=body, headers=head)
    data = response.json()
    # data preprocessing
    step = pd.DataFrame(data[0]['data'])
    step = step.drop(columns='key_time')
    step = step.rename(columns={'GROUP_USER_ID': 'Group_User_ID',
                                'SC': 'Step_Count',
                                'CALORIES': 'Calories'})

    step['TIMESTAMP'] = step['TIMESTAMP'].apply(lambda x: datetime.strptime(x, '%Y-%m-%dT%H:%M:%SZ'))

    step['Date'] = step['TIMESTAMP'].dt.to_period('D')
    step['Time'] = step['TIMESTAMP'].apply(lambda x: x.strftime('%H:%M:%S'))

    step = step.drop(columns='TIMESTAMP')
    return step

