package handler

import (
    "bytes"
    "context"
    "encoding/json"
    "fmt"
    "github.com/aoffy-kku/minemind-backend/model"
    "github.com/labstack/echo/v4"
    "github.com/sirupsen/logrus"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "io/ioutil"
    "net/http"
    "time"
)

const token = "93aafa0e6cd844fd8332ae4027650bb8f08b5b2b4ab7958c5c57a2b5f84b7d84"
const api = "https://info.minemind.net/zensoriumAPI/measurement"

var stepUrl = fmt.Sprintf("%s/%s", api, "getStepData")
var sleepUrl = fmt.Sprintf("%s/%s", api, "getSleepData")
var stressUrl = fmt.Sprintf("%s/%s", api, "getStressData")

type Step struct {
    Calories  int64     `json:"CALORIES"`
    StepCount int64     `json:"SC"`
    Timestamp time.Time `json:"TIMESTAMP"`
}
type StepData struct {
    Data []Step `json:"data"`
}

type Sleep struct {
    Rem       int64     `json:"REM"`
    Timestamp time.Time `json:"TIMESTAMP"`
    Ttb       int64     `json:"TTB"`
}
type SleepData struct {
    Data []Sleep `json:"data"`
}

type Stress struct {
    Quadrant  int64     `json:"QUADRANT"`
    Timestamp time.Time `json:"TIMESTAMP"`
}
type StressData struct {
    Data []Stress `json:"data"`
}

type Mood struct {
    Mood string `bson:"mood" json:"mood"`
    Count int64 `bson:"count" json:"count"`
}

type StepMeasurement struct {
    Value     float64   `json:"value"`
    Timestamp time.Time `json:"timestamp"`
}
type SleepMeasurement struct {
    Rem int64           `json:"rem"`
    Ttb int64           `json:"ttb"`
    Timestamp time.Time `json:"timestamp"`
}
type StressMeasurement struct {
    Normal     int64   `json:"normal"`
    Stress     int64   `json:"stress"`
    Timestamp time.Time `json:"timestamp"`
}
type MoodMeasurement struct {
    Moods []Mood `json:"moods"`
    //Timestamp time.Time `json:"timestamp"`
}

type StepResponse struct {
    Today []StepMeasurement `json:"today"`
    Week  []StepMeasurement `json:"week"`
    Month []StepMeasurement `json:"month"`
}
type SleepResponse struct {
    Today []SleepMeasurement `json:"today"`
    Week  []SleepMeasurement `json:"week"`
    Month []SleepMeasurement `json:"month"`
}
type StressResponse struct {
    Today []StressMeasurement `json:"today"`
    Week  []StressMeasurement `json:"week"`
    Month []StressMeasurement `json:"month"`
}
type MoodResponse struct {
    Today MoodMeasurement `json:"today"`
    Week  MoodMeasurement `json:"week"`
    Month MoodMeasurement `json:"month"`
}
type MeasurementResponse struct {
    Steps  StepResponse `json:"steps"`
    Sleep  SleepResponse `json:"sleep"`
    Stress StressResponse `json:"stress"`
    Mood   MoodResponse `json:"mood"`
}

var client = &http.Client{}

// GetMeasurement godoc
// @tags Measurement
// @Summary Get measurement
// @Accept  json
// @Produce  json
// @Success 200 {object} MeasurementResponse
// @Failure 401 {object} utils.HttpResponse
// @Failure 403 {object} utils.HttpResponse
// @Failure 404 {object} utils.HttpResponse
// @Failure 500 {object} utils.HttpResponse
// @Router /v1/measurement [get]
// @Security ApiKeyAuth
func (h *Handler) GetMeasurement(c echo.Context) error {
    id := c.Get("id").(string)
    user, err := h.userService.GetUserById(id)
    if err != nil {
        return c.JSON(http.StatusNotFound, "user not found")
    }
    clm := user.WatchId
    type Request struct {
        GroupUserIds []string `json:"groupUserIds"`
        FromDateTime string   `json:"fromDateTime"`
        ToDateTime   string   `json:"toDateTime"`
    }
    ts := time.Now()
    y, m, d := time.Unix(ts.Unix()-int64(60*60*24*30), 0).Date()
    req := Request{
        GroupUserIds: []string{
            clm,
        },
        //FromDateTime: "2020-05-07T05:15:59+0000",
        //ToDateTime: "2020-06-07T05:15:59+0000",
        FromDateTime: time.Date(y, m, d, 0, 0, 0, 0, time.Local).Format("2006-01-02T15:04:05.999Z"),
        ToDateTime:   ts.Format("2006-01-02T15:04:05.999Z"),
    }
    body, err := json.Marshal(req)
    if err != nil {
        panic(err)
    }
    return c.JSON(http.StatusOK, MeasurementResponse{
        Steps:  GetStep(body, ts),
        Sleep:  GetSleep(body, ts),
        Stress: GetStress(body, ts),
        Mood: GetMood(h.db, id, ts),
    })
}

