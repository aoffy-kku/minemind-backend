#!/usr/bin/python3

import sys
from src import model_watch, model_cortisol, model_watchcortisol, model_st5, \
    model_phq9, model_all
import json

"""
    parameters
    case = model
    id = analysis id
    clm = group user id ex. CLM_12345678
    st5 = st5 result ex. 20
    phq9 = phq9 result ex. 19
    cotisol = cortisol value ex. 61.51
    timestamp = 31/01/2020 18:30
    birth_date = birth date ex. 31/01/2020
    start_date = start date ex. 2020-01-31
    end_date = start date ex. 2020-02-31
"""
def selecting_model(case, id, clm, st5, phq9, cortisol, timestamp, birth_date, start_date, end_date):
    """
    Depression Classification Model
        case 1: using smart watch data for classify
        case 2: using cortisol data for classify
        case 3: using smart watch and cortisol data for classify
        case 4: using st5 data for classify
        case 5: using phq9 data for classify
        case 6: using All data for classify

    return:
        Class 0 = not depression, 1 = depression
        Score = probability
    """
    if case == 1:
        return model_watch.model(id, clm, birth_date, start_date, end_date)
    if case == 2:
        return model_cortisol.model(id, cortisol, timestamp)
    if case == 3:
        return model_watchcortisol.model(id, clm, cortisol, timestamp, birth_date, start_date, end_date)
    if case == 4:
        return model_st5.model(id, st5)
    if case == 5:
        return model_phq9.model(id, phq9)
    if case == 6:
        return model_all.model(id, clm, st5, phq9, cortisol, timestamp, birth_date, start_date, end_date)


def main():
    prediction = selecting_model(
        case=int(sys.argv[1]),
        id=str(sys.argv[2]),
        clm=str(sys.argv[3]),
        st5=str(sys.argv[4]),
        phq9=str(sys.argv[5]),
        cortisol=str(sys.argv[6]),
        timestamp=str(sys.argv[7]),
        birth_date=str(sys.argv[8]),
        start_date=str(sys.argv[9] + "T00:00:00Z"),
        end_date=str(sys.argv[10] + "T:23:59:00Z")
    )
    
    print("Class {} Score {}".format(prediction['class'], prediction['score']))

    with open('prediction.json', 'w') as f:
        json.dump(prediction, f)


if __name__ == "__main__":
    main()
