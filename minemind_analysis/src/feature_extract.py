import numpy as np
import pandas as pd
import src.utility as utility
from saxpy.sax import sax_via_window


def coef_var(x):
    return x.std()/x.mean()


# ========== feature extraction ==========#
def stat_feature_extraction(df, df2, var=['HR(bpm)', 'Step_Count', 'PSV', 'PSF(mmHg)'], IDCol='Group_User_ID'):
    # ================ tracking data (df) ===============#
    # Allday
    allday = df.copy()
    ad_agg = allday.groupby(IDCol)[var].agg({'mean': np.mean, 'sd': np.std,
                                             'skew': pd.DataFrame.skew, 'kurt': pd.DataFrame.kurt})
    # correcting columns name
    columns = []
    for i in range(len(ad_agg.columns)):
        col = list(ad_agg.columns[i])
        col_name = col[0] + '_' + col[1]
        columns.append(col_name)
    ad_agg.columns = columns
    ad_agg.columns = ['ad_' + str(col) for col in ad_agg.columns]

    # sleep
    sleep_data = df[df.sleep == 1]
    sl_agg = sleep_data.groupby(IDCol)[var].agg({'mean': np.mean, 'sd': np.std,
                                                 'skew': pd.DataFrame.skew, 'kurt': pd.DataFrame.kurt})
    # correcting columns name
    columns = []
    for i in range(len(sl_agg.columns)):
        col = list(sl_agg.columns[i])
        col_name = col[0] + '_' + col[1]
        columns.append(col_name)
    sl_agg.columns = columns
    sl_agg.columns = ['sl_' + str(col) for col in sl_agg.columns]

    # nonsleep
    nonsleep_data = df[df.sleep == 0]
    ns_agg = nonsleep_data.groupby(IDCol)[var].agg({'mean': np.mean, 'sd': np.std,
                                                    'skew': pd.DataFrame.skew, 'kurt': pd.DataFrame.kurt})
    # correcting columns name
    columns = []
    for i in range(len(ns_agg.columns)):
        col = list(ns_agg.columns[i])
        col_name = col[0] + '_' + col[1]
        columns.append(col_name)
    ns_agg.columns = columns
    ns_agg.columns = ['ns_' + str(col) for col in ns_agg.columns]

    # =============== sleep data (df2) ===============#
    # sleep data
    sleep_frame = df2.groupby(IDCol).agg(
        mean_sleep_session=pd.NamedAgg(column='SleepSession(min)', aggfunc=np.mean),
        sd_sleep_session=pd.NamedAgg(column='SleepSession(min)', aggfunc=np.std),
        skew_sleep_session=pd.NamedAgg(column='SleepSession(min)', aggfunc=pd.DataFrame.skew),
        kurt_sleep_session=pd.NamedAgg(column='SleepSession(min)', aggfunc=pd.DataFrame.kurt),
        mean_REM=pd.NamedAgg(column='REM(%)', aggfunc=np.mean),
        sd_REM=pd.NamedAgg(column='REM(%)', aggfunc=np.std),
        skew_REM=pd.NamedAgg(column='REM(%)', aggfunc=pd.DataFrame.skew),
        kurt_REM=pd.NamedAgg(column='REM(%)', aggfunc=pd.DataFrame.kurt),
    )
    # Merge
    new_df = pd.concat([ad_agg, sl_agg, ns_agg, sleep_frame], axis=1)
    return new_df


# ===== sleep quality score ===== #
# calculate score from HR(bpm) during sleep
# require Birth Date, sleep=1, HR(bpm) data
def sleep_quality_score(tracking, user, IDCol='Group_User_ID'):
    tracking = tracking[tracking['sleep'] == 1]
    tracking = tracking[tracking['HR(bpm)'].notnull()]
    # Add Age column
    data = pd.merge(tracking, user, how='left', left_on=IDCol, right_on='GROUP_USER_ID')
    data['Date'] = pd.to_datetime(data['Date'])
    data['Birth_Date'] = pd.to_datetime(data['Birth_Date'], errors='coerce', format='%d/%m/%Y')
    data['Age'] = (data['Date'] - data['Birth_Date']).apply(lambda x: int(x.days/365))
    # sleep quality
    data['range'] = data.apply(lambda x: 1 if ( (x['HR(bpm)'] > 59)&
                                                (x['HR(bpm)'] < (191.5-0.07*x['Age']*x['Age'])*0.5*0.9)
                                                ) else 0,
                               axis=1)
    sleep_score = data.groupby(IDCol)['range'].agg({'sleep_score': np.mean})
    return sleep_score