func GetStep(body []byte, timestamp time.Time) StepResponse {
    httpReq, err := http.NewRequest("POST", stepUrl, bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
    httpReq.Header.Add("Content-Type", "application/json")
    resp, err := client.Do(httpReq)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    res, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    var data []StepData
    json.Unmarshal(res, &data)
    if len(data) > 0 {
        return StepResponse{
            Today: GetTodayStep(data[0], timestamp),
            Week:  GetWeekStep(data[0], timestamp),
            Month: GetMonthStep(data[0], timestamp),
        }
    }
    return StepResponse{
        Today: []StepMeasurement{},
        Week:  []StepMeasurement{},
        Month: []StepMeasurement{},
    }
}

func GetTodayStep(steps StepData, timestamp time.Time) []StepMeasurement {
    y, m, d := timestamp.Date()
    var data []StepMeasurement
    for i := 0; i < 24; i++ {
        sc := int64(0)
        date := time.Date(y, m, d, i, 0, 0, 0, time.Local)
        for _, step := range steps.Data {
            y2, m2, d2 := step.Timestamp.Date()
            h, _, _ := step.Timestamp.Clock()
            if y == y2 && m == m2 && d == d2 && h == i {
                sc = sc + step.StepCount
            }
        }
        data = append(data, StepMeasurement{
            Value:     float64(sc),
            Timestamp: date,
        })
    }
    return data
}

func GetWeekStep(steps StepData, timestamp time.Time) []StepMeasurement {
    y, m, d := timestamp.Date()
    var data []StepMeasurement
    for i := int64(0); i < 7; i++ {
        sc := int64(0)
        date := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix() - (int64(60*60*24) * i)
        y3, m3, d3 := time.Unix(date, 0).Date()
        for _, step := range steps.Data {
            y2, m2, d2 := step.Timestamp.Date()
            if y3 == y2 && m3 == m2 && d3 == d2 {
                sc = sc + step.StepCount
            }
        }
        data = append(data, StepMeasurement{
            Value:     float64(sc),
            Timestamp: time.Unix(date, 0),
        })
    }
    return data
}

func GetMonthStep(steps StepData, timestamp time.Time) []StepMeasurement {
    y, m, d := timestamp.Date()
    var data []StepMeasurement
    for i := int64(0); i < 30; i++ {
        sc := int64(0)
        date := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix() - (int64(60*60*24) * i)
        y3, m3, d3 := time.Unix(date, 0).Date()
        for _, step := range steps.Data {
            y2, m2, d2 := step.Timestamp.Date()
            if y3 == y2 && m3 == m2 && d3 == d2 {
                sc = sc + step.StepCount
            }
        }
        data = append(data, StepMeasurement{
            Value:     float64(sc),
            Timestamp: time.Unix(date, 0),
        })
    }
    return data
}

func GetSleep(body []byte, timestamp time.Time) SleepResponse {
    httpReq, err := http.NewRequest("POST", sleepUrl, bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
    httpReq.Header.Add("Content-Type", "application/json")
    resp, err := client.Do(httpReq)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    res, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    var data []SleepData
    json.Unmarshal(res, &data)
    if len(data) > 0 {
        return SleepResponse{
            Today: GetTodaySleep(data[0], timestamp),
            Week:  GetWeekSleep(data[0], timestamp),
            Month: GetMonthSleep(data[0], timestamp),
        }
    }
    return SleepResponse{
        Today: []SleepMeasurement{},
        Week:  []SleepMeasurement{},
        Month: []SleepMeasurement{},
    }
}

func GetTodaySleep(sleeps SleepData, timestamp time.Time) []SleepMeasurement {
    y, m, d := timestamp.Date()
    var data []SleepMeasurement
    for i := 0; i < 24; i++ {
        ttb := int64(0)
        rem := int64(0)
        date := time.Date(y, m, d, i, 0, 0, 0, time.Local)
        count := int64(0)
        for _, sleep := range sleeps.Data {
            y2, m2, d2 := sleep.Timestamp.Date()
            h, _, _ := sleep.Timestamp.Clock()
            if y == y2 && m == m2 && d == d2 && h == i {
                ttb = ttb + sleep.Ttb
                rem = rem + sleep.Rem
                count = count + 1
            }
        }
        if count == 0 {
            data = append(data, SleepMeasurement{
                Rem: rem,
                Ttb:     ttb,
                Timestamp: date,
            })
        } else {
            data = append(data, SleepMeasurement{
                Rem: rem / count,
                Ttb:     ttb,
                Timestamp: date,
            })
        }

    }
    return data
}

func GetWeekSleep(sleeps SleepData, timestamp time.Time) []SleepMeasurement {
    y, m, d := timestamp.Date()
    var data []SleepMeasurement
    for i := int64(0); i < 7; i++ {
        rem := int64(0)
        ttb := int64(0)
        date := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix() - (int64(60*60*24) * i)
        y3, m3, d3 := time.Unix(date, 0).Date()
        count := int64(0)
        for _, sleep := range sleeps.Data {
            y2, m2, d2 := sleep.Timestamp.Date()
            if y3 == y2 && m3 == m2 && d3 == d2 {
                rem = rem + sleep.Rem
                ttb = ttb + sleep.Ttb
                count = count + 1
            }
        }
        if count == 0 {
            data = append(data, SleepMeasurement{
                Rem: rem,
                Ttb: ttb,
                Timestamp: time.Unix(date, 0),
            })
        } else {
            data = append(data, SleepMeasurement{
                Rem: rem / count,
                Ttb: ttb,
                Timestamp: time.Unix(date, 0),
            })
        }

    }
    return data
}

func GetMonthSleep(sleeps SleepData, timestamp time.Time) []SleepMeasurement {
    y, m, d := timestamp.Date()
    var data []SleepMeasurement
    for i := int64(0); i < 30; i++ {
        rem := int64(0)
        ttb := int64(0)
        date := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix() - (int64(60*60*24) * i)
        y3, m3, d3 := time.Unix(date, 0).Date()
        count := int64(0)
        for _, sleep := range sleeps.Data {
            y2, m2, d2 := sleep.Timestamp.Date()
            if y3 == y2 && m3 == m2 && d3 == d2 {
                rem = rem + sleep.Rem
                ttb = ttb + sleep.Ttb
                count = count + 1
            }
        }
        if count == 0 {
            data = append(data, SleepMeasurement{
                Rem: rem,
                Ttb: ttb,
                Timestamp: time.Unix(date, 0),
            })
        } else {
            data = append(data, SleepMeasurement{
                Rem: rem / count,
                Ttb: ttb,
                Timestamp: time.Unix(date, 0),
            })
        }
    }
    return data
}

func GetStress(body []byte, timestamp time.Time) StressResponse {
    httpReq, err := http.NewRequest("POST", stressUrl, bytes.NewBuffer(body))
    if err != nil {
        panic(err)
    }
    httpReq.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
    httpReq.Header.Add("Content-Type", "application/json")
    resp, err := client.Do(httpReq)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()
    res, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }
    var data []StressData
    json.Unmarshal(res, &data)
    if len(data) > 0 {
        return StressResponse{
            Today: GetTodayStress(data[0], timestamp),
            Week:  GetWeekStress(data[0], timestamp),
            Month: GetMonthStress(data[0], timestamp),
        }
    }
    return StressResponse{
        Today: []StressMeasurement{},
        Week:  []StressMeasurement{},
        Month: []StressMeasurement{},
    }
}

func GetTodayStress(stresses StressData, timestamp time.Time) []StressMeasurement {
    y, m, d := timestamp.Date()
    var data []StressMeasurement
    for i := 0; i < 24; i++ {
        normal := int64(0)
        stressCount := int64(0)
        date := time.Date(y, m, d, i, 0, 0, 0, time.Local)
        for _, stress := range stresses.Data {
            y2, m2, d2 := stress.Timestamp.Date()
            h, _, _ := stress.Timestamp.Clock()
            if y == y2 && m == m2 && d == d2 && h == i {
                if stress.Quadrant == 3 {
                    normal = normal + 1
                } else if stress.Quadrant == 4 {
                    stressCount = stressCount + 1
                }
            }
        }
        data = append(data, StressMeasurement{
            Normal: normal,
            Stress: stressCount,
            Timestamp: date,
        })
    }
    return data
}

func GetWeekStress(stresses StressData, timestamp time.Time) []StressMeasurement {
    y, m, d := timestamp.Date()
    var data []StressMeasurement
    for i := int64(0); i < 7; i++ {
        normal := int64(0)
        stressCount := int64(0)
        date := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix() - (int64(60*60*24) * i)
        y3, m3, d3 := time.Unix(date, 0).Date()
        for _, stress := range stresses.Data {
            y2, m2, d2 := stress.Timestamp.Date()
            if y3 == y2 && m3 == m2 && d3 == d2 {
                if stress.Quadrant == 3 {
                    normal = normal + 1
                } else if stress.Quadrant == 4 {
                    stressCount = stressCount + 1
                }
            }
        }
        data = append(data, StressMeasurement{
            Normal: normal,
            Stress: stressCount,
            Timestamp: time.Unix(date, 0),
        })
    }
    return data
}

func GetMonthStress(stresses StressData, timestamp time.Time) []StressMeasurement {
    y, m, d := timestamp.Date()
    var data []StressMeasurement
    for i := int64(0); i < 30; i++ {
        normal := int64(0)
        stressCount := int64(0)
        date := time.Date(y, m, d, 0, 0, 0, 0, time.Local).Unix() - (int64(60*60*24) * i)
        y3, m3, d3 := time.Unix(date, 0).Date()
        for _, stress := range stresses.Data {
            y2, m2, d2 := stress.Timestamp.Date()
            if y3 == y2 && m3 == m2 && d3 == d2 {
                if stress.Quadrant == 3 {
                    normal = normal + 1
                } else if stress.Quadrant == 4 {
                    stressCount = stressCount + 1
                }
            }
        }
        data = append(data, StressMeasurement{
            Normal: normal,
            Stress: stressCount,
            Timestamp: time.Unix(date, 0),
        })
    }
    return data
}

func GetMood(db *mongo.Database, id string,  timestamp time.Time) MoodResponse {
    y,m,d := timestamp.Date()
    ts := time.Date(y, m, d, 0, 0, 0, 0, time.Local)
    month := time.Unix(ts.Unix()-int64(60*60*24*30), 0).Unix()
    week := time.Unix(ts.Unix()-int64(60*60*24*7), 0).Unix()
    today := ts.Unix()
    logrus.Println(month, week, today)
    return MoodResponse{
        Today: QueryMood(db, id, today),
        Week:  QueryMood(db, id, week),
        Month: QueryMood(db, id, month),
    }
}

func QueryMood(db *mongo.Database, id string, start int64) MoodMeasurement {
    ts := time.Unix(start, 0)
    pipe := []bson.M{
        {
            "$match": bson.M{
                "user_id": id,
                "moods": bson.M{
                    "$ne": nil,
                },
                "created_at": bson.M{
                    "$gte": ts,
                },
            },
        },
        {
            "$project": bson.M{
                "moods": 1,
                "created_at": 1,
            },
        },
        {
            "$unwind": bson.M{
                "path": "$moods",
                "preserveNullAndEmptyArrays": true,
            },
        },
        {
            "$group": bson.M{
                "_id": "$moods",
                "count": bson.M{
                    "$sum": 1,
                },
            },
        },
        {
            "$project": bson.M{
                "_id": 0,
                "moods": "$_id",
                "count": 1,
            },
        },
        {
            "$match": bson.M{
                "count": bson.M{
                    "$gt": 0,
                },
                "moods": bson.M{
                    "$ne": nil,
                },
            },
        },
        {
            "$lookup": bson.M{
                "from": "user_mood",
                "let": bson.M{
                    "mid": "$moods",
                },
                "pipeline": []bson.M{
                    {
                        "$match": bson.M{
                            "$expr": bson.M{
                                "$eq": []interface{}{
                                    "$_id",
                                    "$$mid",
                                },
                            },
                        },
                    },
                },
                "as": "mood",
            },
        },
        {
            "$unwind": bson.M{
                "path": "$mood",
                "preserveNullAndEmptyArrays": true,
            },
        },
        {
            "$sort": bson.M{
              "count": -1,
            },
        },
    }
    col := db.Collection("user_diary")
    cur, err := col.Aggregate(context.Background(), pipe)
    if err != nil {
        return MoodMeasurement{}
    }
    var moods []Mood
    for cur.Next(context.Background()) {
        type Template struct {
            Mood  model.UserMood `bson:"mood" json:"mood"`
            Count int64   `bson:"count" json:"count"`
        }
        var m Template
        if err := cur.Decode(&m); err != nil {
          return MoodMeasurement{}
        }
        moods = append(moods, Mood{
          Mood:  m.Mood.Name,
          Count: m.Count,
        })
    }
    return MoodMeasurement{Moods: moods}
}
