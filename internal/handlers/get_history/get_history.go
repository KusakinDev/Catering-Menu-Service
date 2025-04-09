package gethistory

import (
	historymodel "github.com/KusakinDev/Catering-Menu-Service/internal/models/history_model"
	"github.com/gin-gonic/gin"
)

func GetHistory(c *gin.Context) (int, []historymodel.History) {
	var history historymodel.History
	code, histories := history.GetAllFromTable()
	if code != 200 {
		return code, []historymodel.History{}
	}
	return 200, histories
}