# ===== sleep detail =====#
def sleep_detail(sleep, IDCol='Group_User_ID'):
    # Add columns
    sleep['start_hour'] = pd.to_datetime(sleep['Start_Time']).dt.hour
    sleep['start_day'] = sleep.apply(lambda x: 1 if ((x['start_hour'] >= 2) & (x['start_hour'] < 16)) else 0, axis=1)
    sleep['REM_0'] = sleep.apply(lambda x: 1 if x['REM(%)'] == 0 else 0, axis=1)
    sleep['REM_20_25'] = sleep.apply(lambda x: 1 if ((x['REM(%)'] >= 20) & (x['REM(%)'] <= 25)) else 0, axis=1)
    sleep['REM_Gt25'] = sleep.apply(lambda x: 1 if x['REM(%)'] > 25 else 0, axis=1)
    sleep['REM_Lt20'] = sleep.apply(lambda x: 1 if ((x['REM(%)'] < 20) & (x['REM(%)'] > 0)) else 0, axis=1)
    sleep['SSGT8'] = sleep.apply(lambda x: 1 if x['SleepSession(min)'] >= 480 else 0, axis=1)
    sleep['SS58'] = sleep.apply(lambda x: 1 if ((x['SleepSession(min)'] < 480) & (x['SleepSession(min)'] > 300)) else 0,
                                axis=1)
    sleep['SSLT3'] = sleep.apply(lambda x: 1 if x['SleepSession(min)'] <= 200 else 0, axis=1)

    # Aggregate data
    var = ['start_day', 'REM_0', 'REM_20_25', 'REM_Gt25', 'REM_Lt20', 'SSGT8', 'SS58', 'SSLT3']
    agg = sleep.groupby(IDCol)[var].agg({'(%)': np.mean})

    # correcting columns name
    columns = []
    for i in range(len(agg.columns)):
        col = list(agg.columns[i])
        col_name = col[1] + col[0]
        columns.append(col_name)
    agg.columns = columns
    return agg


#===== SAX BOP Prop =====#
def sax_freq_dist(dat, win_size=12, paa_size=3, alphabet_size=3, z_threshold=0.05):
    # SAX
    sax = sax_via_window(dat, win_size, paa_size, alphabet_size, z_threshold)
    # Count frequency
    sax_freq = dict()
    for key, value in sax.items():
        freq = len(value)
        sax_freq[key] = np.round(freq/len(dat), 2)
    return sax_freq


def sax_bop_prop(df, var='HR(bpm)', IDCol='Group_User_ID'):
    # Preprocessing
    df['Date'] = df['timestamp'].dt.to_period('D')
    df['Date'] = df['Date'].astype('str')
    df['ID-MD'] = df[IDCol] + '@' + df['Date']

    # pattern during sleep
    sleep_data = df[df['sleep'] == 1]
    idx = sleep_data['ID-MD'].unique()
    sax_freq_list = []
    for i in idx:
        # get time series data as array
        sample = sleep_data[sleep_data['ID-MD'] == i]
        sample = sample.sort_index()
        data = sample[var].values
        data = utility.forward_fill(data)
        # SAX
        sax_freq = sax_freq_dist(data)
        sax_freq['ID-MD'] = i
        sax_freq_list.append(sax_freq)

    # Create DataFrame
    col_name = ['ID-MD',
                'aaa', 'aac', 'abb', 'abc', 'acc', 'acb', 'aca',
                'bab', 'bcb', 'bba', 'bca', 'bbb', 'bbc', 'bac',
                'cba', 'cca', 'caa', 'cbb', 'cab', 'cac']
    bop = pd.DataFrame(sax_freq_list, columns=col_name)
    bop = bop.fillna(0)
    bop['ID'] = bop['ID-MD'].str.split('@', expand=True)[0]

    # Calculate Proportion by ID
    bop_prop = bop.groupby('ID').mean()

    return bop_prop


