package pj

import (
	"github.com/revel/revel"
	"PJApp/app/models"
	"PJApp/app/controllers/login"
	"PJApp/app/services"
	"PJApp/app/controllers/common"
	"time"
	"PJApp/app/constant"
	"PJApp/app/forms"
	"fmt"
)

type PJ struct {
	login.Login
}

/* 初期表示処理 */
func (c PJ) Index() revel.Result {
	common.WriteLog("Top", "pushLoginButton", "Start")

	// 取得範囲計算
	startRange := time.Now().AddDate(0, 0, -1 * constant.RANGE_REPORT_LIST_BEFORE_DAY)
	endRange := time.Now().AddDate(0, 0, constant.RANGE_REPORT_LIST_AFTER_DAY)

	// 進捗一覧取得
	reportList := services.GetReportList(c.Connected().Id, startRange, endRange)

	// 進捗一覧整形処理
	outputReports := c.arrangeReports(*reportList, startRange, endRange)

	common.WriteLog("Top", "pushLoginButton", "End")
	return c.Render(outputReports)
}

func (c PJ) arrangeReports(reportList []models.Report, startRange, endRange time.Time) []forms.Top {
	rtnList := []forms.Top{}

	// 最遅出力日算出
	d := (int)(endRange.Weekday())
	switch (d) {
	case 0:
		d += 6
	default:
		d -= 1
	}
	seekDay := endRange.AddDate(0, 0, d * -1)

	// 整形処理
	newRecord := forms.Top{}

	// 日付が範囲外になるまでレコードを作成する
	for seekDay.After(startRange) {
		newRecord.StartDay = seekDay
		newRecord.EndDay = seekDay.AddDate(0, 0, 6)
		// DBからの取得結果に、該当日付が存在する場合
		i := c.seekSlice(reportList, seekDay);
		if i != -1 {
			fmt.Println("seekDay => " + seekDay.Format("20060102") + " YES")
			newRecord.Status = "作成済み"
			// newRecord.UpdateTime = common.ConvertStrToDay(reportList[i].UpdateDay)
			rtnList = append(rtnList, newRecord)
		} else {
			fmt.Println("seekDay => " + seekDay.Format("20060102") + " NO")
			newRecord.Status = "未作成"
			rtnList = append(rtnList, newRecord)
		}

		seekDay = seekDay.AddDate(0 ,0 ,-7)
	}

	common.WriteLog("Top", "Seikei", "End")

	return rtnList
}

func (c PJ) seekSlice(reportList []models.Report, seekDay time.Time) int {
	for i, value := range reportList {
		fmt.Println("seekSlice " + value.StartDay)
		fmt.Println("seekDay " + seekDay.Format(constant.STR_TYPE_LAYOUT))
		if (value.StartDay == seekDay.Format(constant.STR_TYPE_LAYOUT)) {
			fmt.Println("seekSlice " + value.StartDay + " BINGO")
			return i
		}
	}
	return -1
}

func (c PJ) Edit(inputDay string) revel.Result {
	common.WriteLogStr("Top", "PushEditButton", "Start", "startDay = " + inputDay)
	startDay := common.ConvertStrToDay(inputDay)

	// 該当の進捗情報取得
	report := services.GetReport(c.Connected().Id, startDay)

	if report == nil {
		// 前回の進捗を取得
		lastDay := startDay.AddDate(0, 0, -7)
		lastReport := services.GetReport(c.Connected().Id, lastDay)

		// 取得できた場合は、情報をセット
		if (lastReport != nil) {
			report = c.setLastReport(lastReport)
		}

		fmt.Print("------- Report.Kotei => " + report.Kotei)
	}

	// 日付設定
	days := c.setDays(startDay)

	common.WriteLogStr("Top", "PushEditButton", "End", "startDay = " + startDay.Format(constant.STR_TYPE_LAYOUT))
	return c.Render(report, days, inputDay)
}

// 前週の進捗データを今週の進捗に設定
func (c PJ) setLastReport(lastReport *models.Report) *models.Report{
	report := models.Report{}
	report.Kotei = lastReport.Kotei
	report.Hokoku = lastReport.Hokoku
	report.Kadai = lastReport.Kadai
	report.Other = lastReport.Other
	report.Sagyo8 = lastReport.Sagyo1
	report.Sagyo9 = lastReport.Sagyo2
	report.Sagyo10 = lastReport.Sagyo3
	report.Sagyo11 = lastReport.Sagyo4
	report.Sagyo12 = lastReport.Sagyo5
	report.Sagyo13 = lastReport.Sagyo6
	report.Sagyo14 = lastReport.Sagyo7
	report.Kadotime8 = lastReport.Kadotime1
	report.Kadotime9 = lastReport.Kadotime2
	report.Kadotime10 = lastReport.Kadotime3
	report.Kadotime11 = lastReport.Kadotime4
	report.Kadotime12 = lastReport.Kadotime5
	report.Kadotime13 = lastReport.Kadotime6
	report.Kadotime14 = lastReport.Kadotime7

	return &report
}



// 画面に各週の日付を設定
func (c PJ) setDays (startDay time.Time) []string {

	days := make([]string, 0)

	for index := 0; index < 7; index++ {
		addDay := startDay.AddDate(0, 0, index).Format(constant.DATE_TYPE_LAYOUT)
		days = append(days, addDay)
		fmt.Println("====day" + " ->  " + addDay)
	}

	for index := 7; index < 14; index++ {
		addDay := startDay.AddDate(0, 0, index - 14).Format(constant.DATE_TYPE_LAYOUT)
		days = append(days, addDay)
	}
	return days
}

func (c PJ) Register(inputDay string, report models.Report) revel.Result {

	startDay := common.ConvertStrToDay(inputDay)
	fmt.Println("====inputDay ->  " + inputDay)

	// 曜日チェック
	if (startDay.Weekday() != 1) {
		fmt.Println("===曜日チェックエラー!!!===")
	}

	// 該当の進捗情報取得
	userId := c.Connected().Id
	getReport := services.GetReport(userId, startDay)

	if (getReport == nil) {
		// 登録処理
		fmt.Println("===登録処理===")
		fmt.Println("koutei => " + report.Kotei + "  hokoku => " + report.Hokoku)
		services.InsertReport(userId, inputDay, report)

	} else {
		// 更新処理
		fmt.Println("===更新処理===")
		services.UpdateReport(getReport.Id , userId, inputDay, report)
	}

	// 日にち設定
	days := c.setDays(startDay)

	return c.Render(report, days)

}
