#Author Eban
#Get card is open on a week or your time input
URL: "/b/cards/review"
    Method: GET
Input:
    query:list string
    query:time init
Output: [[{
        ID                   string
        Name                 string
        ListName             string
        IdList               string
        TimeGuessForDone     int
        TimeRealForDone      int
        DateLastActivity     *time.Time
        Due                  *time.Time
        ChangeDueDate        bool
        HistoryChangeDueDate []*time.Time
    }],
    [{
        ID                   string
        Name                 string
        ListName             string
        IdList               string
        TimeGuessForDone     int
        TimeRealForDone      int
        DateLastActivity     *time.Time
        Due                  *time.Time
        ChangeDueDate        bool
        HistoryChangeDueDate []*time.Time
    }]]


#Get card is change due on a week or your time input
URL: "/b/cards/change-due"
    Method: GET
Input:
    query:time init
Output: [{
    ID                   string
    Name                 string
    ListName             string
    IdList               string
    TimeGuessForDone     int
    TimeRealForDone      int
    DateLastActivity     *time.Time
    Due                  *time.Time
    ChangeDueDate        bool
    HistoryChangeDueDate []*time.Time
}]