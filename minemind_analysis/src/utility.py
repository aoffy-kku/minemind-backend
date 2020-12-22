import pandas as pd
import datetime

def adj_timestamp(data, timestamp='timestamp', Date='Date', Time='Time'):
    data = data.drop_duplicates()
    data[Date] =  data[Date].astype('str')
    data[Time] = data[Time].astype('str')

    tmp = data[Date] + ' ' + data[Time]
    tmp = pd.to_datetime(tmp)
    hour = tmp.dt.hour
    minute = (tmp.dt.minute // 5) * 5
    data[timestamp] = data[Date] + ' ' + hour.astype(str) + ':' + minute.astype(str)
    data[timestamp] = pd.to_datetime(data[timestamp]) + datetime.timedelta(hours=7)

    data = data.drop(columns=[Date, Time])
    return data


def selecting_time(df, meetingdate, IDCol='Group_User_ID', start='firstdate', end='meet4'):
    new_df = pd.DataFrame()
    idx = meetingdate['ID'].unique()
    for i in idx:
        # data
        data = df[df[IDCol] == i]
        # define date
        smp_md = meetingdate[meetingdate.ID == i]
        start_date = pd.to_datetime(smp_md[start].values[0])
        end_date = pd.to_datetime(smp_md[end].values[0])

        # filter
        smp = data[(pd.to_datetime(data['Date'])>=start_date)&(pd.to_datetime(data['Date'])<end_date)]
        new_df = new_df.append(smp)
    return new_df


def forward_fill(arr):
    df = pd.DataFrame(arr)
    df.fillna(method='ffill', axis=0, inplace=True)
    out = df[0].values
    return out
