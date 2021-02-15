package service

import (
    "context"
    "github.com/aoffy-kku/minemind-backend/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type AdminService struct {
    db *mongo.Database
}

func (a *AdminService) GetUsersAnalysis() ([]*model.AnalysisJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("user_analysis")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.AnalysisJSON
    cur, err := col.Find(ctx, bson.M{
    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.Analysis
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        var questions []model.QuestionJSON
        for _, q := range m.UserEvaluation.Questions {
            var opts []model.OptionJSON
            for _, o := range q.Options {
                opts = append(opts, model.OptionJSON{
                    Id:    o.Id,
                    Title: o.Title,
                    Value: o.Value,
                })
            }
            questions = append(questions, model.QuestionJSON{
                Id:      q.Id,
                Title:   q.Title,
                Options: opts,
            })
        }
        docs = append(docs, &model.AnalysisJSON{
            Id:               m.Id,
            Mode:             m.Mode,
            UserId:           m.UserId,
            UserEvaluationId: m.UserEvaluationId,
            UserEvaluation: model.UserEvaluationJSON{
                Id:           m.UserEvaluation.Id,
                EvaluationId: m.UserEvaluation.EvaluationId,
                UserId:       m.UserEvaluation.UserId,
                Name:         m.UserEvaluation.Name,
                Description:  m.UserEvaluation.Description,
                Questions:    questions,
                CreatedAt:    m.CreatedAt,
                CreatedBy:    m.CreatedBy,
            },
            CortisolId: m.CortisolId,
            Cortisol: model.CortisolJSON{
                Id:        m.Cortisol.Id,
                UserId:    m.Cortisol.UserId,
                Cortisol:  m.Cortisol.Cortisol,
                Timestamp: m.Cortisol.Timestamp,
                CreatedAt: m.Cortisol.CreatedAt,
                CreatedBy: m.Cortisol.CreatedBy,
            },
            Status:    m.Status,
            Class:     m.Class,
            Score:     m.Score,
            CreatedAt: m.CreatedAt,
            CreatedBy: m.CreatedBy,
            UpdatedAt: m.UpdatedAt,
            UpdatedBy: m.UpdatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUserAnalysis(id string) ([]*model.AnalysisJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("user_analysis")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.AnalysisJSON
    cur, err := col.Find(ctx, bson.M{
        "user_id": bson.M{
            "$eq": id,
        },
    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.Analysis
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        var questions []model.QuestionJSON
        for _, q := range m.UserEvaluation.Questions {
            var opts []model.OptionJSON
            for _, o := range q.Options {
                opts = append(opts, model.OptionJSON{
                    Id:    o.Id,
                    Title: o.Title,
                    Value: o.Value,
                })
            }
            questions = append(questions, model.QuestionJSON{
                Id:      q.Id,
                Title:   q.Title,
                Options: opts,
            })
        }
        docs = append(docs, &model.AnalysisJSON{
            Id:               m.Id,
            Mode:             m.Mode,
            UserId:           m.UserId,
            UserEvaluationId: m.UserEvaluationId,
            UserEvaluation: model.UserEvaluationJSON{
                Id:           m.UserEvaluation.Id,
                EvaluationId: m.UserEvaluation.EvaluationId,
                UserId:       m.UserEvaluation.UserId,
                Name:         m.UserEvaluation.Name,
                Description:  m.UserEvaluation.Description,
                Questions:    questions,
                CreatedAt:    m.CreatedAt,
                CreatedBy:    m.CreatedBy,
            },
            CortisolId: m.CortisolId,
            Cortisol: model.CortisolJSON{
                Id:        m.Cortisol.Id,
                UserId:    m.Cortisol.UserId,
                Cortisol:  m.Cortisol.Cortisol,
                Timestamp: m.Cortisol.Timestamp,
                CreatedAt: m.Cortisol.CreatedAt,
                CreatedBy: m.Cortisol.CreatedBy,
            },
            Status:    m.Status,
            Class:     m.Class,
            Score:     m.Score,
            CreatedAt: m.CreatedAt,
            CreatedBy: m.CreatedBy,
            UpdatedAt: m.UpdatedAt,
            UpdatedBy: m.UpdatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUserDiary(id string) ([]*model.UserDiaryJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("user_diary")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.UserDiaryJSON
    cur, err := col.Find(ctx, bson.M{
        "user_id": bson.M{
            "$eq": id,
        },
    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.UserDiary
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        docs = append(docs, &model.UserDiaryJSON{
            Id:        m.Id,
            UserId:    m.UserId,
            Moods:     m.Moods,
            Content:   m.Content,
            CreatedAt: m.CreatedAt,
            CreatedBy: m.CreatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUserEvaluation(id string) ([]*model.UserEvaluationJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("user_evaluation")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.UserEvaluationJSON
    cur, err := col.Find(ctx, bson.M{
        "user_id": bson.M{
            "$eq": id,
        },
    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.UserEvaluation
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        var questions []model.QuestionJSON
        for _, q := range m.Questions {
            var opts []model.OptionJSON
            for _, o := range q.Options {
                opts = append(opts, model.OptionJSON{
                    Id:    o.Id,
                    Title: o.Title,
                    Value: o.Value,
                })
            }
            questions = append(questions, model.QuestionJSON{
                Id:      q.Id,
                Title:   q.Title,
                Options: opts,
            })
        }
        docs = append(docs, &model.UserEvaluationJSON{
            Id:           m.Id,
            EvaluationId: m.EvaluationId,
            UserId:       m.UserId,
            Name:         m.Name,
            Description:  m.Description,
            Questions:    questions,
            CreatedAt:    m.CreatedAt,
            CreatedBy:    m.CreatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUserCortisol(id string) ([]*model.CortisolJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("cortisol")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.CortisolJSON
    cur, err := col.Find(ctx, bson.M{
        "user_id": bson.M{
            "$eq": id,
        },
    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.Cortisol
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        docs = append(docs, &model.CortisolJSON{
            Id:        m.Id,
            UserId:    m.UserId,
            Cortisol:  m.Cortisol,
            Timestamp: m.Timestamp,
            CreatedAt: m.CreatedAt,
            CreatedBy: m.CreatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUsers() ([]*model.UserJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("user")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
    })
    var docs []*model.UserJSON
    cur, err := col.Find(ctx, bson.M{
    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.User
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        docs = append(docs, &model.UserJSON{
            Email:       m.Email,
            DisplayName: m.DisplayName,
            WatchId:     m.WatchId,
            Roles:       m.Roles,
            Begin:       m.Begin,
            End:         m.End,
            BirthDate:   m.BirthDate,
            CreatedAt:   m.CreatedAt,
            CreatedBy:   m.CreatedBy,
            UpdatedAt:   m.UpdatedAt,
            UpdatedBy:   m.UpdatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUsersDiary() ([]*model.UserDiaryJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("user_diary")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.UserDiaryJSON
    cur, err := col.Find(ctx, bson.M{

    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.UserDiary
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        docs = append(docs, &model.UserDiaryJSON{
            Id:        m.Id,
            UserId:    m.UserId,
            Moods:     m.Moods,
            Content:   m.Content,
            CreatedAt: m.CreatedAt,
            CreatedBy: m.CreatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUsersEvaluation() ([]*model.UserEvaluationJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("user_evaluation")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.UserEvaluationJSON
    cur, err := col.Find(ctx, bson.M{

    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.UserEvaluation
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        var questions []model.QuestionJSON
        for _, q := range m.Questions {
            var opts []model.OptionJSON
            for _, o := range q.Options {
                opts = append(opts, model.OptionJSON{
                    Id:    o.Id,
                    Title: o.Title,
                    Value: o.Value,
                })
            }
            questions = append(questions, model.QuestionJSON{
                Id:      q.Id,
                Title:   q.Title,
                Options: opts,
            })
        }
        docs = append(docs, &model.UserEvaluationJSON{
            Id:           m.Id,
            EvaluationId: m.EvaluationId,
            UserId:       m.UserId,
            Name:         m.Name,
            Description:  m.Description,
            Questions:    questions,
            CreatedAt:    m.CreatedAt,
            CreatedBy:    m.CreatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) GetUsersCortisol() ([]*model.CortisolJSON, error) {
    ctx := context.Background()

    col := a.db.Collection("cortisol")
    opts := &options.FindOptions{}
    opts.SetSort(bson.M{
        "user_id": -1,
        "created_at": -1,
    })
    var docs []*model.CortisolJSON
    cur, err := col.Find(ctx, bson.M{

    }, opts)
    if err != nil {
        return nil, err
    }
    for cur.Next(ctx) {
        var m model.Cortisol
        if err := cur.Decode(&m); err != nil {
            return nil, err
        }
        docs = append(docs, &model.CortisolJSON{
            Id:        m.Id,
            UserId:    m.UserId,
            Cortisol:  m.Cortisol,
            Timestamp: m.Timestamp,
            CreatedAt: m.CreatedAt,
            CreatedBy: m.CreatedBy,
        })
    }
    return docs, nil
}

func (a *AdminService) UpdateUser(request model.UpdateUserRequestJSON) (*model.UserJSON, error) {
    ctx := context.Background()
    col := a.db.Collection("user")
    _, err := col.UpdateOne(ctx, bson.M{
        "_id": bson.M{
            "$eq": request.Email,
        },
    }, bson.M{
        "$set": bson.M{
            "watch_id": request.WatchId,
            "display_name": request.DisplayName,
            "begin": request.Begin,
            "end": request.End,
        },
    })
    if err != nil {
        return nil, err
    }
    var doc model.User
    if err := col.FindOne(ctx, bson.M{
        "_id": bson.M{
            "$eq": request.Email,
        },
    }).Decode(&doc); err != nil {
        return nil, err
    }
    return &model.UserJSON{
        Email:       doc.Email,
        DisplayName: doc.DisplayName,
        WatchId:     doc.WatchId,
        Roles:       doc.Roles,
        Begin:       doc.Begin,
        End:         doc.End,
        BirthDate:   doc.BirthDate,
        CreatedAt:   doc.CreatedAt,
        CreatedBy:   doc.CreatedBy,
        UpdatedAt:   doc.UpdatedAt,
        UpdatedBy:   doc.UpdatedBy,
    }, nil
}

func NewAdminService(db * mongo.Database) *AdminService {
    return &AdminService{
        db: db,
    }
}
