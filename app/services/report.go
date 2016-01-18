package services

import (
	"PJApp/app/models"

	_ "github.com/go-sql-driver/mysql"
	"PJApp/app/constant"
	"time"
)

// 一覧画面での取得処理
func GetReportList(userId int, start, end time.Time) *[]models.Report {

	reportList := []models.Report{}
	DB.Debug().Select([]string{"Id","start_day", "update_day"}).
	Where("user_id = ? and start_day between ? and ? and Status = ?",
		userId, start.Format(constant.STR_TYPE_LAYOUT), end.Format("20060102"), constant.STATUS_ARI).
	Order("start_day desc").Find(&reportList)

	return &reportList
}

// 該当する進捗情報取得
func GetReport(userId int, startDay time.Time) *models.Report {

	startDayStr := startDay.Format(constant.STR_TYPE_LAYOUT)
	reportList := []models.Report{}

	DB.Debug().Where("user_id = ? and start_day = ? and Status = ?",
		userId, startDayStr, constant.STATUS_ARI).Find(&reportList)

	if (len(reportList) == 0) {
		return nil
	}

	return &reportList[0]
}

// Insert
func InsertReport(userId int, startDay string, report models.Report) error {

	report.UserId = userId
	report.StartDay = startDay
	report.Status = constant.STATUS_ARI
	return DB.Debug().Create(&report).Error
}

// Update
func UpdateReport(id, userId int, startDay string, report models.Report)  {

	report.Id = id
	report.UserId = userId
	report.StartDay = startDay
	report.Status = constant.STATUS_ARI

	DB.Debug().Where("Status = ?", constant.STATUS_ARI).Save(&report)
}
