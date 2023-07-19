package handle

import "tabManager/internal/define"

func CreateTaskTable() {
	db.ExecNoTran(define.CREATE_TASK_TABLE)
}
