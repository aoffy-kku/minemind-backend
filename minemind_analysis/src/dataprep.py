import pandas as pd
from src.utility import adj_timestamp


def tracking_prep(stress, step, sleep):
    # ===== preprocessing ===== #
    stress = adj_timestamp(stress)
    step = adj_timestamp(step)
    sleep = adj_timestamp(sleep, timestamp='start_timestamp', Date='Start_Date', Time='Start_Time')
    sleep = adj_timestamp(sleep, timestamp='end_timestamp', Date='End_Date', Time='End_Time')

    # ===== add columns ===== #
    # Merge Stress, Step and add Active column
    stress['Active'] = 1
    df = pd.merge(stress, step, on=['Group_User_ID', 'timestamp'], how='right')

    # Add sleep column
    df['sleep'] = 0
    for sample in df['Group_User_ID'].unique():
        sleep_sample = sleep[sleep['Group_User_ID'] == sample]
        sleep_start = sleep_sample['start_timestamp']
        sleep_end = sleep_sample['end_timestamp']
        for i in range(0, len(sleep_sample)):
            sleep_bool = (df['Group_User_ID'] == sample) & (df['timestamp'] >= sleep_start.iloc[i]) & (
                    df['timestamp'] <= sleep_end.iloc[i])
            df.loc[sleep_bool, ['sleep']] = 1

    # preprocessing
    df['timestamp'] = pd.to_datetime(df['timestamp'])
    df['Date'] = df['timestamp'].dt.to_period('D').astype('str')
    df.index = df['timestamp']

    return df

